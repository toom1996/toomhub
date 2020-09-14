//index.js
//获取应用实例
const app = getApp()


Page({
  data: {
    skeletonShow: true,
    list:[],
    refreshTag: '下拉刷新',
    threshold: 70,
    scrollViewHeight: 0,
    triggered: true
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
    let that = this
    wx.getSystemInfo({
      success: function (res) {
        var scrollViewHeight = 750 * res.windowHeight / res.windowWidth; //rpx
        console.log(res.windowWidth)
        var scrollTop = res.windowWidth * 400 / 750; //矢量转换后的高度
        that.setData({
          scrollViewHeight: scrollViewHeight,
          scrollTop: scrollTop,
          fixedTop: false
        });
      }
    });
    
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
    console.log('ddddd')
  },
  // 图片点击事件
  previewImage: function (event) {
    var src = event.currentTarget.dataset.src;//获取data-src
    var imgList = event.currentTarget.dataset.list;//获取data-list
    //图片预览
    wx.previewImage({
      current: src, // 当前显示图片的http链接
      urls: [
        'http://toomhub.image.23cm.cn/006APoFYly1fowt3eeuk6g306o08g4q3.gif?imageMogr2/auto-orient/format/webp',
        'http://qeiwdcsh5.bkt.clouddn.com/152170904610306200.gif'
      ] // 需要预览的图片http链接列表
    })
  },
  onReachBottom: function () {
    console.log(1111111111111)
  },
  onRefresh : function () {
    app.httpClient.post('/v1/mini/sq/create', {
      'content': this.data.content,
      'image_list': JSON.stringify(obj),
      'tag': this.data.tag == defaultTag ? '' : this.data.tag,
    }).then(res=>{
      let response = res.data
      Toast.clear();
      if (response.code == 200) {
        app.redirectToIndex();
        Toast('发布成功');
      }
      console.log(11111111)
    });

    this.setData({
      triggered: false
    });
  },
  onRestore(e) {
    console.log('onRestore:', e)
  },
  onPulling: function(e) {
    var p = Math.min(e.detail.dy / this.data.threshold, 1)
    
    if (p == 1) {
      this.setData({
        refreshTag: '松开刷新'
      })
    }else{
      this.setData({
        refreshTag: '下拉刷新'
      })
    }
  }
})
