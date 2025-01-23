package dsr_test

import (
	"testing"

	"github.com/samjtro/go-dsr"
)

func TestChatCompletion(t *testing.T) {
	c := dsr.NewChatClient()
	c.AddMessage("user", "hello!")
	_, _ = c.GetNextChatCompletion()
}
