/*
	Author: Cleiton Viana
*/

package main

import (
	"fmt"

	"github.com/cleitonSilvaViana/social-go/api/router"
	"github.com/cleitonSilvaViana/social-go/config"
)

func main() {
	config.GetEnv()

	fmt.Println("http://localhost" + config.API_PORT)

	router.InitRouter(config.API_PORT)
}
