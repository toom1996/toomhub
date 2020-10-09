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
    console.log(app.globalData.userInfo)
    if (app.globalData.userInfo != null) {
      app.httpClient.get(app.getApi('getUserInfo')).then(res=> {
        console.log(res.data.data.likes_count)
        console.log(app.globalData.userInfo)
        app.globalData.userInfo.likes_count = res.data.data.likes_count

        this.setData({
          userInfo: app.globalData.userInfo
        })
      })
    }
    this.setData({
      userInfo: app.globalData.userInfo
    })
    console.log(this.data.userInfo)
  }
})