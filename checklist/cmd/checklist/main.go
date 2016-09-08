// Package main contains an example of QRA in action.
// CHECK LIST between multiple users.
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
package main

import (
	"flag"
	"log"
	"net/http"

	// import driver SQLite
	_ "github.com/mattn/go-sqlite3"

	// import driver PostgreSQL
	_ "github.com/lib/pq"

	"github.com/jimmy-go/qra-examples/checklist/auth"
	"github.com/jimmy-go/qra-examples/checklist/dai"
	"github.com/jimmy-go/qra-examples/checklist/list"
	"github.com/jimmy-go/qra-examples/checklist/session"
	"github.com/jimmy-go/qra-examples/checklist/users"
	"github.com/jimmy-go/qra/rawmanager"
	"gopkg.in/jimmy-go/srest.v0"
)

var (
	port      = flag.Int("port", 5050, "Listen port.")
	templates = flag.String("templates", "", "Templates dir.")
	static    = flag.String("static", "", "Static dir.")
)

func main() {
	flag.Parse()
	log.SetFlags(0)
	log.Printf("templates [%v]", *templates)
	log.Printf("static [%v]", *static)
	log.SetFlags(log.Lshortfile)

	// qra logic

	// register qra/rawmanager as qra.DefaultManager
	err := rawmanager.Connect()
	P(err)

	// business logic

	err = srest.LoadViews(*templates, srest.DefaultFuncMap)
	P(err)

	m := srest.New(nil)
	m.Debug(true)
	m.Get("/static/", srest.Static("/static/", *static))
	m.Get("/", http.HandlerFunc(users.Index), auth.MW)
	m.Get("/login", http.HandlerFunc(session.Index))
	m.Post("/login", http.HandlerFunc(session.Login))
	m.Get("/logout", http.HandlerFunc(session.Logout), auth.MW)
	m.Get("/users", http.HandlerFunc(users.Index), auth.MW)
	m.Get("/checklist", http.HandlerFunc(list.Index), auth.MW)
	<-m.Run(*port)
	dai.Close()
}

// P panics if error is present.
func P(err error) {
	if err != nil {
		panic(err)
	}
}
