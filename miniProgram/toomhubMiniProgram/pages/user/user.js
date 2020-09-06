//index.js
//获取应用实例
const app = getApp()

Page({
  data: {
    userInfo: app.globalData.userInfo
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
    console.log('21212121212')
    console.log(app.globalData.userInfo)
    this.setData({
      userInfo: app.globalData.userInfo
    })
  }
})