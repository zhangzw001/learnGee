package gee

import (
	"strings"
)

// 采用 前缀树 方式来实现动态路由
type node struct {
	pattern string //待匹配路由, 例如  /p/:lang , 是 req.URL.Path
	part string // 路由中一部分, 例如 :lang, 是按照 / 分割的部分
	children []*node // 子节点, 例如 [ doc, tutorial, intro ]
	isWild bool // 是否精确匹配, part含有: 或 * 时为true
}



// 第一个匹配成功的节点, 用于插入 insert 方法
func (n *node ) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// 所有匹配成功的节点，用于查找 search 方法
//  例如 /test/v1/add
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	// 查询方法就是从trie树中查询, 如果查到了添加到列表, 返回查询到的列表
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
			// 这里加break是让查到的第一个就结束, 不希望路由匹配到多个结果
			//break
		}
	}

	return nodes
}

// 插入方法就是 判断 part路由是否已经存在路由节点中 如果是模糊匹配 直接返回
//   if child.part == part || child.isWild {
// 如果没有找到且是精确匹配, 那么就添加到子节点中
//   n.children = append(n.children, child)
//
//  pattern 是完整路由url ,parts 是将 pattern 按照 / 拆分的 每一部分 , 然后height 从0 递归对 parts 进行查询或插入
//
//  这里是一个逆向看的, 可以先去查询 parsePattern 方法
//  例如 /test/v1/add
func (n *node) insert(pattern string, parts []string, height int) {
	// 在node中生成一条 test(node) -> v2(node) -> add(node) 一条trie链表 之后
	// len(parts == height == 3 , 可以退出
	if len(parts) == height {
		n.pattern = pattern
		return
	}
	//
	part := parts[height]
	// 首先取 test 从node中查询, 如果查到了 那么目前查到路由 : /test
	// 那么继续递归insert, 取v1 去查询,
	// 如果查到了说明路由存在, 如果没有查到 那么新建这个node,添加到当前node的children, 目前查到路由: /test/v1,
	// 同样最后可以 在node中生成一条 test(node) -> v2(node) -> add(node) 一条trie链表
	child := n.matchChild(part)
	if child == nil {
		child = &node{
			part:part,
			isWild:part[0] == ':' || part[0] == '*',
		}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

// 从node中查询节点是否存在, 如果存在就返回
//  例如 /test/v1/add

func (n *node) search(parts []string, height int ) *node {
	if len(parts) == height || strings.HasPrefix(n.part,"*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}
	part := parts[height]
	// 这里查询同样是 先查询 test, 如果找到trie树的第一个节点 children = [ test(node) ]
	// 然后在出现v1, 根据trie结构, v1只能是在test(node) 这个节点的
	// 这里需要for循环 为什么不直接取 children[0] 呢?
	// 1. 有可能是因为v1 可能匹配到 /test/v1 也可能匹配到 /test/:version, 所以返回的是2个
	// 2. 如果让每次matchChildren查询返回一定是唯一呢? 那如果先有路由1:"/hello/geektutu", 在注册了路由2:"/hello/*name", 当请求"/hello/geektutu/a/b"的时候就会无法匹配路由2
	children := n.matchChildren(part)
	for _, child := range children {
		//log.Println("child:",child.pattern)
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}
