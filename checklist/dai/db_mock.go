// Package dai contains database mock for example.
// You can have a real database connection here.
//
// Purpose of dai is simulate the business logic of your site.
package dai

import (
	"errors"
	"log"
	"sync"
)

var (
	// Db database connection.
	Db *DbMock
)

// Configure prepares mockup database.
func Configure() error {
	var err error
	Db = &DbMock{
		Data: map[string][]Item{
			"admin@mail.com": []Item{},
			"user1@mail.com": []Item{
				Item{
					Who:    "user1",
					Task:   "car whashing",
					Status: "started",
				},
				Item{
					Who:    "user1",
					Task:   "god walk",
					Status: "started",
				},
			},
			"user2@mail.com": []Item{},
		},
	}
	return err
}

// Close closes database connections.
func Close() {
}

// DbMock is a simulation of database.
type DbMock struct {
	Data map[string][]Item

	sync.RWMutex
}

// Item simulates ToDo item.
type Item struct {
	Who    string
	Task   string
	Status string
}

// List returns user checklist.
func (db *DbMock) List(userID string) ([]Item, error) {
	db.RLock()
	defer db.RUnlock()

	list, ok := db.Data[userID]
	if !ok {
		log.Printf("List : err : user [%s] len [%v]", userID, len(list))
		return []Item{}, errors.New("user not found")
	}
	log.Printf("List : user [%s] len [%v]", userID, len(list))
	return list, nil
}
