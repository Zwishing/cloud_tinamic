package model

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

func (v *Vector) AfterCreate(tx *gorm.DB) (err error) {
	// 直接更新 Vector 表中的几何范围和中心点
	err = tx.Table(VectorTable).
		Where("source_key = ?", v.SourceKey). // 根据 SourceKey 进行过滤
		Updates(map[string]interface{}{
			"bbox":   gorm.Expr("ST_Extent(geom)"),   // 更新边界框
			"center": gorm.Expr("ST_Centroid(geom)"), // 更新中心点
		}).Error

	if err != nil {
		klog.Errorf("Failed to update bbox and center for source_key %s: %v", v.SourceKey, err)
		return err
	}

	klog.Infof("Successfully updated bbox and center for source_key %s", v.SourceKey)
	return nil
}
