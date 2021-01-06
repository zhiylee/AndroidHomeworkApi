package main

import (
	"androidHomeworkApi/pkg/setting"
	"androidHomeworkApi/router"
	"fmt"
	"net/http"
)

func main() {
	r := router.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        r,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	_ = s.ListenAndServe()
}

func test()  {

}