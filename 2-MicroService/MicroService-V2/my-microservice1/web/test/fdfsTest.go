package main

import (
	"fmt"

	"github.com/tedcy/fdfs_client"
)

func main() {
	clt, err := fdfs_client.NewClientWithConfig("/etc/fdfs/client.conf")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 存储凭证：组名/M00/00/00/XXXXX
	resp, err := clt.UploadByFilename("xx.jpg")
	fmt.Println(resp, err)
}
