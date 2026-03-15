import request from '@/utils/request'
// 获取文章频道列表
export const artGetChannelsService = () => request.get('/my/cate/list')
// 添加文章频道
export const artAddChannelsService = (data) =>
  request.post('/my/cate/add', data)
//编辑文章频道
export const artEditChannelsService = (data) =>
  request.put('/my/cate/info', data)
// 删除文章频道
export const artDelChannelsService = (id) =>
  request.delete(`/my/cate/del`, { params: { id } })
export const artGetListService = (params) =>
  request.get('/my/article/list', { params })
export const artPubArticleService = (data) =>
  request.post('/my/article/add', data)
export const artGetDetailService = (id) =>
  request.get('my/article/info', { params: { id } })
export const artEditService = (data) => request.put('my/article/info', data)
export const artDelService = (id) =>
  request.delete('my/article/info', { params: { id } })
