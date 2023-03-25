// 自动生成模板Type
package green

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// Type 结构体
type Type struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:;"`
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName Type 表名
func (Type) TableName() string {
  return "type"
}

