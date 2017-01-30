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

	// import driver PostgreSQL
	_ "github.com/lib/pq"

	"gopkg.in/jimmy-go/qra.v0/pgmanager"
	srest "gopkg.in/jimmy-go/srest.v0"

	"github.com/gorilla/context"
	"github.com/jimmy-go/qra-examples/checklist/auth"
	"github.com/jimmy-go/qra-examples/checklist/dai"
	"github.com/jimmy-go/qra-examples/checklist/list"
	"github.com/jimmy-go/qra-examples/checklist/session"
	"github.com/jimmy-go/qra-examples/checklist/sessions"
	"github.com/jimmy-go/qra-examples/checklist/users"
)

var (
	port       = flag.Int("port", 5050, "Listen port.")
	templates  = flag.String("templates", "", "Templates dir.")
	static     = flag.String("static", "", "Static dir.")
	cookskey   = flag.String("cookies-key", "0Kq3cmiTZNbUPZcybgdc", "Gorilla sessions cookie secret.")
	connectURL = flag.String("db-url", "", "PostgreSQL connection url.")
)

func main() {
	flag.Parse()
	log.SetFlags(log.Lshortfile)

	// register qra/pgmanager as qra.DefaultManager
	if err := pgmanager.Connect("postgres", *connectURL); err != nil {
		log.Fatal(err)
	}

	// business logic
	if err := dai.Configure(); err != nil {
		log.Fatal(err)
	}

	if err := srest.LoadViews(*templates, srest.DefaultFuncMap); err != nil {
		log.Fatal(err)
	}

	if err := sessions.Configure(*cookskey); err != nil {
		log.Fatal(err)
	}

	m := srest.New(nil)
	m.Debug(true)
	m.Get("/static/", srest.Static("/static/", *static))
	m.Get("/", http.HandlerFunc(users.Index), auth.Handler, context.ClearHandler)
	m.Get("/login", http.HandlerFunc(session.Index), context.ClearHandler)
	m.Post("/login", http.HandlerFunc(session.Login), context.ClearHandler)
	m.Get("/logout", http.HandlerFunc(session.Logout), auth.Handler, context.ClearHandler)
	m.Get("/users", http.HandlerFunc(users.Index), auth.Handler, context.ClearHandler)
	m.Get("/checklist", http.HandlerFunc(list.Index), auth.Handler, context.ClearHandler)
	<-m.Run(*port)
	dai.Close()
}
