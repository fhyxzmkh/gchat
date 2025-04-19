package main

import (
	"fmt"
	"gchat/internal/config"
	"gchat/internal/dao"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"time"
)

func initDB() (*gorm.DB, error) {
	// 构建 MySQL DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.databaseName"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	if err := dao.InitTable(db); err != nil {
		return nil, fmt.Errorf("failed to init tables: %w", err)
	}

	return db, nil
}

func initWebServer() *gin.Engine {
	server := gin.Default()

	server.Use(cors.New(cors.Config{
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"X-Request-ID"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") {
				return true
			}
			return strings.Contains(origin, "sunyongan.top")
		},
		MaxAge: 12 * time.Hour,
	}))

	return server
}

func main() {
	// 初始化配置
	if err := config.NewViperConfig(); err != nil {
		panic(fmt.Sprintf("failed to load config: %v", err))
	}

	// 初始化数据库
	_, err := initDB()
	if err != nil {
		panic(fmt.Sprintf("failed to init DB: %v", err))
	}

	// 初始化 Web 服务
	server := initWebServer()

	// 注册路由
	// registerRoutes(server, db)

	// 启动服务
	addr := fmt.Sprintf("%s:%d",
		viper.GetString("main.host"),
		viper.GetInt("main.port"))

	if err := server.Run(addr); err != nil {
		panic(fmt.Sprintf("failed to start server: %v", err))
	}
}
