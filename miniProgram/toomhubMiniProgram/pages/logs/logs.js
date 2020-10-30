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
          // var videoContextPrev = wx.createVideoContext('这里写videoid');
          // videoContextPrev.pause();
          let query = wx.createSelectorQuery()
            query.select('#'+ res.id).boundingClientRect( (rect) => {
                console.log(rect.top)
            }).exec(function(res) {
              console.log(res)
            })
        }
      })
  },
  onUnload() {
    if (this._observer) this._observer.disconnect()
  }
})
