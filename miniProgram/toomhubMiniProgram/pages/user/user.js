//index.js
//获取应用实例
const app = getApp()

Page({
  data: {
    userInfo: app.globalData.userInfo,
    tag: {}
  },
  onReady() {
    //是否登陆
    app.isLogin()
    // 隐藏小房子
    wx.hideHomeButton();
  },
  //跳转到登陆界面
  userLogin: function (event) {
    wx.navigateTo({
      url: '/pages/login/login'
    })
  },

  onShow: function () {
    if (app.globalData.userInfo != null) {
      app.httpClient.get(app.getApi('getUserInfo')).then(res=> {
        app.globalData.userInfo.likes_count = res.data.data.likes_count
        this.setData({
          userInfo: app.globalData.userInfo,
          tag: res.data.data.tag
        })
      })
    }
  }
})