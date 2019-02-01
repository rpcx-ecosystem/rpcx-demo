package product

import (
	"context"
	"errors"
	"io/ioutil"
	"mime"
	"path/filepath"

	"github.com/rpcx-ecosystem/rpcx-demo/service/product/model"
)

// ProductService 产品图片服务
type ProductService struct {
	Dir string
}

// New 新建一个产品图片服务对象.
func New(d string) *ProductService {
	return &ProductService{Dir: d}
}

// Get 提供指定的图片.
// 仅用作演示，并没有进行缓存.
func (s ProductService) Get(ctx context.Context, req model.ImageRequest, res *model.ImageResponse) error {
	name := string(req)
	ctype := mime.TypeByExtension(filepath.Ext(name))
	if ctype == "" {
		return errors.New("unknown mime type ")
	}

	data, err := ioutil.ReadFile(filepath.Join(s.Dir, name))
	if err != nil {
		return err
	}

	*res = model.ImageResponse{
		ContentType:   ctype,
		ContentLength: len(data),
		Content:       data,
	}

	return nil
}
