Page({
  data: {
    appear: false
  },
  onLoad() {
    this._observer = wx.createIntersectionObserver(this, {observeAll:true})
    this._observer
      .relativeTo('.scroll-view')
      .observe('.ball', (res) => {
        // console.log(res);
        if (res.intersectionRatio > 0) {
          console.log(res.id, '出现');
          // var videoContextPrev = wx.createVideoContext('这里写videoid');
          // videoContextPrev.pause();
          // let query = wx.createSelectorQuery()
          //   query.select('#'+ res.id).boundingClientRect( (rect) => {
          //       console.log(res.id)
          //       console.log(rect.top)
          //   }).exec(function(res) {
          //     console.log(res)
          //   })
        }else{
          console.log(res.id, '消失');
        }
      })
  },
  onUnload() {
    if (this._observer) this._observer.disconnect()
  }
})
