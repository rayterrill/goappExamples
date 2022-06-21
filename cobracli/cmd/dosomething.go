/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var (
	port int
)

// dosomethingCmd represents the dosomething command
var dosomethingCmd = &cobra.Command{
	Use:   "dosomething",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dosomething called")
		fmt.Printf("port was :%s\n", strconv.Itoa(port))
	},
}

func init() {
	rootCmd.AddCommand(dosomethingCmd)

	dosomethingCmd.Flags().IntVarP(&port, "port", "p", 8080, "port of the web server")
}
