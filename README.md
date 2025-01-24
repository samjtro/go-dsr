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
c.AddMessage(res.Choices[0].Message)
c.AddUserMessage("can you help me with a math problem?")
res, _ = c.GetNextChatCompletion()
c.AddMessage(res.Choices[0].Message)
fmt.Println(c.Messages)
```

output:

```
[{user hello! }

{assistant Hello! How can I assist you today? 😊 Okay, the user just said "hello!" That's a friendly greeting.

I should respond in a warm and welcoming manner.

Maybe say something like "Hello! How can I assist you today?" to keep the conversation going.

Keep it simple and open-ended so they feel comfortable asking for help.}

{user can you help me with a math problem? }

{assistant Of course! I'd be happy to help with your math problem. Please go ahead and share the details, and I'll do my best to guide you through it step by step. 😊 Okay, the user asked if I can help with a math problem. Let me respond positively first. I need to be encouraging.

They might have a specific problem in mind. I should ask them to provide the details so I can understand what they need. Maybe it's algebra, calculus, geometry? I shouldn't assume.

I should also mention that they can type the problem as they have it, even if it's not perfectly formatted. That way they feel comfortable sharing whatever they have. Maybe add an emoji to keep the tone friendly. Let me check the previous message for the emoji usage. The user used a smiley in their first message, so it's okay to use one here too.

Wait, the last response from the assistant had a 😊. Repeating emojis might be okay, but maybe a different one? Or stick to the same. Hmm. Probably better to use a similar friendly emoji. Let me go with 😊 again.

Need to make sure the user feels supported. Also, avoid technical jargon in the request. Keep it simple. Let them know I can guide them through step by step. That's important because they might need explanation, not just the answer.

Alright, the response should be something like: "Of course! I'd be happy to help with your math problem. Please go ahead and share the details, and I'll do my best to guide you through it step by step. 😊"}]
```
