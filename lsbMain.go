package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"lsb/utils"
)

var pictureInputFile string
var pictureOutputFile string
var messageInputFile string
var messageOutputFile string
var printLen bool
var read bool
var write bool
var help bool

func init() {
	flag.BoolVar(&read, "r", false, "指定是否要从给定的图片文件中读取消息")
	flag.BoolVar(&write, "w", false, "指定是否要将消息写入给定的图片文件")

	flag.StringVar(&pictureInputFile, "imgsrc", "", "输入图像的路径")
	flag.StringVar(&pictureOutputFile, "imgout", "", "输出图像的路径")

	flag.StringVar(&messageInputFile, "msgsrc", "", "输入消息的路径")
	flag.StringVar(&messageOutputFile, "msgout", "", "输出消息的路径")

	flag.BoolVar(&help, "help", false, "帮助")

	flag.Parse()
}

func main() {
	if (!read && !write) || help {
		if help {
			fmt.Println("----------帮助文档----------")
			fmt.Println("两种模式: write and read:")
			fmt.Println("- Write: 把一条信息写到指定的图片")
			fmt.Println("    例如: ./lsbMain -w -msgsrc [消息文件] -imgsrc [需要编码的图片] -imgout [编码后的图片]")
			fmt.Println("- Read: 获取一张照片并阅读其中的信息")
			fmt.Println("    例如，指定输出文件: ./lsbMain -r -imgsrc [需要读取信息的图片] -msgout [从图片中读取的消息文件]")
			fmt.Println("    例如，不指定输出文件，将直接输出到控制台: ./lsbMain -r -imgsrc [需要读取信息的图片]")
		} else if !read || !write {
			fmt.Println("必须指定读、写，有关更多信息参照 -help 指令给出的信息\n")
			fmt.Println("\t+ EX: ./lsbMain -help")
		}
		return
	}

	if write {
		if pictureInputFile != "" { // 如果文件名不为空
			utils.GetLength(pictureInputFile) // 执行检查操作
		}
		if messageInputFile == "" || pictureInputFile == "" || pictureOutputFile == "" {
			fmt.Println("Error: 要在写入模式下运行lsbMain，必须指定: ")
			fmt.Println("-imgsrc: 需要编码的图像")
			fmt.Println("-imgout: 编码图像存储在哪里")
			fmt.Println("-msgsrc: 要嵌入到图像中的消息")
			return
		}

		message, err := ioutil.ReadFile(messageInputFile) // 从消息文件中读取消息
		if err != nil {
			fmt.Println("Error: -msgsrc参数错误，从消息文件中读取失败")
			return
		}

		utils.LsbWrite(pictureInputFile, pictureOutputFile, message) // 将消息编码到图像文件中
	}

	if read {
		if pictureInputFile == "" { //
			fmt.Println("Error: 要在读入模式下运行lsbMain，必须指定: ")
			fmt.Println("-imgsrc: 包含嵌入消息的图片")
			return
		}
		utils.LsbRead(pictureInputFile, messageOutputFile) // 从指定图片中读取消息
	}
}
