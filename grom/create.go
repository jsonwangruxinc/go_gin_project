package grom

import (
	"fmt"
	"time"
)

var teacherTemp = Techer{
	Name:         "wang",
	Email:        "wangruxin996@gmail.com",
	Age:          30,
	WorkingYears: 10,
	Birthday:     time.Now().Unix(),
	StuNumber: struct {
		String string
		Valid  bool
	}{String: "10", Valid: true},
	Roles: []string{"普通用户", "讲师"},
	JobInfo: Job{
		Title:    "讲师",
		Location: "北京市",
	},
	JobInfo2: Job{
		Title:    "讲师",
		Location: "北京市",
	},
}

// 创建记录
func CreateRecord() {
	t := teacherTemp
	res := DB.Create(&t)
	fmt.Println(res.RowsAffected, res.Error, t)
}

// 创建记录并为指定字段分配值
func SelcetCreate() {
	t := teacherTemp
	res := DB.Select("Name", "Age").Create(&t)
	fmt.Println(res.RowsAffected, res.Error, t)
}

// 创建一条记录并忽略传递给省略的字段的值
func OmitCreate() {
	t := teacherTemp
	res := DB.Omit("Name", "Age").Create(&t)
	fmt.Println(res.RowsAffected, res.Error, t)
}
func BatchCreate() {
	var teachers = []Techer{{Name: "wang", Age: 40}, {Name: "ru", Age: 50}, {Name: "xin", Age: 50}}
	DB.Create(teachers)
	for _, t := range teachers {
		fmt.Println(t.ID)
	}
}
