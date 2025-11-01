package main

import (
	"embed"
	"net/http"

	"github.com/a-h/templ/examples/integration-gin/gintemplrenderer"
	"github.com/gin-gonic/gin"
	"github.com/namnd/go-ssr/ui"
)

//go:embed assets/*
var f embed.FS

func main() {
	r := gin.Default()

	r.StaticFS("/public", http.FS(f))

	r.GET("/", Home)

	r.Run(":8080")
}

func Home(c *gin.Context) {
	p := gintemplrenderer.New(
		c.Request.Context(),
		http.StatusOK,
		ui.Home(),
	)

	c.Render(http.StatusOK, p)

}
