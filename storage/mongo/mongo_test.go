package mongo

import (
	"fmt"
	"testing"
)

func MongoTest(t *testing.T) {
	var c chan int
	db := NewMongo(&MongoOptions{
		IP:   "127.0.0.1",
		Port: "27017",
	})

	fmt.Print(*db.Opts)

	select {
	case <-c:
		break
	}
}
