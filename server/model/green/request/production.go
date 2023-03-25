package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/green"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type ProductionSearch struct{
    green.Production
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
