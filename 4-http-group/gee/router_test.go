package gee

import (
	"reflect"
	"testing"
)

func TestParsePattern(t *testing.T) {
	tests := []struct {
		patterns string
		parts    []string
		ok bool
	}{
		{"/test/v1/add", []string{"test", "v1", "add"},true },
		{"/test/*name/add", []string{"test", "*name"},true},
		{"/test/:name", []string{"test", ":name"},true },
		{"/test/:name/abc", []string{"test", ":name", "abc"},true},
		{"/test/v1/add", []string{"testv1", "add"},false },
		{"/test/*name/add", []string{"test", "*name","add"},false},
		{"/test/:name", []string{"test", ":","name"},false },
		{"/test/:name/abc", []string{"test", ":name"},false},
	}

	m := make(map[string][]string)

	for _, test := range tests {
		m[test.patterns] = test.parts
		p := parsePattern(test.patterns)

		if test.ok {
			if !reflect.DeepEqual(p, m[test.patterns]) {
				t.Errorf("expect %v, but actual %v\n", m[test.patterns], p)
			}
		}else {
			if reflect.DeepEqual(p, m[test.patterns]) {
				t.Errorf("not expect %v, but actual %v\n", m[test.patterns], p)
			}
		}


	}

}


func TestRouterGroup_GET(t *testing.T) {
	tests := []struct {
		method  string
		pattern string
		parts   []string
		ok      bool
	}{
		{"GET", "/test1/v1/add", []string{"/test1/v1/add"}, true},
		{"GET", "/test2/*name", []string{"/test2/name1", "/test2/name1/name2"}, true},
		{"GET", "/test3/:file/add", []string{"/test3/file123/add", "/test3/file_!@#$/add", "/test3/123431/add"}, true},
		{"GET", "/test4/:file/add", []string{"/test4/file123", "/test3/file_!@#$/add123", "/test3/123431/123_@#"}, false},
		// 这两条 应该怎么走呢?
		{"GET", "/hello/*name", []string{"/hello/geektutu/a/b","hello/geektutu"}, true},
		{"GET", "/hello/geektutu", []string{"hello/geektutu"}, true},

	}

	r := newRouter()
	for _, test := range tests {
		r.addRoute(test.method, test.pattern, nil)
	}

	for _, test := range tests {
		for _, part := range test.parts {
			if test.ok {
				node, _ := r.getRoute(test.method, part)
				if node == nil {
					t.Errorf("url Path %s, expected : %s, but actual : %s\n", part, test.pattern, "")
					continue
				}
				if node.pattern != test.pattern && test.ok {
					t.Errorf("url Path %s, expected : %s, but actual : %s\n", part, test.pattern, node.pattern)
				}
			}else {
				node, _ := r.getRoute(test.method, part)
				if node == nil {
					continue
				}
				if node.pattern == test.pattern {
					t.Errorf("url Path %s, not expected : %s, but actual : %s\n", part, test.pattern, node.pattern)
				}
			}

		}
	}

}
