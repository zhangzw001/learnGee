package gee

import (
	"log"
	"net/http"
	"path"
)

type RouterGroup struct {
	prefix      string
	handlers []HandlerFunc
	parent      *RouterGroup
	engine      *Engine // 这里是为了 group直接访问 router
}

// Group is defined to create a new RouterGroup
// remember all groups share the same Engine instance
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix:      group.prefix + prefix,
		parent:      group,
		engine:      engine,
	}
	engine.groups = append(engine.groups,newGroup)
	return newGroup
}

//可以仔细观察下addRoute函数，调用了group.engine.router.addRoute来实现了路由的映射。
//由于Engine从某种意义上继承了RouterGroup的所有属性和方法，
//因为 (*Engine).engine 是指向自己的。这样实现，我们既可以像原来一样添加路由，也可以通过分组添加路由。
func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.router.addRoute(method, pattern, handler)
}

// GET defines the method to add GET request
func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST request
func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

func (group *RouterGroup) Header(pattern string, handler HandlerFunc) {
	group.addRoute(http.MethodHead, pattern, handler)
}

func (group *RouterGroup) Any(pattern string, handler HandlerFunc) {
	group.addRoute(http.MethodHead, pattern, handler)
	group.addRoute(http.MethodGet, pattern, handler)
	group.addRoute(http.MethodPost, pattern, handler)
}

// Use is defined to add middleware to the group
func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
	// 只需要将这些需要执行的 中间件方法 加入到列表即可
	group.handlers = append(group.handlers,middlewares...)
}


// create static handler
func (group *RouterGroup) createStaticHandler(relativePath string, fs http.FileSystem) HandlerFunc {
	absolutePath := path.Join(group.prefix, relativePath)
	//log.Printf("RouterGroup.createStaticHandler absolutePath : %s\n",absolutePath)
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))
	return func(c *Context) {
		file := c.Param("filepath")
		if _, err := fs.Open(file); err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		fileServer.ServeHTTP(c.Writer,c.Req)

	}
}

// save static files
func (group *RouterGroup) Static(relativePath string, root string) {
	// relativePath  = /assets
	handler := group.createStaticHandler(relativePath,http.Dir(root))
	// 设置一个通配符的路由规则
	urlPattern := path.Join(relativePath, "/*filepath")
	log.Printf("RouterGroup.Static urlPattern: %s\n",urlPattern)
	// 根据通配符的路由规则 绑定 handler 方法
	group.Any(urlPattern,handler)
}
