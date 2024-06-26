package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wcygan/yap/yap-cli/internal/application"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "yap-cli",
	Short: "Yap with people on the internet!",
	Run: func(cmd *cobra.Command, args []string) {
		host := cmd.Flag("host").Value.String()

		initialModel, err := application.InitialModel(host)
		if err != nil {
			fmt.Printf("could not create initial model: %s\n", err)
			os.Exit(1)
		}

		if _, err := tea.NewProgram(initialModel).Run(); err != nil {
			fmt.Printf("could not start program: %s\n", err)
			os.Exit(1)
		}
	},
}

var host string

func init() {
	rootCmd.PersistentFlags().StringVar(&host, "host", "localhost:50050", "Host server address")
}
