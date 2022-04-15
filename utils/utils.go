package utils

import (
	"bufio"
	"image"
	"image/draw"
	"log"
	"os"
)

// 将给定图像转换为RGBA图像
func imageToRGBA(src image.Image) *image.RGBA {
	b := src.Bounds()

	var m *image.RGBA
	var width, height int

	width = b.Dx()
	height = b.Dy()

	m = image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(m, m.Bounds(), src, b.Min, draw.Src)
	return m
}

// 读取并返回给定路径上的解码图像
func decodeImage(filename string) image.Image {
	inFile, err := os.Open(filename)

	if err != nil {
		log.Fatalf("Error: 不能打开文件 %s: %v", filename, err)
	}

	defer inFile.Close()
	reader := bufio.NewReader(inFile)
	img, _, errs := image.Decode(reader)
	if errs != nil {
		log.Println(errs)
	}
	return img
}

// 给定一幅图像，返回使用最低有效位编码可以在该图像中存储多少字节
func maxEncodeSize(img image.Image) uint32 {
	width := img.Bounds().Dx()  //宽
	height := img.Bounds().Dy() // 高
	return uint32((width*height*3)/8) - 4
}
