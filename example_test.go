package tlru_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/heucuva/go-tlru"
)

func TestExample(t *testing.T) {
	m := tlru.TLRU{}

	// add "hello" -> "world"
	m.SetString("hello", "world")
	// wait a second
	time.Sleep(1 * time.Second)

	// add "hello2" -> "world2"
	m.SetString("hello2", "world2")
	// wait a second
	time.Sleep(1 * time.Second)

	// add "hello1" -> "world1"
	m.SetString("hello1", "world1")
	// update "hello1" -> "world13"
	m.SetString("hello1", "world13")
	// wait a second
	time.Sleep(1 * time.Second)

	// expire anything older than 2 seconds
	// this should expire "hello" and "hello2"
	m.Expire(2 * time.Second)

	// iterate over the entries (in age order)
	// we should only see "hello1" with a value of "world13" and an age >= 1s
	m.Iterate(func(key, value interface{}, age time.Duration) bool {
		fmt.Printf("key[%v] value[%v] age[%v]\n", key, value, age)
		return true
	})
}

func TestExample2(t *testing.T) {
	m := tlru.TLRU{}

	// add "hello" -> "world"
	m.SetString("hello", "world")
	// wait a second
	time.Sleep(1 * time.Second)

	// add "hello2" -> "world2"
	m.SetString("hello2", "world2")
	// wait a second
	time.Sleep(1 * time.Second)

	// add "hello1" -> "world1"
	m.SetString("hello1", "world1")
	// update "hello1" -> "world13"
	m.SetString("hello1", "world13")
	// wait a second
	time.Sleep(1 * time.Second)

	// expire anything older than 3 seconds
	// this should expire "hello"
	m.Expire(3 * time.Second)

	// iterate over the entries (in age order)
	// we should see:
	// - "hello2" with a value of "world2" and an age >= 2s
	// - "hello1" with a value of "world13" and an age >= 1s
	m.Iterate(func(key, value interface{}, age time.Duration) bool {
		fmt.Printf("key[%v] value[%v] age[%v]\n", key, value, age)
		return true
	})
}
