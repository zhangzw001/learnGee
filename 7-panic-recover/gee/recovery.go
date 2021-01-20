package gee

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"
)

// 如果不加recover web服务也不会宕机, 因为http的serve 实现了recover 的捕获
func Recovery() HandlerFunc {
	return func(c *Context) {
		//Recovery 的实现非常简单，使用 defer 挂载上错误恢复的函数，在这个函数中调用 recover()*，捕获 panic，并且将堆栈信息打印在日志中，向用户返回 *Internal Server Error。
		defer func() {
			if err := recover(); err != nil {
				msg := fmt.Sprintf("%s",err)
				log.Printf("%s\n\n",trace(msg))
				c.Fail(http.StatusInternalServerError,"Internal Server Error ")
			}
		}()
		// 执行下一个中间件
		c.Next()
	}
}


// trace() 函数是用来获取触发 panic 的堆栈信息
func trace(msg string) string {
	var pcs [32]uintptr
	//在 trace() 中，调用了 runtime.Callers(3, pcs[:])，Callers 用来返回调用栈的程序计数器,
	//第 0 个 Caller 是 Callers 本身，
	//第 1 个是上一层 trace，
	//第 2 个是再上一层的 defer func。
	//因此，为了日志简洁一点，我们跳过了前 3 个 Caller。
	n := runtime.Callers(2,pcs[:])	// skip first 3 caller

	var str strings.Builder
	str.WriteString(msg+"\nTraceBack:")
	counts := 0
	for _, pc := range pcs[:n] {
		//通过 runtime.FuncForPC(pc) 获取对应的函数
		fn := runtime.FuncForPC(pc)
		//fn.FileLine(pc) 获取到调用该函数的文件名和行号，打印在日志中。
		file, line := fn.FileLine(pc)
		fname := fn.Name()
		str.WriteString(fmt.Sprintf("\n#%d %s: \t%s:%d",counts,fname,file,line))
		counts ++

	}
	return str.String()
}
