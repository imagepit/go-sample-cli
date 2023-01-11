/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var hogeCmd = &cobra.Command{
	Use:   "hoge",
	Short: "hoge command",
	Long: `hoge command`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hoge called!!!!")
	},
}

func init() {
	rootCmd.AddCommand(hogeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hogeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hogeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
