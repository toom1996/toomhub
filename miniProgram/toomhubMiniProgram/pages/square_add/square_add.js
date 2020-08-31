const app = getApp()
// pages/square_add/square_add.js
import Toast from '../../miniprogram_npm/@vant/weapp/toast/toast';
Page({

  /**
   * 页面的初始数据
   */
  data: {
    imageList: [],
    autosize: {
      maxHeight: 350,
      minHeight: 150
    }, //文字内容输入框高度
    addSize: {
      maxHeight: 30,
      minHeight: 20
    }, //添加标签输入框高度
    tag: '添加标签', //添加进来的标签数组
    tagShow: false, //是否显示添加标签dialog
    mainActiveIndex: 0,
    activeId: null,
  },

  //显示添加标签的dialog
  onTagAddDialogShow() {
    console.log('onTagAddDialogShow')
    this.setData({
      tagShow: true,
    });
  },

  //显示添加标签的dialog
  onTagAddDialogClose() {
    console.log('onTagAddDialogClose');
    this.setData({
      tagShow: false,
    });
  },

  tagClick(event) {
    this.setData({
      tag: event.currentTarget.dataset.value
    });
    this.setData({
      tagShow: false
    });
  },


  //上传图片触发事件
  afterRead(event) {
    console.log(event.detail.file)
    let list = event.detail.file;
    let _this = this;

    var promise = Promise.all(list.map((item, index) => {
      return new Promise(function(resolve, reject) {
        //先插入一个空图片
        let tmp = _this.data.imageList
        // let baseData = ''
        let index = tmp.push({
          deletable: false,
          status: 'uploading',
          message: '上传中',
        });
        _this.setData({
          imageList: tmp
        });

        //转成base64, 这里本来想做base64展示, 省点cdn流量, 但是组件好像不支持:-(
        // wx.getFileSystemManager().readFile({
        //   filePath: item.path, //选择图片返回的相对路径
        //   encoding: 'base64', //编码格式
        //   success: (resBaseData) => {
        //     wx.getImageInfo({
        //       src: item.path,
        //       success: function (resImageInfo) {
        //         baseData = 'data:image/' + resImageInfo.type + ';base64,' + resBaseData.data;
        //       }
        //     }) 
        //   }
        // })



        let uploadTask = wx.uploadFile({
          url: 'http://127.0.0.1:8080/c/image/upload',
          filePath: item.path,
          name: 'file',
          success: (res) => {
            console.log(res)

            //上传成功后更换空图的内容
            let data = JSON.parse(res.data)
            console.log(data.data)
            let tmp = _this.data.imageList
            tmp[index - 1].url = data.data;
            tmp[index - 1].deletable = true;
            tmp[index - 1].status = 'done'
            tmp[index - 1].message = '0%'
            _this.setData({
              imageList: tmp
            });
          },
          fail: (res) => {
            //上传失败的处理
            console.log(data.data)
            let tmp = _this.data.imageList
            tmp[index - 1].deletable = true;
            tmp[index - 1].status = 'fail'
            tmp[index - 1].message = ''
            _this.setData({
              imageList: tmp
            });
          }
        });

        uploadTask.onProgressUpdate((res) => {
          tmp[index - 1].message = res.progress + '%'
        })
      })
    }))
  },

  imageDelete: function(event) {
    console.log(event)
    let tmp = this.data.imageList;
    tmp.splice(event.detail.index, 1)
    console.log(event)
    this.setData({
      imageList: tmp
    })
  },

  send() {
    console.log(111111111)
    Toast.loading({
      message: '发布中...',
      forbidClick: true,
      duration: 0
    });

    app.httpClient.post('/v1/mini/sq/create').then(res=>{
      Toast.clear();
      console.log(11111111)
    })
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function(options) {

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
    // app.isLogin()
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

  }
})