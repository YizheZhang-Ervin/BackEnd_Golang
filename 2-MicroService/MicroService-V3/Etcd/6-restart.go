package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/jpillora/overseer"
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
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		dbUser := cfg.Section("db").Key("db_user").Value()
		dbPass := cfg.Section("db").Key("db_pass").Value()
		writer.Write([]byte("<h1>" + dbUser + "</h1>"))
		writer.Write([]byte("<h1>" + dbPass + "</h1>"))
	})
	mux.HandleFunc("/reload", func(writer http.ResponseWriter, request *http.Request) {
		newCfg, _ := ini.Load("my.ini")
		cfg = newCfg
	})

	server := &http.Server{
		// Addr:":"+strconv.Itoa(*port),
		Handler: mux,
	}
	prog := func(state overseer.State) { //state这个参数是官方的不用改
		server.Serve(state.Listener) //使用overseer去启动服务
	}

	errChan := make(chan error)
	go (func() {
		overseer.Run(overseer.Config{
			Program:          prog,
			TerminateTimeout: time.Second * 2, //如果配置更改了需要等待当前请求全部结束结束再重启，加上这个最多等待2秒
			Address:          ":" + strconv.Itoa(*port),
		})
	})()

	//监听信号
	go (func() {
		sig_c := make(chan os.Signal)
		signal.Notify(sig_c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-sig_c)
	})()
	//监控配置文件变化
	go (func() {
		fileMd5, err := getFileMD5("my.ini")
		if err != nil {
			log.Println(err)
			return
		}
		for {
			newMd5, err := getFileMD5("my.ini")
			if err != nil {
				log.Println(err)
				break
			}
			if strings.Compare(newMd5, fileMd5) != 0 {
				fileMd5 = newMd5
				fmt.Println("文件发生了变化")
				overseer.Restart() //使用overseer平滑重启服务
			}
			time.Sleep(time.Second * 2)
		}
	})()

	getErr := <-errChan

	log.Println(getErr)
}
func getFileMD5(filePath string) (string, error) {
	file, err := os.Open(filePath)
	hash := md5.New()
	if _, err = io.Copy(hash, file); err != nil {
		return "", err
	}
	hashInBytes := hash.Sum(nil)[:16]
	return hex.EncodeToString(hashInBytes), nil
}
