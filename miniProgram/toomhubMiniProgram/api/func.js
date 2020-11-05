const app = getApp()
// 生成缩略图
export const getThumbnail = (options, callback) => {
    wx.showLoading({
      title: '分享图片生成中',
    })

    let title = options.target.dataset.title;
    let list = options.target.dataset.list;
    let id = options.target.dataset.id;
    let type = options.target.dataset.type;
    let cover = options.target.dataset.cover;
    let imageUrl = '';
    let createdBy = 'nooooooooooooooooooooob';
    let avatar = 'https://thirdwx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTIAqVYWPkszmEb057h6SXMRVictUKPOW2ajvwSrfnlc77yHAFibdCwIic7lZxAM0X1h3SOleFwgQTdgA/132';
    let internetImgData = [options.target.dataset.cover, 'https://thirdwx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTIAqVYWPkszmEb057h6SXMRVictUKPOW2ajvwSrfnlc77yHAFibdCwIic7lZxAM0X1h3SOleFwgQTdgA/132']
    let imgData = [];
    for (let i = 0; i < internetImgData.length; i++) {
      wx.getImageInfo({
        src: internetImgData[i],
        success (res) {
          console.log(res)
          imgData[i] = {
            path: res.path,
            width: res.width,
            height: res.height,
          };
          console.log(imgData)
          if (imgData.length === internetImgData.length) {
            createPhoto(options, title, id, type, createdBy, imgData, callback)
          }
        }
      })
      console.log('666666')
    }
}



function createPhoto (options, title, id, type, createdBy, imgData, callback) {
  let imageUrl =  ''
  if (type == 1) {
    imageUrl = imgData[0];

    const context = wx.createCanvasContext('videoCanvas', self);

    //截取用户名
    console.log(strlen(createdBy))
    if (strlen(createdBy) > 14) {
      createdBy = createdBy.substring(0,14) + '...';
    }

    var w = imgData[1].width
    var h = imgData[1].height
    var dw = 500/w          //canvas与图片的宽高比
    var dh = 400/h

    //绘制头像
    var avatarurl_width = 50;    //绘制的头像宽度
    var avatarurl_heigth = 50;   //绘制的头像高度
    var avatarurl_x = 0;   //绘制的头像在画布上的位置
    var avatarurl_y = 0;   //绘制的头像在画布上的位置
    context.save();
    context.beginPath(); //开始绘制
    context.arc(avatarurl_width / 2 + avatarurl_x, avatarurl_heigth / 2 + avatarurl_y, avatarurl_width / 2, 0, Math.PI * 2, false);
    context.clip()	//裁剪
    context.drawImage(imgData[1].path, avatarurl_x, avatarurl_y, avatarurl_width, avatarurl_heigth); // 推进去图片，必须是https图片
    context.restore(); //恢复之前保存的绘图上下文 恢复之前保存的绘图上下午即状态 还可以继续绘制

    //绘制用户昵称
    context.setFontSize(20)
    context.fillText(createdBy, 55, 30)

    console.log(imgData[0].path)
     // canvas宽高
    let cWidth = 500;
    let cHeight = 400;

    let imgWidth = imgData[0].width;
    let imgHeight = imgData[0].height;

    let dWidth = cWidth/imgWidth;  // canvas与图片的宽度比例
    let dHeight = cHeight/imgHeight;  // canvas与图片的高度比例
    if (imgWidth > cWidth && imgHeight > cHeight || imgWidth < cWidth && imgHeight < cHeight) {
      if (dWidth > dHeight) {
        console.log(11111)
        context.drawImage(imgData[0].path, 0, (imgHeight - cHeight/dWidth)/2, imgWidth, cHeight/dWidth, 0, 50, cWidth, cHeight)
      } else {
        console.log(222222)
        context.drawImage(imgData[0].path, (imgWidth - cWidth/dHeight)/2, 50, cWidth/dHeight, imgHeight, 0, 0, cWidth, cHeight)
      }
    } else {
      if (imgWidth < cWidth) {
        console.log(3333)
        context.drawImage(imgData[0].path, 0, (imgHeight - cHeight/dWidth)/2, imgWidth, cHeight/dWidth, 0, 0, cWidth, cHeight)
      } else {
        console.log(4444)
        context.drawImage(imgData[0].path, (imgWidth - cWidth/dHeight)/2, 20, cWidth/dHeight, imgHeight, 0, 0, cWidth, cHeight)
      }
    }

    context.drawImage('/static/icon/play.png', 0, 50, 200, 400); // 推进去图片，必须是https图片


    context.draw(false, ()=> {
      wx.canvasToTempFilePath({
        canvasId: 'videoCanvas',
        x: 0,
        y: 0,
        width: 500,
        height: 400,
        complete: (tmpRes) => {
          callback({
            src: tmpRes.tempFilePath,
            id: id,
            title: title
          });
          wx.hideLoading({
          })
        }
      }, this);
    })
  } else {
    imageUrl = list[0] + app.globalData.imageThumbnailParam;
  }
}


export const strlen = (str) => {
  var len = 0;
  for (var i = 0; i < str.length; i++) {
    var c = str.charCodeAt(i);
    //单字节加1
    if ((c >= 0x0001 && c <= 0x007e) || (0xff60 <= c && c <= 0xff9f)) {
      len++;
    }
    else {
      len += 2;
    }
  }
  return len;
}

