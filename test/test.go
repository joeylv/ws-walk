package main

import (
	"../models"
	"fmt"
)

func tysw() {
	//var i interface{} = "sssss"
	//switch i.(type){
	//case int:
	//	fmt.Println("it's int")
	//case string:
	//	fmt.Println("it's string")
	//default:
	//	fmt.Println("unknown type")
	//}
	var x interface{} = models.PreBook{}
	x = models.Member{}
	var list interface{}
	switch a := x.(type) {
	case models.PreBook:
		fmt.Println("PreBook", a.Search(nil))
	case models.Employee:
		fmt.Println("Employee", a.Search(0))
	case models.Member:
		fmt.Println("Member", a.Search(0))
		list = a.Search(0)
	default:
		fmt.Println("unknown type", a)

	}
	//list = list.([]models.Member)
	for i, j := range list.([]models.Member) {
		fmt.Println(i, j.Name)
	}
	//switch a := list.(type) {
	//default:
	//	fmt.Println(a.([]models.Member)[0].Name)
	//}
	//for i, j := range list {
	//	fmt.Println(i.j)
	//}
}

func main() {
	tysw()
	//fmt.Println(x.(models.PreBook).Search())
	//e1, ok := x.(models.Employee)
	//fmt.Println(e1, ok)
	//fmt.Println(e1.Search())
}
