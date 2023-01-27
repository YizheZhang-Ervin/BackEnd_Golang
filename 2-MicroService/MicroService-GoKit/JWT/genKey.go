package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

func GenRSAPubAndPri(bits int, filepath string) error {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	priBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}

	err = ioutil.WriteFile(filepath+"/private.pem", pem.EncodeToMemory(priBlock), 0644)
	if err != nil {
		return err
	}
	fmt.Println("=======私钥文件创建成功========")
	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	publicBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}

	err = ioutil.WriteFile(filepath+"/public.pem", pem.EncodeToMemory(publicBlock), 0644)
	if err != nil {
		return err
	}
	fmt.Println("=======公钥文件创建成功=========")

	return nil
}

// func main() {
// 	err := GenRSAPubAndPri(1024, "./pem") //1024是长度，长度越长安全性越高，但是性能也就越差
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	//执行完生成公钥和私钥，公钥给别人私钥给自己
// }
