package cmd

import (
	"fmt"
	"github.com/alexpfx/go-dotfiles/common/util"
	"github.com/alexpfx/go-dotfiles/dotfile"
	"github.com/spf13/cobra"
	"log"
)

// cfgCmd represents the cfg command
var cfgCmd = &cobra.Command{
	Use:                "cfg",
	Short:              "Comando para interagir com o diretório existente",
	Args:               cobra.MinimumNArgs(1),
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {
		if update {
			updateCfg(dryRun)
			return
		}

		conf := dotfile.LoadConfig(alias)
		aliasArgs := []string{
			"--git-dir=" + conf.GitDir + "/",
			"--work-tree=" + conf.WorkTree,
		}

		if len(args) == 0 {
			return
		}

		out, stderr, err := util.ExecCmd(git, append(aliasArgs, args...), dryRun)
		util.CheckFatal(err, stderr)
		fmt.Println(out)
	},
}

func updateCfg(dryRun bool) {
	fmt.Println(gitDir)
	fmt.Println(workTree)
	fmt.Println(alias)
	checkArgs(gitDir, workTree, alias)
	conf := dotfile.Config{
		WorkTree: workTree,
		GitDir:   gitDir,
	}

	if !dryRun {
		dotfile.WriteConfig(alias, &conf)
	}

}

func checkArgs(args ...string) {
	for _, s := range args {
		if s == "" {
			log.Fatal("all parameters must be provided")
		}
	}

}

func init() {
	rootCmd.AddCommand(cfgCmd)
	cfgCmd.Flags().BoolVar(&update, "update", false, "atualiza as configurações")

}
