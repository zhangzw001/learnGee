package main

import (
	"github.com/zhangzw001/learnGee/7-panic-recover/routers"
	"log"
	"math"
)

//https://geektutu.com/post/gee-day2.html
func main() {
	log.Println(math.MaxInt8 / 2)

	e := routers.InitRouter()
	e.Run("localhost:8000")
	//curl -XPOST "localhost:8000/admin/login" -d'username=admin&password=123'
	//{"password":"123","username":"admin"}
	//curl "localhost:8000/admin/logout?username=admin"
	//{"username":"admin"}
}
