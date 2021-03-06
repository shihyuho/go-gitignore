package formatter

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
)

var ln = fmt.Sprintln()

// PlainFormatter 代表什麼都不 format 的 formatter
type PlainFormatter struct {
}

// Format 將傳入的 entry 轉換成要寫 log 的文字
func (f *PlainFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var buf *bytes.Buffer
	if entry.Buffer != nil {
		buf = entry.Buffer
	} else {
		buf = &bytes.Buffer{}
	}
	if entry.Message != "" {
		buf.WriteString(entry.Message)
	}
	if !strings.HasSuffix(entry.Message, ln) {
		buf.WriteString(ln)
	}
	return buf.Bytes(), nil
}
