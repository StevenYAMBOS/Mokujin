# Mokujin

<div align="center"> <img src="https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white" alt="Golang"> <img src="https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white" alt="Docker"> <img src="https://img.shields.io/badge/git-%23F05032.svg?style=for-the-badge&logo=git&logoColor=white" alt="Git"> </div> <div align="center"> <h3>CLI App pour les documentations techniques</h3> <p> <strong>Mokujin</strong> est une application en ligne de commande moderne développée avec <a href="https://github.com/spf13/cobra">Cobra</a> qui permet de générer des fichiers de documentation technique avancés. </p> </div>

---

## Description

Mokujin est un outil en ligne de commande simple et léger, permettant de générer automatiquement un dossier de documentation contenant plusieurs fichiers Markdown prédéfinis. Il s'adresse aux développeurs, rédacteurs techniques ou équipes souhaitant structurer rapidement la documentation d'un projet.

### Fonctionnalités

- Génération rapide d'un dossier `documentation` structuré
- Templates Markdown prédéfinis et personnalisables
- Interface en ligne de commande intuitive
- Développé avec Cobra pour une expérience utilisateur optimale

### Templates générés

Le dossier `documentation` créé contient les fichiers suivants :

- **README.md** - Documentation générale du projet
- **Base de données.md** - Spécifications et schémas de base de données
- **Organisation du dépôt.md** - Structure et conventions du projet

## Structure du projet

```
mokujin/
├── cmd/
│   ├── root.go          # Commande racine et configuration
│   └── version.go       # Gestion des versions
├── templates/
│   ├── documentation.md # Template de documentation générale
│   └── specification.md # Template de spécifications
├── go.mod
├── go.sum
├── LICENSE
├── main.go
└── README.md
```

## Installation

### Prérequis

- Go 1.19+ installé sur votre système
- Git pour cloner le projet

### Installation depuis les sources

```bash
# Cloner le dépôt
git clone https://github.com/StevenYAMBOS/mokujin.git
cd mokujin

# Compiler l'application
go build -o mokujin .

# Rendre l'exécutable accessible (optionnel)
sudo mv mokujin /usr/local/bin/
```

### Installation avec Go

```bash
go install github.com/StevenYAMBOS/mokujin@latest
```

## Utilisation

### Commandes disponibles

```bash
# Générer un dossier de documentation
mokujin generate

# Afficher la version
mokujin version

# Afficher l'aide
mokujin --help
```

### Exemple d'utilisation

```bash
# Se placer dans le répertoire de votre projet
cd mon-projet

# Générer la documentation
mokujin generate

# Le dossier 'documentation' est créé avec les templates
ls documentation/
# README.md  Base de données.md  Organisation du dépôt.md
```

## Docker

### Construction de l'image

```bash
docker build -t mokujin .
```

### Utilisation avec Docker

```bash
# Monter le répertoire courant et générer la documentation
docker run --rm -v $(pwd):/workspace mokujin generate
```

## Développement

### Prérequis de développement

- Go 1.19+
- Git
- Make (optionnel)

### Lancer en mode développement

```bash
# Cloner le projet
git clone https://github.com/StevenYAMBOS/mokujin.git
cd mokujin

# Installer les dépendances
go mod download

# Lancer l'application
go run main.go generate
```

### Tests

```bash
# Lancer les tests
go test ./...

# Tests avec couverture
go test -cover ./...
```

### Contribution

Les contributions sont les bienvenues ! Voici comment procéder :

1. Fork le projet
2. Créez votre branche de fonctionnalité (`git checkout -b feature/amazing-feature`)
3. Committez vos changements (`git commit -m 'Add amazing feature'`)
4. Push vers la branche (`git push origin feature/amazing-feature`)
5. Ouvrez une Pull Request

## Licence

Ce projet est sous licence MIT. Voir le fichier LICENSE pour plus de détails.

## Support

Si vous rencontrez des problèmes ou avez des questions :

- [Signaler un bug](https://github.com/StevenYAMBOS/mokujin/issues/new?template=bug_report.md)
- [Proposer une fonctionnalité](https://github.com/StevenYAMBOS/mokujin/issues/new?template=feature_request.md)
- Email : [stevenyambos@gmail.com](mailto:stevenyambos@gmail.com)

## Remerciements

- [Cobra](https://github.com/spf13/cobra) - Framework CLI pour Go
- [Viper](https://github.com/spf13/viper) - Gestion de la configuration
- La communauté Go pour leurs outils et ressources

---

<div align="center"> Développé par <a href="https://github.com/StevenYAMBOS">Steven YAMBOS</a> <br><br> <a href="https://x.com/StevenYambos">Twitter</a> • <a href="https://www.linkedin.com/in/steven-yambos">LinkedIn</a> • <a href="mailto:stevenyambos@gmail.com">Email</a> </div>
