package repo

import (
	"cloud_tinamic/pkg"
	"cloud_tinamic/rpc/service_collection/model"
	"cloud_tinamic/rpc/vector_service/repo"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

type ServiceCollectionRepo interface {
	AddCollection(cloudOptimizedKey string, serviceKey string, title string) (string, error)
	UpdateThumbnail(serviceKey string, thumbnail []byte) (bool, error)
}
type ServiceCollectionRepoImpl struct {
	DB *gorm.DB
}

func NewServiceCollectionRepoImpl(db *gorm.DB) ServiceCollectionRepo {
	return &ServiceCollectionRepoImpl{
		DB: db,
	}
}

func (s *ServiceCollectionRepoImpl) GetCollections(page, pageSize int) {
}

func (s *ServiceCollectionRepoImpl) AddCollection(cloudOptimizedKey string, serviceKey string, title string) (string, error) {
	// 开始事务
	tx := s.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 定义矢量信息结构
	type VectorInfo struct {
		RecordCount  int64  `gorm:"column:record_count"`
		GeometryType string `gorm:"column:geometry_type"`
		SRID         int32  `gorm:"column:srid"`
	}

	// 查询矢量信息
	var vectorInfo VectorInfo
	query := fmt.Sprintf(`
		SELECT 
			COUNT(*) AS record_count
		FROM %s."%s"`, pkg.VectorSchema, cloudOptimizedKey)

	row := s.DB.Raw(query).Row()
	if err := row.Scan(&vectorInfo.RecordCount); err != nil {
		klog.Errorf("查询矢量数据时出错: %v", err)
		return "", tx.Rollback().Error
	}

	query = fmt.Sprintf(`
		SELECT  
			GeometryType(geom) AS geometry_type, 
			ST_SRID(geom) AS srid 
		FROM %s."%s"`, pkg.VectorSchema, cloudOptimizedKey)

	row = s.DB.Raw(query).Row()
	if err := row.Scan(&vectorInfo.GeometryType, &vectorInfo.SRID); err != nil {
		klog.Errorf("查询矢量数据时出错: %v", err)
		return "", tx.Rollback().Error
	}

	// 获取字段信息
	fieldCount, properties, err := s.GetFieldInfo(cloudOptimizedKey)
	if err != nil {
		klog.Errorf("查询字段信息时出错: %v", err)
		return "", tx.Rollback().Error
	}

	// 构造信息模型
	infos := model.Info{
		ServiceKey:        serviceKey,
		Title:             title,
		CloudOptimizedKey: cloudOptimizedKey,
		SourceSchema:      pkg.VectorSchema,
		SourceCategory:    1,
		Srid:              vectorInfo.SRID,
	}

	// 构造集合模型
	collection := model.Collection{
		ServiceKey:      serviceKey,
		Title:           title,
		ServiceCategory: 0,
		Srid:            vectorInfo.SRID,
	}

	// 构造矢量数据模型
	vector := model.Vector{
		CloudOptimizedKey: cloudOptimizedKey,
		Title:             title,
		GeometryField:     pkg.GeometryFieldName,
		GeometryCategory:  vectorInfo.GeometryType,
		RecordCount:       vectorInfo.RecordCount,
		FieldCount:        fieldCount,
		Properties:        properties,
	}

	// 插入信息记录
	if err := tx.Table(pkg.ServiceInfoTable).Create(&infos).Error; err != nil {
		klog.Errorf("插入信息记录时出错: %v", err)
		return "", tx.Rollback().Error
	}
	// 插入集合记录
	if err := tx.Table(pkg.ServiceCollectionTable).Create(&collection).Error; err != nil {
		klog.Errorf("插入集合记录时出错: %v", err)
		return "", tx.Rollback().Error
	}
	// 插入向量记录
	if err := tx.Table(pkg.ServiceVectorTable).Create(&vector).Error; err != nil {
		klog.Errorf("插入向量记录时出错: %v", err)
		return "", tx.Rollback().Error
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		klog.Errorf("提交事务时出错: %v", err)
		return "", tx.Rollback().Error
	}

	klog.Infof("成功添加集合，服务键为: %s", serviceKey)
	return serviceKey, nil
}

func (s *ServiceCollectionRepoImpl) GetCollectionByServiceKey(serviceKey string) repo.Layer {

	return &repo.LayerTable{
		ID:             serviceKey,
		Schema:         pkg.VectorSchema,
		Table:          "",
		Description:    "",
		Properties:     nil,
		GeometryType:   "",
		IDColumn:       "service_key",
		GeometryColumn: "geom",
		Srid:           0,
	}
}

func (s *ServiceCollectionRepoImpl) GetFieldInfo(cloudOptimizedKey string) (int64, map[string]string, error) {
	var fieldCount int64
	// 定义一个 processor.thrift 来存储字段名称和类型
	fieldInfo := make(map[string]string)

	// 查询字段名称和类型
	rows, err := s.DB.Table("information_schema.columns").
		Select("column_name, data_type").
		Where("table_name = ? AND table_schema = ?", cloudOptimizedKey, pkg.VectorSchema).
		Rows()
	if err != nil {
		klog.Errorf("查询字段信息时出错: %v", err) // 使用 klog 记录错误
		return 0, nil, err
	}
	defer rows.Close()

	// 统计字段数量并填充字段信息
	for rows.Next() {
		var columnName string
		var dataType string
		if err := rows.Scan(&columnName, &dataType); err != nil {
			klog.Errorf("扫描字段信息时出错: %v", err) // 使用 klog 记录错误
			return 0, nil, err
		}
		fieldInfo[columnName] = dataType // 将字段名称和类型存储到 processor.thrift 中
		fieldCount++                     // 统计字段数量
	}

	// 检查是否有错误发生
	if err = rows.Err(); err != nil {
		klog.Errorf("查询结果集时出错: %v", err) // 使用 klog 记录错误
		return 0, nil, err
	}

	return fieldCount, fieldInfo, nil
}

func (s *ServiceCollectionRepoImpl) UpdateThumbnail(serviceKey string, thumbnail []byte) (bool, error) {
	result := s.DB.Table(pkg.ServiceCollectionTable).
		Where("service_key = ?", serviceKey).
		Update("thumbnail", thumbnail)

	if result.Error != nil {
		klog.Errorf("更新缩略图失败, serviceKey: %s, 错误: %v", serviceKey, result.Error)
		return false, result.Error
	}

	if result.RowsAffected == 0 {
		klog.Warnf("未找到匹配的记录, serviceKey: %s", serviceKey)
		return false, nil
	}

	klog.Infof("成功更新缩略图, serviceKey: %s", serviceKey)
	return true, nil
}
