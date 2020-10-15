//index.js
//获取应用实例
const app = getApp()
import Toast from '../../miniprogram_npm/@vant/weapp/toast/toast';
let myStyle = `
--tooom__tag-top:
`

Page({
  data: {
    // 自定义顶部导航
    navHeight: app.globalData.navHeight,
    navTop: app.globalData.navTop,
    viewData: {
      style: myStyle //顶部搜索栏样式
    },
    skeletonShow: true, //是否展示骨架图
    data:[], //页面数据
    likeHandle: true, //是否加载点赞处理器, 防止连续点击出现问题
    page: 1, //上拉页码
    loadingText: "正在加载更多....", //上拉加载文字
    showPubButtom: false
  },
  navigationSwitch: function(event) {
    wx.navigateTo({
      url: '../user/user'
    })
  },
  //事件处理函数
  bindViewTap: function() {
    wx.navigateTo({
      url: '../logs/logs'
    })
  },
  onLoad: function () {
    let that = this
    wx.getSystemInfo({
      success: function (res) {
        var scrollViewHeight = 750 * res.windowHeight / res.windowWidth; //rpx
        console.log(res.windowWidth)
        var scrollTop = res.windowWidth * 400 / 750; //矢量转换后的高度
        that.setData({
          scrollViewHeight: scrollViewHeight,
          scrollTop: scrollTop,
          fixedTop: false
        });
      }
    });
    this.setData({ 'viewData.style': myStyle + '40px;' })
    this.refreshIndex(1, true);
  },
  getUserInfo: function(e) {
    console.log(e)
    app.globalData.userInfo = e.detail.userInfo
    this.setData({
      userInfo: e.detail.userInfo,
      hasUserInfo: true
    })
  },
  addHandle: function () {
    wx.navigateTo({
      url: '../square_add/square_add'
    })
  },
  // 图片点击事件
  previewImage: function (event) {
    console.log('212112')
    wx.navigateTo({
      url: '../components/image_preview/image_preview?list=' + event.currentTarget.dataset.list + '&index=' + event.currentTarget.dataset.index
    })
    // var src = event.currentTarget.dataset.src;//获取data-src
    // console.log(src)
    // var imgList = event.currentTarget.dataset.list;//获取data-list
    // console.log(imgList)
    // //图片预览
    // wx.previewImage({
    //   current: src, // 当前显示图片的http链接
    //   urls: imgList // 需要预览的图片http链接列表
    // })
  },
  onReachBottom: function () {
    console.log('000000')
    this.refreshIndex(this.data.page);
  },

  onRestore(e) {
    console.log('onRestore:', e)
  },
  refreshIndex: function (page, refresh) {
    this.setData({
      skeletonShow: false
    })
    if (!page) {
      page = 1
    }
    //下拉刷新
    if (refresh === true) {
      console.log("sdfsdsf")
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
              d.push(item);
            })
          }
          let newPage = page + 1
          this.setData({
            data: d,
            page: newPage,
            skeletonShow: false,
          })
        }
        
      }

      
    })
  },
  onPullDownRefresh: function () {
    console.log('onPullDownRefresh')
    this.refreshIndex(1, true);
    wx.stopPullDownRefresh();
  },

  onShow: function () {
    if (app.globalData.forceRefresh == true) {
      app.globalData.forceRefresh = false;
      this.refreshIndex(1, true);
    }

    if (app.globalData.userInfo !== null) {
      if (app.globalData.userInfo.MiniId == 123162) {
        this.setData({
          showPubButton: true
        })
      }
    }
  },
  goBack: function () {
    let pages = getCurrentPages();   //获取小程序页面栈
    let beforePage = pages[pages.length - 2];  //获取上个页面的实例对象
    beforePage.setData({      //直接修改上个页面的数据（可通过这种方式直接传递参数）
      txt: '修改数据了'
    })
    beforePage.goUpdate();   //触发上个页面自定义的go_update方法
    wx.navigateBack({         //返回上一页  
      delta: 1
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
    let isLike = e.target.dataset.like;
    if (isLike === 0) {
      isLike = 1;
    }else {
      isLike = 0;
    }
    
    app.httpClient.post(app.getApi('squareLike'), {
      'id': e.target.dataset.id,
      'o': isLike,
      'page':this.data.page
    }).then(res => {
      let response = res.data
      if (response.code == 200) {
        newList[e.target.dataset.index].is_like = isLike
        if (isLike === 1) {
          newList[e.target.dataset.index].like_count += 1
        }else{
          newList[e.target.dataset.index].like_count -= 1
        }
        this.setData({
          data: newList
        })
        wx.showToast({
          title: '操作成功',
          icon: 'none',
          duration: 1000,
        })
      }else if(response.code != 401){
        wx.showToast({
          title: '操作失败',
          icon: 'none',
          duration: 1000,
        })
      }
      console.log('set like handle')
      this.setData({
        likeHandle: true
      })
    })
  }
})
