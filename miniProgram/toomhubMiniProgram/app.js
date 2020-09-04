
//app.js
const REQUEST_HOST = 'http://127.0.0.1:8080'

App({
  onLaunch: function () {
    // 展示本地存储能力
    var logs = wx.getStorageSync('logs') || []
    logs.unshift(Date.now())
    wx.setStorageSync('logs', logs)
    let userInfo = wx.getStorageSync('userInfo');
    if (userInfo !== '') {
      this.globalData.userInfo = userInfo
    }

    wx.setStorage({
      data: false,
      key: 'is_refresh',
    })
  },
  globalData: {
    userInfo: null, //全局用户信息,
    request_host: 'http://127.0.0.1:8080'
  },
  httpClient: {
    request: function (method, url, data) {

      //返回一个promise实例
      let _this = this
      return new Promise((resolve, reject) => {
        wx.request({
          url: getApp().globalData.request_host + url,
          data: data,
          method: method,
          header: {
            'Content-Type': 'application/x-www-form-urlencoded', // 默认值
            'Toomhub-Token': wx.getStorageSync('userInfo').token
          },
          success(res) {
            console.log(res)
            //token 过期
            if (res.data.code == 401) {

              console.log(wx.getStorageSync('is_refresh'))
              if (wx.getStorageSync('is_refresh') == false) {
                //TODO 阻塞所有请求
                //设置正在刷新token变量为真
                wx.setStorage({
                  data: true,
                  key: 'is_refresh',
                })
                wx.showLoading({
                  title: '加载中...',
                })

                _this.tokenRefresh();
              }
            }
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

    //刷新token
    tokenRefresh: function () {
      console.log('tokenRefresh')
      let token = wx.getStorageSync('userInfo').token
      let refreshToken = wx.getStorageSync('userInfo').refreshToken

      //没有token直接跳转到登陆页面
      if (!token || !refreshToken) {
        this.redirectToLogin();
      }

      wx.request({
        url: getApp().globalData.request_host + '/v1/mini/user/refresh',
        data: {token: token, refreshToken:refreshToken},
        method: 'POST',
        success: function (res) {
          console.log(res)
        }
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
    if (wx.getStorageSync('userInfo') == '') {
      wx.navigateTo({
        url: '/pages/login/login'
      })
    }
    wx.checkSession({
      success() {
        console.log('success')
        return true;
        //session_key 未过期，并且在本生命周期一直有效
      },
      fail() {
        console.log('fail')
        wx.navigateTo({
          url: '/pages/login/login'
        })
        return false;
      }
    }) 
  },

  //跳转到登陆页面
  redirectToLogin: function () {
    wx.navigateTo({
      url: '/pages/login/login'
    })
  },

  //设置缓存
  setCache: function(key, value) {
    wx.setStorage({
      key: key,
      data: value,
    })
  }
})