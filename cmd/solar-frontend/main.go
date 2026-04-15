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

		dayName := now.Weekday().String()
		date := now.Format("02 january 2006")
		timerStr := now.Format("15:04:05")

		ctx.HTML(200, "base.html", gin.H{
			"DayName":  dayName,
			"FullDate": date,
			"Time":     timerStr,
		})
	})

	r.Run(":8080")
}
