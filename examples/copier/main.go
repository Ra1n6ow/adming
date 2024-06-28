package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jinzhu/copier"
)

func (menu *Menu) ParentMenu(parentmenu string) {
	// if parentmenu == "" || parentmenu == "0" {
	// 	menu.ParentID = nil
	// } else {
	var v string
	if strings.Contains(parentmenu, "-") {
		parts := strings.SplitN(parentmenu, "-", 3) // 最多分隔2次，即最多3个子串
		v = parts[len(parts)-1]
	} else {
		v = parentmenu
	}
	iv, _ := strconv.Atoi(v)
	ivp := uint(iv)
	menu.ParentID = &ivp
}

type CreateMenuRequest struct {
	MenuName   string
	ParentMenu string
}

type Menu struct {
	MenuName string
	ParentID *uint
	Type     string
}

func main() {
	var (
		r = CreateMenuRequest{MenuName: "Joy", ParentMenu: "1-2-5"}
		m = Menu{}
	)
	err := copier.Copy(&m, &r)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(*m.ParentID)
	fmt.Println(m.MenuName)
}
