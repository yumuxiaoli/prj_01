package main

import (
	"prj_01/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/go_im?charset=utf8mb4&parseTime=True&loc=Local"))
	// 	// 迁移 schema
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.UserBasic{})
}

// 	// Create
// 	user := &models.UserBasic{
// 		Name: "yumu",
// 	}
// 	db.Create(user)

// 	// Read
// 	fmt.Println(db.First(user, 1)) // 根据整型主键查找
// 	// db.First(user, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

// 	// Update - 将 product 的 price 更新为 200
// 	db.Model(user).Update("Password", "123456")
// 	// Update - 更新多个字段
// 	// db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
// 	// db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

// 	// Delete - 删除 product
// 	// db.Delete(&product, 1)
// }
