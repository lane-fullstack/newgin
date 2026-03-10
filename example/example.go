package main

import (
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/lane-fullstack/newgin"
	"github.com/lane-fullstack/newgin/ctx"
)

func main() {
	// Create a new engine instance with default middleware (Logger and Recovery)
	r := newgin.Default()
	r.GET("/ping", func(c *ctx.Context) {
		c.Success("测试的")
	})

	// Or create a blank engine
	// r := newgin.New()

	// ... define routes ...

	r.Run(":8080")
}
func main2() {

	r := newgin.New()

	r.GET("/ping", func(c *ctx.Context) {
		c.Success("测试的")
	})

	adminSrv := &http.Server{
		Addr:           ":8081",
		Handler:        r,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := adminSrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
