package utils

import (
	"fmt"
	"image/color"
	"io/ioutil"
)

func LsbRead(pictureInputFile string, messageOutputFile string) {
	// 从图像中编码的前四个字节获取消息的大小
	sizeOfMessage := getSizeOfMessageFromImage(pictureInputFile)

	// 从图片文件中读取消息
	msg := decodeMessageFromPicture(pictureInputFile, 4, sizeOfMessage)

	// 如果指定要将消息写入的位置
	if messageOutputFile != "" {
		// 将消息写入指定的输出文件
		err := ioutil.WriteFile(messageOutputFile, msg, 0644)
		if err != nil {
			fmt.Println("写入文件时出错: ", messageOutputFile)
		}
	} else { // 否则，将消息打印到控制台
		fmt.Printf("%s\n", string(msg))
	}
}

// 从图像中编码的前四个字节获取消息的大小
func getSizeOfMessageFromImage(pictureInputFile string) (size uint32) {
	// 使用LSB隐写术，解码图片中的消息，并将其作为字节序列返回
	sizeAsByteArray := decodeMessageFromPicture(pictureInputFile, 0, 4)

	// 给定四个字节，将返回32位无符号整数，这是这四个字节的组成部分（一个是MSB）
	size = combineToInt(sizeAsByteArray[0], sizeAsByteArray[1], sizeAsByteArray[2], sizeAsByteArray[3])
	return
}

// 使用LSB隐写术，解码图片中的消息，并将其作为字节序列返回
func decodeMessageFromPicture(pictureInputFile string, startOffset uint32, msgLen uint32) (message []byte) {
	var byteIndex uint32 = 0
	var bitIndex uint32 = 0

	// 将给定图像转换为RGBA图像
	rgbIm := imageToRGBA(decodeImage(pictureInputFile))

	// 获取宽高
	width := rgbIm.Bounds().Dx()
	height := rgbIm.Bounds().Dy()

	var c color.RGBA
	var lsb byte

	message = append(message, 0)

	// 遍历图像中的每个像素，并将消息一点一点地缝合在一起
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			// 获取像素的颜色
			c = rgbIm.RGBAAt(x, y)

			/* 红色 */
			//从该像素的红色分量中获取最低有效位
			lsb = getLSB(c.R)
			// 将此位添加到消息中
			message[byteIndex] = setBitInByte(message[byteIndex], bitIndex, lsb)
			bitIndex++

			if bitIndex > 7 { // 填充完一个字节后，继续下一个字节
				bitIndex = 0
				byteIndex++

				if byteIndex >= msgLen+startOffset {
					return message[startOffset : msgLen+startOffset]
				}

				message = append(message, 0)
			}

			/* 绿色 */
			//从该像素的绿色分量中获取最低有效位
			lsb = getLSB(c.G)
			// 将此位添加到消息中
			message[byteIndex] = setBitInByte(message[byteIndex], bitIndex, lsb)
			bitIndex++

			if bitIndex > 7 { // 填充完一个字节后，继续下一个字节

				bitIndex = 0
				byteIndex++

				if byteIndex >= msgLen+startOffset {
					return message[startOffset : msgLen+startOffset]
				}

				message = append(message, 0)
			}

			/* 蓝色 */
			//从该像素的蓝色分量中获取最低有效位
			lsb = getLSB(c.B)
			// 将此位添加到消息中
			message[byteIndex] = setBitInByte(message[byteIndex], bitIndex, lsb)
			bitIndex++
			// 填充完一个字节后，继续下一个字节
			if bitIndex > 7 {
				bitIndex = 0
				byteIndex++

				if byteIndex >= msgLen+startOffset {
					return message[startOffset : msgLen+startOffset]
				}

				message = append(message, 0)
			}
		}
	}
	return
}

// 给定四个字节，将返回32位无符号整数，这是这四个字节的组成部分（一个是MSB）
func combineToInt(one, two, three, four byte) (ret uint32) {
	ret = uint32(one)
	ret = ret << 8
	ret = ret | uint32(two)
	ret = ret << 8
	ret = ret | uint32(three)
	ret = ret << 8
	ret = ret | uint32(four)
	return
}

// 给定一个字节，将返回该字节的最低有效位
func getLSB(b byte) byte {
	if b%2 == 0 {
		return 0
	} else {
		return 1
	}
	return b
}

// 将字节中的特定位设置为给定值并返回新字节
func setBitInByte(b byte, indexInByte uint32, bit byte) byte {
	var mask byte = 0x80
	mask = mask >> uint(indexInByte)

	if bit == 0 {
		mask = ^mask
		b = b & mask
	} else if bit == 1 {
		b = b | mask
	}
	return b
}
