package apiserver

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	v1 "github.com/jdxj/sign/internal/apiserver/v1"
	"github.com/jdxj/sign/internal/pkg/config"
	"github.com/jdxj/sign/internal/pkg/logger"
)

var (
	srv *http.Server
)

func Start(conf config.APIServer) {
	r := gin.Default()
	registerRouter(r, conf)

	addr := fmt.Sprintf("%s:%s", conf.Host, conf.Port)
	srv = &http.Server{
		Addr:    addr,
		Handler: r,
	}
	go func() {
		err := srv.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Errorf("stop api server err: %s", err)
		}
	}()
}

func Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Errorf("shutdown api server err: %s", err)
		return
	}
	logger.Infof("api server already stop")
}

func registerRouter(r gin.IRouter, conf config.APIServer) {
	r.GET("/", hello)
	r.GET("/healthz", healthz)

	api := r.Group("/api")
	{
		v1.RegisterV1(api)
	}
}

func hello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
	})
}

func healthz(ctx *gin.Context) {
	logger.Debugf("receive healthz")
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
	})
}
