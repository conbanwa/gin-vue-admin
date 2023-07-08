// 自动生成模板Orn
package grn

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	"gorm.io/datatypes"
)

// Orn 结构体
type Orn struct {
      global.GVA_MODEL
      Pic  string `json:"pic" form:"pic" gorm:"column:pic;comment:;"`
      File  datatypes.JSON `json:"file" form:"file" gorm:"column:file;comment:;"`
}


// TableName Orn 表名
func (Orn) TableName() string {
  return "orn"
}

