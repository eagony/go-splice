package main

import (
	"github.com/nfnt/resize"
	"image"
	"image/draw"
	"log"
	"os"
)

// SeqSpl 顺序拼接
func SeqSpl(imgFiles []*os.File) (*image.NRGBA, error) {
	n := len(imgFiles)
	var bar Bar
	bar.NewOption(0, int64(n))

	// 最终高度
	finalImgHeight := 0
	// 最小宽度
	minImageWidth := 1<<15 - 1
	// 解析后的图片
	imgs := []image.Image{}

	for _, file := range imgFiles {
		if img, _, err := image.Decode(file); err != nil {
			log.Fatal("Error on decoding image: ", err)
			return nil, err
		} else {
			if img.Bounds().Max.X < minImageWidth {
				minImageWidth = img.Bounds().Max.X
			}
			finalImgHeight += img.Bounds().Max.Y
			imgs = append(imgs, img)
		}
	}
	finished := int64(1)
	//创建一个新RGBA图像
	newImg := image.NewNRGBA(image.Rect(0, 0, minImageWidth, finalImgHeight))
	// 先画上第一张
	firstImg := imgs[0]
	imgs = imgs[1:]
	draw.Draw(newImg, newImg.Bounds(), firstImg, firstImg.Bounds().Min, draw.Over)
	bar.Play(finished)
	// 轮流处理剩下的
	currentHeight := firstImg.Bounds().Max.Y
	for _, img := range imgs {
		fixSizedImg := resize.Resize(uint(minImageWidth), 0, img, resize.Lanczos3)
		//画上缩放后的图片（这里需要注意Y值的起始位置）
		draw.Draw(newImg, newImg.Bounds(), fixSizedImg, fixSizedImg.Bounds().Min.Sub(image.Pt(0, currentHeight)), draw.Over)
		currentHeight += fixSizedImg.Bounds().Max.Y
		finished++
		bar.Play(finished)
	}
	bar.Finish()

	return newImg, nil
}

// TODO: 按自定义分辨率拼接
// CusSolSpl ...
func CusSolSpl() {

}
