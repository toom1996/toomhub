// pages/components/tabbar/tabbar.js
Component({
  /**
   * 组件的属性列表
   */
  properties: {
  },

  /**
   * 组件的初始数据
   */
  data: {
    active: 'pages/index/index'
  },
  // 组件生命周期：一打开页面就执行
  attached:function (){
    console.log('tabbar')
    let pages = getCurrentPages()
    let currentPages = pages[pages.length - 1]
    console.log('currentpages', currentPages.route);
    this.setData({
      active: currentPages.route
    })
  },
  /**
   * 组件的方法列表
   */
  methods: {
    navigationSwitch:function (event) {
      console.log('event.detail', event.detail)
      wx.reLaunch({
        url: '/' + event.detail,
      })
    }
  }
})
