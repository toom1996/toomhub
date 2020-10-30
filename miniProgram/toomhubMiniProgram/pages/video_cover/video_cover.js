const app = getApp();
Page({

  /**
   * 页面的初始数据
   */
  data: {
    host: '',
    videoUrl: '',
    duration: 0,
    defaultCover: '',
    checkedCover: 0,
    coverHeight: ( app.globalData.windowWidth - 40 ) / 4 
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {
    console.log(options.host + options.name + '?vframe/jpg/offset/' + this.data.checkedCover)
    this.setData({
      videoUrl: options.host + options.name,
      duration: parseInt(options.duration),
      defaultCover: options.host + options.name + '?vframe/jpg/offset/' + this.data.checkedCover
    })
    console.log(this.data.duration)
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
  onShareAppMessage: function () {

  },
  selectCoverHandle(event) {
    console.log(event)
    let index = event.currentTarget.dataset.index;
    this.setData({
      checkedCover: index,
    })
    this.setData({
      defaultCover: this.data.videoUrl + '?vframe/jpg/offset/' + this.data.checkedCover
    })
  },
  checkedCoverHandel() {
    app.globalData.videoCover = this.data.videoUrl + '?vframe/jpg/offset/' + this.data.checkedCover;
    wx.navigateBack({
      delta: -1
    })
  }
})