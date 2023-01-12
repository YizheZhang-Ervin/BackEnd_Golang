package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql" //"_" 代码不直接使用包, 底层链接要使用!
	"fmt"
	"time"
)

// 创建全局连接池句柄
var GlobalConn *gorm.DB

func main() {
	// 链接数据库--获取连接池句柄 格式: 用户名:密码@协议(IP:port)/数据库名
	conn, err := gorm.Open("mysql",
		"root:123456@tcp(127.0.0.1:3306)/test?parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("gorm.Open err:", err)
		return
	}
	//defer conn.Close()

	GlobalConn = conn;

	// 初始数
	GlobalConn.DB().SetMaxIdleConns(10)
	// 最大数
	GlobalConn.DB().SetMaxOpenConns(100)

	// 不要复数表名
	GlobalConn.SingularTable(true)

	// 借助 gorm 创建数据库表.
	fmt.Println(GlobalConn.AutoMigrate(new(Student)).Error)

	// 插入数据
	//InsertData()

	// 查询数据
	//SearchData()

	// 更新数据
	// UpdateData()

	// 删除数据
	//DeleteData()
}

// 创建全局结构体
type Student struct {
	Id    int
	Name  string    `gorm:"size:100;default:'xiaoming'"`
	Age   int
	Class int       `gorm:"not null"`
	Join  time.Time `gorm:"type:timestamp"`
}

func InsertData() {
	// 先创建数据 --- 创建对象
	var stu Student
	stu.Name = "zhangsan"
	stu.Age = 100

	// 插入(创建)数据
	fmt.Println(GlobalConn.Create(&stu).Error)
}

func SearchData() {

	// GlobalConn.First(&stu)   查询 第一条全部信息.
	// GlobalConn.Select("name, age").First(&stu)   查询第一条 name 和 age
	// GlobalConn.Select("name, age").Find(&stu)   查询所有条记录的  name 和 age
	// GlobalConn.Select("name, age").Where("name = ?", "lisi").Find(&stu)  查询姓名为lisi的 name 和 age
	// GlobalConn.Select("name, age").Where("name = ?", "lisi").Where("age = ?", 22).Find(&stu)
	//GlobalConn.Select("name, age").Where("name = ? and age = ?", "lisi", 22).Find(&stu)
	//GlobalConn.Where("name = ?", "lisi").Select("name, age").Find(&stu)
	//GlobalConn.Select("name, age").Where("name = ?", "lisi").Find(&stu)

	var stu []Student
	GlobalConn.Unscoped().Find(&stu)
	fmt.Println(stu)
}

func UpdateData() {
	/*	var stu Student

		stu.Name = "wangwu"
		stu.Age = 77
		stu.Id = 4*/

	/*	fmt.Println(GlobalConn.Model(new(Student)).Where("name = ?", "zhaoliu").
			Update("name", "lisi").Error)*/

	fmt.Println(GlobalConn.Model(new(Student)).Where("age = ?", 77).
		Updates(map[string]interface{}{"name": "lisi", "age": 119}).Error)
}

func DeleteData() {
	fmt.Println(GlobalConn.Unscoped().Where("name = ?", "lisi").Delete(new(Student)).Error)
}
