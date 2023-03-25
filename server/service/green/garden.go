package green

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/green"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    greenReq "github.com/flipped-aurora/gin-vue-admin/server/model/green/request"
    "gorm.io/gorm"
)

type GardenService struct {
}

// CreateGarden 创建Garden记录
// Author [piexlmax](https://github.com/piexlmax)
func (gardenService *GardenService) CreateGarden(garden green.Garden) (err error) {
	err = global.GVA_DB.Create(&garden).Error
	return err
}

// DeleteGarden 删除Garden记录
// Author [piexlmax](https://github.com/piexlmax)
func (gardenService *GardenService)DeleteGarden(garden green.Garden) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&green.Garden{}).Where("id = ?", garden.ID).Update("deleted_by", garden.DeletedBy).Error; err != nil {
              return err
        }
        if err = tx.Delete(&garden).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteGardenByIds 批量删除Garden记录
// Author [piexlmax](https://github.com/piexlmax)
func (gardenService *GardenService)DeleteGardenByIds(ids request.IdsReq,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&green.Garden{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", ids.Ids).Delete(&green.Garden{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateGarden 更新Garden记录
// Author [piexlmax](https://github.com/piexlmax)
func (gardenService *GardenService)UpdateGarden(garden green.Garden) (err error) {
	err = global.GVA_DB.Save(&garden).Error
	return err
}

// GetGarden 根据id获取Garden记录
// Author [piexlmax](https://github.com/piexlmax)
func (gardenService *GardenService)GetGarden(id uint) (garden green.Garden, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&garden).Error
	return
}

// GetGardenInfoList 分页获取Garden记录
// Author [piexlmax](https://github.com/piexlmax)
func (gardenService *GardenService)GetGardenInfoList(info greenReq.GardenSearch) (list []green.Garden, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&green.Garden{})
    var gardens []green.Garden
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&gardens).Error
	return  gardens, total, err
}
