package model

// ImageRequest 请求图片的名称
type ImageRequest string

// ImageResponse 图片信息
type ImageResponse struct {
	ContentType   string
	ContentLength int
	Content       []byte
}
