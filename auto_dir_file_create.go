package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

//手动创建go目录和文件累死，决定写个自动创建的，节约一点时间

const (
	CRACK        = 0
	CSTR         = "cracking_the_code_interview"
	LEETCODEMAIN = 1
	LSTR         = "leetcode_main"
	EASY         = 0
	MEDIUM       = 1
	DIFFICULT    = 2
)

func AutoDirFileCreate(mainPath int, difficulty int, index string, url string) {
	var subPathStr, fatherPath string
	urlSplitedList := strings.Split(url, "/")
	length := len(urlSplitedList)
	var urlSplited string
	if urlSplitedList[length-1] == "" {
		urlSplited = urlSplitedList[length-2]
	} else {
		urlSplited = urlSplitedList[length-1]
	}
	urlSplitedS := strings.Split(urlSplited, "-")
	urlSplited = strings.Join(urlSplitedS, "_")
	switch difficulty {
	case EASY:
		subPathStr = "easy_" + index + "_" + urlSplited
	case MEDIUM:
		subPathStr = "medium_" + index + "_" + urlSplited
	case DIFFICULT:
		subPathStr = "difficult_" + index + "_" + urlSplited
	}
	switch mainPath {
	case CRACK:
		fatherPath = CSTR

	case LEETCODEMAIN:
		fatherPath = LSTR
		switch difficulty {
		case EASY:
			fatherPath += "/easy"
		case MEDIUM:
			fatherPath += "/medium"
		case DIFFICULT:
			fatherPath += "/difficult"
		}
	}
	finalPath := "./" + fatherPath + "/" + subPathStr
	if _, err := os.Stat(finalPath); os.IsNotExist(err) {
		err = os.MkdirAll(finalPath, 0777)
		if err != nil {
			fmt.Println(err)
		}
	}
	file, er := os.Open(finalPath + "/solution.go")
	defer func() { _ = file.Close() }()
	if er != nil && os.IsNotExist(er) {
		file, _ = os.Create(finalPath + "/solution.go")
	}
	getDate := time.Now().Format("2006-01-02")
	writeString := "package " + subPathStr + "\n\n//" + url + "\n" + "//" + getDate + "\n"
	_, _ = file.WriteString(writeString)

	fmt.Println(finalPath)

}
