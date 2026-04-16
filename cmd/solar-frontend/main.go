package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.LoadHTMLGlob("web/templates/*.html")
	r.Static("/static", "./web/static")

	r.GET("/", func(ctx *gin.Context) {
		now := time.Now()
		ctx.HTML(200, "index.html", gin.H{
			"DayName":  now.Weekday().String(),
			"FullDate": now.Format("02 January 2006"),
			"Time":     now.Format("15:04"),
		})
	})

	r.Run(":8080")
}
