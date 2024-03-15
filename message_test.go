package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMessage(t *testing.T) {

	data := make(map[string]any)
	data["hi"] = "isgood"
	data["fsdf"] = 123

	m := &Message{
		Action: "go",
		Data:   data,
	}

	j, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(j))
}
