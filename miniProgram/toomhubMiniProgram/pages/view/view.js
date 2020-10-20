// pages/view/view.js
Page({

  /**
   * 页面的初始数据
   */
  data: {
    image:[
    ]
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {
    wx.setNavigationBarTitle({
      title: '详情',
      }) 
    console.log(options)
    this.setData({
      image: options.list.split(",")
    })
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