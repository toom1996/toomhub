// pages/view/view.js
const app = getApp()

import Toast from '@vant/weapp/toast/toast';
Page({

  /**
   * 页面的初始数据
   */
  data: {
    avatar_url: '',
    created_by: '',
    created_at: '',
    content: '',
    argument_count: 0,
    collect_count: 0,
    like_count: 0,
    image:[
    ]
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {
    console.log(options)
    let id = options.id
    Toast.loading({
      message: '加载中...',
      forbidClick: true,
      loadingType: 'spinner',
    });
    wx.setNavigationBarTitle({
      title: '详情',
    })
    app.httpClient.get(app.getApi('getView') + '?id=' + id).then(res => {
      console.log(res.data.code)
      if (res.data.code == 200) {
        Toast.clear()
      }

      let response = res.data.data;
      console.log(response)
      this.setData({
        avatar_url: response.avatar_url,
        created_by: response.created_by,
        created_at: response.created_at,
        content: response.content,
        tag: response.tag,
        image: response.image,
        argument_count: response.argument_count,
        collect_count: response.collect_count,
        like_count: response.like_count,
      })
    })

    // this.setData({
    //   image: options.list.split(",")
    // })
  },

  /**
   * 生命周期函数--监听页面初次渲染完成
   */
  onReady: function () {

  },

  /**
   * 生命周期函数--监听页面显示
   */
  onShow: function () {

  },

  /**
   * 生命周期函数--监听页面隐藏
   */
  onHide: function () {

  },

  /**
   * 生命周期函数--监听页面卸载
   */
  onUnload: function () {

  },

  /**
   * 页面相关事件处理函数--监听用户下拉动作
   */
  onPullDownRefresh: function () {

  },

  /**
   * 页面上拉触底事件的处理函数
   */
  onReachBottom: function () {

  },

  /**
   * 用户点击右上角分享
   */
  onShareAppMessage: function (options) {
    let title = options.target.dataset.title;
    let list = options.target.dataset.list;

    if (app.strlen(title) > 14 ) {
      title = title.substring(0,14) + '...';
    }
    return {
      title: title,
      path: '/pages/view/view',
      imageUrl: list[0] + app.globalData.imageThumbnailParam
    }
  },
})