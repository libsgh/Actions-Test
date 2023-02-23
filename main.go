package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	gorm_logrus "github.com/onrik/gorm-logrus"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func main() {
	r := gin.Default()
	database, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{
		Logger: gorm_logrus.New(),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(fmt.Sprintf("Got error when connect database, the error is '%v'", err))
	} else {
		log.Info("[Boot]Sqlite >> Database connection succeeded")
	}
	database.AutoMigrate(&User{})
	r.GET("/", func(c *gin.Context) {
		database.Create(User{"张三"})
		u := &User{}
		database.Find(u)
		c.String(200, "Hello, %s", u.Name)
	})
	r.Run()
}

type User struct {
	Name string `json:"name"`
}
