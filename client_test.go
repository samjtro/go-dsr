package dsr_test

import (
	"fmt"
	"testing"

	"github.com/samjtro/go-dsr"
)

func TestChatCompletion(t *testing.T) {
	c := dsr.NewChatClient()
	c.AddMessage("user", "hello!", "")
	res, _ := c.GetNextChatCompletion()
	fmt.Println(res.Choices[0].Message.Content)
}
