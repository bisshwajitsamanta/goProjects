// Package cmd /*
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	name, greeting  string
	preview, prompt bool
	debug           bool = false
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "motd",
	Short: "It is a greeting cli",
	Long:  `It is a greeting cli which accepts name and greet them`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if !prompt && (name == "" || greeting == "") {
			err := cmd.Usage()
			if err != nil {
				return
			}
			os.Exit(1)
		}
		if debug {
			fmt.Println("Name:", name)
			fmt.Println("Greeting:", greeting)
			fmt.Println("Prompt:", prompt)
			fmt.Println("Preview:", preview)
			os.Exit(0)
		}
		// Conditionally read from stdin
		if prompt {
			name, greeting = renderPrompt()
		}
		//Generate message
		message := buildMessage(name, greeting)

		// Either preview the message or write to file
		if preview {
			fmt.Println(message)
		} else {
			//write content
			f, err := os.OpenFile("motd.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
			if err != nil {
				fmt.Println("Error unable to open motd.txt")
				os.Exit(1)
			}
			defer f.Close()
			_, err = f.Write([]byte(message))
			if err != nil {
				fmt.Println("Failed to write in motd.txt")
				os.Exit(1)
			}
		}
	},
}

func buildMessage(name, greeting string) string {
	return fmt.Sprintf("%s, %s", name, greeting)
}

func renderPrompt() (name, greeting string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Your Greeting: ")
	greeting, _ = reader.ReadString('\n')
	greeting = strings.TrimSpace(greeting)
	fmt.Print("Your Name: ")
	name, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)
	return
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&preview, "preview", "v", false, "Preview Message")
	rootCmd.Flags().BoolVarP(&prompt, "prompt", "p", false, "Prompt for name and greeting")
	rootCmd.Flags().StringVarP(&greeting, "greeting", "g", "", "Greeting to use in the message")
	rootCmd.Flags().StringVarP(&name, "name", "n", "", "Name to use in the message")

	if os.Getenv("DEBUG") != "" {
		debug = true
	}

}
