package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/hhxsv5/go-redis-memory-analysis"  //blog http://blog.sina.com.cn/s/blog_9bbafb790102x2sd.html
)

// redis写入不同数量不同长度的value, 分析内存占用, 将结果导出到csv文件中
var rdb *redis.Client
var ctx context.Context

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		PoolSize: 100,
	})
	ctx = context.Background()
}

func main() {
	arrOne := [3]int{10, 1000, 5000}
	arrTwo := [4]int{10, 100, 250, 500}

	for i := range arrOne {
		for j := range arrTwo {
			key := fmt.Sprintf("key%vk_g%v", j, i)
			setLoop(j*1000, key, generateValue(i))
		}
	}
	analysis()
}

func setLoop(num int, key string, value string) {
	for i := 0; i < num; i++ {
		k := fmt.Sprintf("%s:%v", key, i)
		cmd := rdb.Set(ctx, k, value, -1)
		err := cmd.Err()
		if err != nil {
			fmt.Println("set error:", err)
		}
	}
}

func analysis() {
	analysis, err := gorma.NewAnalysisConnection("localhost", 6379, "")
	if err != nil {
		fmt.Println("analysis wrong,error info", err)
		return
	}
	defer analysis.Close()
	analysis.Start([]string{":"})
	err = analysis.SaveReports("./results")
	if err != nil {
		fmt.Println("save reports error:", err)
	}
}

func generateValue(size int) string {
	arr := make([]byte, size)
	for i := 0; i < size; i++ {
		arr[i] = '1'
	}
	return string(arr)
}
