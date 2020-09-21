//index.js
//获取应用实例
const app = getApp()
let myStyle = `
--tooom__tag-top:
`

Page({
  data: {
    // 自定义顶部导航
    navHeight: app.globalData.navHeight,
    navTop: app.globalData.navTop,
    viewData: {
      style: myStyle
    },
    skeletonShow: true,
    data:[],
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
    this.setData({ 'viewData.style': myStyle + '40px;' })
    this.refreshIndex();
  },
  getUserInfo: function(e) {
    console.log(e)
    app.globalData.userInfo = e.detail.userInfo
    this.setData({
      userInfo: e.detail.userInfo,
      hasUserInfo: true
    })
  },
  addHandle: function () {
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
    console.log(src)
    var imgList = event.currentTarget.dataset.list;//获取data-list
    console.log(imgList)
    //图片预览
    wx.previewImage({
      current: src, // 当前显示图片的http链接
      urls: imgList // 需要预览的图片http链接列表
    })
  },
  onReachBottom: function () {
    console.log(1111111111111)
  },
  onRefresh : function () {
    this.refreshIndex()
    this.setData({
      triggered: false
    });
  },
  onRestore(e) {
    console.log('onRestore:', e)
  },
  onPulling: function(e) {
    var p = Math.min(e.detail.dy / this.data.threshold, 1)
    
    if (p >= 1) {
      this.setData({
        refreshTag: '松开刷新'
      })
    }else{
      this.setData({
        refreshTag: '下拉刷新'
      })
    }
  },
  refreshIndex: function () {
    this.setData({
      skeletonShow: false
    })
    //请求首页接口
    app.httpClient.get(app.getApi('SQ_INDEX') + '?last_id=100&page=1').then(res => {
      var responseData = res.data.data
      this.setData({
        data: responseData.list,
      })
      console.log(this.data.list)
    })
  },
  onPullDownRefresh: function () {
    console.log('onPullDownRefresh')
  },

  goBack: function () {
    let pages = getCurrentPages();   //获取小程序页面栈
    let beforePage = pages[pages.length - 2];  //获取上个页面的实例对象
    beforePage.setData({      //直接修改上个页面的数据（可通过这种方式直接传递参数）
      txt: '修改数据了'
    })
    beforePage.goUpdate();   //触发上个页面自定义的go_update方法
    wx.navigateBack({         //返回上一页  
      delta: 1
    })
  },
  /**
   * 获取顶部固定高度
   */
  attached: function () {
    this.setData({
      navHeight: App.globalData.navHeight,
      navTop: App.globalData.navTop,
    })
  },

  test: function (e) {
    let newList = this.data.data;
    let isLike = e.target.dataset.like;
    if (isLike === 0) {
      isLike = 1;
    }else {
      isLike = 0;
    }
    newList[e.target.dataset.index].is_like = isLike
    this.setData({
      data: newList
    })
    
  }
})
