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
    
    //请求首页接口
    app.httpClient.get('/v1/mini/sq/index?last_id=100&page=10').then(res=>{
      var responseData = res.data.data
      this.setData({
        list: responseData,
        skeletonShow: false
      })
      console.log(this.data.list)
    })
  },
  getUserInfo: function(e) {
    console.log(e)
    app.globalData.userInfo = e.detail.userInfo
    this.setData({
      userInfo: e.detail.userInfo,
      hasUserInfo: true
    })
  },
  add: function () {
    wx.navigateTo({
      url: '../square_add/square_add'
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
        'http://qeiwdcsh5.bkt.clouddn.com/11eaca620fff1761c041093c9484a6b9.gif',
        'http://qeiwdcsh5.bkt.clouddn.com/152170904610306200.gif'
      ] // 需要预览的图片http链接列表
    })
  }
})
