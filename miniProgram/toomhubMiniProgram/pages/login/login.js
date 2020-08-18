const app = getApp()

Page({
  data: {
    //判断小程序的API，回调，参数，组件等是否在当前版本可用。
    canIUse: wx.canIUse('button.open-type.getUserInfo'),
    isHide: false
  },

  // onShow: function (options) {
  //   let pages = getCurrentPages();
  //   var that = this;
  //   // 查看是否授权
  //   wx.getSetting({
  //     success: function (res) {
  //       console.log(res)
  //       //已经登陆
  //       if (res.authSetting['scope.userInfo']) {
  //         //登陆了取消授权页面
  //         wx.navigateBack({
  //           delta: 1
  //         })
  //         wx.getUserInfo({
  //           success: function (res) {
  //             console.log(res)
  //             // 根据自己的需求有其他操作再补充
  //             // 我这里实现的是在用户授权成功后，调用微信的 wx.login 接口，从而获取code
  //             wx.login({
  //               success: res => {
  //                 console.log(res)
  //                 // 获取到用户的 code 之后：res.code
  //                 console.log("用户的code:" + res.code);
  //                 // 可以传给后台，再经过解析获取用户的 openid
  //                 // 或者可以直接使用微信的提供的接口直接获取 openid ，方法如下：
  //                 // wx.request({
  //                 //  // 自行补上自己的 APPID 和 SECRET
  //                 //  url: 'https://api.weixin.qq.com/sns/jscode2session?appid=自己的APPID&secret=自己的SECRET&js_code=' + res.code + '&grant_type=authorization_code',
  //                 //  success: res => {
  //                 //   // 获取到用户的 openid
  //                 //   console.log("用户的openid:" + res.data.openid);
  //                 //  }
  //                 // });
  //               }
  //             });
  //           }
  //         });
  //       }
  //     }
  //   });
  // },

  bindGetUserInfo: function (e) {
    if (e.detail.userInfo) {
      //用户按了允许授权按钮
      var that = this;
      // 获取到用户的信息
      app.globalData.userInfo = e.detail.userInfo
      //获取已经授权的权限
      wx.getUserInfo({
        success: function (userinfo) {
          wx.checkSession({
            success() {
              console.log('未过期')
              //session_key 未过期，并且在本生命周期一直有效
            },
            fail() {
              wx.login({
                success: function (loginRes) {
                  console.log('wx.login----->', loginRes)
                  app.httpClient.post('/v1/mini/login', {
                    code: loginRes.code,
                    userInfo: JSON.stringify(userinfo.userInfo)
                  }).then(res => {
                    console.log(res.data.code)
                    if (res.data.code == 200) {
                      wx.setStorage({
                        key: 'userinfo',
                        data: userinfo,
                      })
                    }
                  })
                }
              })
            }
          }) 
        }
      })
      wx.navigateBack({
        delta: 1
      })
    } else {
      //用户按了拒绝按钮
      wx.showModal({
        title: '警告',
        content: '您点击了拒绝授权，将无法进入小程序，请授权之后再进入!!!',
        showCancel: false,
        confirmText: '返回授权',
        success: function (res) {
          // 用户没有授权成功，不需要改变 isHide 的值
          if (res.confirm) {
            console.log('用户点击了“返回授权”');
          }
        }
      });
    }
  }
})