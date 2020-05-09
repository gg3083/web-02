package main

import "web-02/route"

// @title Golang Gin API
// @version 1.0
// @description An example of gin
// @termsOfService https://github.com/EDDYCJY/go-gin-example
// @license.name MIT
// @license.url https://Pay/blob/master/LICENSE
func main() {
	r := route.InitRoute()
	r.Run(":8090") // listen and serve on 0.0.0.0:808
}
