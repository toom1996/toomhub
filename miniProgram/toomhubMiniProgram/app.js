
//app.js
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
    } else {

    }


    var _this = this;
    let menuButtonObject = wx.getMenuButtonBoundingClientRect();
    wx.getSystemInfo({
      success: res => {
        //导航高度
        let statusBarHeight = res.statusBarHeight,
          navTop = menuButtonObject.top,
          navHeight = statusBarHeight + menuButtonObject.height + (menuButtonObject.top - statusBarHeight) * 2;
        this.globalData.navHeight = navHeight;
        this.globalData.navTop = navTop;        //navTop
        this.globalData.windowHeight = res.windowHeight;
        this.globalData.windowWidth = res.windowWidth;
      },
      fail(err) {
      }
    })
  },
  globalData: {
    userInfo: null, //全局用户信息,
    forceRefresh: false,
    imageThumbnailParam: '?imageMogr2/auto-orient/format/webp'
  },
  httpClient: {
    request: function (method, url, data) {
      //返回一个promise实例
      let _this = this
      return new Promise((resolve, reject) => {
        wx.request({
          url: toomhubApi.requestHost + url,
          data: data,
          method: method,
          header: {
            'Content-Type': 'application/x-www-form-urlencoded', // 默认值
            'Toomhub-Token': wx.getStorageSync('userInfo').token != undefined ? wx.getStorageSync('userInfo').token : ''
          },
          success: (res) => {
            if (res.data.code == 401) {
              wx.navigateTo({
                url: '/pages/login/login'
              })
            }
            resolve(res)
          },
          fail(res) {
            reject(res)
          },
          complete(res) {
          }
        })
      })
    },
    get: function (url, data) {
      return this.request('GET', url);
    },
    post: function (url, data) {
      return this.request('POST', url, data);
    }
  },
  //判断用户是否登陆, 未登陆则跳转到登陆界面
  isLogin: function () {
    if (this.globalData.userInfo == null) {
      wx.navigateTo({
        url: '/pages/login/login'
      })
    }
  },

  //跳转到登陆页面
  redirectToLogin: function () {
    wx.navigateTo({
      url: '/pages/login/login'
    })
  },

  //跳转到首页
  redirectToIndex: function () {
    wx.switchTab({
      url: '/pages/index/index'
    })
  },

  //设置缓存
  setCache: function (key, value) {
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
    this.httpClient.post(toomhubApi.tokenCheck, {
      token: p.token,
      r: p.refreshToken
    }).then(res => {
      if (res.data.code != 200) {
        this.redirectToLogin();
      }
    })
  },

  //获取api
  getApi: function (name) {
    return toomhubApi[name]
  },

  //获取字符串长度
  strlen: function (str) {
    var len = 0;
    for (var i = 0; i < str.length; i++) {
      var c = str.charCodeAt(i);
      //单字节加1
      if ((c >= 0x0001 && c <= 0x007e) || (0xff60 <= c && c <= 0xff9f)) {
        len++;
      }
      else {
        len += 2;
      }
    }
    return len;
  }
})