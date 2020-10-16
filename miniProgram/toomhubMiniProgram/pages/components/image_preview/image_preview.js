const app = getApp()
let myStyle = `
--tooom__tag-top:
`

Page({
  onShareAppMessage() {
    return {
      title: 'swiper',
      path: 'page/component/pages/swiper/swiper'
    }
  },
  onLoad(options) {
    console.log(options.index)

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

    this.setData({
      index: options.index,
      image:options.list.split(",")
    })
    
  },

  /**
   * 获取顶部固定高度
   */
  attached: function () {
    this.setData({
      navHeight: app.globalData.navHeight,
      navTop: app.globalData.navTop,
    })
  },
  data: {
    flag: false,
    // 自定义顶部导航
    navHeight: app.globalData.navHeight,
    navTop: app.globalData.navTop,

    show: false,
    actions: [
      {
        name: '发送给好友',
        openType: 'shareToFriendHandle',
      },
      {
        name: '选项',
      },
      {
        name: '选项',
        subname: '副文本',
        openType: 'share',
      },
      
    ],
    interval: 2000,
    duration: 200,
    index:0,
    image:[
    ]
  },
  //长按
  imageLongTapHandle(){
    this.setData({
      flag: true,
      show: true,
    })
  },

  onClose() {
    this.setData({
      show: false,
    })
  },

  onSelect(event) {
    console.log(event.detail);
  },

  imageClickHandle() {
    if (this.data.flag == false) {
      wx.navigateBack({
        delta: -1
      })
    }else{
      this.setData({
        flag: false,
      })
    }
  },
})