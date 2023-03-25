import service from '@/utils/request'

// @Tags Type
// @Summary 创建Type
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Type true "创建Type"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /type_/createType [post]
export const createType = (data) => {
  return service({
    url: '/type_/createType',
    method: 'post',
    data
  })
}

// @Tags Type
// @Summary 删除Type
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Type true "删除Type"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /type_/deleteType [delete]
export const deleteType = (data) => {
  return service({
    url: '/type_/deleteType',
    method: 'delete',
    data
  })
}

// @Tags Type
// @Summary 删除Type
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Type"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /type_/deleteType [delete]
export const deleteTypeByIds = (data) => {
  return service({
    url: '/type_/deleteTypeByIds',
    method: 'delete',
    data
  })
}

// @Tags Type
// @Summary 更新Type
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Type true "更新Type"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /type_/updateType [put]
export const updateType = (data) => {
  return service({
    url: '/type_/updateType',
    method: 'put',
    data
  })
}

// @Tags Type
// @Summary 用id查询Type
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Type true "用id查询Type"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /type_/findType [get]
export const findType = (params) => {
  return service({
    url: '/type_/findType',
    method: 'get',
    params
  })
}

// @Tags Type
// @Summary 分页获取Type列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Type列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /type_/getTypeList [get]
export const getTypeList = (params) => {
  return service({
    url: '/type_/getTypeList',
    method: 'get',
    params
  })
}
