package main

import (
	"bytes"
	"fmt"
	"github.com/shihyuho/go-gitignore/pkg/formatter"
	"github.com/shihyuho/go-gitignore/pkg/gitignoreio"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"text/template"
)

const (
	help = `
The command line interface creates or adds .gitignore 

To fetches a .gitignore from gitignore.io.

	$ gitignore go intellij+all 

Alternatively, running the command without arguments gives you a interactive 
interface to choose what to ignore.

	$ gitignore

To save the content to a file instead of stdout, use '--outout'. 
If the output already exist, it will append the content to the end of the file.

	$ gitignore -o /path/to/.gitignore

If the output is a directory, the default filename to save will be '.gitignore'.

	$ gitignore -o .

Environment:

  ${{.ENV}}	set the api url for integrating with gitignore.io
`
)

var (
	optionSize int
	output     string
	verbose    bool
	client     = &gitignoreio.Client{}
)

func main() {
	if err := newRootCmd(os.Args[1:]).Execute(); err != nil {
		os.Exit(1)
	}
}

func newRootCmd(args []string) *cobra.Command {
	cmd := &cobra.Command{
		Use:          "gitignore [TYPE...]",
		Short:        "Fetches a .gitignore from gitignore.io",
		Long:         format(help),
		SilenceUsage: true,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			client.Log = logrus.New()
			client.Log.SetOutput(cmd.OutOrStdout())
			client.Log.SetFormatter(&formatter.PlainFormatter{})
			if verbose {
				client.Log.SetLevel(logrus.DebugLevel)
			}
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				types, err := client.Prompt(optionSize)
				if err != nil {
					return err
				}
				args = append(args, types...)
			}
			content, err := client.Retrieve(args)
			if err != nil {
				return err
			}
			return out(content)
		},
	}

	f := cmd.Flags()
	f.BoolVarP(&verbose, "verbose", "v", verbose, "enable verbose output")
	f.StringVar(&client.API, "api", gitignoreio.GetAPI(), "specify api url for integrating with gitignore.io, Overrides $"+gitignoreio.ENV)
	f.StringVarP(&output, "output", "o", "", "specify the output file or directory to save instead of stdout")
	f.IntVar(&optionSize, "option-size", gitignoreio.DefaultPromptOptions, "specify the showing option size for prompt")
	f.Parse(args)

	return cmd
}

func out(content []byte) error {
	if output != "" {
		return client.Save(content, output)
	}
	fmt.Println(string(content))
	return nil
}

func format(tpl string) string {
	var buf bytes.Buffer
	parsed := template.Must(template.New("").Parse(tpl))
	data := make(map[string]string)
	data["ENV"] = gitignoreio.ENV
	err := parsed.Execute(&buf, data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return buf.String()
}
