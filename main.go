package main

import (
	"image/jpeg"
	"log"
	"os"
)

const (
	// IMAGEDIR 图片根目录
	IMAGEDIR = "images/"
)

func main() {
	// 获取所有图片的文件名
	imgFileNames := GetImgFileNames(IMAGEDIR)

	// 添加文件流
	var imgFiles []*os.File
	for _, fileName := range imgFileNames {
		file, err := os.Open(IMAGEDIR + fileName)
		if err != nil {
			log.Fatal("Error on opening file: ", err)
		} else {
			imgFiles = append(imgFiles, file)
		}
	}
	defer func() {
		for _, file := range imgFiles {
			file.Close()
		}
	}()

	// 顺序拼接
	newImg, err := SeqSpl(imgFiles)
	if nil != err {
		log.Fatal("Error on sequentially splice images, ", err)
	}

	// 保存文件
	FinalImage, _ := os.Create("final.jpeg")
	defer FinalImage.Close()
	if err := jpeg.Encode(FinalImage, newImg, &jpeg.Options{100}); err != nil {
		panic(err)
	}
}
