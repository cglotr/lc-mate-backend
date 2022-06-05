package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/arikama/go-arctic-tern/arctictern"
	"github.com/cglotr/lc-mate-backend/controller"
	"github.com/cglotr/lc-mate-backend/dao"
	"github.com/cglotr/lc-mate-backend/leetcode"
	"github.com/cglotr/lc-mate-backend/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	mysqlUsername := os.Getenv("MYSQL_USERNAME")
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	mysqlIp := os.Getenv("MYSQL_IP")
	mysqlPort := os.Getenv("MYSQL_PORT")
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")

	dataSourceName := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", mysqlUsername, mysqlPassword, mysqlIp, mysqlPort, mysqlDatabase)
	dataSourceName += "?charset=utf8mb4"
	dataSourceName += "&collation=utf8mb4_unicode_ci"

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatalln(err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln(err.Error())
	}

	migrationDir := "./migration"
	err = arctictern.Migrate(db, migrationDir)
	if err != nil {
		log.Fatalln(err.Error())
	}

	userDaoImpl := dao.NewUserDaoImpl(db)
	leetcodeApiImpl := leetcode.NewLeetcodeApiImpl(leetcode.BASE_URL)
	userServiceImpl := service.NewUserServiceImpl(userDaoImpl, leetcodeApiImpl)

	s := gocron.NewScheduler(time.UTC)
	s.Every(10).Second().Do(func() {
		user, err := userServiceImpl.UpdateMostOutdatedUser()
		if err != nil {
			log.Printf("error updating user: %v\n", user.Username)
		}
	})
	s.StartAsync()

	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"https://leetcode.com"}

	r.Use(cors.New(config))
	r.GET("/ping", controller.GetPingController())
	r.GET("/user", controller.GetUserController(userServiceImpl))
	r.GET("/users", controller.GetUsersController(userServiceImpl))
	r.Run()
}
