package main_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	bytes, _ := json.Marshal(time.Now())
	fmt.Println(string(bytes))
}
