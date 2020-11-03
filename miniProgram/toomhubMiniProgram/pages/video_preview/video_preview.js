// pages/components/video_preview/video_preview.js
Page({

  /**
   * 页面的初始数据
   */
  data: {
    videoSrc: '',
    videoInitTime: '',
    videoContext: {},
    videoPoster: ''
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {
    console.log(options)
    this.data.videoContext = wx.createVideoContext("video")
    let videoContext = this.data.videoContext
    this.setData({
      videoSrc: options.src,
      videoInitTime: options.time,
      videoPoster: options.poster.replace("*", "?")
    });
    videoContext.seek(options.time);
    console.log(this.data.videoInitTime)
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
  
  //模拟器上不会触发
  videoSeekHandel() {
    let videoContext = this.data.videoContext;
    videoContext.play();
  }
})