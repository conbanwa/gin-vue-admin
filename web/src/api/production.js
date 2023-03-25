import service from '@/utils/request'

// @Tags Production
// @Summary 创建Production
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Production true "创建Production"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /production/createProduction [post]
export const createProduction = (data) => {
  return service({
    url: '/production/createProduction',
    method: 'post',
    data
  })
}

// @Tags Production
// @Summary 删除Production
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Production true "删除Production"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /production/deleteProduction [delete]
export const deleteProduction = (data) => {
  return service({
    url: '/production/deleteProduction',
    method: 'delete',
    data
  })
}

// @Tags Production
// @Summary 删除Production
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Production"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /production/deleteProduction [delete]
export const deleteProductionByIds = (data) => {
  return service({
    url: '/production/deleteProductionByIds',
    method: 'delete',
    data
  })
}

// @Tags Production
// @Summary 更新Production
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Production true "更新Production"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /production/updateProduction [put]
export const updateProduction = (data) => {
  return service({
    url: '/production/updateProduction',
    method: 'put',
    data
  })
}

// @Tags Production
// @Summary 用id查询Production
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Production true "用id查询Production"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /production/findProduction [get]
export const findProduction = (params) => {
  return service({
    url: '/production/findProduction',
    method: 'get',
    params
  })
}

// @Tags Production
// @Summary 分页获取Production列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Production列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /production/getProductionList [get]
export const getProductionList = (params) => {
  return service({
    url: '/production/getProductionList',
    method: 'get',
    params
  })
}
