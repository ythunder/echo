package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		// 解析 GET 请求参数
		queryParams := r.URL.Query()
		name := queryParams.Get("name")
		// 获取当前时间戳
		timestamp := time.Now().Unix()
		envName := os.Getenv("ENV_NAME")
		// 构建响应内容
		response := fmt.Sprintf("Hello, %s! Current timestamp: %d, cluster: %v", name, timestamp, envName)

		// 将响应内容写回响应
		_, err := w.Write([]byte(response))
		if err != nil {
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
			return
		}
	})

	// 启动 HTTP 服务器
	log.Println("Server listening on port 8080...")
	go func() {
		time.Sleep(30*time.Second)
		for {
			log.Println("CURRENT LOG ", time.Now())
		}
	}()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
