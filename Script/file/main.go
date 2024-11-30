package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	BaseDir string
	Op      int64
)

func init() {

}
func parseFlag() {
	flag.StringVar(&BaseDir, "dir", ".", "文件路径")
	flag.Int64Var(&Op, "op", 1, "操作类型:\n1. 输出整理结果\n2. 聚合视频到一起(文件名加路径前缀)")

	flag.Parse()
}

func main() {
	parseFlag()
	fmt.Println("main start", BaseDir, Op)

	var err error

	switch Op {
	case 1:
		err = statics()
	case 2:

	default:
		err = fmt.Errorf("无效的操作类型Op")
	}

	if err != nil {
		fmt.Printf("操作失败, 失败原因是: %s", err.Error())
	}
}

type VideoCategory struct {
	CategoryName string
	Count        int32
}

// 输出整理结果
// 1. 按照视频分类 输出各分类的视频数量
//
// 视频文件名格式为 名称-编号-简略内容说明
//
// 如果视频不知道名称, 即名称为不知名, 或者特殊的视频资源(比如顶级资源)
//
// 则文件名格式为 不知名/顶级-名称/分类-编号-简略内容说明
func statics() (err error) {

	var head, name, category, order, content string

	files, err := os.ReadDir(BaseDir)
	if err != nil {
		return
	}
	for _, file := range files {
		head, name, category, order, content = "", "", "", "", ""
		if !file.IsDir() {
			names := strings.Split(file.Name(), "-")
			if len(names) == 3 {
				head = ""
				name = names[0]
				category = ""
				order = names[1]
				content = names[2]

			} else if len(names) == 4 {
				flag := names[0]
				if flag == "顶级" || flag == "不知名" {
					head = flag
				} else {
					return errors.New("123")
				}

			}
		}
	}

	_, _, _, _, _ = head, name, category, order, content

	return
}
