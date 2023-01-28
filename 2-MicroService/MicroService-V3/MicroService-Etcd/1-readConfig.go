package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"

	"gopkg.in/ini.v1"
)

func main() {
	port := flag.Int("p", 8080, "服务端口")
	flag.Parse()
	if *port == 0 {
		log.Fatal("请指定端口")
	}
	cfg, err := ini.Load("my.ini")
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		dbUser := cfg.Section("db").Key("db_user").Value()
		dbPass := cfg.Section("db").Key("db_pass").Value()
		writer.Write([]byte("<h1>" + dbUser + "</h1>"))
		writer.Write([]byte("<h1>" + dbPass + "</h1>"))
	})
	http.ListenAndServe(":"+strconv.Itoa(*port), nil)
}
