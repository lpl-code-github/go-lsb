#!/bin/bash
echo "操作: 把msg.txt的内容写到图片test.jpg中，并将包含msg的图片输出为lsb_test.jpg"
echo "命令: ./lsbMain -w -msgsrc ../test_file/msg.txt -imgsrc ../test_file/test.jpg -imgout ../test_file/lsb_test.jpg"
./lsbMain -w -msgsrc ../test_file/msg.txt -imgsrc ../test_file/test.jpg -imgout ../test_file/lsb_test.jpg
