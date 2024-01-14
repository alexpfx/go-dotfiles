package cmd

import (
	"fmt"
	"github.com/alexpfx/go-dotfiles/common/util"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

const git = "/usr/bin/git"
const defaultAlias = "cfg"
const defaultRepo = "https://github.com/alexpfx/linux_wayland_dotfiles.git"
const defaultGitdir = ".cfg"

var (
	update    bool
	dryRun    bool
	gitDir    string
	alias     string
	workTree  string
	version   = "development"
	buildTime = "N\\A"
	homeDir   string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-dot",
	Short: "",
	Long:  ``,
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
	var err error
	homeDir, err = os.UserHomeDir()
	util.CheckFatal(err, "")

	rootCmd.PersistentFlags().BoolVar(&dryRun, "dry-run", false, "mostra o resultado sem efetivar os comandos")
	rootCmd.PersistentFlags().StringVarP(&alias, "alias", "a", defaultAlias, "alias")
	rootCmd.PersistentFlags().StringVarP(&gitDir, "git-dir", "d", filepath.Join(homeDir, defaultGitdir), "git dir")
	rootCmd.PersistentFlags().StringVarP(&workTree, "work-tree", "t", homeDir, "work tree")

}

func printVersionAndExit() {
	fmt.Printf("	Version: %s\n	Build time: %s", version, buildTime)
	os.Exit(0)
}
