package main

import (
	"flag"
	"fmt"
	"github.com/lizongying/go-crawler/pkg/utils"
	"github.com/lizongying/go-crawler/static"
	"io/fs"
	"net/http"
)

func main() {
	uiPortPtr := flag.Int("ui-port", 8091, "-ui-port 8091")
	flag.Parse()

	fmt.Printf("ui at http://%s:%d http://%s:%d http://%s:%d\n", "localhost", *uiPortPtr, utils.LanIp(), *uiPortPtr, utils.InternetIp(), *uiPortPtr)

	files, _ := fs.Sub(static.Dist, "dist")
	http.HandleFunc("/", http.StripPrefix("/", http.FileServer(http.FS(files))).ServeHTTP)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", *uiPortPtr), nil); err != nil {
		panic(err)
	}
}
