package main

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql" //"_" 代码不直接使用包, 底层链接要使用(导入后会在main前调用init)
	// _ "my-microservice1/web/test2" // 测试会调用init
	"github.com/jinzhu/gorm"
)

// 创建全局连接池句柄
var GlobalConn *gorm.DB

func main() {
	// 连接数据库--获取连接池句柄 格式: 用户名:密码@协议(IP:port)/数据库名
	// parseTime和loc用于指定时区
	conn, err := gorm.Open("mysql",
		"root:123456@tcp(127.0.0.1:3306)/test?parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("gorm.Open err:", err)
		return
	}
	// 连接池不用关
	//defer conn.Close()

	GlobalConn = conn

	// 初始数
	GlobalConn.DB().SetMaxIdleConns(10)
	// 最大数
	GlobalConn.DB().SetMaxOpenConns(100)

	// 不要复数表名
	GlobalConn.SingularTable(true)

	// 借助 gorm 创建数据库表(自动生成，不设置则默认表为复数类型，即会加s)=>打印错误信息
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

// 创建全局结构体，设属性
// mysql时间(date/datetime/timeStamp) => gorm默认只有timeStamp
type Student struct {
	Id    int
	Name  string `gorm:"size:100;default:'xiaoming'"`
	Age   int
	Class int       `gorm:"not null"`
	Join  time.Time `gorm:"type:timestamp"` // 可以用type设置mysql特有类型，例如type:datetime
}

// Student继承Model =>用于软删除
// type Student struct {
// 	gorm.Model
// 	Name string `gorm:"size:100;default:'xiaoming'"`
// 	Age  int
// }

// gorm自动维护 =>用于软删除
// type Model struct {
// 	ID        uint `gorm:"primary_key"`
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// 	DeletedAt *time.Time `sql:"index"`
// }

func InsertData() {
	// 先创建数据 --- 创建对象
	var stu Student
	stu.Name = "zhangsan"
	stu.Age = 100

	// 插入(创建)数据
	fmt.Println(GlobalConn.Create(&stu).Error)
}

func SearchData() {

	// GlobalConn.First(&stu)   查询 第一条全部信息(按主键排序)
	// GlobalConn.Last(&stu)  查最后一条信息
	// GlobalConn.Find(&stu)  查所有信息
	// GlobalConn.Select("name, age").First(&stu)   查询第一条 name 和 age
	// GlobalConn.Select("name, age").Find(&stu)   查询所有条记录的  name 和 age
	// GlobalConn.Select("name, age").Where("name = ?", "lisi").Find(&stu)  查询姓名为lisi的 name 和 age
	// GlobalConn.Select("name, age").Where("name = ?", "lisi").Where("age = ?", 22).Find(&stu)
	// GlobalConn.Select("name, age").Where("name = ? and age = ?", "lisi", 22).Find(&stu)
	// GlobalConn.Where("name = ?", "lisi").Select("name, age").Find(&stu)
	// GlobalConn.Select("name, age").Where("name = ?", "lisi").Find(&stu)

	var stu []Student
	GlobalConn.Unscoped().Find(&stu) // Unscoped可查询到软删除的数据
	fmt.Println(stu)
}

func UpdateData() {
	/*	var stu Student

		stu.Name = "wangwu"
		stu.Age = 77
		stu.Id = 4*/

	// 更新单个字段
	//fmt.Println(GlobalConn.Model(new(Student)).Where("name = ?", "zhaoliu").Update("name", "lisi").Error)

	// 更新多个字段
	fmt.Println(GlobalConn.Model(new(Student)).
		Where("age = ?", 77).
		Updates(map[string]interface{}{"name": "lisi", "age": 119}).Error)
}

func DeleteData() {
	// 软删除，delete中指定要删的数据所在的表
	// fmt.Println(GlobalConn.Where("name=?", "lisi").Delete(new(Student)).Error)

	// 物理删除
	fmt.Println(GlobalConn.Unscoped().Where("name = ?", "lisi").Delete(new(Student)).Error)
}
