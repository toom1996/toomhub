var api = {
  requestHost: 'http://192.168.31.88:8080', //host
  tokenCheck: '/v1/mini/user/token-checker', //token过期验证
  squareIndex: '/v1/mini/sq/index', //说说首页
  squareCreate: '/v1/mini/sq/create', //发布一条说说
  squareTagSearch: '/v1/mini/sq/tag-search', //标签搜索
  squareLike: '/v1/mini/sq/like', //说说点赞
  getSession: '/v1/mini/user/get-session', //获取小程序session
  login: '/v1/mini/user/login', //用户登陆
  getUserInfo: '/v1/mini/user/get-info', //刷新用户粉丝数点赞数啥的
}

module.exports.api = api;