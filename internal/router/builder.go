package router

import "github.com/gin-gonic/gin"

type Builder struct {
	engine *gin.Engine
}

func NewBuilder() *Builder {
	return &Builder{
		engine: gin.Default(),
	}
}

func (b *Builder) WithMiddleware(middleware ...gin.HandlerFunc) *Builder {
	b.engine.Use(middleware...)
	return b
}

func (b *Builder) Build() *gin.Engine {
	return b.engine
}
