package model

import (
	"github.com/gomodule/redigo/redis"
	"fmt"
)

// 存储图片id 到redis 数据库
func SaveImgCode(code, uuid string) error {
	// 1. 链接数据库
	conn, err := redis.Dial("tcp", "192.168.6.108:6379")
	if err != nil {
		fmt.Println("redis Dial err:", err)
		return err
	}
	defer conn.Close()

	// 2. 写数据库  --- 有效时间 5 分钟
	_, err = conn.Do("setex", uuid, 60*5, code)

	return err  // 不需要回复助手!
}
