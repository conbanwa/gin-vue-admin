import service from '@/utils/request'

// @Tags Garden
// @Summary 创建Garden
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Garden true "创建Garden"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /garden/createGarden [post]
export const createGarden = (data) => {
  return service({
    url: '/garden/createGarden',
    method: 'post',
    data
  })
}

// @Tags Garden
// @Summary 删除Garden
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Garden true "删除Garden"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /garden/deleteGarden [delete]
export const deleteGarden = (data) => {
  return service({
    url: '/garden/deleteGarden',
    method: 'delete',
    data
  })
}

// @Tags Garden
// @Summary 删除Garden
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Garden"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /garden/deleteGarden [delete]
export const deleteGardenByIds = (data) => {
  return service({
    url: '/garden/deleteGardenByIds',
    method: 'delete',
    data
  })
}

// @Tags Garden
// @Summary 更新Garden
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Garden true "更新Garden"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /garden/updateGarden [put]
export const updateGarden = (data) => {
  return service({
    url: '/garden/updateGarden',
    method: 'put',
    data
  })
}

// @Tags Garden
// @Summary 用id查询Garden
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Garden true "用id查询Garden"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /garden/findGarden [get]
export const findGarden = (params) => {
  return service({
    url: '/garden/findGarden',
    method: 'get',
    params
  })
}

// @Tags Garden
// @Summary 分页获取Garden列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Garden列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /garden/getGardenList [get]
export const getGardenList = (params) => {
  return service({
    url: '/garden/getGardenList',
    method: 'get',
    params
  })
}
