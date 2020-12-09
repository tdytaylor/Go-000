package main

import (
	"context"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	group, ctx := errgroup.WithContext(context.Background())

	srv := &http.Server{
		Addr: ":5000",
	}

	group.Go(func() error {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
		return nil
	})

	// 2. 通过信道阻塞主线程，监听信号
	group.Go(func() error {
		sigs := make(chan os.Signal)
		// //监听 Ctrl+C 信号
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		sig := <-sigs
		log.Fatalf("Shutdown Server by sig: %v", sig)
		// 关闭服务器
		return srv.Shutdown(ctx)
	})
	// 4. 退出
	group.Wait()
	log.Fatalln("Server exiting")
}
