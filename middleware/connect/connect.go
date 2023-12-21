package connect

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // PostgreSQL 驅動程序
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
	models "todoList/entity/db/list"
)

var DB *gorm.DB
var rdb *redis.Client

func MigrateDatabase(db *gorm.DB) {
	db.AutoMigrate(&models.Table{})
	fmt.Println("Migration completed successfully!")
}

func ConnectToPostgres() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	portStr := os.Getenv("DB_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		panic(err)
	}
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), port, os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	MigrateDatabase(DB)
	if err != nil {
		panic(err)
	}
}

func ConnectToRedis() *redis.Client {
	// 載入環境變數
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("無法載入 .env 檔案")
	}

	//// 解析連接 Redis 所需的資訊
	redisAddr := os.Getenv("REDIS_ADDR")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisDB, _ := strconv.Atoi(os.Getenv("REDIS_DB"))

	// 建立 Redis 客戶端連接
	rdb = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       redisDB,
	})

	// 測試與 Redis 的連接
	_, err = rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("無法連接到 Redis：%v", err)
	}
	return rdb
}
