package dsr_test

import (
	"fmt"
	"testing"

	"github.com/samjtro/go-dsr"
)

func TestChatCompletion(t *testing.T) {
	c := dsr.NewChatClient()
	c.AddMessage("user", "hello!", "")
	choices, _ := c.GetNextChatCompletion()
	fmt.Println(choices)
}
