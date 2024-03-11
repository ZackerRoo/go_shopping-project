package graceful

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// 优雅的停止server，当收到一个 os.Interrupt 或者 syscall.SIGTERM 信号.
func ShutdownGin(instance *http.Server, timeout time.Duration) {

	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("关闭 Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	if err := instance.Shutdown(ctx); err != nil {
		log.Fatal("Server 关闭:", err)
	}
	// 超时5秒 ctx.Done().
	select {
	case <-ctx.Done():
		log.Println("超时5秒.")
	}
	log.Println("Server 退出")
}
