const app = getApp()

Page({
  data: {
    //判断小程序的API，回调，参数，组件等是否在当前版本可用。
    canIUse: wx.canIUse('button.open-type.getUserInfo'),
    isHide: false
  },

  bindGetUserInfo: function (e) {
    if (e.detail.userInfo) {
      //用户按了允许授权按钮
      var that = this;
      // 获取到用户的信息
      // app.globalData.userInfo = e.detail.userInfo
      //获取已经授权的权限
      wx.getUserInfo({
        success: function (userinfo) {
          wx.login({
            success: function (loginRes) {
              console.log('wx.login----->', loginRes)
              app.httpClient.post('/v1/mini/user/login', {
                code: loginRes.code,
                userInfo: JSON.stringify(userinfo.userInfo)
              }).then(res => {
                let data = res.data
                console.log(data)
                if (data.code == 200) {
                  //登陆成功后缓存token, refreshToken, nickname
                 app.setCache('userInfo', {
                   'avatarUrl': data.data.avatar_url,
                   'nickName': data.data.nick_name,
                   'token': data.data.token,
                   'refreshToken': data.data.refresh_token
                 })
                }
              }) 
            }
          })
        }
      })
      wx.navigateBack({
        delta: 1
      })
    } else {
      //用户按了拒绝按钮
      wx.showModal({
        title: '警告',
        content: '您点击了拒绝授权，将无法进入小程序，请授权之后再进入!!!',
        showCancel: false,
        confirmText: '返回授权',
        success: function (res) {
          // 用户没有授权成功，不需要改变 isHide 的值
          if (res.confirm) {
            console.log('用户点击了“返回授权”');
          }
        }
      });
    }
  }
})