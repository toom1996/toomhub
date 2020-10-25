const app = getApp()
Page({
  data: {
    items: [],
    currentItem: 0,
    vTmpPath:''
  },
  onLoad: function () {
    global.index = {};
  },
  onMediaChanged(e) {

    // 选择视频之后插件提供的回调

    const videoTrack = e.detail.track;

    // 用于数据传递
    global.index.videoTrack = videoTrack;
  },
  test() {
    wx.chooseMedia({
      count: 9,
      mediaType: ['image', 'video'],
      sourceType: ['album', 'camera'],
      maxDuration: 30,
      camera: 'back',
      success: (res => {
        console.log(res)
        console.log(res.tempFiles[0].tempFilePath)//thumbTempFilePath
        
        this.setData({
          vTmpPath: res.tempFiles[0].tempFilePath
        })
        //初始化一个空轨道：
        let mediaTrack = new app.globalData.trackTypes.Track({
          type: "media" // 可以是支持的任何类型
        })

        let exampleClip1 = new app.globalData.trackTypes.Clip({
          trackId: mediaTrack.id,
          type: 'video',
          info: {
            tempFilePath: res.tempFiles[0].tempFilePath,
            tempThumbPath: res.tempFiles[0].thumbTempFilePath,
            width: '1280',
            height: '720',
            duration: res.tempFiles[0].duration
          },
          section: new app.globalData.trackTypes.ClipSection({
            start: 0,
            end: res.tempFiles[0].duration,
            duration: res.tempFiles[0].duration
          }),
          startAt: 0
        })

        mediaTrack.clips = [exampleClip1]
        mediaTrack.duration = res.tempFiles[0].duration
        mediaTrack.filterArea = {
          height: 1334,
          width: 750,
        }
        console.log(mediaTrack)

        const videoTrack = mediaTrack;

        // 用于数据传递
        global.index.videoTrack = videoTrack;
      })
    })
  },
  eidtorHandle(){
    console.log(111)
    wx.openVideoEditor({
      filePath: this.data.vTmpPath,
      success:function(res){
        console.log(res)
      }
    })
  }
});