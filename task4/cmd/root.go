package cmd

import (
	"fmt"
	"strings"
	// "time"
	"os"
	"github.com/spf13/cobra"

        "github.com/go-git/go-git/v5"
        . "github.com/go-git/go-git/v5/_examples"
//        "github.com/go-git/go-git/v5/plumbing/object"
        //"github.com/go-git/go-git/v5/storage/memory"
)

var (
	version = "0.0.1"
	gitRepo string
	localDirectory string
)

func cloneRepo(url string,dir string) {
	fmt.Println("Cloning repository...")
	// Clones the given repository, creating the remote, the local branches
        // and fetching the objects, everything in memory:
        Info("git clone " + url)
	   //_ , err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
	   _, err := git.PlainClone(dir, false, &git.CloneOptions{
	   URL: url,
	   Progress: os.Stdout,
        })
        CheckIfError(err)

}

// Main command for the application
var rootCmd = &cobra.Command{
	Use:   "GitHubRepoAnalyzer [command] <github repo> [directory]",
	Version: version,
	Short: "GitHubRepoAnalyzer analyzes GitHub repositories for code quality and security.",
	Long:  `A Fast and Flexible Static Analysis CLI built with love by Gophers for GitHub repositories.`,
	/* Run: func(cmd *cobra.Command, args []string) {
		// This function will be called when the root command runs
		fmt.Println("Analyzing repository...")
		//cloneRepo("https://github.com/src-d/go-siva", "/tmp/foo")
	},
	*/
}

// Command to analyze code quality
var codeQualityCmd = &cobra.Command{
	Use:   "quality",
	Short: "Analyze code quality of the repository.",
	Aliases: []string{"q"},
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
	// Implementation for code quality analysis
	   for i := 0; i < len(args); i++ {
              fmt.Println("arg: " + strings.Join(args, " "))
           }
	},
}

// Command to check for security vulnerabilities
var securityCheckCmd = &cobra.Command{
	Use:   "security",
	Short: "Check the repository for security vulnerabilities.",
	Aliases: []string{"s"},
	Run: func(cmd *cobra.Command, args []string) {
		// Implementation for security checks
	},
}

// Command to ensure compliance with best practices
var complianceCheckCmd = &cobra.Command{
	Use:   "compliance",
	Short: "Ensure the repository complies with best practices.",
	Run: func(cmd *cobra.Command, args []string) {
		// Implementation for compliance checks
	},
}

// Initialize the commands
func init() {
	rootCmd.PersistentFlags().StringVar(&gitRepo, "repo", "r", "git repo to analyze")
	rootCmd.PersistentFlags().StringVar(&localDirectory, "directory", "current directory", "path to clone directory")
	rootCmd.MarkFlagRequired("repo")
	rootCmd.MarkFlagRequired("directory")
	rootCmd.AddCommand(codeQualityCmd)
	rootCmd.AddCommand(securityCheckCmd)
	rootCmd.AddCommand(complianceCheckCmd)
}

// Entry point of the application
//func main() {
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
	//if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

// Analyze the code quality of a GitHub repository
func analyzeCodeQuality(repoUrl string) {
	// Placeholder for code quality analysis logic
}

// Check a GitHub repository for security vulnerabilities
func checkSecurity(repoUrl string) {
	// Placeholder for security check logic
}

// Ensure a GitHub repository is compliant with best practices
func checkCompliance(repoUrl string) {
	// Placeholder for compliance check logic
}

