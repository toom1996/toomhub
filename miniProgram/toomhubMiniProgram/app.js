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
    userInfo: null //全局用户信息
  },
  httpClient: {
    request: function (method, url, data) {
      //返回一个promise实例
      return new Promise((resolve, reject) => {
        console.log('REQUEST_HOST', REQUEST_HOST)
        wx.request({
          url: REQUEST_HOST + url,
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
  isLogin:function () {
    console.log(wx.getStorageSync('userInfo'))
    if (this.globalData.userInfo === null && wx.getStorageSync('userInfo') === '') {
      wx.navigateTo({
        url: '/pages/login/login'
      })
    }
    // wx.getSetting({
    //   success: function (settingRes) {
    //     if (settingRes.authSetting['scope.userInfo']) {
    //       wx.checkSession({
    //         success() {
    //           console.log('未过期')
    //           return false;
    //           //session_key 未过期，并且在本生命周期一直有效
    //         },
    //         fail() {
    //           console.log('过期')
    //           // session_key 已经失效，需要重新执行登录流程
    //           wx.login({
    //             success: function (loginRes) {
    //               console.log('wx.login----->', loginRes)
    //               wx.getUserInfo({
    //                 success: function (data) {
    //                   app.httpClient.post('/v1/mini/login', {
    //                     code: loginRes.code
    //                   }).then(res => {
    //                     console.log(res.data.code)
    //                     if (res.data.code == 200) {

    //                     }
    //                   })
    //                 }
    //               })
    //             }
    //           })
    //         }
    //       })
    //     }
    //   }
    // });
  }
})