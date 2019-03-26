package gitignoreio

import (
	"gopkg.in/AlecAivazis/survey.v1"
)

const (
	// DefaultPromptOptions specify the showing options to the prompt
	DefaultPromptOptions = 10
)

// Prompt asks interactive question
func (c *Client) Prompt(optionSize int) (answer []string, err error) {
	types, err := c.list()
	if err != nil {
		return
	}
	var defaults []string
	var done bool
	for {
		if answer, err = multiQs(types, defaults, optionSize); err != nil {
			return nil, err
		}
		if done, err = confirm(); err != nil {
			return nil, err
		}
		if done {
			break
		}
		defaults = answer
	}
	return
}

func multiQs(options []string, defaults []string, size int) (chose []string, err error) {
	prompt := &survey.MultiSelect{
		Message:  "What environments would your .gitignore to ignore?",
		Options:  options,
		Default:  defaults,
		PageSize: size,
	}
	err = survey.AskOne(prompt, &chose, survey.Required)
	return
}

func confirm() (ok bool, err error) {
	prompt := &survey.Confirm{
		Message: "Do you want to continue?",
		Default: true,
	}
	err = survey.AskOne(prompt, &ok, nil)
	return
}
