package main

import (
	"fmt"
	"image"
	_ "image/gif" // 以空导入方式注入gif 图片格式驱动
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func main() {
	// 支持png，jpeg， gif
	width, height, err := imageSize(os.Args[1]) // 获取传入的图片文件的宽与高
	if err != nil {
		fmt.Println("get image size error:", err)
		return
	}
	fmt.Printf("image size: [%d, %d]\n", width, height)
}

func imageSize(imageFile string) (int, int, error) {
	f, _ := os.Open(imageFile) // 打开图文文件
	defer f.Close()

	img, _, err := image.Decode(f) // 对文件进行解码， 得到图片实例
	if err != nil {
		return 0, 0, err
	}

	b := img.Bounds() // 返回图片区域
	return b.Max.X, b.Max.Y, nil
}
