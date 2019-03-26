package gitignoreio

import (
	"fmt"
	"gopkg.in/resty.v1"
	"strings"
)

func (c *Client) list() ([]string, error) {
	path := fmt.Sprintf("%s/list", c.API)
	resp, err := resty.R().Get(path)
	if err != nil {
		return nil, err
	}
	data := strings.Replace(resp.String(), "\n", ",", -1)
	return strings.Split(data, ","), nil
}
