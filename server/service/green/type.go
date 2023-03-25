package green

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/green"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    greenReq "github.com/flipped-aurora/gin-vue-admin/server/model/green/request"
    "gorm.io/gorm"
)

type TypeService struct {
}

// CreateType 创建Type记录
// Author [piexlmax](https://github.com/piexlmax)
func (type_Service *TypeService) CreateType(type_ green.Type) (err error) {
	err = global.GVA_DB.Create(&type_).Error
	return err
}

// DeleteType 删除Type记录
// Author [piexlmax](https://github.com/piexlmax)
func (type_Service *TypeService)DeleteType(type_ green.Type) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&green.Type{}).Where("id = ?", type_.ID).Update("deleted_by", type_.DeletedBy).Error; err != nil {
              return err
        }
        if err = tx.Delete(&type_).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteTypeByIds 批量删除Type记录
// Author [piexlmax](https://github.com/piexlmax)
func (type_Service *TypeService)DeleteTypeByIds(ids request.IdsReq,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&green.Type{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", ids.Ids).Delete(&green.Type{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateType 更新Type记录
// Author [piexlmax](https://github.com/piexlmax)
func (type_Service *TypeService)UpdateType(type_ green.Type) (err error) {
	err = global.GVA_DB.Save(&type_).Error
	return err
}

// GetType 根据id获取Type记录
// Author [piexlmax](https://github.com/piexlmax)
func (type_Service *TypeService)GetType(id uint) (type_ green.Type, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&type_).Error
	return
}

// GetTypeInfoList 分页获取Type记录
// Author [piexlmax](https://github.com/piexlmax)
func (type_Service *TypeService)GetTypeInfoList(info greenReq.TypeSearch) (list []green.Type, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&green.Type{})
    var type_s []green.Type
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&type_s).Error
	return  type_s, total, err
}
