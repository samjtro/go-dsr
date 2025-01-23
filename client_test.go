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
	c.AddSystemMessage(res.Choices[0].Message.Content, res.Choices[0].Message.ReasoningContent)
	fmt.Println(c.Messages)
}
