package gee

import "log"

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

// Use is defined to add middleware to the group
func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
	// 只需要将这些需要执行的 中间件方法 加入到列表即可
	group.handlers = append(group.handlers,middlewares...)
}
