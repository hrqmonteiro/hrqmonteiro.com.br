package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"hrqmonteiro.com.br/templates/pages"
)

func IndexHandler(c *gin.Context) {
	component := pages.Index()
	c.Header("Content-Type", "text/html")
	component.Render(context.Background(), c.Writer)
}
