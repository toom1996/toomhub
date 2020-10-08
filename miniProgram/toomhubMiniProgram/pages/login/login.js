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
      wx.login({
        success: function (loginRes) {
          wx.showToast({
            title: '加载中...',
            icon: 'loading',
            mask: true,
          });
          app.httpClient.get(app.getApi('getSession') + '?code=' + loginRes.code).then(res => {
            res = res.data;
            if (res.code == 200 && res.errcode == 0) {
              wx.getUserInfo({
                success: function (userinfo) {
                  console.log(userinfo)
                  app.httpClient.post(app.getApi('login'), {
                    rawData: userinfo.rawData,
                    signature: userinfo.signature,
                    encryptedData: userinfo.encryptedData,
                    iv: userinfo.iv,
                    authKey: res.data.authKey,
                  }).then(res => {
                    //登陆成功后缓存token, refreshToken, nickname
                    let info = res.data.data
                    console.log(info)
                    app.setCache('userInfo', info)
                    app.globalData.userInfo = info
                    wx.hideToast();
                    app.globalData.forceRefresh = true
                    wx.navigateBack({
                      delta: 1
                    })
                  })
                }
              })
            } else {
              console.log(res.message)
            }
          })
        }
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