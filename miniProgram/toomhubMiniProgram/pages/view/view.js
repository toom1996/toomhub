// pages/view/view.js
const app = getApp()

import Toast from '@vant/weapp/toast/toast';
Page({

  /**
   * 页面的初始数据
   */
  data: {
    likeHandle: true, //是否加载点赞处理器, 防止连续点击出现问题
    avatar_url: '',
    created_by: '',
    created_at: '',
    content: '',
    argument_count: 0,
    collect_count: 0,
    like_count: 0,
    image:[
    ]
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {
    console.log(options)
    let id = options.id
    wx.showLoading({
      title: '加载中...',
    })
    wx.setNavigationBarTitle({
      title: '详情',
    })
    app.httpClient.get(app.getApi('getView') + '?id=' + id).then(res => {
      console.log(res.data.code)
      if (res.data.code == 200) {
        wx.hideLoading();
      }

      let response = res.data.data;
      console.log(response)
      this.setData({
        avatar_url: response.avatar_url,
        created_by: response.created_by,
        created_at: response.created_at,
        content: response.content,
        tag: response.tag,
        image: response.image,
        argument_count: response.argument_count,
        collect_count: response.collect_count,
        like_count: response.like_count,
        list: response.list,
        id: response.id,
        is_like: response.is_like
      })
    })

    // this.setData({
    //   image: options.list.split(",")
    // })
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
  onShareAppMessage: function (options) {
    
    let title = this.data.content;
    let list = this.data.list;

    if (app.strlen(title) > 14 ) {
      title = title.substring(0,14) + '...';
    }
    console.log('/pages/view/view?id=' + this.data.id)
    return {
      title: title,
      path: '/pages/view/view?id=' + this.data.id,
      imageUrl: list[0] + app.globalData.imageThumbnailParam
    }
  },
  // 图片点击事件
  previewImage: function (event) {
    console.log(event)
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

  //点赞处理函数
  likeHandle: function (e) {
    console.log(e)
    this.setData({
      likeHandle: false
    })

    let likeCount = this.data.like_count;
    let isLike = e.currentTarget.dataset.like;
    if (isLike === 0) {
      isLike = 1;
    } else {
      isLike = 0;
    }

    app.httpClient.post(app.getApi('squareLike'), {
      'id': e.currentTarget.dataset.id,
      'o': isLike,
    }).then(res => {
      let response = res.data
      if (response.code == 200) {
        this.setData({
          is_like: isLike
        })
        if (isLike === 1) {
          likeCount += 1
        } else {
          likeCount -= 1
        }
        this.setData({
          like_count: likeCount
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
      console.log('set like handle')
      this.setData({
        likeHandle: true
      })
    })
  }
})
