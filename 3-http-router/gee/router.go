package gee

import (
	"net/http"
	"strings"
)

// 这里包装router
type router struct {
	roots map[string]*node
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		handlers: make(map[string]HandlerFunc),
		roots: make(map[string]*node ),
	}
}

// Only one * is allowed
// 将 /test/v1/add -> [test v1 add]
// /test/*name/add -> [test *name]
func parsePattern(pattern string ) []string {
	vs := strings.Split(pattern,"/")
	parts := make([]string,0)
	for _ , item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

//
func (r *router) addRoute(method string, pattern string, handler HandlerFunc ) {
	parts := parsePattern(pattern)
	key := method+"-"+pattern
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern,parts,0)
	r.handlers[key] = handler
}


func (r *router) handle(c *Context) {

	n, params := r.getRoute(c.Method,c.Path)
	if n != nil  {
		c.Params = params
		key := c.Method + "-" + n.pattern
		r.handlers[key](c)
	}else {
		c.String(http.StatusNotFound,"404 not found : %s\n",c.Path)
	}
}

func (r *router) getRoute(method string , path string ) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	root, ok := r.roots[method]

	if !ok {
		return nil, nil
	}
	n := root.search(searchParts, 0)
	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}

	return nil, nil
}

// 这里GET方法写在 router 里面
// engine 是包含 router
func (r *router) GET(pattern string, handler HandlerFunc) {
	r.addRoute(http.MethodGet,pattern,handler)
}

func (r *router) POST(pattern string, handler HandlerFunc) {
	r.addRoute(http.MethodPost,pattern,handler)
}
