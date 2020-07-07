package main

import (
	"fmt"
	_ "github.com/moonrailgun/PUBGo/server/config/db"
	_ "github.com/moonrailgun/PUBGo/server/config/db/migration"
	"github.com/moonrailgun/PUBGo/server/routes"
)

func main() {
	r := routes.SetupRouter()
	fmt.Println("Service is listening: 9091")
	if err := r.Run(":9091"); err != nil {
		fmt.Println("Service start error", err)
	}
}
