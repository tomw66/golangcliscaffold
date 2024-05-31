/*
Copyright © 2024 Tom Wilson <t.wilson6@exeter.ac.uk>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tomw66/golangcliscaffold/common"
	"github.com/tomw66/golangcliscaffold/dirs"
)

// filesCmd represents the files command
var dirsCmd = &cobra.Command{
	Use:   "dirs",
	Short: "Show the largest directories in the given path.",
	Long: `Quickly scan a directory and find large directories. Use the flags below to target the output.`,
	Run: func(cmd *cobra.Command, args []string) {
		if Debug {
			common.LogFlags()
		}
		dirsFound, _ := dirs.ReadDirDepth(Path)
		dirs.PrintResults(dirsFound)
	},
}

var Depth int
var Mindirsize int

func init() {
	rootCmd.AddCommand(dirsCmd)
   
	dirsCmd.PersistentFlags().IntVarP(&Depth, "depth", "", 2, "Depth of directory tree to display")
	viper.BindPFlag("depth", dirsCmd.PersistentFlags().Lookup("depth"))
   
	dirsCmd.PersistentFlags().IntVarP(&Mindirsize, "mindirsize", "", 100, "Only display directories larger than this threshold in MB.")
	viper.BindPFlag("mindirsize", dirsCmd.PersistentFlags().Lookup("mindirsize"))
   }