//app.js
const REQUEST_HOST = 'http://127.0.0.1:8080'

App({
  onLaunch: function () {
    // 展示本地存储能力
    var logs = wx.getStorageSync('logs') || []
    logs.unshift(Date.now())
    wx.setStorageSync('logs', logs)
  },
  globalData: {
    userInfo: null, //全局用户信息,
    request_host: 'http://127.0.0.1:8080'
  },
  httpClient: {
    request: function (method, url, data) {
      //返回一个promise实例
      return new Promise((resolve, reject) => {
        wx.request({
          url: getApp().globalData.request_host + url,
          data: data,
          method: method,
          header: {
            'content-type': 'application/x-www-form-urlencoded' // 默认值
          },
          success(res) {
            resolve(res)
          },
          fail(res) {
            reject(res)
          },
          complete(res) {
            console.log('1223143234')
          }
        })
      })
    },
    get: function (url) {
      return this.request('GET', url);
    },
    post: function (url, data) {
      return this.request('POST', url, data);
    }
  },
  //判断用户是否登陆, 未登陆则跳转到登陆界面
  isLogin:function () {
    console.log(wx.getStorageSync('userInfo'))
    if (this.globalData.userInfo === null && wx.getStorageSync('userInfo') === '') {
      wx.navigateTo({
        url: '/pages/login/login'
      })
      return false;
    }
    return true;
  }
})