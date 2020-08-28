const app = getApp()
// pages/square_add/square_add.js
import Toast from '../../miniprogram_npm/@vant/weapp/toast/toast';
Page({

  /**
   * 页面的初始数据
   */
  data: {
    imageList: [
    ],
    autosize:{ maxHeight: 350, minHeight: 150 },  //文字内容输入框高度
    addSize:{ maxHeight: 30, minHeight: 20 }, //添加标签输入框高度
    tagArr : [
    ],//添加进来的标签数组
    hotTagArr: [
      '#热门1#',
      '#热门2#',
      '#热门3#'
    ],//热门标签数组
    hotTagArr2: [
      '#热门1#',
      '#热门2#',
      '#热门3#'
    ],//COPY的热门标签数组
    tagAddButtonShow: true, //是否显示添加标签的按钮
    tagInputContent:'', //添加的标签内容
    tagShow: false //是否显示添加标签dialog
  },

  //显示添加标签的dialog
  onTagAddDialogShow() {
    this.setData({
       tagShow: true,
      });
  },

  //剔除添加的标签
  tagClose(event) {

    let v = event.target.dataset.value

    let hotIndex = this.data.tagArr.indexOf(v)

    //热门标签剔除后返回列表
    if(this.data.hotTagArr2.indexOf(v) !== -1) {
      let hotTagTmp = []
      hotTagTmp.push(v)
      this.setData({
        hotTagArr:  this.data.hotTagArr.concat(hotTagTmp),
      }, ()=> {
        console.log(this.data.tagArr)
      })
    }

    //剔除标签
    let tagTmp = this.data.tagArr
    tagTmp.splice(hotIndex, 1)

    this.setData({
      tagArr: tagTmp
    })

    if (this.getTagLength() <= 5) {
      this.setData({
        tagAddButtonShow: true
      })
      return false;
    }

  },

  //上传图片触发事件
  afterRead(event) {
    console.log(event)
    let list = event.detail.file;
    let _this = this;

    var promise = Promise.all(list.map((item, index) => {

      return new Promise(function (resolve, reject) {
        //先插入一个空图片
        let tmp = _this.data.imageList
        let index = tmp.push({ 
          deletable: false ,
          status: 'uploading',
          message: '上传中',
        });
        console.log(index)
        _this.setData({ imageList: tmp });

        let uploadTask = wx.uploadFile({
            url: 'http://127.0.0.1:8080/c/image/upload',
            filePath: item.path,
            name: 'file',
            success: (res)=> {
              //上传成功后更换空图的内容
              let data = JSON.parse(res.data)
              console.log(data.data)
              let tmp = _this.data.imageList
              tmp[index - 1].url = data.data;
              tmp[index - 1].deletable = true;
              tmp[index - 1].status = 'done'
              tmp[index - 1].message = '0%'
              _this.setData({ imageList: tmp });
            },
            fail: (res) => {
              //上传失败的处理
              console.log(data.data)
              let tmp = _this.data.imageList
              tmp[index - 1].deletable = true;
              tmp[index - 1].status = 'fail'
              tmp[index - 1].message = ''
              _this.setData({ imageList: tmp });
            }
          });

          uploadTask.onProgressUpdate((res) => {
            tmp[index - 1].message = res.progress +  '%'
          })
        })
    }))



    // // 当设置 mutiple 为 true 时, file 为数组格式，否则为对象格式
    // wx.uploadFile({
    //   url: 'http://127.0.0.1:8080/c/image/upload', // 仅为示例，非真实的接口地址
    //   filePath: file[0].path,
    //   name: 'file',
    //   success: (res)=> {
    //     let data = JSON.parse(res.data)
    //     // 上传完成需要更新 fileList
    //     let tmp = this.data.imageList
    //     tmp.push({ url: data.data });
    //     this.setData({ imageList: tmp });
    //     console.log(this.data.imageList)
    //   },
    //   fail(res) {
    //     console.log(res)
    //   }
    // });
    // console.log(6666)
  },

  //热门标签被点击
  hotTagClick: function (event) {
    if (this.getTagLength() >= 5) {
      Toast('标签不能超过5个哦~');
      return false;
    }
    let v = event.target.dataset.value
    let hotIndex = this.data.hotTagArr.indexOf(v)

    console.log(hotIndex)
    let hotTmp = []

    hotTmp = this.data.hotTagArr;
    hotTmp.splice(hotIndex, 1)
    console.log(hotTmp)
    let tmp = [];
    if (this.data.tagArr.indexOf(v)  === -1) {
      tmp.push(v);
      this.setData({
        tagArr: this.data.tagArr.concat(tmp),
        hotTagArr: hotTmp
      }, ()=> {
        console.log(this.data.tagArr)
      })
    }
  },

  //添加标签
  addTag: function (event) {
    if (this.getTagLength() >= 5) {
      Toast('标签不能超过5个哦~');
      return false;
    }

    let tmp = [];
    //剔除重复标签
    if (this.data.tagArr.indexOf('#' + this.data.tagInputContent + '#')  === -1) {
      tmp.push('#' + this.data.tagInputContent + '#');
      this.setData({
        tagArr: this.data.tagArr.concat(tmp)
      }, ()=> {
        console.log(this.data.tagArr)
      })
    }
    //设置标签输入框内容为空
    this.setData({
      tagInputContent: ''
    })

    //大于5个标签隐藏添加标签按钮
    if (this.getTagLength() >= 5) {
      this.setData({
        tagAddButtonShow: false
      })
    }
  },

  //获取添加标签的数量
  getTagLength() {
    return this.data.tagArr.length
  },

  //添加标签输入框输入事件
  tagInput: function (data) {
    console.log(data.value)
    this.setData({
      tagInputContent: data.detail
    })
  },


  imageDelete: function (event) {
    console.log(event)
    let tmp = this.data.imageList;
    tmp.splice(event.detail.index, 1)
    console.log(event)
    this.setData({
      imageList: tmp
    })
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {
    
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
    // app.isLogin()
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

  }
})