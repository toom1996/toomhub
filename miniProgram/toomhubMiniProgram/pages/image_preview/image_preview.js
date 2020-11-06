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
    loadedImageList: []
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
      imageList: options.list.split(",")
    })
    for (let i = 0; i < this.data.imageList.length; i++) {
      image.push('');
      loaded.push({});
    }
    //修改当前索引及上下一个图片
    image[swiperIndex] = this.data.imageList[swiperIndex] + '?imageMogr2/auto-orient/format/webp';
    loaded[swiperIndex] =  {
      is_load: 1
    };
    if (swiperIndexLast >= 0) {
      image[swiperIndex - 1] = this.data.imageList[swiperIndex - 1];
      loaded[swiperIndex - 1] =  {
        is_load: 1
      };
    }
    if (swiperIndexNext <= this.data.imageList.length - 1) {
      image[swiperIndex + 1] = this.data.imageList[swiperIndex + 1];
      loaded[swiperIndex + 1] =  {
        is_load: 1
      };
    }
    this.setData({
      image: image,
      tmpImage: image,
      loadedImageList: loaded
    })
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
    let height = event.detail.height;
    let width = event.detail.width;

    let widthP = (width / app.globalData.windowWidth).toFixed(2);
    
    tmp[loadedIndex] = {
      is_load: 1,
      is_overflow: height / widthP > app.globalData.windowHeight ? true : false 
    };
    this.setData({
      loadedImageList: tmp
    });
  },
  //长按事件, 抑制退出事件触发
  imageLongTapHandle() {
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
  },

  imageClickHandle() {
    if (this.data.flag == false) {
      wx.navigateBack({
        delta: -1
      })
    } else {
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
      loaded[swiperIndex - 1]['is_load'] = 1;
    }
    if (swiperIndexNext <= this.data.imageList.length - 1) {
      image[swiperIndex + 1] = this.data.imageList[swiperIndex + 1];
      loaded[swiperIndex + 1]['is_load'] = 1;
    }
    this.setData({
      image: image,
      loadedImageList: loaded
    })
  }
})