package stg

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"io"
)

type PicCache map[string]*ebiten.Image

var defaultPicCache = make(PicCache)

func GetPic(s string) *ebiten.Image {
	return defaultPicCache[s]
}

func InitPicReader(s string, reader io.Reader) error {
	return defaultPicCache.InitReader(s, reader)
}
func (cache PicCache) InitReader(s string, reader io.Reader) error {
	img, _, err := image.Decode(reader)
	if err != nil {
		return err
	}
	cache.Init(s, img)
	return nil
}

func InitPic(s string, img image.Image) {
	defaultPicCache.Init(s, img)
}
func (cache PicCache) Init(s string, img image.Image) {
	cache[s] = ebiten.NewImageFromImage(img)
}
