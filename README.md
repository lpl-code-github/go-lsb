# Go-LSB

基于包"github.com/gzcharleszhang/stego"

项目源码：https://github.com/gzcharleszhang/stego

# Directory meaning

- test_file：提供了一些测试文件
  - msg.txt：需要写入图像的消息数据
  - test.jpg：需要写入数据的图像
- test_shell：提供了一些**测试脚本**（详见Quick Start）
- utils：一些function
  - getLength.go：将图像转为RGBA格式，返回使用最低有效位编码可以在该图像中存储多少字节
  - LsbRead.go：读取图像文件中编码的消息数据
  - LsbWrite.go：将消息数据写入到图像文件中
  - utils.go：一些工具方法
- lsbMain.go：初始化命令行参数以及参数对应执行的函数

# Quick Start

## clone代码

## 终端进入根目录下，删除根目录下的go.mod，重新初始化项目

```bash
rm -rf go.mod
go mod init lsb
```

## 进入test_shell文件

```bash
cd test_shell
```

## 执行脚本init.sh编译lsbMain.go

```bash
bash init.sh
```

![image-20220415225414282](https://pic.imgdb.cn/item/625986f3239250f7c55dc0ad.png)

## 查看帮助文档

```bash
bash help.sh
```

![image-20220415225443981](https://pic.imgdb.cn/item/62598711239250f7c55df91d.png)

## 使用测试文件测试

### 测试一：将消息数据写入图片中，并导出新的图片

```bash
bash write_lsb.sh
```

![image-20220415225600512](https://pic.imgdb.cn/item/6259875d239250f7c55eab2b.png)

### 测试二：读取图片中的消息数据，以控制台方式输出

```bash
bash read_lsb_console.sh
```

![image-20220415225649986](https://pic.imgdb.cn/item/6259878e239250f7c55f0726.png)

### 测试三：读取图片中的消息数据，以文件的形式输出

```bash
bash read_lsb_file.sh
```

![image-20220415225748873](https://pic.imgdb.cn/item/625987c9239250f7c55f92d5.png)

除了测试脚本及测试文件以外，可以根据bash help输出的帮助文档自行选择需要操作的图像文件和消息数据文件