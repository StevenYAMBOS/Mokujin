/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "Mokujin",
	Short: "Génèrez vos documents techniques en un clin d'oeil",
	Long: `Mokujin est un outil en ligne de commande simple et léger, permettant de générer automatiquement un dossier de documentation contenant plusieurs fichiers Markdown prédéfinis. Il s’adresse aux développeurs, rédacteurs techniques ou équipes souhaitant structurer rapidement la documentation d’un projet.
	Le dossier 'documentation' généré contient les templates suivants :
	- README.md
	- Base de données.md
	- Organisation du dépôt.md`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.Mokujin.yaml)")
	rootCmd.PersistentFlags().StringP("author", "a", "Steven YAMBOS", "Les droits du projet reviennent à Steven YAMBOS")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
