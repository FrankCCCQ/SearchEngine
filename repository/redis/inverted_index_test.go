package redis

import (
	"SearchEngine/config"
	logs "SearchEngine/pkg/logger"
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	// 这个文件相对于config.yaml的位置
	re := config.ConfigReader{FileName: "../../config/config.yaml"}
	config.InitConfigForTest(&re)
	logs.InitLog()
	InitRedis()
	fmt.Println("Write tests on values: ", config.Conf)
	m.Run()
}
