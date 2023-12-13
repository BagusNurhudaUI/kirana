package database

import (
	"fmt"
	"kirana/model"
	"sync"
)

var (
	db []*model.User
	mu sync.Mutex
)

// Connect with database
func Connect() error {
	db = make([]*model.User, 0)
	fmt.Println("Connected with Database")
	return nil
}

func Insert(user *model.User) {
	mu.Lock()
	db = append(db, user)
	mu.Unlock()
}

func Get() (res []*model.User, err error) {
	// return nil, errors.New("db not found")
	return db, nil
}
