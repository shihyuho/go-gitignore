package gitignoreio

import (
	"github.com/sirupsen/logrus"
	"os"
)

const (
	// ENV is the Env key for GITIGNORE_API
	ENV             = `GITIGNORE_API`
	api             = `https://www.gitignore.io/api`
	defaultFilename = ".gitignore"
)

// GetAPI returns GitIgnore Api
func GetAPI() string {
	if val, exist := os.LookupEnv(ENV); exist {
		return val
	}
	return api
}

// Client wrap the GitIgnore.io api
type Client struct {
	Log *logrus.Logger
	API string
}
