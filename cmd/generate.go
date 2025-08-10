package cmd

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// copyFile copie un fichier source vers une destination
func copyFile(src, dst string) error {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s n'est pas un fichier régulier", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	return err
}

// copyTemplates copie tous les fichiers du dossier templates vers documentation
func copyTemplates(templatesDir, docDir string) error {
	// Lire le contenu du dossier templates
	entries, err := os.ReadDir(templatesDir)
	if err != nil {
		return fmt.Errorf("impossible de lire le dossier templates: %v", err)
	}

	// Copier chaque fichier
	for _, entry := range entries {
		if entry.IsDir() {
			continue // Ignorer les sous-dossiers pour le moment
		}

		srcPath := filepath.Join(templatesDir, entry.Name())
		dstPath := filepath.Join(docDir, entry.Name())

		if err := copyFile(srcPath, dstPath); err != nil {
			return fmt.Errorf("erreur lors de la copie de %s: %v", entry.Name(), err)
		}

		fmt.Printf("✓ %s copié avec succès\n", entry.Name())
	}

	return nil
}

// generateCmd commande pour générer le dossier documentation
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Générer le dossier 'documentation'",
	Long: `Génère un dossier 'documentation' à la racine du projet courant
et y copie tous les templates de documentation disponibles.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Obtenir le répertoire de travail courant
		workingDir, err := os.Getwd()
		if err != nil {
			fmt.Printf("❌ Erreur: impossible de déterminer le répertoire courant: %v\n", err)
			return
		}

		// Chemins absolus
		docDir := filepath.Join(workingDir, "documentation")

		// Déterminer le chemin des templates
		// Si on lance depuis le binaire compilé, les templates sont dans le même dossier
		// Si on lance depuis le code source, ils sont dans ./templates
		var templatesDir string

		// Essayer d'abord le chemin relatif au binaire
		execPath, err := os.Executable()
		if err == nil {
			execDir := filepath.Dir(execPath)
			templatesDir = filepath.Join(execDir, "templates")
		}

		// Si le dossier templates n'existe pas, essayer le chemin relatif au code source
		if _, err := os.Stat(templatesDir); os.IsNotExist(err) {
			templatesDir = "./templates"
		}

		// Vérifier que le dossier templates existe
		if _, err := os.Stat(templatesDir); os.IsNotExist(err) {
			fmt.Printf("❌ Erreur: dossier templates introuvable\n")
			fmt.Printf("   Recherché dans: %s\n", templatesDir)
			return
		}

		// Créer le dossier documentation s'il n'existe pas
		if err := os.MkdirAll(docDir, os.ModePerm); err != nil {
			fmt.Printf("❌ Erreur: impossible de créer le dossier documentation: %v\n", err)
			return
		}

		fmt.Printf("📁 Dossier 'documentation' créé dans: %s\n", docDir)

		// Copier tous les templates
		if err := copyTemplates(templatesDir, docDir); err != nil {
			fmt.Printf("❌ Erreur lors de la copie des templates: %v\n", err)
			return
		}

		fmt.Println("✅ Documentation générée avec succès!")
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
