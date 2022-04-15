#!/bin/bash
echo "操作: 读取lsb_test.jpg图像中的消息并输出到lsb_msg.txt文件中"
echo "命令: ./lsbMain -r -imgsrc ../test_file/lsb_test.jpg -msgout ../test_file/lsb_msg.txt"
./lsbMain -r -imgsrc ../test_file/lsb_test.jpg -msgout ../test_file/lsb_msg.txt
