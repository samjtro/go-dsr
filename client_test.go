package dsr_test

import (
	"fmt"
	"testing"

	"github.com/samjtro/go-dsr"
)

func TestChatCompletion(t *testing.T) {
	c := dsr.NewChatClient()
	c.AddUserMessage("hello!")
	res, _ := c.GetNextChatCompletion()
	c.AddMessage(res.Choices[0].Message)
	c.AddUserMessage("can you help me with a math problem?")
	res, _ = c.GetNextChatCompletion()
	c.AddMessage(res.Choices[0].Message)
	fmt.Println(c.Messages)
}
