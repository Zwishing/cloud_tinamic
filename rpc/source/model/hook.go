package model

import (
	"cloud_tinamic/pkg"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

// gorm的Hook 钩子函数

// AfterCreate 插入文件时，更新size
func (s *Original) AfterCreate(tx *gorm.DB) (err error) {
	// 这里的1代表文件
	if s.StorageCategory == 1 && s.ParentKey != "" {
		parentKey := s.ParentKey
		sizeToAdd := s.Size
		// 使用 WITH RECURSIVE 递归查找父级文件夹并更新其大小
		query := fmt.Sprintf(`
			WITH RECURSIVE parent_folders AS (
				SELECT key, parent_key
				FROM %[1]s
				WHERE key = ?
				UNION ALL
				SELECT s.key, s.parent_key
				FROM %[1]s s
				INNER JOIN parent_folders pf ON s.key = pf.parent_key
				WHERE s.storage_category = 2
			)
			UPDATE %[1]s
			SET size = size + ?
			WHERE key IN (SELECT key FROM parent_folders);
		`, pkg.SourceOriginalTable)

		err = tx.Exec(query, parentKey, sizeToAdd).Error

		if err != nil {
			klog.Errorf("", err)
			return err
		}
	}
	return nil
}

// BeforeDelete 前端删除文件夹前，钩子函数将文件里的子文件及子文件夹里的内容删除
//func (s *Original) BeforeDelete(tx *gorm.DB) (err error) {
//	// 判断当前 key 是否是文件夹
//	fmt.Println(s)
//
//	if s.IsDirectory() {
//		query := fmt.Sprintf(`
//			-- 创建临时表用于存储递归查询结果
//			CREATE TEMP TABLE temp_folder_tree AS
//			WITH RECURSIVE folder_tree AS (
//				SELECT key
//				FROM %[1]s
//				WHERE storage_category=2 AND parent_key = ?
//				UNION ALL
//				SELECT s.key
//				FROM %[1]s s
//				INNER JOIN folder_tree ft ON s.parent_key = ft.key
//			)
//			SELECT key FROM folder_tree;
//
//			-- 使用临时表更新 storage 表中的 deleted_at
//			UPDATE %[1]s
//			SET deleted_at = NOW()
//			WHERE key IN (SELECT key FROM temp_folder_tree);
//
//			-- 使用临时表更新 base_info 表中的 deleted_at
//			UPDATE %[2]s
//			SET deleted_at = NOW()
//			WHERE key IN (SELECT key FROM temp_folder_tree);
//			`, pkg.SourceOriginalTable, pkg.SourceInfoTable)
//		// 执行软删除操作
//		result := tx.Exec(query, s.Key)
//		if result.Error != nil {
//			klog.Errorf("Failed to soft delete contents of folder %s: %v", s.Key, result.Error)
//			return result.Error
//		}
//	}
//
//	return nil
//}

// IsDirectory 判断是否为文件夹的方法
func (s *Original) IsDirectory() bool {
	// 假设有一个字段表示是否为文件夹
	return s.StorageCategory == 2
}

func (s *Original) IsFile() bool {
	// 假设有一个字段表示是否为文件
	return s.StorageCategory == 1
}
