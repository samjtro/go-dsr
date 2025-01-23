package dsr

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"
	"strings"

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
		Messages []Message
	}

	Response struct {
		ID      string   `json:"id"`
		Object  string   `json:"object"`
		Created int      `json:"created"`
		Model   string   `json:"model"`
		Choices []Choice `json:"choices"`
		Usage   struct {
			PromptTokens        int `json:"prompt_tokens"`
			CompletionTokens    int `json:"completion_tokens"`
			TotalTokens         int `json:"total_tokens"`
			PromptTokensDetails struct {
				CachedTokens int `json:"cached_tokens"`
			} `json:"prompt_tokens_details"`
			CompletionTokensDetails struct {
				ReasoningTokens int `json:"reasoning_tokens"`
			} `json:"completion_tokens_details"`
			PromptCacheHitTokens  int `json:"prompt_cache_hit_tokens"`
			PromptCacheMissTokens int `json:"prompt_cache_miss_tokens"`
		} `json:"usage"`
		SystemFingerprint string `json:"system_fingerprint"`
	}

	Message struct {
		Role             string `json:"role"`
		Content          string `json:"content"`
		ReasoningContent string `json:"reasoning_content,omitempty"`
	}

	Choice struct {
		Index        int     `json:"index"`
		Message      Message `json:"message"`
		Logprobs     string  `json:"logprobs"`
		FinishReason string  `json:"finish_reason"`
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
		o.Log,
		o.Key,
	}
}

func NewChatClient(opts ...ClientOption) *ChatClient {
	m := []Message{}
	return &ChatClient{
		Initiate(opts...),
		m,
	}
}

func (c *ChatClient) AddUserMessage(content string) {
	c.Messages = append(c.Messages, Message{
		"user",
		content,
		"",
	})
}

func (c *ChatClient) AddSystemMessage(content, reasoningContent string) {
	c.Messages = append(c.Messages, Message{
		"system",
		content,
		reasoningContent,
	})
}

func (c *ChatClient) GetNextChatCompletion() (*Response, error) {
	m, err := sonic.MarshalString(c.Messages)
	if err != nil {
		return nil, err
	}
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
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var res Response
	s := strings.ReplaceAll(string(body), "null", `"hi"`)
	err = sonic.UnmarshalString(s, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
