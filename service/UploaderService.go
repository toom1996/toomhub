// @Description
// @Author    2020/8/27 17:01
package service

import "mime/multipart"

//  加载文件接口 mime/multipart 包的基础是post请求，即基于post请求来实现的
type Uploader interface {
	Upload(file multipart.File, fileHeader *multipart.FileHeader) (string, error)
}
