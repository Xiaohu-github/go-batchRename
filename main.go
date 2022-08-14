package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

//获取指定路径下的文件
func getAllFile(path string, s []string) ([]string, error) {
	//ioutil 是io的工具包
	//ioutil.ReadDir ：读取一个目录下子内容：子文件和子目录，但是仅有一层
	rd, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println("read dir fail:", err)
		return s, err
	}
	for _, file := range rd {
		if !file.IsDir() {
			fullName := path + file.Name()
			s = append(s, fullName)
		}
	}

	return s, nil

}

func main() {
	fmt.Println("已启动...")
	flag.Parse()
	var pathName string = "D:/***/***/****"
	var files []string
	files, err := getAllFile(pathName, files)
	if err != nil {
		fmt.Println("read dir fail:", err)
		return
	}
	for i, value := range files {
		// fileNameWithSuffix := path.Base(value)
		fileType := path.Ext(value)
		// fmt.Println("fileType=>", fileType)
		newName := fmt.Sprintf("%s%d%s", pathName, i, fileType)
		err := os.Rename(value, newName)
		if err != nil {
			fmt.Println("文件名修改失败：", err)
			break
		}
		fmt.Printf("%s===>%s\n", value, newName)
	}
	fmt.Println("已结束...")
}
