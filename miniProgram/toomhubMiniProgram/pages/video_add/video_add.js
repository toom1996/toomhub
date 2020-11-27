const app = getApp()
const device = wx.getSystemInfoSync();
const device_width = device.windowWidth;
const device_height = device_width / 2;
import { calculateVideoSize } from '../../api/func'
const  Uploader  = require('../../miniprogram_npm/miniprogram-file-uploader/index.js');
Page({

  /**
   * 页面的初始数据
   */
  data: {
    content: '',
    autosize: {
      maxHeight: 350,
      minHeight: 150
    }, //文字内容输入框高度
    addSize: {
      maxHeight: 30,
      minHeight: 20
    }, //添加标签输入框高度
    tag: '', //添加进来的标签数组
    tagList: '', //标签数组
    tagShow: false, //是否显示添加标签dialog
    content_show: false,
    keyword: '',
    videoSrc: '',
    isHiddenvideoContainer: true, //是否隐藏视频播放容器
    isHiddenUploader: false, //是否显示上传组件
    duration: 0, //视频时长
    videoCdnSrc: '', //视频cdn链接
    host: '',
    size: 0,
    videoHeight: 0,
    videoWidth: 0,
    videoContainerHidden: false,
    videoPoster: '', //视频封面
    videoContainerWidth: 0,
    videoContainerHeight: 0
  },


  conetentHandel (value) {
    this.setData({
      content: value.detail
    })
  },

  //显示添加标签的dialog
  onTagAddDialogShow() {
    this.setData({
      videoContainerHidden: true,
      tagShow: true,
    });
    this.search(this.data.tag);
  },
  search: function (s) {
    app.httpClient.get(app.getApi('squareTagSearch') + '?k=' + s).then(res=>{
      this.setData({
        tagList:res.data.data
      });
    })
  },
  onChange: function(e) {
    this.setData({
      tag: e.detail
    });
    this.search(e.detail);
  },
  
  //显示添加标签的dialog
  onTagAddDialogClose() {
    this.setData({
      videoContainerHidden: false,
      tagShow: false,
    });
  },

  tagClick(event) {
    this.setData({
      tag: event.currentTarget.dataset.value
    });
    this.setData({
      tagShow: false,
      videoContainerHidden: false
    });
  },


  //上传视频触发事件
  afterRead(event) {

      console.log('event', event)
    let file = event.detail.file;
    console.log(file)
    let size = calculateVideoSize(file.width, file.height, 40)
    this.setData({
      videoContainerHeight: size.height,
      videoContainerWidth: size.width,
      videoHeight: file.height,
      videoWidth:file.width
    })
    if (file.size > 1024 * 1024 * 50) {
      Toast('文件太大啦~~~');
    }

    
    // wx.openVideoEditor({
    //   filePath: file.url,
    //   complete: editorRes => {
        console.log('Uploader', Uploader)

        
        // let uploadTask = wx.uploadFile({
        //   url: app.getApi('requestHost') + app.getApi('videoUpload'),
        //   filePath: file.url,
        //   name: 'file',
        //   success: (res) => {
        //       console.log('path->',file.tempFilePath);
            if (Uploader.default.isSupport()) {
              const uploader = new Uploader.default({
                tempFilePath: file.url,
                totalSize: file.size,
                uploadUrl: 'http://192.168.10.113:8080/video/upload1/',
                mergeUrl: 'http://www.google.com/',
                maxConcurrency: 2, // 并发上传数，默认 5，最大不超过 10
                maxChunkRetries: 1, // 请求失败时最大重试次数，默认为 0
                verbose: true, // 是否输出开始日志，默认 false
                fileName: '11111',
              })
              console.log('uploader', uploader)
              uploader.upload()

              uploader.on('complete', (res) => {
                console.log('upload complete', res)
              })
              
              // 文件上传成功
              uploader.on('success', (res) => {
                console.log('upload success', res)
              })
              
              // 文件上传失败
              uploader.on('fail', (res) => {
                console.log('fail', res)
              })
              
              // 文件进度变化
              uploader.on('progress', (res) => {
                this.setData({
                  progress: res.progress,
                  uploadedSize: parseInt(res.uploadedSize / 1024),
                  averageSpeed: parseInt(res.averageSpeed / 1000),
                  timeRemaining: res.timeRemaining
                })
              })
            }
            // try {
            //   let data = JSON.parse(res.data)
            // console.log('data->', data)
            //   if (res.statusCode == 200 && data.code == 200) {
            //     this.setData({
            //       videoSrc: file.tempFilePath,
            //       isHiddenvideoContainer: false,
            //       isHiddenUploader: true,
            //       videoCdnSrc: data.data.name,
            //       duration: file.duration,
            //       host: data.data.request_host,
            //       size: data.data.size,
            //     })
            //     console.log('this.data ->',this.data)
            //   }
            // }catch(e) {
            //   wx.showToast({
            //     title: '上传文件发生错误..',
            //     icon: 'fail'
            //   })
            // }
        //   },
        //   fail: (res) => {
        //     console.log(res)
        //   }
        // });

        // uploadTask.onProgressUpdate((res) => {
        //   console.log('上传进度', res.progress + " " + new Date())
        //   console.log('已经上传的数据长度', res.totalBytesSent)
        //   console.log('预期需要上传的数据总长度', res.totalBytesExpectedToSend)

        // })
    //   }
    // })
  },
  send() {
    
    if (this.data.content == '') {
      Toast.fail("文字内容不能为空哦")
      return 
    }

    Toast.loading({
      message: '发布中...',
      forbidClick: true,
      duration: 0
    });
    
    let obj = {...this.data.imageList}
    
    app.httpClient.post(app.getApi('videoCreate'), {
      'duration': this.data.duration,
      'host': this.data.host,
      'name': this.data.videoCdnSrc,
      'cover': app.globalData.videoCover,
      'content': this.data.content,
      'tag': this.data.tag,
      'size': this.data.size,
      'height': this.data.videoHeight,
      'width': this.data.videoWidth
    }).then(res=>{
      let response = res.data
      Toast.clear();
      if (response.code == 200) {
        wx.showToast({
          title: '发布成功',
          icon: 'success',
          duration: 2000,
          success: function () {
            app.globalData.forceRefresh = true;
            setTimeout(function () {
              app.redirectToIndex();
            }, 2000);
          }
        })
      }
    })
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function(options) {
    app.globalData.videoCover = '';
    console.log(options)
    this.setData({
      device_height: device_height,
      device_width: device_width
    })
  },

  /**
   * 生命周期函数--监听页面初次渲染完成
   */
  onReady: function() {

  },

  /**
   * 生命周期函数--监听页面显示
   */
  onShow: function() {
    app.isLogin()
    console.log(app.globalData.videoCover)
    wx.getStorage({
      key: 'userInfo',
      success (res) {
      }
      })
    if (app.globalData.userInfo) {
      this.setData({
        userInfo: app.globalData.userInfo,
      })
    } else {
      this.setData({
        userInfo: null,
      })
    }
  },

  onClose: function () {
    this.setData({
      tag:''
    })
  },

  /**
   * 生命周期函数--监听页面隐藏
   */
  onHide: function() {

  },

  /**
   * 生命周期函数--监听页面卸载
   */
  onUnload: function() {

  },

  /**
   * 页面相关事件处理函数--监听用户下拉动作
   */
  onPullDownRefresh: function() {

  },

  /**
   * 页面上拉触底事件的处理函数
   */
  onReachBottom: function() {

  },

  /**
   * 用户点击右上角分享
   */
  onShareAppMessage: function() {

  },
  maskTouchendCallback: function () {
    this.setData({
      videoSrc: '',
      isHiddenvideoContainer: true,
      isHiddenUploader: false
    })
  },
  redirectToVideoCoverHandle: function () {
    wx.navigateTo({
      url: '../video_cover/video_cover?duration=' + this.data.duration + '&name=' + this.data.videoCdnSrc +'&host=' + this.data.host
    })
  },
  test: function() {
    console.log('test')
  }
})