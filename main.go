package main

import (
	"flag"
	"github.com/luisfsantos/thysis/daemon"
	"log"
	"net/http"
)

var assetsPath string

func processFlags() *daemon.Configuration {
	config := &daemon.Configuration{}

	flag.StringVar(
		&config.BindingAddress,
		"bind",
		"127.0.0.1:8181",
		"HTTP address and port to bind to.")

	flag.StringVar(
		&config.DBconfig.DBConnection,
		"postgres-connect",
		"host=localhost dbname=thysis_db user=thysis password=password sslmode=disable",
		"DBconfig Connection String")
	flag.StringVar(&assetsPath, "assets-path", "grayscale", "Path to assets dir")

	flag.Parse()
	return config
}

func setupHttpAssets(config *daemon.Configuration) {
	log.Printf("Assets served from %q.", assetsPath)
	config.UI.Assets = http.Dir(assetsPath)
}

func setupLogger() {
	//f, err := os.Create("./log.out")
	//if err != nil {
	//	panic(err)
	//}
	//defer f.Close()
	//w := bufio.NewWriter(f)

	log.SetFlags(log.LstdFlags | log.Llongfile)
	//log.SetOutput(w)
}

func main() {
	setupLogger()
	cfg := processFlags()

	setupHttpAssets(cfg)

	if err := daemon.Run(cfg); err != nil {
		log.Printf("Error in main(): %v", err)
	}
}
