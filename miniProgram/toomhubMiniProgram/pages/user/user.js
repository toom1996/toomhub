//index.js
//获取应用实例
const app = getApp()

Page({
  data: {
    userInfo : null,
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
    if (app.globalData.userInfo) {
      this.setData({
        userInfo: app.globalData.userInfo,
      })
    } else {
      this.setData({
        userInfo: null,
      })
    }
  }
})