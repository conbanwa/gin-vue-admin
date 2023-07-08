package grn

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/grn"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    grnReq "github.com/flipped-aurora/gin-vue-admin/server/model/grn/request"
)

type OrnService struct {
}

// CreateOrn 创建Orn记录
// Author [piexlmax](https://github.com/piexlmax)
func (ornnService *OrnService) CreateOrn(ornn *grn.Orn) (err error) {
	err = global.GVA_DB.Create(ornn).Error
	return err
}

// DeleteOrn 删除Orn记录
// Author [piexlmax](https://github.com/piexlmax)
func (ornnService *OrnService)DeleteOrn(ornn grn.Orn) (err error) {
	err = global.GVA_DB.Delete(&ornn).Error
	return err
}

// DeleteOrnByIds 批量删除Orn记录
// Author [piexlmax](https://github.com/piexlmax)
func (ornnService *OrnService)DeleteOrnByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]grn.Orn{},"id in ?",ids.Ids).Error
	return err
}

// UpdateOrn 更新Orn记录
// Author [piexlmax](https://github.com/piexlmax)
func (ornnService *OrnService)UpdateOrn(ornn grn.Orn) (err error) {
	err = global.GVA_DB.Save(&ornn).Error
	return err
}

// GetOrn 根据id获取Orn记录
// Author [piexlmax](https://github.com/piexlmax)
func (ornnService *OrnService)GetOrn(id uint) (ornn grn.Orn, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&ornn).Error
	return
}

// GetOrnInfoList 分页获取Orn记录
// Author [piexlmax](https://github.com/piexlmax)
func (ornnService *OrnService)GetOrnInfoList(info grnReq.OrnSearch) (list []grn.Orn, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&grn.Orn{})
    var ornns []grn.Orn
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&ornns).Error
	return  ornns, total, err
}
