package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
        "github.com/go-git/go-git/v5"
        . "github.com/go-git/go-git/v5/_examples"
	//"github.com/go-git/go-git/v5/plumbing/object"
        //"github.com/go-git/go-git/v5/storage/memory"
)

var (
	version = "0.0.1"
	gitRepo string
	repoDirectory string
)

func cloneRepo(url ,dir string) {
	fmt.Println("Cloning repository...")
	// If no directory provided then
	// clone into /tmp/foo
	if dir != "" {
	   var err error
	   repoDirectory = dir   
	   if err != nil {
	      fmt.Println("Error getting current directory:", err)
	      return
	   }
	} else {
	   repoDirectory = "/tmp/foo"
	}
	fmt.Println("Cloning into " + repoDirectory)
        Info("git clone %s %s", url,repoDirectory)
	   _, err := git.PlainClone(dir, false, &git.CloneOptions{
	   URL: url,
	   Progress: os.Stdout,
        })
        CheckIfError(err)

}

// Main command for the application
var rootCmd = &cobra.Command{
	Use:   "gitHub",
	Version: version,
	Short: "gitHub analyzes GitHub repositories for code quality and security.",
	Long:  `A Fast and Flexible Static Analysis CLI built with love by Gophers for GitHub repositories.`,
}

// Command to analyze code quality
var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Clone a github repo",
	Aliases: []string{"c"},
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
	   gitRepo := args[0]
	   repoDirectory := ""
	   if len(args) > 1 {
	      repoDirectory = args[1]
	   }
	   cloneRepo(gitRepo, repoDirectory)
	},
}

// Initialize the commands
func init() {
	rootCmd.AddCommand(cloneCmd)
}

// Entry point of the application
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
	}
}
