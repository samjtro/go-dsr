package dsr_test

import (
	"fmt"
	"testing"

	"github.com/samjtro/go-dsr"
)

func TestChatCompletion(t *testing.T) {
	c := dsr.NewChatClient()
	c.AddUserMessage("can you help me with a math problem?")
	res, _ := c.GetNextChatCompletion()
	fmt.Println(res)
}
