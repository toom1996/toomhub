
//app.js
const REQUEST_HOST = 'http://127.0.0.1:8080'

var toomhubApi = require("./api.js");
var _this = this;
App({
  onLaunch: function () {
    toomhubApi = toomhubApi.api;
    
    //初始化用户信息
    let userInfo = wx.getStorageSync('userInfo');
    if (userInfo !== '') {
      //赋值
      this.globalData.userInfo = userInfo
      //设置定时器
      this.setTokenCheckTimer();
    }else{

    }
  },
  globalData: {
    userInfo: null, //全局用户信息,
  },
  httpClient: {
    request: function (method, url, data) {
      //返回一个promise实例
      let _this = this
      console.log()
      return new Promise((resolve, reject) => {
        wx.request({
          url: toomhubApi.REQUEST_HOST + url,
          data: data,
          method: method,
          header: {
            'Content-Type': 'application/x-www-form-urlencoded', // 默认值
            'Toomhub-Token': wx.getStorageSync('userInfo').token
          },
          success(res) {
            console.log(res)
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
    console.log(this.globalData.userInfo)
    if (this.globalData.userInfo == null) {
      wx.navigateTo({
        url: '/pages/login/login'
      })
    }
    // wx.checkSession({
    //   success() {
    //     console.log('success')
    //     return true;
    //     //session_key 未过期，并且在本生命周期一直有效
    //   },
    //   fail() {
    //     console.log('fail')
    //     wx.navigateTo({
    //       url: '/pages/login/login'
    //     })
    //     return false;
    //   }
    // }) 
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
  },
  //设置验证定时器
  setTokenCheckTimer: function () {
    //每小时获取一次新的token
    let _this = this;
    _this.tokenRefresh()
    var i = setInterval(() => {
      _this.tokenRefresh()
    }, 1000 * 60 * 60)
  },

  //验证token接口
  tokenRefresh: function () {
    let p = this.globalData.userInfo;
    console.log(p)
    this.httpClient.post(toomhubApi.TOKEN_CHECK, {
      token: p.token,
      r: p.refreshToken
    }).then(res => {
        if (res.data.code != 200) {
          this.redirectToLogin();
        }
    })
    console.log('refreshToken')
  },

  //获取api
  getApi: function(name) {
    return toomhubApi[name]
  }
})