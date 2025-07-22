package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Afficher la version
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Afficher la version actuelle de l'outil",
	Long:  `Mokujin a plusieurs versions, voici sa version actuelle`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Mokujin version 1.0.0 !")
	},
}

// Lance la commande
func init() {
	rootCmd.AddCommand(versionCmd)
}
