package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/green"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type TypeSearch struct{
    green.Type
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
