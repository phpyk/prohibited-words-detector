package Detector

import (
	"fmt"
	"testing"
)

func TestDetector_Lookup(t *testing.T) {
	dt := New()
	err := dt.Init("../banned_words.txt")

	if err != nil {
		t.Error(err.Error())
	}

	//fmt.Printf("%+v", dt)
	words := "XO"

	if dt.Lookup(words) {
		fmt.Println("find")
	}

}
