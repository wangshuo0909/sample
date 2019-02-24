package main

import (
	"log"
	"os"

	_ "github.com/wangshuo0909/sample/matchers"
	"github.com/wangshuo0909/sample/search"
)

func init() {
	log.SetOutPut(os.Stdout)
}
func main() {
	search.Run("president")
}
