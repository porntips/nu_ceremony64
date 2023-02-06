package main

import (
	"fmt"
	"log"
	"net/http"
	 "nu_ceremony/connected"
	"nu_ceremony/controller"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
	socketio "github.com/googollee/go-socket.io"

)

func GinMiddleware(allowOrigin string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", allowOrigin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, X-CSRF-Token, Token, session, Origin, Host, Connection, Accept-Encoding, Accept-Language, X-Requested-With")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Request.Header.Del("Origin")

		c.Next()
	}
}

func main() {
	connected.InitDB()

	router := gin.New()
	router.Use(cors.Default())

	server := socketio.NewServer(nil)

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		log.Println("connected:", s.ID())

		s.Join("ceremonyg")
		
		return nil
	})

	server.OnEvent("/", "ceremony", func(s socketio.Conn, group string) {
		g, _ := strconv.Atoi(group)
		res := controller.GetAllCeremony(&gin.Context{}, g)
		log.Println("ceremony : ", group)
		server.BroadcastToRoom("/", "ceremonyg", "graduate", res)
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		log.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Println("closed", reason)
	})

	go func() {
		if err := server.Serve(); err != nil {
			log.Fatalf("socketio listen error: %s\n", err)
		}
	}()
	defer server.Close()

	router.POST("/ceremony", controller.AddDefaultDb)
	router.GET("/ceremonyall", controller.GetAllGrad)
	router.GET("/ceremony/:group", func(c *gin.Context) {
		group := c.Param("group")
		g, _ := strconv.Atoi(group)
		res := controller.GetAllCeremony(c, g)

		c.JSON(http.StatusOK, res)
	})
	router.PUT("/ceremony/:studentcode/:ceremony", controller.RunningCeremony)

	// router.Use(GinMiddleware("http://localhost/index.html"))
	router.Use(GinMiddleware(fmt.Sprintf("http://%s:4200", os.Getenv("APP_HOST"))))

	router.GET("/socket.io/*any", gin.WrapH(server))
	router.POST("/socket.io/*any", gin.WrapH(server))
	router.StaticFS("/public", http.Dir("./"))

	if err := router.Run(":8080"); err != nil {
		log.Fatal("failed run app: ", err)
	}
}
