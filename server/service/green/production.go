package green

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/green"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    greenReq "github.com/flipped-aurora/gin-vue-admin/server/model/green/request"
    "gorm.io/gorm"
)

type ProductionService struct {
}

// CreateProduction 创建Production记录
// Author [piexlmax](https://github.com/piexlmax)
func (productionService *ProductionService) CreateProduction(production green.Production) (err error) {
	err = global.GVA_DB.Create(&production).Error
	return err
}

// DeleteProduction 删除Production记录
// Author [piexlmax](https://github.com/piexlmax)
func (productionService *ProductionService)DeleteProduction(production green.Production) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&green.Production{}).Where("id = ?", production.ID).Update("deleted_by", production.DeletedBy).Error; err != nil {
              return err
        }
        if err = tx.Delete(&production).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteProductionByIds 批量删除Production记录
// Author [piexlmax](https://github.com/piexlmax)
func (productionService *ProductionService)DeleteProductionByIds(ids request.IdsReq,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&green.Production{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", ids.Ids).Delete(&green.Production{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateProduction 更新Production记录
// Author [piexlmax](https://github.com/piexlmax)
func (productionService *ProductionService)UpdateProduction(production green.Production) (err error) {
	err = global.GVA_DB.Save(&production).Error
	return err
}

// GetProduction 根据id获取Production记录
// Author [piexlmax](https://github.com/piexlmax)
func (productionService *ProductionService)GetProduction(id uint) (production green.Production, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&production).Error
	return
}

// GetProductionInfoList 分页获取Production记录
// Author [piexlmax](https://github.com/piexlmax)
func (productionService *ProductionService)GetProductionInfoList(info greenReq.ProductionSearch) (list []green.Production, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&green.Production{})
    var productions []green.Production
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&productions).Error
	return  productions, total, err
}
