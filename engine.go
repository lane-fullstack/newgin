package newgin

import (
	"github.com/gin-gonic/gin"
	"github.com/lane-fullstack/newgin/ctx"
)

type HandlerFunc func(*ctx.Context)

type Engine struct {
	*gin.Engine
}

type RouterGroup struct {
	*gin.RouterGroup
}

func New(opts ...gin.OptionFunc) *Engine {
	engine := gin.New(opts...)
	return &Engine{
		Engine: engine,
	}
}

func Default(opts ...gin.OptionFunc) *Engine {
	engine := gin.Default(opts...)
	return &Engine{
		Engine: engine,
	}
}

func adapt(h HandlerFunc) gin.HandlerFunc {

	return func(c *gin.Context) {

		cc := ctx.New(c)

		h(cc)
	}
}

func (r *Engine) GET(path string, h HandlerFunc) {
	r.Engine.GET(path, adapt(h))
}

func (r *Engine) POST(path string, h HandlerFunc) {
	r.Engine.POST(path, adapt(h))
}

func (r *Engine) PUT(path string, h HandlerFunc) {
	r.Engine.PUT(path, adapt(h))
}

func (r *Engine) DELETE(path string, h HandlerFunc) {
	r.Engine.DELETE(path, adapt(h))
}

func (r *Engine) PATCH(path string, h HandlerFunc) {
	r.Engine.PATCH(path, adapt(h))
}

func (r *Engine) OPTIONS(path string, h HandlerFunc) {
	r.Engine.OPTIONS(path, adapt(h))
}

func (r *Engine) HEAD(path string, h HandlerFunc) {
	r.Engine.HEAD(path, adapt(h))
}

func (r *Engine) Any(path string, h HandlerFunc) {
	r.Engine.Any(path, adapt(h))
}

func (r *Engine) Handle(method string, path string, h HandlerFunc) {
	r.Engine.Handle(method, path, adapt(h))
}

func (r *Engine) Group(path string) *RouterGroup {

	g := r.Engine.Group(path)

	return &RouterGroup{
		RouterGroup: g,
	}
}

func (g *RouterGroup) GET(path string, h HandlerFunc) {
	g.RouterGroup.GET(path, adapt(h))
}

func (g *RouterGroup) POST(path string, h HandlerFunc) {
	g.RouterGroup.POST(path, adapt(h))
}

func (g *RouterGroup) PUT(path string, h HandlerFunc) {
	g.RouterGroup.PUT(path, adapt(h))
}

func (g *RouterGroup) DELETE(path string, h HandlerFunc) {
	g.RouterGroup.DELETE(path, adapt(h))
}

func (g *RouterGroup) PATCH(path string, h HandlerFunc) {
	g.RouterGroup.PATCH(path, adapt(h))
}

func (g *RouterGroup) OPTIONS(path string, h HandlerFunc) {
	g.RouterGroup.OPTIONS(path, adapt(h))
}

func (g *RouterGroup) HEAD(path string, h HandlerFunc) {
	g.RouterGroup.HEAD(path, adapt(h))
}

func (g *RouterGroup) Any(path string, h HandlerFunc) {
	g.RouterGroup.Any(path, adapt(h))
}

func (g *RouterGroup) Handle(method string, path string, h HandlerFunc) {
	g.RouterGroup.Handle(method, path, adapt(h))
}

func (g *RouterGroup) Group(path string) *RouterGroup {
	return &RouterGroup{
		RouterGroup: g.RouterGroup.Group(path),
	}
}
