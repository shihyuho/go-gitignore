package gitignoreio

import (
	"fmt"
	"gopkg.in/resty.v1"
	"strings"
)

// Retrieve get the content of .gitignore from api
func (c *Client) Retrieve(types []string) ([]byte, error) {
	if l := len(types); l == 0 {
		return nil, fmt.Errorf("requires at least 1 type(s), only received %v", l)
	}
	path := fmt.Sprintf("%s/%s", c.API, strings.Join(types, ","))
	c.Log.Debugf("retrieving ignore content through: %s\n", path)
	resp, err := resty.R().Get(path)
	return resp.Body(), err
}
