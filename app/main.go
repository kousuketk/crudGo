package main

import (
	"github.com/kousuketk/crudGo/app/controllers"
)

func main() {
	controllers.Router().Run()
}
