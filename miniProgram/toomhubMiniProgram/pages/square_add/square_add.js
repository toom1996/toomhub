const app = getApp()

// pages/square_add/square_add.js
import Toast from '../../miniprogram_npm/@vant/weapp/toast/toast';
Page({

  /**
   * 页面的初始数据
   */
  data: {
    imageList: [],
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
    mainActiveIndex: 0,
    activeId: null,
    content_show: false,
    keyword: ''
  },


  conetentHandel (value) {
    this.setData({
      content: value.detail
    })
  },

  //显示添加标签的dialog
  onTagAddDialogShow() {
    console.log('onTagAddDialogShow')
    this.setData({
      // content_show: true,
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
    console.log('onTagAddDialogClose');
    this.setData({
      // content_show: false,
      tagShow: false,
    });
  },

  tagClick(event) {
    console.log(event.currentTarget.dataset)
    this.setData({
      tag: event.currentTarget.dataset.value
    });
    this.setData({
      tagShow: false
    });
  },


  //上传图片触发事件
  afterRead(event) {
    console.log('所有文件', event.detail.file)
    let list = event.detail.file;
    for (let i = 0; i < list.length; i++) {
      if (list[i].size > 1024 * 1024) {
        list.splice(i, 1)
        Toast('部分文件超出大小限制,自动忽略');
      }
    }
    let _this = this;

    var promise = Promise.all(list.map((item, index) => {
      return new Promise(function(resolve, reject) {
        //先插入一个空图片
        let tmp = _this.data.imageList
        let baseData = ''
        let index = tmp.push({
          deletable: false,
          status: 'uploading',
          message: '上传中',
        });
        _this.setData({
          imageList: tmp
        });

        // // 转成base64, 这里本来想做base64展示, 省点cdn流量, 但是组件好像不支持:-(
        // wx.get.readFile({FileSystemManager()
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

        console.log(baseData)

        let uploadTask = wx.uploadFile({
          url: app.getApi('requestHost') + '/c/image/upload',
          filePath: item.path,
          name: 'file',
          success: (res) => {
            console.log(res)
            try{
              let data = JSON.parse(res.data)
              if (res.statusCode == 200 && data.code == 200) {
                //上传成功后更换空图的内容
                console.log(data.data)
                let tmp = _this.data.imageList
                tmp[index - 1].url = data.data.url;
                tmp[index - 1].deletable = true;
                tmp[index - 1].status = 'done'
                tmp[index - 1].message = '0%'
                tmp[index - 1].ext = data.data.extension
                tmp[index - 1].size = data.data.size
                tmp[index - 1].name = data.data.name
                tmp[index - 1].host = data.data.request_host
                tmp[index - 1].param = data.data.param
                _this.setData({
                  imageList: tmp
                });
              } else {
                _this.removeImage(index - 1)
                Toast.fail(data.message);
              }
            }catch(e){
              _this.removeImage(index - 1)
              Toast.fail('上传失败');
            }
          },
          fail: (res) => {
            _this.removeImage(index - 1)
            Toast.fail(res);
          }
        });

        uploadTask.onProgressUpdate((res) => {
          tmp[index - 1].message = res.progress + '%'
        })
      })
    }))
  },

  imageDelete: function(event) {
    this.removeImage(event.detail.index)
  },

  removeImage: function (index) {
    let tmp = this.data.imageList;
    tmp.splice(index, 1)
    this.setData({
      imageList: tmp
    })
  },

  send() {
    
    if (this.data.content == '') {
      Toast.fail("文字内容不能为空哦")
      return 
    }

    if (this.data.imageList.length == 0) {
      Toast.fail("最少上传一张图片")
      return 
    }

    Toast.loading({
      message: '发布中...',
      forbidClick: true,
      duration: 0
    });
    
    let obj = {...this.data.imageList}
    
    app.httpClient.post(app.getApi('squareCreate'), {
      'content': this.data.content,
      'image_list': JSON.stringify(obj),
      'tag': this.data.tag,
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
    app.isLogin()
    wx.getStorage({
      key: 'userInfo',
      success (res) {
      console.log(res.data)
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

  }
})