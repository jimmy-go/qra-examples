// Package list contains TO-DO list by User.
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
package list

import (
	"log"
	"net/http"

	"github.com/jimmy-go/qra-examples/checklist/dai"
	"github.com/jimmy-go/qra-examples/checklist/menu"
	"github.com/jimmy-go/qra-examples/checklist/sessions"

	"gopkg.in/jimmy-go/srest.v0"
)

// Index endpoint /list GET
func Index(w http.ResponseWriter, r *http.Request) {
	v := map[string]interface{}{}
	userID, err := sessions.UserID(w, r)
	if err != nil {
		log.Printf("Index : cooksess : err [%s]", err)
	}
	log.Printf("Index : userID [%s]", userID)

	menus := menu.UserMenus(userID)
	log.Printf("Index : user menus [%v][%v]", len(menus), menus)
	v["menus"] = menus

	list, err := dai.Db.List(userID)
	if err != nil {
		log.Printf("Index : dai : err [%s]", err)
	}

	v["user_checklist"] = list

	err = srest.Render(w, "list/index.html", v)
	if err != nil {
		log.Printf("Index : Render : err [%s]", err)
	}
}
