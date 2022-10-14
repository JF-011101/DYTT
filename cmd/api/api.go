/*
 * @Author: jf-011101 2838264218@qq.com
 * @Date: 2022-08-02 14:03:24
 * @LastEditors: jf-011101 2838264218@qq.com
 * @LastEditTime: 2022-08-19 22:33:22
 * @FilePath: \dytt\cmd\api\main.go
 * @Description: Use the API service provided by gin to
 send the HTTP request to the RPC micro server
*/

package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/jf-011101/dytt/internal/api/handlers"
	"github.com/jf-011101/dytt/internal/api/rpc"
	"github.com/jf-011101/dytt/internal/pkg/ilog"
	"github.com/jf-011101/dytt/internal/pkg/ttviper"
)

var (
	Config      = ttviper.ConfigInit("TIKTOK_API", "apiConfig")
	ServiceName = Config.Viper.GetString("Server.Name")
	ServiceAddr = fmt.Sprintf("%s:%d", Config.Viper.GetString("Server.Address"), Config.Viper.GetInt("Server.Port"))
	EtcdAddress = fmt.Sprintf("%s:%d", Config.Viper.GetString("Etcd.Address"), Config.Viper.GetInt("Etcd.Port"))
)

func main() {
	fmt.Print("begin\n")
	rpc.InitRPC(&Config)
	fmt.Print("initrpc over")

	r := gin.New()
	r.Static("/static", "static")
	r.LoadHTMLGlob("templates/*")

	// Add a ginzap middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format.
	r.Use(ginzap.Ginzap(zap.L(), time.RFC3339, false))

	// Logs all panic to error log
	//   - stack means whether output the stack info.
	r.Use(ginzap.RecoveryWithZap(zap.L(), true))

	pir := r.Group("/pir")
	pir.POST("/refresh/", handlers.Refresh)
	pir.GET("/query/", handlers.QueryUserBoundary)
	pir.POST("/query/", handlers.QueryUser)
	// douyin := r.Group("/douyin")

	// user := douyin.Group("/user")

	// user.POST("/login/", handlers.Login)
	// user.POST("/register/", handlers.Register)
	// user.GET("/", handlers.GetUserById)

	// video := douyin.Group("/feed")
	// video.GET("/", handlers.GetUserFeed)

	// publish := douyin.Group("/publish")
	// publish.POST("/action/", handlers.PublishAction)
	// publish.GET("/list/", handlers.PublishList)

	// favorite := douyin.Group("/favorite")
	// favorite.POST("/action/", handlers.FavoriteAction)
	// favorite.GET("/list/", handlers.FavoriteList)

	// comment := douyin.Group("/comment")
	// comment.POST("/action/", handlers.CommentAction)
	// comment.GET("/list/", handlers.CommentList)

	// relation := douyin.Group("/relation")
	// relation.POST("/action/", handlers.RelationAction)
	// relation.GET("/follow/list/", handlers.RelationFollowList)
	// relation.GET("/follower/list/", handlers.RelationFollowerList)

	srv := &http.Server{
		Addr:    ":8088",
		Handler: r,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			ilog.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ilog.Info("Shutting down http server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		ilog.Fatal("Server forced to shutdown:", err)
	}

	ilog.Info("Server exiting")
}
