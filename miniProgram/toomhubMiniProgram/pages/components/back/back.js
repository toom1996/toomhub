// pages/components/back/back.js
const app = getApp()
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
    navTop: app.globalData.navTop
  },

  /**
   * 组件的方法列表
   */
  methods: {
    //跳转到首页
    redirectToIndex() {
      app.redirectToIndex()
    }
  }
})
