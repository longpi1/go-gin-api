package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name     string
	Age      int8
	Birthday time.Time
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/go_gin_api?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//add
/*	user1 := User{Name: "lp1", Age: 18, Birthday: time.Now()}
	//db.AutoMigrate(&User{})
	_ = db.Create(&user1)*/
	//search
	user := new(User)
	//db.Debug().First(user)
	db.Debug().Last(user)
	db.Debug().Find(user)
	// 随机获取一条记录
	db.Take(&user)
	db.Debug().Take(user)
	fmt.Println(user)
	// 有效，因为通过 `db.Model()` 指定了 model
	result := map[string]interface{}{}
	db.Model(&User{}).First(&result)
	fmt.Println(result)

	// 无效
	result1 := map[string]interface{}{}
	db.Table("users").Debug().First(&result1)
	fmt.Println(result1)
	// 配合 Take 有效
	result2 := map[string]interface{}{}
	db.Table("users").Debug().Take(&result2)
	fmt.Println(result2)

	db.Debug().First(&user, 10)
	// SELECT * FROM users WHERE id = 10;

	db.First(&user, "10")
	// SELECT * FROM users WHERE id = 10;

	type Student struct {
		ID     uint
		Name   string
		Age    int
		Gender string
		// 假设后面还有几百个字段...
	}

	type APIUser struct {
		ID   uint
		Name string
	}

	// 查询时会自动选择 `id`, `name` 字段
	db.Debug().Model(&Student{}).Limit(10).Find(&APIUser{})
	// SELECT `id`, `name` FROM `users` LIMIT 10

	db.Debug().Model(&User{}).Where("name = ?", "lp").First(user)
	fmt.Println(user)

	db.Debug().Model(&User{}).Not("name = ?", "lp").First(user)
	fmt.Println(user)
	db.Debug().Delete(user)
}
