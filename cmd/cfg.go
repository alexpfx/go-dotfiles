package cmd

import (
	"fmt"
	"github.com/alexpfx/go-dotfiles/dotfile"
	"github.com/spf13/cobra"
	"log"
)

// cfgCmd represents the cfg command
var cfgCmd = &cobra.Command{
	Use:   "cfg",
	Short: "Comando para interagir com o diretório existente",
	Run: func(cmd *cobra.Command, args []string) {

		if update {
			updateCfg(dryRun)
			return
		}

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
	cfgCmd.Flags().BoolVarP(&update, "update", "u", false, "atualiza as configurações")

}
