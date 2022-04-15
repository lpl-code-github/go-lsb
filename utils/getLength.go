package utils

import (
	"fmt"
	_ "image/jpeg"
	_ "image/png"
)

func GetLength(pictureInputFile string) {
	if pictureInputFile == "" {
		fmt.Println("Error: 文件路径为空，无法检查其最大编码长度的图像")
		return
	}

	fmt.Printf("检查文件：%s\n", pictureInputFile)
	rgbIm := imageToRGBA(decodeImage(pictureInputFile)) // 将路径中的图片转换为RGBA图像

	var sizeInBytes = maxEncodeSize(rgbIm) // 检查大小
	fmt.Println("单位B:\t", sizeInBytes)
	fmt.Println("单位KB:\t", float64(sizeInBytes)/1000)
	fmt.Println("单位MB:\t", (float64(sizeInBytes)/1000)/1000)

}
