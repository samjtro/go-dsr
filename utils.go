package dsr

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/bytedance/sonic"
)

func Handler(resp *http.Response) (*Response, error) {
	var res Response
	switch resp.StatusCode {
	case 200:
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		s := strings.ReplaceAll(string(body), `"logprobs": null`, fmt.Sprintf(`"logprobs": %v`, Logprobs{}))
		err = sonic.UnmarshalString(s, &res)
		if err != nil {
			return nil, err
		}
	default:
		fmt.Println(resp)
	}
	return &res, nil
}
