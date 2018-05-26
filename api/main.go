package api

import (
	"log"

	_ "github.com/astaxie/beego"
	_ "gopkg.in/redis.v4"
	"fmt"
	"time"
	"os"
)

func main()  {

fmt.Println("[Server Starting]...")
	gin.SetMode(gin.ReleaseMode)
	router := routers.InitRouter()
	server := endless.NewServer("127.0.0.1:8080", router)
	server.ReadTimeout = 3 * time.Second
	server.WriteTimeout = 3 * time.Second

	server.BeforeBegin = func(add string) {
		log.Printf(add)
	}

	err := server.ListenAndServe()

	if err != nil {
	log.Println(err)
	}

	log.Println("Server stopped")
	os.Exit(0)
}

