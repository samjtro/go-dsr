package dsr

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/bytedance/sonic"
)

func (c *Client) Handler(resp *http.Response) (*Response, error) {
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	var res Response
	switch resp.StatusCode {
	case 200:
		if err != nil {
			return nil, err
		}
		s := strings.ReplaceAll(string(body), `"logprobs": null`, fmt.Sprintf(`"logprobs": %v`, Logprobs{}))
		err = sonic.UnmarshalString(s, &res)
		if err != nil {
			return nil, err
		}
	default:
		c.log.Error(string(body))
		return nil, fmt.Errorf("%d: %s", resp.StatusCode, string(body))
	}
	c.log.Info(string(body))
	return &res, nil
}
