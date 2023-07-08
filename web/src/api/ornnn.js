import service from '@/utils/request'

// @Tags Orn
// @Summary 创建Orn
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Orn true "创建Orn"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ornn/createOrn [post]
export const createOrn = (data) => {
  return service({
    url: '/ornn/createOrn',
    method: 'post',
    data
  })
}

// @Tags Orn
// @Summary 删除Orn
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Orn true "删除Orn"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ornn/deleteOrn [delete]
export const deleteOrn = (data) => {
  return service({
    url: '/ornn/deleteOrn',
    method: 'delete',
    data
  })
}

// @Tags Orn
// @Summary 删除Orn
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Orn"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ornn/deleteOrn [delete]
export const deleteOrnByIds = (data) => {
  return service({
    url: '/ornn/deleteOrnByIds',
    method: 'delete',
    data
  })
}

// @Tags Orn
// @Summary 更新Orn
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Orn true "更新Orn"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ornn/updateOrn [put]
export const updateOrn = (data) => {
  return service({
    url: '/ornn/updateOrn',
    method: 'put',
    data
  })
}

// @Tags Orn
// @Summary 用id查询Orn
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Orn true "用id查询Orn"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ornn/findOrn [get]
export const findOrn = (params) => {
  return service({
    url: '/ornn/findOrn',
    method: 'get',
    params
  })
}

// @Tags Orn
// @Summary 分页获取Orn列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Orn列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ornn/getOrnList [get]
export const getOrnList = (params) => {
  return service({
    url: '/ornn/getOrnList',
    method: 'get',
    params
  })
}
