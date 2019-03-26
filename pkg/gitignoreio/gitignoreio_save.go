package gitignoreio

import (
	"github.com/mitchellh/go-homedir"
	"github.com/shihyuho/go-gitignore/pkg/paths"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Save create or append content to file
func (c *Client) Save(content []byte, path string) error {
	path, _ = homedir.Expand(path)
	abs, err := filepath.Abs(path)
	if err != nil {
		return err
	}
	exist, err := paths.Exists(abs)
	if err != nil {
		return err
	}
	if !exist {
		return c.create(content, abs)
	}

	c.Log.Debugf("file or directory already exist: %s", abs)
	if isFile, err := paths.IsRegularFile(abs); err != nil {
		return err
	} else if isFile {
		return c.append(content, abs)
	}

	c.Log.Debugf("searching %s in %s", defaultFilename, abs)
	abs = filepath.Join(abs, defaultFilename)
	exist, err = paths.Exists(abs)
	if err != nil {
		return err
	}
	if !exist {
		return c.create(content, abs)
	}
	return c.append(content, abs)
}

func (c *Client) create(content []byte, path string) error {
	c.Log.Debugf("create %s", path)
	c.Log.Debugf("ensuring intermediate directories exist...")
	if err := paths.EnsureDirectory(logrus.StandardLogger(), filepath.Dir(path)); err != nil {
		return err
	}
	return ioutil.WriteFile(path, content, 0755)
}

func (c *Client) append(content []byte, path string) error {
	c.Log.Debugf("appending content to the end of the file...")
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(content)
	return err
}
