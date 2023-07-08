package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/grn"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type OrnSearch struct{
    grn.Orn
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
