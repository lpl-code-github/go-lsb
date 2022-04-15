package utils

import (
	"bufio"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func LsbWrite(pictureInputFile string, pictureOutputFile string, message []byte) {
	rgbIm := imageToRGBA(decodeImage(pictureInputFile)) // 将指定路径的图片转为RGBA格式

	var messageLength = uint32(len(message)) // 获得message字节数组的长度，指定一个int占32位，32/8=4个字节

	var width = rgbIm.Bounds().Dx()  // 宽
	var height = rgbIm.Bounds().Dy() // 高
	var c color.RGBA
	var bit byte
	var ok bool

	if maxEncodeSize(rgbIm) < messageLength+4 { // 最低有效位编码在该图像中存储多少字节 < 消息内容长度+4
		print("Error: 您试图编码的消息太大")
		return
	}

	one, two, three, four := splitToBytes(messageLength) //给定一个无符号整数，将该整数拆分为四个字节

	message = append([]byte{four}, message...)
	message = append([]byte{three}, message...)
	message = append([]byte{two}, message...)
	message = append([]byte{one}, message...)

	// 创建一个有100个缓冲的channel 当接收channel的时候,如果channel的缓冲区为空,则会阻塞当前goroutine,直到有新数据可取
	ch := make(chan byte, 100)
	go getNextBitFromString(message, ch) // 返回字符串中的下一个后续位

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {

			c = rgbIm.RGBAAt(x, y) // 获取这个像素的颜色

			/*  RED  */
			bit, ok = <-ch
			if !ok { // 如果message中没有更多的信息
				rgbIm.SetRGBA(x, y, c)
				encodePNG(pictureOutputFile, rgbIm) // 把编码过的图像写出到指定文件
				return
			}
			setLSB(&c.R, bit)

			/*  GREEN  */
			bit, ok = <-ch
			if !ok {
				rgbIm.SetRGBA(x, y, c)
				encodePNG(pictureOutputFile, rgbIm)
				return
			}
			setLSB(&c.G, bit)

			/*  BLUE  */
			bit, ok = <-ch
			if !ok {
				rgbIm.SetRGBA(x, y, c)
				encodePNG(pictureOutputFile, rgbIm)
				return
			}
			setLSB(&c.B, bit)

			rgbIm.SetRGBA(x, y, c)
		}
	}

	encodePNG(pictureOutputFile, rgbIm) // 将给定的图像写入文件名中的给定路径
}

// 每次调用都将返回字符串中的下一个后续位
func getNextBitFromString(byteArray []byte, ch chan byte) {
	var offsetInBytes int = 0
	var offsetInBitsIntoByte int = 0
	var choiceByte byte

	lenOfString := len(byteArray)

	for {
		if offsetInBytes >= lenOfString {
			close(ch)
			return
		}

		choiceByte = byteArray[offsetInBytes]
		ch <- getBitFromByte(choiceByte, offsetInBitsIntoByte)

		offsetInBitsIntoByte++

		if offsetInBitsIntoByte >= 8 {
			offsetInBitsIntoByte = 0
			offsetInBytes++
		}
	}
}

// 将给定的图像写入文件名中的给定路径
func encodePNG(filename string, img image.Image) {
	fo, err := os.Create(filename)

	if err != nil {
		log.Fatalf("Error creating file %s: %v", filename, err)
	}

	defer fo.Close()
	defer fo.Sync()

	writer := bufio.NewWriter(fo)
	defer writer.Flush()

	err = png.Encode(writer, img)
}

// 给定的位将从该字节返回一位
func getBitFromByte(b byte, indexInByte int) byte {
	b = b << uint(indexInByte)
	var mask byte = 0x80

	var bit byte = mask & b

	if bit == 128 {
		return 1
	}
	return 0
}

// 给定字节将该字节的最低有效位设置为给定值（其中true为1，false为0）
func setLSB(b *byte, bit byte) {
	if bit == 1 {
		*b = *b | 1
	} else if bit == 0 {
		var mask byte = 0xFE
		*b = *b & mask
	}
}

// 给定一个无符号整数，将该整数拆分为四个字节
func splitToBytes(x uint32) (one, two, three, four byte) {
	one = byte(x >> 24)
	var mask uint32 = 255

	two = byte((x >> 16) & mask)
	three = byte((x >> 8) & mask)
	four = byte(x & mask)
	return
}
