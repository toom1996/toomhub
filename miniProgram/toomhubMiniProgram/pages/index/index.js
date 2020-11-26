//index.js
//获取应用实例
import { getThumbnail, calculateVideoSize } from '../../api/func'
const app = getApp()
import Toast from '../../miniprogram_npm/@vant/weapp/toast/toast';
let myStyle = `
--tooom__tag-top:
`

Page({
  /**
   * 用户点击右上角分享
   */
  onShareAppMessage: function (options) {
    return new Promise((resolve, reject) => {
      getThumbnail(options, data => {
        console.log('图片', data.title, data.id);
        resolve({
          title: data.title,
          path: '/pages/view/view?id=' + data.id,
          imageUrl: data.src
        })
      })
    })
  },
  data: {
    // 自定义顶部导航
    navHeight: app.globalData.navHeight,
    navTop: app.globalData.navTop,
    screenWidth: app.globalData.windowWidth,
    screenHeight: app.globalData.windowHeight,
    viewData: {
      style: myStyle //顶部搜索栏样式
    },
    skeletonShow: true, //是否展示骨架图
    data: [], //页面数据
    likeHandle: true, //是否加载点赞处理器, 防止连续点击出现问题
    page: 1, //上拉页码
    loadingText: "正在加载更多....", //上拉加载文字
    showPubButtom: false, //是否显示发布按钮
    sheetShow: false, //是否显示发布sheet
    actions: [ //发布sheet选项
      { name: '发图片', openType: 'imageAddHandle' },
      { name: '发视频', openType: 'videoAddHandle' },
    ],
    videoTime: {}, //video当前播放时间
    videoShareTmpImage: '',
    list: [
      { "name": "发现" },
      { "name": "关注" },
    ]
  },
  navigationSwitch: function (event) {
    wx.navigateTo({
      url: '../user/user'
    })
  },
  //事件处理函数
  bindViewTap: function () {
    wx.navigateTo({
      url: '../logs/logs'
    })
  },
  viewHandle(event) {
    wx.navigateTo({
      url: '../view/view?id=' + event.currentTarget.dataset.id
    })
  },

  //空事件
  emptyHandle() {
  },

  //视频点击事件
  videoContainerClickHandle(e) {
    console.log(e)
    let videoId = e.currentTarget.id;
    let videoTime = this.data.videoTime[videoId] == undefined ? 0 : this.data.videoTime[videoId];
    let videoSrc = e.currentTarget.dataset.src;
    let videoPoster = e.currentTarget.dataset.poster.indexOf('?') !== -1 ? e.currentTarget.dataset.poster.replace('?', '*') : e.currentTarget.dataset.src + '?vframe/jpg/offset/10';
    console.log('videoPoster', videoPoster)
    console.log('../video_preview/video_preview?time=' + videoTime + '&src=' + videoSrc + '&poster=' + videoPoster)
    wx.navigateTo({
      url: '../video_preview/video_preview?time=' + videoTime + '&src=' + videoSrc + '&poster=' + videoPoster
    })
  },

  //发布按钮sheet取消事件
  sheetOnCancleHandle() {
    this.setData({
      sheetShow: false
    })
  },

  //发布按钮sheet关闭事件
  sheetOnCloseHandle() {
    this.setData({
      sheetShow: false
    })
  },

  //video进度条change事件
  videoTimeUpdateHandle(e) {
    let videoId = e.currentTarget.id;
    let videoTime = this.data.videoTime;
    videoTime[videoId] = e.detail.currentTime;
    this.setData({
      videoTime: videoTime
    })
  },
  onLoad: function () {
    console.log(this.data.screenHeight)
    this._observer = wx.createIntersectionObserver(this, { observeAll: true })
    this.setData({ 'viewData.style': myStyle + '40px;' })
    this.refreshIndex(1, true);
  },
  getUserInfo: function (e) {
    app.globalData.userInfo = e.detail.userInfo
    this.setData({
      userInfo: e.detail.userInfo,
      hasUserInfo: true
    })
  },
  SheetSelectHandle(event) {
    console.log(event)
    let openType = event.detail.openType;
    switch (openType) {
      case "imageAddHandle":
        this.imageAddHandle();
        break;
      case "videoAddHandle":
        this.videoAddHandle();
        break;
    }
  },
  //发图片点击事件
  imageAddHandle() {
    wx.navigateTo({
      url: '../image_add/image_add'
    })
  },
  //发布按钮点击事件
  addHandle: function () {
    this.setData({
      sheetShow: true
    })
  },

  //点击发布视频处理事件
  videoAddHandle() {
    wx.navigateTo({
      url: '../video_add/video_add'
    })
  },

  // 图片点击事件
  previewImage: function (event) {
    let data = event.currentTarget.dataset;

    wx.navigateTo({
      url: '/pages/image_preview/image_preview?list=' + data.list + '&index=' + data.index + '&param=' + data.param
    })
  },

  //滑动到底部刷新事件
  onReachBottom: function () {
    this.refreshIndex(this.data.page);
  },

  onRestore(e) {
  },
  //页面数据刷新函数
  refreshIndex: function (page, refresh) {
    this.setData({
      skeletonShow: false
    })
    if (!page) {
      page = 1
    }
    //下拉刷新
    if (refresh === true) {
      this.setData({
        skeletonShow: true,
        data: [],
        page: 0,
        loadingText: '正在加载中...'
      })
    }
    //请求首页接口
    app.httpClient.get(app.getApi('squareIndex') + '?page=' + page).then(res => {
      if (res.data.code == 200) {
        if (res.data.data.count == 0) {
          this.setData({
            loadingText: '已经到底啦~~'
          })
        } else {
          var responseData = res.data.data
          let d = this.data.data

          if (responseData.list.length > 0) {
            responseData.list.forEach(item => {
              //视频文件控制大小
              if (item.type == 1) {
                let size = calculateVideoSize(item.width, item.height );
                item.width = size.width;
                item.height = size.height;
              }
              d.push(item);
            })
          }
          let newPage = page + 1
          this.setData({
            data: d,
            page: newPage,
            skeletonShow: false,
          })
          this.bindObserver()
        }
      }
    })
  },

  //计算视频宽高
  calculateVideoSize(width, height) {
    console.log('宽',width)
    console.log('高',height)
    return {
      width: width,
      height:height
    }
  },

  //下拉刷新接口
  onPullDownRefresh: function () {
    this.refreshIndex(1, true);
    wx.stopPullDownRefresh();
  },

  onShow: function () {
    
    if (app.globalData.forceRefresh == true) {
      app.globalData.forceRefresh = false;
      this.refreshIndex(1, true);
    }

    this.setData({
      userInfo: app.globalData.userInfo
    })
  },
  /**
   * 获取顶部固定高度
   */
  attached: function () {
    this.setData({
      navHeight: App.globalData.navHeight,
      navTop: App.globalData.navTop,
    })
  },

  //点赞处理函数
  likeHandle: function (e) {
    this.setData({
      likeHandle: false
    })

    let newList = this.data.data;
    let isLike = e.currentTarget.dataset.like;
    if (isLike === 0) {
      isLike = 1;
    } else {
      isLike = 0;
    }

    app.httpClient.post(app.getApi('squareLike'), {
      'id': e.currentTarget.dataset.id,
      'o': isLike,
      'page': this.data.page
    }).then(res => {
      let response = res.data
      if (response.code == 200) {
        newList[e.currentTarget.dataset.index].is_like = isLike
        if (isLike === 1) {
          newList[e.currentTarget.dataset.index].like_count += 1
        } else {
          newList[e.currentTarget.dataset.index].like_count -= 1
        }
        this.setData({
          data: newList
        })
        wx.showToast({
          title: '操作成功',
          icon: 'none',
          duration: 1000,
        })
      } else if (response.code != 401) {
        wx.showToast({
          title: '操作失败',
          icon: 'none',
          duration: 1000,
        })
      }
      this.setData({
        likeHandle: true
      })
    })
  },

  bindObserver() {
    let viewportBottom = app.globalData.windowHeight / 2 * - 1
    let viewportTop = app.globalData.windowHeight / 2.5 * - 1
    console.log(viewportBottom)
    console.log(viewportTop)
    wx.createIntersectionObserver(this, { observeAll: true }).relativeTo('.scroll-view').relativeToViewport({ top: viewportTop, bottom: viewportBottom }).observe('.video', (res) => {
      this.videoContext = wx.createVideoContext(res.id)
      if (res.intersectionRatio > 0) {
        console.log(res.id, '播放')
        this.videoContext.play()//开始播放
      } else {
        this.videoContext.pause()//开始播放
        console.log(res.id, '暂停')
      }
    })
  }
})
