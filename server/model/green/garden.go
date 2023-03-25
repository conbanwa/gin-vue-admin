// 自动生成模板Garden
package green

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// Garden 结构体
type Garden struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:;"`
      Type  *int `json:"type" form:"type" gorm:"column:type;comment:;"`
      Spec  string `json:"spec" form:"spec" gorm:"column:spec;comment:;"`
      Power  *int `json:"power" form:"power" gorm:"column:power;comment:;"`
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName Garden 表名
func (Garden) TableName() string {
  return "garden"
}

