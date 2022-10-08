package main

import (
	"context"
	"fmt"
	"github.com/shallwepoker/ggpoker-hands-converter/internal/configs"
	"github.com/shallwepoker/ggpoker-hands-converter/internal/dao"
	"github.com/shallwepoker/ggpoker-hands-converter/internal/routers"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	dao.InitTables()

	port := configs.Config.Port
	router := routers.NewRouter()


	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(fmt.Sprintf("Listen: %s\n", err))
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		panic(fmt.Sprintf("Server forced to shutdown: %+v", err))
	}

	fmt.Println("Server exited")
}
