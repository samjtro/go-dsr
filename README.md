# go-dsr

[![GoDoc](https://pkg.go.dev/badge/github.com/samjtro/go-dsr)](https://pkg.go.dev/github.com/samjtro/go-dsr)
[![Go Report Card](https://goreportcard.com/badge/github.com/samjtro/go-dsr)](https://goreportcard.com/report/github.com/samjtro/go-dsr)

⚠️ **This is currently in development**. Things will probably break, but existing functionality is usable. ⚠️

```shell
go get github.com/samjtro/go-dsr
```

built by: [rizome labs](https://rizome.dev)

contact us: [hi (at) rizome.dev](mailto:hi@rizome.dev)

## quick start

0. create a `.env`, formatted:
```
DEEPSEEK_API_KEY="KEY"
```

## example

```go
c := dsr.NewChatClient()
c.AddUserMessage("hello!")
res, _ := c.GetNextChatCompletion()
c.AddSystemMessage(res.Choices[0].Message.Content, res.Choices[0].Message.ReasoningContent)
fmt.Println(c.Messages)
```

output:

```
[{user hello! } {system Hello! How can I assist you today? 😊 Okay, the user said "hello!" That's a friendly greeting. I should respond in a welcoming manner. Maybe say "Hello! How can I assist you today?" to keep the conversation going. Need to make sure the tone is approachable and helpful. Let me check for any typos. Yep, that looks good. Ready to send.}]
```
