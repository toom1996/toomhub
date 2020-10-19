const app = getApp()
let myStyle = `
--tooom__tag-top:
`

Page({
  data: {
    flag: false,
    // 自定义顶部导航
    navHeight: app.globalData.navHeight,
    navTop: app.globalData.navTop,
    show: false,
    index: 0,
    imageList: [],
    image: [],
    tmpImage: [],
    loadedImageList:[]
  },

  onLoad(options) {
    let swiperIndex = parseInt(options.index);
    let swiperIndexLast = this.getLastIndex(swiperIndex);
    let swiperIndexNext = this.getNextIndex(swiperIndex);

    this.setData({
      index: options.index
    })

    //插入空数组
    let image = this.data.tmpImage;
    let loaded = this.data.loadedImageList;
    this.setData({
      imageList:options.list.split(",")
    })
    console.log(this.data.imageList)
    for (let i = 0; i < this.data.imageList.length; i ++) {
      image.push('');
      loaded.push('');
    }
    //修改当前索引及上下一个图片
    image[swiperIndex] = this.data.imageList[swiperIndex];
    loaded[swiperIndex] = 1;
    if (swiperIndexLast >= 0) {
      image[swiperIndex - 1] = this.data.imageList[swiperIndex - 1];
      loaded[swiperIndex - 1] = 1;
    }
    if (swiperIndexNext <= this.data.imageList.length - 1) {
      image[swiperIndex + 1] = this.data.imageList[swiperIndex + 1];
      loaded[swiperIndex + 1] = 1;
    }
    this.setData({
      image: image,
      tmpImage: image,
      loadedImageList: loaded
    })
  },


  redirectToIndex() {
    app.redirectToIndex()
  },

  /**
   * 获取顶部固定高度
   */
  attached: function () {
    this.setData({
      navHeight: app.globalData.navHeight,
      navTop: app.globalData.navTop,
    })
  },
  imageLoadedHandle(event) {
    let loadedIndex = event.currentTarget.dataset.index;
    let tmp = this.data.loadedImageList;
    tmp[loadedIndex] = 1;
    this.setData({
      loadedImageList: tmp
    });
    console.log(this.data.loadedImageList)
  },
  //长按事件, 抑制退出事件触发
  imageLongTapHandle(){
    console.log(121211)
    this.setData({
      flag: true,
    })
    var _this = this
    setTimeout(function () {
      _this.setData({
        flag: false,
      })
    }, 300) //延迟时间 这里是1秒
  },

  onClose() {
    this.setData({
      show: false,
    })
  },

  onSelect(event) {
    console.log(event.detail);
  },

  imageClickHandle() {
    if (this.data.flag == false) {
      wx.navigateBack({
        delta: -1
      })
    }else{
      this.setData({
        flag: false,
      })
    }
  },
  getLastIndex(index) {
    return index - 1;
  },
  getNextIndex(index) {
    return index + 1;
  },
  swiperChangeHandle(event) {
    let swiperIndex = event.detail.current
    let swiperIndexLast = this.getLastIndex(swiperIndex);
    let swiperIndexNext = this.getNextIndex(swiperIndex);
    let image = this.data.tmpImage;
    let loaded = this.data.loadedImageList;

    if (swiperIndexLast >= 0) {
      image[swiperIndex - 1] = this.data.imageList[swiperIndex - 1];
      loaded[swiperIndex - 1] = 1;
    }
    if (swiperIndexNext <= this.data.imageList.length - 1) {
      image[swiperIndex + 1] = this.data.imageList[swiperIndex + 1];
      loaded[swiperIndex + 1] = 1;
    }
    this.setData({
      image: image,
      loadedImageList: loaded
    })
  }
})