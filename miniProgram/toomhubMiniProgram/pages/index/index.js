//index.js
//获取应用实例
const app = getApp()

Page({
  data: {
    skeletonShow: true,
    list:[]
  },
  navigationSwitch: function(event) {
    wx.navigateTo({
      url: '../user/user'
    })
  },
  //事件处理函数
  bindViewTap: function() {
    wx.navigateTo({
      url: '../logs/logs'
    })
  },
  onLoad: function () {
    console.log(123123)
    
    // //请求首页接口
    // app.httpClient.get('http://192.168.10.207:9501/mini/index').then(res=>{
    //   var responseData = res.data.data
    //   this.setData({
    //     list: responseData,
    //     skeletonShow: true
    //   })
    //   console.log(this.data.list)
    // })
  },
  getUserInfo: function(e) {
    console.log(e)
    app.globalData.userInfo = e.detail.userInfo
    this.setData({
      userInfo: e.detail.userInfo,
      hasUserInfo: true
    })
  },
  // 滚动至低端事件
  ScrollLower: function () {
    console.log(21212121)
  },
  // 图片点击事件
  previewImage: function (event) {
    var src = event.currentTarget.dataset.src;//获取data-src
    var imgList = event.currentTarget.dataset.list;//获取data-list
    //图片预览
    wx.previewImage({
      current: src, // 当前显示图片的http链接
      urls: [
        'http://qeiwdcsh5.bkt.clouddn.com/006APoFYly1fowt3eeuk6g306o08g4q3.gif',
        'http://qeiwdcsh5.bkt.clouddn.com/152170904610306200.gif'
      ] // 需要预览的图片http链接列表
    })
  }
})
