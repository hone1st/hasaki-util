package int

import (
	"fmt"
	"log"
	"testing"
)

func TestBcAdd(t *testing.T) {
	result, err := BcAdd("1231.2313", "2313.42442")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
func TestBcSub(t *testing.T) {
	result, err := BcSub("1231.2313", "2313.42442", 2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
func TestBcMul(t *testing.T) {
	result, err := BcMul("1231.2313", "2313.42442", 2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
func TestBcDiv(t *testing.T) {
	result, err := BcDiv("1231.2313", "2313.42442", 2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
