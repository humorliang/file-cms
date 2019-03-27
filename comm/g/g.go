package g

import (
	"fmt"
	"github.com/humorliang/file-cms/comm/setting"
	"flag"
)

const (
	AesKey   = "h123h3h23n42nw"
	PageSize = 3
)

//加载配置
func init() {
	mode := flag.String("mode", "dev", "this is dev mode options")
	flag.Parse()
	fmt.Println("	   ***************加载配置***************")
	fmt.Printf("	   ***************当前为：%s 模式**********", *mode)
	//初始化配置
	setting.ConfigInit(mode)
}
