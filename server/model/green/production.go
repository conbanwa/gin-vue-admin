// 自动生成模板Production
package green

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// Production 结构体
type Production struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:;"`
      Type  *int `json:"type" form:"type" gorm:"column:type;comment:;"`
      Spec  string `json:"spec" form:"spec" gorm:"column:spec;comment:;"`
      Detail  string `json:"power" form:"power" gorm:"column:power;comment:;"`
      Pic  string `json:"pic" form:"pic" gorm:"column:pic;comment:;"`
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName Production 表名
func (Production) TableName() string {
  return "production"
}

