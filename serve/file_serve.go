package serve

import (
	"fmt"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type FileResult struct {
	FileName string `json:"fileName"`
	IsDir    bool   `json:"isDir"`
}

func FileServe(root string) func(host string, port int) {
	app := gin.Default()
	abs, err := filepath.Abs(root)
	if err != nil {
		panic(err)
	}
	app.StaticFS("/f", gin.Dir(abs, true))
	return func(host string, port int) {
		app.Run(fmt.Sprint(host, ":", port))
	}
}
