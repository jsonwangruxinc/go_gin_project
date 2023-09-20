package grom

import (
	"fmt"
)

func GetOnce() {
	t := Techer{}
	res := DB.First(&t)
	fmt.Println(res.RowsAffected, res.Error, t)

	t = Techer{}
	res = DB.Take(&t)
	fmt.Println(res.RowsAffected, res.Error, t)

	t = Techer{}
	res = DB.Last(&t)
	fmt.Println(res.RowsAffected, res.Error, t)

	result := map[string]interface{}{}
	res = DB.Model(&Techer{}).Omit("Birthday", "roles", "job_info2").First(&result)
	fmt.Println(res.RowsAffected, res.Error, result)

	result = map[string]interface{}{}
	res = DB.Model("techers").Omit("Birthday", "roles", "job_info2").Take(&result)
	fmt.Println(res.RowsAffected, res.Error, result)

	result = map[string]interface{}{}
	res = DB.Take("tachers")
}
func GetByPrimarKey() {
	t := Techer{}
	res := DB.First(&t, 2)
	fmt.Println(res.RowsAffected, res.Error, t)
	var t2 = Techer{}
	res1 := DB.Model(t2).Where("id = ?", 6).First(&t2)
	if res1.Error != nil {
		fmt.Println("Error:", res1.Error)
		return
	}
	fmt.Println("RowsAffected:", res1.RowsAffected)
	fmt.Println("Techer:", t2)
	var techers []Techer
	res = DB.Find(&techers, []int{1, 2, 3, 4, 5, 6})
	fmt.Println(res.RowsAffected, res.Error, techers)
}
func GetByConnt() {
	t := Techer{}
	res := DB.Where("techer_name IN ?", []string{"wang", "ru"}).Find(&t)
	if res.Error != nil {
		fmt.Println("resError", res.Error)
		return
	}
	fmt.Println("RowsAffected:", res.RowsAffected)
	fmt.Println("Techer:", t)
}
