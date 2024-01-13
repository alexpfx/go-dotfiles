package cmd

import (
	"fmt"
	"github.com/alexpfx/go-dotfiles/common/util"
	"github.com/alexpfx/go-dotfiles/dotfile"
	"os"

	"github.com/spf13/cobra"
)

var (
	force      bool
	repository string
)

// repoCmd represents the repo command
var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Inicializa um novo repositório",
	Run: func(cmd *cobra.Command, args []string) {
		initRepoCmd(
			repository,
			gitDir,
			workTree,
			alias,
			force,
			dryRun,
		)

	},
}

func initRepoCmd(repo, gitDir, workTree, alias string, force bool, dryRun bool) {
	fmt.Printf("repo: %s gitDir: %s workTree: %s alias: %s force:%v\n", repo, gitDir, workTree, alias, force)

	conf := dotfile.Config{
		WorkTree: workTree,
		GitDir:   gitDir,
	}

	if !dryRun && force && util.DirExists(gitDir) {
		err := os.RemoveAll(gitDir)
		util.CheckFatal(err, "cannot remove gitDir")
	}

	cloneCmd := []string{"clone", "--bare", repo, gitDir}
	_, serr, err := util.ExecCmd(git, cloneCmd, dryRun)
	util.CheckFatal(err, serr)

	aliasArgs := []string{
		"--git-dir=" + conf.GitDir + "/",
		"--work-tree=" + conf.WorkTree,
	}

	configColorCmd := append(aliasArgs, "config", "--global", "color.ui", "always")
	configShowUntrackedFilesCmd := append(aliasArgs, "config", "--local", "status.showUntrackedFiles", "no")

	_, serr, err = util.ExecCmd(git, configColorCmd, dryRun)
	util.CheckFatal(err, serr)
	_, serr, err = util.ExecCmd(git, configShowUntrackedFilesCmd, dryRun)
	util.CheckFatal(err, serr)

	dotfile.WriteConfig(alias, &conf)

	checkout(alias, aliasArgs, workTree)
}

func checkout(alias string, aliasArgs []string, workTree string) {
	var existUntracked []string
	_, serr, err := util.ExecCmd(git, append(aliasArgs, "checkout"), dryRun)

	if err != nil {
		existUntracked = util.ParseExistUntracked(workTree, serr)
		if len(existUntracked) == 0 {
			util.CheckFatal(err, err.Error())
		}

		dotfile.BackupFiles(fmt.Sprintf(".%s%s_bkp/", workTree, alias), existUntracked)

		for _, untracked := range existUntracked {
			os.RemoveAll(untracked)
		}

		_, serr, err = util.ExecCmd(git, append(aliasArgs, "checkout"), dryRun)
		util.CheckFatal(err, serr)
	}
}

func init() {
	rootCmd.AddCommand(repoCmd)

	repoCmd.Flags().BoolVarP(&force, "force", "f", false, "force")
	repoCmd.Flags().StringVarP(&repository, "repository", "r", defaultRepo, "url do repositório")
}
