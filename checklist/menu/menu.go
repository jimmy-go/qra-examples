// Package menu contains menu validation for users.
//
// MIT License
//
// Copyright (c) 2016 Angel Del Castillo
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
package menu

import (
	"bytes"
	"log"

	"github.com/jimmy-go/qra"
)

// Menu struct represents UI menu.
type Menu struct {
	Name  string
	Role  string
	Icon  string
	Link  string
	Badge int
}

var (
	// Menus and roles required.
	Menus = map[string]Menu{
		"admin": Menu{
			Name:  "Users",
			Role:  "admins",
			Icon:  "fa fa-users",
			Link:  "/users",
			Badge: 0,
		},
		"users": Menu{
			Name:  "TO-DOs",
			Role:  "users",
			Icon:  "fa fa-list",
			Link:  "/checklist",
			Badge: 0,
		},
	}
)

// UserMenus returns user available menus.
func UserMenus(userID string) []Menu {
	var list []Menu
	buf := bytes.NewBuffer([]byte{})
	for ii := range Menus {
		menu := Menus[ii]
		log.Printf("UserMenus : menu [%v]", menu)

		err := qra.Search(Ctx{userID}, buf, "read:"+menu.Role)
		if err != nil {
			log.Printf("UserMenus : don't menu [%v] menuRole [%s]", menu, menu.Role)
			continue
		}
		log.Printf("UserMenus : append menu [%v] menuRole [%s]", menu, menu.Role)
		list = append(list, menu)
	}

	log.Printf("UserMenus : buf [%v]", buf.String())
	return list
}

// Ctx satisfies Identity interface.
type Ctx struct {
	UserID string
}

// Me method.
func (c Ctx) Me() string {
	return c.UserID
}

// Session method.
func (c Ctx) Session(dst interface{}) error {
	// DO nothing
	return nil
}
