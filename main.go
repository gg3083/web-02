package main

import "web-02/route"

func main() {
	r := route.InitRoute()
	r.Run(":8090") // listen and serve on 0.0.0.0:808
}