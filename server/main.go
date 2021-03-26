package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const VersionNumber = "v1"
func main() {
	router := mux.NewRouter()
	homeRouter := router.Methods(http.MethodGet).Subrouter()
	homeRouter.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		output := fmt.Sprintf("Welcome! The entry endpoint is %s/%s/", request.Host, VersionNumber)
		writer.Write([]byte(output))
	})
	server := http.Server{
		Addr: ":7777",
		Handler: router,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			fmt.Println("server error:", err)
			os.Exit(1)
		}
	}()
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)
	sig := <- c
	fmt.Println("Shutting down server with signal: ", sig)
	//定义30秒后关闭服务，这段时间用于处理正在进行的请求。不再接受新请求。
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(ctx)
}
