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

0. create a `.env`, `DEEPSEEK_API_KEY="KEY"`

## example

```go
c := dsr.NewChatClient()
c.AddMessage("user", "hello!", "")
res, _ := c.GetNextChatCompletion()
fmt.Println(res.Choices[0].Message.Content)
```
