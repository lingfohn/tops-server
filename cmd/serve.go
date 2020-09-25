/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	auth2 "github.com/lingfohn/lime/auth"
	"github.com/lingfohn/lime/handler"
	"github.com/lingfohn/lime/middleware"
	"github.com/lingfohn/lime/storage"
	"github.com/spf13/cobra"
	"log"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: serve,
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func serve(cmd *cobra.Command, args []string) {
	storage.InitClient()
	defer storage.CloseClient()
	r := gin.Default()
	r.Use(middleware.RequestId())
	r.Use(middleware.PanicHandler())
	r.GET("/ping", middleware.JSONRenderWrap(ping))

	if err := auth2.InitAuthMiddleware(); err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	authMiddleware := auth2.AuthMiddleware()

	r.POST("/api/auth/login", authMiddleware.LoginHandler)
	r.GET("/api/user/nav", middleware.JSONRenderWrap(handler.UserNav))

	// menu

	mh := handler.NewMenuHandler()

	r.POST("/api/auth/menus", middleware.JSONRenderWrap(mh.CreateMenu))
	r.GET("/api/auth/menus", middleware.JSONRenderWrap(mh.QueryMenu))
	r.PUT("/api/auth/menus/:id", middleware.JSONRenderWrap(mh.ModifyMenu))
	r.DELETE("/api/auth/menus/:id", middleware.JSONRenderWrap(mh.DeleteMenu))
	// auth
	auth := r.Group("/api/auth")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/refresh_token", authMiddleware.RefreshHandler)
		auth.GET("/ping/:id", middleware.JSONRenderWrap(ping))
		auth.POST("/logout", authMiddleware.LogoutHandler)
	}

	user := r.Group("/api/user")
	user.Use(authMiddleware.MiddlewareFunc())
	{
		user.POST("/info", handler.UserInfo)
	}
	// 用户管理路由
	h := handler.NewUserHandler()
	r.POST("/api/users", middleware.JSONRenderWrap(h.CreateUser))
	r.GET("/api/users", middleware.JSONRenderWrap(h.ListUser))
	r.PUT("/api/users/:id", middleware.JSONRenderWrap(h.ModifyUser))
	r.DELETE("/api/users/:id", middleware.JSONRenderWrap(h.DeleteUser))

	// 权限管理
	ph := handler.NewPermissionHandler()
	r.POST("/api/auth/permissions",
		middleware.JSONRenderWrap(ph.CreatePermission))
	r.DELETE("/api/auth/permissions/:id", middleware.JSONRenderWrap(ph.DeletePermission))
	r.GET("/api/auth/permissions", middleware.JSONRenderWrap(ph.QueryPermission))
	r.PUT("/api/auth/permissions/:id", middleware.JSONRenderWrap(ph.ModifyPermission))

	// 角色管理
	rh := handler.NewRoleHandler()
	r.POST("/api/auth/roles", middleware.JSONRenderWrap(rh.CreateRole))
	r.DELETE("/api/auth/roles/:id", middleware.JSONRenderWrap(rh.DeleteRole))
	r.GET("/api/auth/roles", middleware.JSONRenderWrap(rh.RoleList))
	r.PUT("/api/auth/roles/:id", middleware.JSONRenderWrap(rh.ModifyRole))
	// 404 路由
	r.NoRoute(func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	r.Run()
}
func ping(c *gin.Context) error {
	c.Set("message", "pong")
	return nil
}
