# Base de données

<!-- Remplacez le badge par celui correspondant à votre SGBD -->
![MongoDB](https://img.shields.io/badge/MongoDB-%234ea94b.svg?style=for-the-badge&logo=mongodb&logoColor=white)
<!-- Autres exemples de badges :
![PostgreSQL](https://img.shields.io/badge/postgresql-%23336791.svg?style=for-the-badge&logo=postgresql&logoColor=white)
![MySQL](https://img.shields.io/badge/mysql-%234479A1.svg?style=for-the-badge&logo=mysql&logoColor=white)
![SQLite](https://img.shields.io/badge/sqlite-%2307405e.svg?style=for-the-badge&logo=sqlite&logoColor=white)
![Redis](https://img.shields.io/badge/redis-%23DD0031.svg?style=for-the-badge&logo=redis&logoColor=white)
-->

## Vue d'ensemble

<!-- Décrivez brièvement votre architecture de base de données -->
Ce document décrit l'architecture et la structure de la base de données utilisée pour [nom du projet]. La base de données [nom du SGBD] est utilisée pour [objectif principal, ex: stocker les données utilisateurs, gérer les transactions, etc.].

### Architecture générale

```
[Diagramme ASCII de votre architecture]
Exemple :
Application
    ↓
Base de données principale ([SGBD])
    ├── Collection/Table 1
    ├── Collection/Table 2
    └── Collection/Table N

Cache (Redis) [si applicable]
    ├── Sessions utilisateurs
    └── Données temporaires
```

### Informations techniques

- **SGBD** : [MongoDB/PostgreSQL/MySQL/autre] version [X.Y]
- **Environnement** : [Production/Développement/Staging]
- **Taille estimée** : [Estimation de la taille des données]
- **Nombre d'enregistrements** : [Estimation du volume]
- **Backup** : [Fréquence et stratégie de sauvegarde]
- **Réplication** : [Configuration de réplication si applicable]

## Conventions et bonnes pratiques

### Conventions de nommage

#### Collections/Tables
- Les collections/tables sont nommées au **singulier** et en **minuscule**
- Utiliser des noms explicites : `user` plutôt que `u`
- Pas d'espaces ni de caractères spéciaux
- Exemple : `user`, `product`, `order`, `category`

#### Champs/Colonnes
- Les champs de référence (clés étrangères) commencent par une **majuscule**
- Les champs simples sont en **camelCase**
- Les champs de date se terminent par `At` : `createdAt`, `updatedAt`
- Exemple : `UserId`, `ProductId`, `username`, `createdAt`

#### Index
- Préfixe : `idx_[table]_[champ]`
- Exemple : `idx_user_email`, `idx_product_category`

### Bonnes pratiques de conception

- **Normalisation** : [Niveau de normalisation utilisé et justification]
- **Index** : Index sur les champs fréquemment requêtés
- **Contraintes** : Utilisation de contraintes d'intégrité
- **Types de données** : Choix optimal des types pour chaque champ
- **Sécurité** : Chiffrement des données sensibles
- **Performance** : Optimisation des requêtes fréquentes

## Schéma de base de données

### Diagramme relationnel

```
[Insérez ici un diagramme ERD ou un schéma de vos collections/tables]
<!-- Vous pouvez utiliser des outils comme draw.io, dbdiagram.io, ou créer un diagramme ASCII -->

Exemple MongoDB :
user
├── _id (ObjectId)
├── username (String)
├── email (String)
└── ProfileId (ObjectId) → profile

product
├── _id (ObjectId)
├── name (String)
├── CategoryId (ObjectId) → category
└── UserId (ObjectId) → user
```

### Relations principales

<!-- Décrivez les relations entre vos entités -->
- **user** → **profile** : Relation 1:1 (un utilisateur a un profil)
- **user** → **order** : Relation 1:N (un utilisateur peut avoir plusieurs commandes)
- **product** → **category** : Relation N:1 (plusieurs produits dans une catégorie)
- **order** → **product** : Relation N:N (une commande peut contenir plusieurs produits)

## Collections/Tables

<!-- Répétez cette section pour chaque collection/table -->

### Collection `user`

**Description :** Les utilisateurs constituent le cœur de la plateforme. Cette collection stocke toutes les informations relatives aux comptes utilisateurs, leurs préférences et leurs relations avec les autres entités.

**Utilisation principale :** Authentification, autorisation, gestion des profils utilisateurs.

#### Schéma

```json
{
  "_id": "ObjectId",                     // Identifiant unique MongoDB
  "Bookmarks": ["ObjectId"],             // Références à la collection `product`
  "OrdersHistory": ["ObjectId"],         // Références à la collection `order`
  "CreatedProducts": ["ObjectId"],       // Références à la collection `product`
  "Reviews": ["ObjectId"],               // Références à la collection `review`
  "username": "string",                  // Pseudo unique de l'utilisateur
  "email": "string",                     // Adresse email (unique)
  "password": "string",                  // Mot de passe haché (bcrypt/scrypt)
  "firstName": "string",                 // Prénom
  "lastName": "string",                  // Nom de famille
  "role": "string",                      // Rôle : "admin", "moderator", "user"
  "status": "string",                    // Statut : "active", "suspended", "deleted"
  "profilePicture": "string",            // URL vers l'image de profil
  "isEmailVerified": "boolean",          // Email vérifié ou non
  "preferences": {                       // Préférences utilisateur
    "newsletter": "boolean",
    "notifications": "boolean",
    "language": "string"                 // Code langue : "fr", "en", "es"
  },
  "metadata": {                          // Métadonnées techniques
    "lastLoginIp": "string",
    "userAgent": "string",
    "loginCount": "number"
  },
  "createdAt": "timestamp",              // Date de création du compte
  "updatedAt": "timestamp",              // Dernière mise à jour
  "lastLoginAt": "timestamp"             // Dernière connexion
}
```

#### Explications détaillées des champs

- **`_id`** : Identifiant unique généré automatiquement par MongoDB
- **`Bookmarks`** : Liste des produits mis en favoris par l'utilisateur
- **`OrdersHistory`** : Historique des commandes passées par l'utilisateur
- **`CreatedProducts`** : Produits créés par l'utilisateur (si il est vendeur)
- **`Reviews`** : Avis et commentaires laissés par l'utilisateur
- **`username`** : Pseudo unique, utilisé pour l'affichage public (3-20 caractères)
- **`email`** : Adresse email unique, utilisée pour la connexion et les notifications
- **`password`** : Mot de passe haché avec [bcrypt/scrypt/autre], jamais stocké en clair
- **`firstName/lastName`** : Nom complet de l'utilisateur pour l'identification
- **`role`** : Niveau d'autorisation dans l'application :
  - `admin` : Accès complet à l'administration
  - `moderator` : Peut modérer le contenu
  - `user` : Utilisateur standard
- **`status`** : État du compte utilisateur :
  - `active` : Compte actif et utilisable
  - `suspended` : Compte temporairement suspendu
  - `deleted` : Compte marqué pour suppression
- **`profilePicture`** : URL vers l'image de profil (CDN ou stockage local)
- **`isEmailVerified`** : Indique si l'adresse email a été vérifiée
- **`preferences`** : Paramètres de préférence utilisateur modifiables
- **`metadata`** : Informations techniques pour le debugging et l'analyse
- **`createdAt/updatedAt`** : Horodatage pour l'audit et la traçabilité
- **`lastLoginAt`** : Suivi de l'activité utilisateur

#### Index recommandés

```javascript
// Index unique sur l'email
db.user.createIndex({ "email": 1 }, { unique: true })

// Index unique sur le username
db.user.createIndex({ "username": 1 }, { unique: true })

// Index composé pour les requêtes de recherche
db.user.createIndex({ "status": 1, "role": 1 })

// Index sur la date de création pour les tris chronologiques
db.user.createIndex({ "createdAt": -1 })

// Index sparse sur lastLoginAt (peut être null)
db.user.createIndex({ "lastLoginAt": -1 }, { sparse: true })
```

#### Requêtes fréquentes

```javascript
// Trouver un utilisateur par email
db.user.findOne({ "email": "user@example.com", "status": "active" })

// Utilisateurs actifs par ordre de création
db.user.find({ "status": "active" }).sort({ "createdAt": -1 })

// Utilisateurs avec email non vérifié
db.user.find({ "isEmailVerified": false, "status": "active" })

// Statistiques par rôle
db.user.aggregate([
  { $match: { "status": "active" } },
  { $group: { "_id": "$role", "count": { $sum: 1 } } }
])
```

---

### Collection `product`

**Description :** Catalogue des produits disponibles sur la plateforme. Contient toutes les informations nécessaires à l'affichage et à la gestion des produits.

#### Schéma

```json
{
  "_id": "ObjectId",
  "CategoryId": "ObjectId",              // Référence à la collection `category`
  "SellerId": "ObjectId",                // Référence à la collection `user`
  "Reviews": ["ObjectId"],               // Références à la collection `review`
  "name": "string",                      // Nom de la catégorie
  "description": "string",               // Description de la catégorie
  "slug": "string",                      // Identifiant URL-friendly
  "icon": "string",                      // URL de l'icône de catégorie
  "image": "string",                     // Image de bannière de catégorie
  "level": "number",                     // Niveau dans la hiérarchie (0 = racine)
  "path": "string",                      // Chemin complet : "/electronics/computers/laptops"
  "isActive": "boolean",                 // Catégorie active ou non
  "sortOrder": "number",                 // Ordre d'affichage
  "seo": {
    "metaTitle": "string",
    "metaDescription": "string"
  },
  "productCount": "number",              // Nombre de produits (mis à jour par trigger)
  "createdAt": "timestamp",
  "updatedAt": "timestamp"
}
```

#### Index recommandés

```javascript
// Index unique sur le slug
db.category.createIndex({ "slug": 1 }, { unique: true })

// Index pour la hiérarchie
db.category.createIndex({ "ParentCategoryId": 1, "sortOrder": 1 })

// Index pour les catégories actives
db.category.createIndex({ "isActive": 1, "level": 1 })
```

---

### Collection `order`

**Description :** Commandes passées par les utilisateurs, incluant les détails de facturation et livraison.

#### Schéma

```json
{
  "_id": "ObjectId",
  "UserId": "ObjectId",                  // Référence à la collection `user`
  "orderNumber": "string",               // Numéro de commande unique (ex: ORD-2024-001234)
  "items": [                             // Articles commandés
    {
      "ProductId": "ObjectId",           // Référence au produit
      "name": "string",                  // Nom du produit (snapshot)
      "price": "number",                 // Prix unitaire au moment de la commande
      "quantity": "number",              // Quantité commandée
      "subtotal": "number"               // Prix total pour cet article
    }
  ],
  "pricing": {
    "subtotal": "number",                // Sous-total des articles
    "shippingCost": "number",            // Frais de livraison
    "taxAmount": "number",               // Montant des taxes
    "discountAmount": "number",          // Montant de la remise
    "totalAmount": "number",             // Montant total
    "currency": "string"                 // Devise
  },
  "customer": {                          // Informations client (snapshot)
    "email": "string",
    "firstName": "string",
    "lastName": "string",
    "phone": "string"
  },
  "billingAddress": {
    "company": "string",
    "address1": "string",
    "address2": "string",
    "city": "string",
    "state": "string",
    "postalCode": "string",
    "country": "string"
  },
  "shippingAddress": {
    "company": "string",
    "address1": "string",
    "address2": "string",
    "city": "string",
    "state": "string",
    "postalCode": "string",
    "country": "string"
  },
  "payment": {
    "method": "string",                  // "card", "paypal", "bank_transfer"
    "status": "string",                  // "pending", "completed", "failed", "refunded"
    "transactionId": "string",           // ID de transaction du processeur de paiement
    "paidAt": "timestamp"                // Date de paiement
  },
  "shipping": {
    "method": "string",                  // "standard", "express", "overnight"
    "carrier": "string",                 // "dhl", "ups", "fedex", "colissimo"
    "trackingNumber": "string",          // Numéro de suivi
    "shippedAt": "timestamp",           // Date d'expédition
    "estimatedDelivery": "timestamp",    // Livraison estimée
    "deliveredAt": "timestamp"           // Date de livraison effective
  },
  "status": "string",                    // "pending", "confirmed", "shipped", "delivered", "cancelled"
  "notes": "string",                     // Notes internes
  "customerNotes": "string",             // Notes du client
  "createdAt": "timestamp",
  "updatedAt": "timestamp"
}
```

#### Index recommandés

```javascript
// Index unique sur le numéro de commande
db.order.createIndex({ "orderNumber": 1 }, { unique: true })

// Index pour les commandes utilisateur
db.order.createIndex({ "UserId": 1, "createdAt": -1 })

// Index pour le statut et la date
db.order.createIndex({ "status": 1, "createdAt": -1 })

// Index pour le suivi des expéditions
db.order.createIndex({ "shipping.trackingNumber": 1 }, { sparse: true })
```

---

## Triggers et procédures

### Triggers automatiques

#### Mise à jour des compteurs

```javascript
// Trigger MongoDB pour mettre à jour le compteur de produits dans les catégories
// (À implémenter au niveau applicatif ou avec MongoDB Change Streams)

// Trigger pour mettre à jour la note moyenne des produits
// Exécuté lors de l'ajout/modification/suppression d'un avis
```

#### Audit des modifications

```javascript
// Trigger d'audit pour tracer les modifications importantes
// Log des changements de statut des commandes
// Log des modifications de prix des produits
```

### Procédures stockées

#### Nettoyage des données

```sql
-- Exemple pour SQL (adapter selon votre SGBD)
-- Suppression des sessions expirées
-- Archivage des anciennes commandes
-- Nettoyage des logs anciens
```

## Migrations et versioning

### Structure des migrations

```
migrations/
├── 001_initial_schema.js              // Création du schéma initial
├── 002_add_user_preferences.js        // Ajout des préférences utilisateur
├── 003_update_product_pricing.js      // Modification structure pricing
├── 004_add_category_hierarchy.js      // Ajout hiérarchie catégories
└── 005_add_search_indexes.js          // Ajout index de recherche
```

### Exemple de migration

```javascript
// Migration 002_add_user_preferences.js
exports.up = async function(db) {
  // Ajouter le champ preferences à tous les utilisateurs existants
  await db.collection('user').updateMany(
    { preferences: { $exists: false } },
    {
      $set: {
        preferences: {
          newsletter: true,
          notifications: true,
          language: "fr"
        }
      }
    }
  );
};

exports.down = async function(db) {
  // Supprimer le champ preferences
  await db.collection('user').updateMany(
    {},
    { $unset: { preferences: "" } }
  );
};
```

### Versioning du schéma

- **Version actuelle** : v1.4.2
- **Dernière migration** : 005_add_search_indexes.js
- **Migrations en attente** : Aucune

## Sauvegarde et restauration

### Stratégie de sauvegarde

#### Sauvegarde automatique

```bash
# Sauvegarde quotidienne automatisée
# Script de sauvegarde MongoDB
mongodump --host [HOST] --port [PORT] --db [DATABASE] --out /backups/$(date +%Y%m%d_%H%M%S)/

# Rétention : 30 jours pour les sauvegardes quotidiennes
# Rétention : 12 mois pour les sauvegardes mensuelles
```

#### Sauvegarde manuelle

```bash
# Sauvegarde complète de la base
[commande-sauvegarde-complete]

# Sauvegarde d'une collection spécifique
[commande-sauvegarde-collection]

# Export des données en format JSON
[commande-export-json]
```

### Restauration

#### Restauration complète

```bash
# Restaurer depuis une sauvegarde
[commande-restauration-complete]

# Vérification après restauration
[commandes-verification]
```

#### Restauration partielle

```bash
# Restaurer une collection spécifique
[commande-restauration-collection]

# Restaurer des documents spécifiques
[commande-restauration-selective]
```

## Performance et optimisation

### Monitoring des performances

#### Métriques surveillées

- **Temps de réponse des requêtes** : < 100ms pour 95% des requêtes
- **Utilisation de la mémoire** : < 80% de la RAM disponible
- **Utilisation du CPU** : < 70% en moyenne
- **Taille de la base** : Croissance surveillée mensuellement
- **Index performance** : Analyse des requêtes lentes

#### Outils de monitoring

```bash
# Commandes de monitoring MongoDB
db.runCommand({serverStatus: 1})
db.stats()
db.[collection].getIndexes()

# Profiling des requêtes lentes
db.setProfilingLevel(1, {slowms: 100})
db.system.profile.find().sort({ts: -1}).limit(5)
```

### Optimisations appliquées

#### Index de performance

```javascript
// Index composés pour les requêtes complexes
db.product.createIndex({ "CategoryId": 1, "status": 1, "price": 1 })

// Index partiels pour les données actives uniquement
db.user.createIndex(
  { "lastLoginAt": -1 },
  { partialFilterExpression: { "status": "active" } }
)

// Index TTL pour les données temporaires
db.session.createIndex({ "createdAt": 1 }, { expireAfterSeconds: 86400 })
```

#### Optimisation des requêtes

```javascript
// Utilisation de l'aggregation pipeline pour les statistiques
db.order.aggregate([
  { $match: { "status": "completed" } },
  { $group: {
    "_id": { $dateToString: { format: "%Y-%m", date: "$createdAt" } },
    "totalRevenue": { $sum: "$pricing.totalAmount" },
    "orderCount": { $sum: 1 }
  }},
  { $sort: { "_id": -1 } }
])

// Projection pour limiter les données transférées
db.user.find(
  { "status": "active" },
  { "username": 1, "email": 1, "lastLoginAt": 1 }
)
```

## Sécurité

### Mesures de sécurité

#### Protection des données

- **Chiffrement au repos** : [Méthode de chiffrement utilisée]
- **Chiffrement en transit** : TLS 1.3 pour toutes les connexions
- **Authentification** : [Méthode d'authentification base de données]
- **Autorisation** : Rôles et permissions granulaires
- **Audit** : Logging de toutes les opérations sensibles

#### Configurations de sécurité

```javascript
// Configuration MongoDB sécurisée
// Authentification obligatoire
security:
  authorization: enabled

// Restriction des connexions
net:
  bindIp: 127.0.0.1,10.0.0.5  // IPs autorisées uniquement

// Audit des opérations
auditLog:
  destination: file
  format: JSON
  path: /var/log/mongodb/audit.json
```

### Gestion des accès

#### Rôles définis

```javascript
// Rôle lecture seule pour les rapports
db.createRole({
  role: "reportReader",
  privileges: [
    { resource: { db: "production", collection: "" }, actions: ["find"] }
  ],
  roles: []
})

// Rôle application avec permissions limitées
db.createRole({
  role: "appUser",
  privileges: [
    { resource: { db: "production", collection: "user" }, actions: ["find", "insert", "update"] },
    { resource: { db: "production", collection: "product" }, actions: ["find"] }
  ],
  roles: []
})
```

## Maintenance et surveillance

### Tâches de maintenance

#### Maintenance régulière

```bash
# Hebdomadaire
- Analyse des requêtes lentes
- Vérification de l'utilisation des index
- Nettoyage des logs anciens

# Mensuelle
- Analyse de la croissance des données
- Optimisation des index
- Revue des performances

# Trimestrielle
- Audit de sécurité complet
- Test de restauration des sauvegardes
- Planification de la capacité
```

#### Scripts de maintenance

```javascript
// Nettoyage des sessions expirées
db.session.deleteMany({
  "expiredAt": { $lt: new Date() }
})

// Mise à jour des compteurs de cache
db.category.aggregate([
  { $lookup: { from: "product", localField: "_id", foreignField: "CategoryId", as: "products" } },
  { $addFields: { productCount: { $size: "$products" } } },
  { $merge: { into: "category", whenMatched: "merge" } }
])
```

### Alertes et notifications

#### Seuils d'alerte

- **Espace disque** : Alerte à 80%, critique à 90%
- **Mémoire** : Alerte à 85%, critique à 95%
- **Requêtes lentes** : > 1000ms
- **Connexions** : > 80% du maximum autorisé
- **Réplication lag** : > 10 secondes

#### Canaux de notification

- **Email** : [admin@votre-domaine.com] pour les alertes critiques
- **Slack** : Canal #alerts pour toutes les notifications
- **SMS** : Pour les alertes critiques uniquement

## Documentation technique

### Ressources supplémentaires

- **Guide de développement** : [Lien vers documentation développeur]
- **API Documentation** : [Lien vers documentation API]
- **Procédures d'urgence** : [Lien vers guide incident]
- **Formation équipe** : [Lien vers ressources formation]

### Contacts

- **DBA Principal** : [Nom] - [email] - [téléphone]
- **Équipe DevOps** : [Contact équipe]
- **Support infrastructure** : [Contact support]

---

**Dernière mise à jour** : [Date de dernière modification]
**Version du document** : [Version]
**Révisé par** : [Nom du responsable] du produit
  "description": "string",               // Description détaillée
  "shortDescription": "string",          // Description courte pour les listes
  "price": "number",                     // Prix en centimes (évite les erreurs de virgule flottante)
  "originalPrice": "number",             // Prix original (pour les promotions)
  "currency": "string",                  // Code devise : "EUR", "USD", "GBP"
  "stock": "number",                     // Quantité en stock
  "sku": "string",                       // Code produit unique (Stock Keeping Unit)
  "images": [                            // Images du produit
    {
      "url": "string",                   // URL de l'image
      "alt": "string",                   // Texte alternatif
      "isPrimary": "boolean"             // Image principale ou non
    }
  ],
  "specifications": {                    // Caractéristiques techniques
    "weight": "number",                  // Poids en grammes
    "dimensions": {
      "length": "number",
      "width": "number",
      "height": "number"                 // Dimensions en cm
    },
    "color": "string",
    "material": "string",
    "brand": "string"
  },
  "seo": {                              // Optimisation SEO
    "metaTitle": "string",
    "metaDescription": "string",
    "slug": "string"                     // URL-friendly identifier
  },
  "status": "string",                    // "active", "inactive", "outOfStock", "discontinued"
  "tags": ["string"],                    // Tags pour la recherche et filtrage
  "rating": {
    "average": "number",                 // Note moyenne sur 5
    "count": "number"                    // Nombre total d'avis
  },
  "createdAt": "timestamp",
  "updatedAt": "timestamp",
  "publishedAt": "timestamp"             // Date de publication
}
```

#### Index recommandés

```javascript
// Index unique sur SKU
db.product.createIndex({ "sku": 1 }, { unique: true })

// Index pour les recherches
db.product.createIndex({ "name": "text", "description": "text", "tags": "text" })

// Index composé pour le catalogue
db.product.createIndex({ "status": 1, "CategoryId": 1, "createdAt": -1 })

// Index pour les prix
db.product.createIndex({ "price": 1 })
```

---

### Collection `category`

**Description :** Hiérarchie des catégories de produits pour l'organisation et la navigation.

#### Schéma

```json
{
  "_id": "ObjectId",
  "ParentCategoryId": "ObjectId",        // Référence à la catégorie parente (null pour racine)
  "name": "string",                      // Nom
