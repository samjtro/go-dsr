package dsr

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/bytedance/sonic"
	"github.com/joho/godotenv"
)

const (
	BASE_URL = "https://api.deepseek.com"
)

type (
	Client struct {
		log *slog.Logger
		key string
	}
	ClientOptions struct {
		Key string
		Log *slog.Logger
	}
	ClientOption func(*ClientOptions)

	ChatClient struct {
		*Client
		Messages []*Message
	}

	Message struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	}
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func Initiate(opts ...ClientOption) *Client {
	o := ClientOptions{}
	for _, opt := range opts {
		opt(&o)
	}
	if o.Log == nil {
		o.Log = slog.New(slog.NewTextHandler(io.Discard, nil))
	}
	if o.Key == "" {
		o.Key = os.Getenv("DEEPSEEK_API_KEY")
	}
	return &Client{
		log: o.Log,
		key: o.Key,
	}
}

func NewChatClient(opts ...ClientOption) *ChatClient {
	m := []*Message{}
	return &ChatClient{
		Initiate(opts...),
		m,
	}
}

func (c *ChatClient) AddMessage(role, content string) {
	c.Messages = append(c.Messages, &Message{
		role,
		content,
	})
}

func (c *ChatClient) GetNextChatCompletion() (*Message, error) {
	m, err := sonic.MarshalString(c.Messages)
	if err != nil {
		return nil, err
	}
	fmt.Println(m)
	data := []byte(fmt.Sprintf(`{"model": "deepseek-reasoner", "messages": %s, "stream": false}`, m))
	r, err := http.NewRequest("POST", fmt.Sprintf(BASE_URL+"/chat/completions"), bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.key))
	r.Header.Add("Accept", "application/json")
	r.Header.Add("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	fmt.Println(resp)
	return nil, nil
}
