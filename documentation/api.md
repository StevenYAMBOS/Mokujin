# Documentation API

## Vue d'ensemble

Cette documentation décrit l'API REST de [nom du projet]. L'API permet aux applications clientes d'interagir avec le système pour [décrire l'objectif principal de l'API].

### Informations générales

- **Version de l'API** : v1.0.0
- **URL de base** : `https://api.[votre-domaine].com/v1`
- **Protocole** : HTTPS uniquement
- **Format** : JSON
- **Authentification** : [JWT/OAuth2/API Key/autre]
- **Rate limiting** : [limite de requêtes par heure/minute]

### URLs des environnements

| Environnement | URL de base | Description |
|---------------|-------------|-------------|
| **Production** | `https://api.[domaine].com/v1` | Environnement de production |
| **Staging** | `https://api-staging.[domaine].com/v1` | Tests et validation |
| **Développement** | `https://api-dev.[domaine].com/v1` | Développement |
| **Local** | `http://localhost:3000/api/v1` | Développement local |

## Authentification

### Type d'authentification

L'API utilise [JWT/OAuth2/API Key] pour l'authentification. Toutes les requêtes (sauf indication contraire) nécessitent une authentification valide.

#### Obtenir un token

```http
POST /auth/login
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "motdepasse123"
}
```

**Réponse :**
```json
{
  "success": true,
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "refreshToken": "rt_1234567890abcdef",
    "expiresIn": 3600,
    "user": {
      "id": "user123",
      "email": "user@example.com",
      "role": "user"
    }
  }
}
```

#### Utiliser le token

Inclure le token dans l'en-tête `Authorization` de chaque requête :

```http
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

#### Rafraîchir le token

```http
POST /auth/refresh
Content-Type: application/json

{
  "refreshToken": "rt_1234567890abcdef"
}
```

### Gestion des erreurs d'authentification

| Code | Message | Description |
|------|---------|-------------|
| 401 | `Unauthorized` | Token manquant ou invalide |
| 403 | `Forbidden` | Token valide mais permissions insuffisantes |
| 429 | `Too Many Requests` | Limite de taux dépassée |

## Format des réponses

### Structure standard

Toutes les réponses de l'API suivent cette structure :

```json
{
  "success": boolean,
  "data": object|array|null,
  "message": "string",
  "errors": array|null,
  "meta": {
    "timestamp": "ISO 8601",
    "version": "string",
    "requestId": "string"
  }
}
```

### Exemples de réponses

#### Succès
```json
{
  "success": true,
  "data": {
    "id": "123",
    "name": "Exemple"
  },
  "message": "Ressource récupérée avec succès",
  "errors": null,
  "meta": {
    "timestamp": "2024-03-15T10:30:00Z",
    "version": "1.0.0",
    "requestId": "req_1234567890"
  }
}
```

#### Erreur
```json
{
  "success": false,
  "data": null,
  "message": "Validation échouée",
  "errors": [
    {
      "field": "email",
      "code": "INVALID_FORMAT",
      "message": "Format d'email invalide"
    }
  ],
  "meta": {
    "timestamp": "2024-03-15T10:30:00Z",
    "version": "1.0.0",
    "requestId": "req_1234567890"
  }
}
```

## Codes de statut HTTP

| Code | Signification | Utilisation |
|------|---------------|-------------|
| **200** | OK | Requête réussie |
| **201** | Created | Ressource créée |
| **204** | No Content | Suppression réussie |
| **400** | Bad Request | Erreur de validation |
| **401** | Unauthorized | Authentification requise |
| **403** | Forbidden | Accès interdit |
| **404** | Not Found | Ressource non trouvée |
| **409** | Conflict | Conflit (ressource existe déjà) |
| **422** | Unprocessable Entity | Erreur de validation métier |
| **429** | Too Many Requests | Limite de taux atteinte |
| **500** | Internal Server Error | Erreur serveur |
| **503** | Service Unavailable | Service temporairement indisponible |

## Pagination

### Paramètres de pagination

| Paramètre | Type | Défaut | Description |
|-----------|------|---------|-------------|
| `page` | integer | 1 | Numéro de page (commence à 1) |
| `limit` | integer | 20 | Nombre d'éléments par page (max: 100) |
| `sort` | string | - | Champ de tri (ex: `name`, `-createdAt`) |

### Exemple de requête paginée

```http
GET /users?page=2&limit=10&sort=-createdAt
```

### Format de réponse paginée

```json
{
  "success": true,
  "data": [
    // ... éléments de la page
  ],
  "message": "Utilisateurs récupérés",
  "errors": null,
  "meta": {
    "pagination": {
      "currentPage": 2,
      "totalPages": 15,
      "totalItems": 147,
      "itemsPerPage": 10,
      "hasNextPage": true,
      "hasPreviousPage": true
    },
    "timestamp": "2024-03-15T10:30:00Z",
    "version": "1.0.0",
    "requestId": "req_1234567890"
  }
}
```

## Filtrage et recherche

### Paramètres de filtrage

| Opérateur | Syntaxe | Exemple | Description |
|-----------|---------|---------|-------------|
| Égalité | `field=value` | `status=active` | Valeur exacte |
| Recherche | `search=term` | `search=john` | Recherche textuelle |
| Plage | `field[gte]=value` | `price[gte]=100` | Supérieur ou égal |
| Plage | `field[lte]=value` | `price[lte]=500` | Inférieur ou égal |
| Dans | `field[in]=val1,val2` | `status[in]=active,pending` | Valeurs multiples |
| Date | `createdAt[gte]=date` | `createdAt[gte]=2024-01-01` | Date supérieure |

### Exemples de filtrage

```http
# Utilisateurs actifs créés cette année
GET /users?status=active&createdAt[gte]=2024-01-01

# Produits dans une gamme de prix
GET /products?price[gte]=50&price[lte]=200

# Recherche d'utilisateurs par nom
GET /users?search=martin&sort=lastName
```

## Endpoints

### Authentification

#### POST /auth/login
Authentification utilisateur

**Corps de la requête :**
```json
{
  "email": "string (required)",
  "password": "string (required)",
  "rememberMe": "boolean (optional)"
}
```

**Réponse 200 :**
```json
{
  "success": true,
  "data": {
    "token": "string",
    "refreshToken": "string",
    "expiresIn": "number",
    "user": {
      "id": "string",
      "email": "string",
      "role": "string"
    }
  }
}
```

#### POST /auth/register
Inscription d'un nouvel utilisateur

**Corps de la requête :**
```json
{
  "email": "string (required)",
  "password": "string (required)",
  "firstName": "string (required)",
  "lastName": "string (required)"
}
```

#### POST /auth/logout
Déconnexion (invalide le token)

#### POST /auth/forgot-password
Demande de réinitialisation de mot de passe

#### POST /auth/reset-password
Réinitialisation du mot de passe

---

### Utilisateurs

#### GET /users
Récupérer la liste des utilisateurs

**Paramètres :**
- `page` (integer) : Numéro de page
- `limit` (integer) : Éléments par page
- `search` (string) : Recherche par nom/email
- `role` (string) : Filtrer par rôle
- `status` (string) : Filtrer par statut

**Réponse 200 :**
```json
{
  "success": true,
  "data": [
    {
      "id": "user123",
      "email": "user@example.com",
      "firstName": "John",
      "lastName": "Doe",
      "role": "user",
      "status": "active",
      "createdAt": "2024-01-15T10:30:00Z",
      "lastLoginAt": "2024-03-10T14:22:00Z"
    }
  ]
}
```

#### GET /users/{id}
Récupérer un utilisateur par ID

**Paramètres d'URL :**
- `id` (string, required) : Identifiant de l'utilisateur

**Réponse 200 :**
```json
{
  "success": true,
  "data": {
    "id": "user123",
    "email": "user@example.com",
    "firstName": "John",
    "lastName": "Doe",
    "role": "user",
    "status": "active",
    "preferences": {
      "newsletter": true,
      "notifications": false
    },
    "createdAt": "2024-01-15T10:30:00Z",
    "updatedAt": "2024-03-10T14:22:00Z"
  }
}
```

#### POST /users
Créer un nouvel utilisateur

**Corps de la requête :**
```json
{
  "email": "string (required)",
  "firstName": "string (required)",
  "lastName": "string (required)",
  "role": "string (optional, default: user)",
  "password": "string (required)"
}
```

#### PUT /users/{id}
Mettre à jour un utilisateur

**Corps de la requête :**
```json
{
  "firstName": "string (optional)",
  "lastName": "string (optional)",
  "preferences": {
    "newsletter": "boolean (optional)",
    "notifications": "boolean (optional)"
  }
}
```

#### DELETE /users/{id}
Supprimer un utilisateur

---

### Produits

#### GET /products
Récupérer la liste des produits

**Paramètres :**
- `page`, `limit`, `sort` : Pagination
- `category` (string) : Filtrer par catégorie
- `minPrice`, `maxPrice` (number) : Filtrer par prix
- `status` (string) : Filtrer par statut
- `search` (string) : Recherche textuelle

#### GET /products/{id}
Récupérer un produit par ID

#### POST /products
Créer un nouveau produit

**Corps de la requête :**
```json
{
  "name": "string (required)",
  "description": "string (required)",
  "price": "number (required)",
  "categoryId": "string (required)",
  "stock": "number (required)",
  "sku": "string (required)",
  "images": [
    {
      "url": "string",
      "alt": "string",
      "isPrimary": "boolean"
    }
  ]
}
```

#### PUT /products/{id}
Mettre à jour un produit

#### DELETE /products/{id}
Supprimer un produit

---

### Commandes

#### GET /orders
Récupérer les commandes

#### GET /orders/{id}
Récupérer une commande par ID

#### POST /orders
Créer une nouvelle commande

**Corps de la requête :**
```json
{
  "items": [
    {
      "productId": "string",
      "quantity": "number"
    }
  ],
  "shippingAddress": {
    "address1": "string",
    "city": "string",
    "postalCode": "string",
    "country": "string"
  },
  "paymentMethod": "string"
}
```

#### PUT /orders/{id}/status
Mettre à jour le statut d'une commande

## Webhooks

### Configuration

Les webhooks permettent de recevoir des notifications en temps réel lors d'événements spécifiques.

#### Créer un webhook

```http
POST /webhooks
Content-Type: application/json

{
  "url": "https://votre-app.com/webhook",
  "events": ["order.created", "order.updated", "payment.completed"],
  "secret": "votre-secret-optionnel"
}
```

### Événements disponibles

| Événement | Description | Payload |
|-----------|-------------|---------|
| `user.created` | Nouvel utilisateur | Objet utilisateur |
| `user.updated` | Utilisateur modifié | Objet utilisateur |
| `order.created` | Nouvelle commande | Objet commande |
| `order.updated` | Commande modifiée | Objet commande |
| `payment.completed` | Paiement réussi | Objet paiement |
| `payment.failed` | Paiement échoué | Objet paiement |

### Format des webhooks

```json
{
  "id": "wh_1234567890",
  "event": "order.created",
  "timestamp": "2024-03-15T10:30:00Z",
  "data": {
    // Objet concerné par l'événement
  },
  "signature": "sha256=..."
}
```

### Vérification de signature

```javascript
// Exemple Node.js
const crypto = require('crypto');

function verifyWebhookSignature(payload, signature, secret) {
  const computedSignature = crypto
    .createHmac('sha256', secret)
    .update(payload)
    .digest('hex');

  return `sha256=${computedSignature}` === signature;
}
```

## Rate Limiting

### Limites par défaut

| Niveau | Limite | Fenêtre | Description |
|--------|--------|---------|-------------|
| **Anonyme** | 100 req | 1 heure | Requêtes sans authentification |
| **Utilisateur** | 1000 req | 1 heure | Utilisateurs authentifiés |
| **Premium** | 5000 req | 1 heure | Comptes premium |
| **API Key** | 10000 req | 1 heure | Intégrations API |

### En-têtes de réponse

```http
X-RateLimit-Limit: 1000
X-RateLimit-Remaining: 999
X-RateLimit-Reset: 1640995200
X-RateLimit-RetryAfter: 3600
```

### Gestion des limites dépassées

**Réponse 429 :**
```json
{
  "success": false,
  "message": "Limite de taux dépassée",
  "errors": [
    {
      "code": "RATE_LIMIT_EXCEEDED",
      "message": "Trop de requêtes. Réessayez dans 3600 secondes."
    }
  ],
  "meta": {
    "retryAfter": 3600
  }
}
```

## Gestion des erreurs

### Types d'erreurs

#### Erreurs de validation (400)
```json
{
  "success": false,
  "message": "Erreurs de validation",
  "errors": [
    {
      "field": "email",
      "code": "REQUIRED",
      "message": "L'email est requis"
    },
    {
      "field": "password",
      "code": "MIN_LENGTH",
      "message": "Le mot de passe doit contenir au moins 8 caractères"
    }
  ]
}
```

#### Erreurs d'autorisation (403)
```json
{
  "success": false,
  "message": "Accès refusé",
  "errors": [
    {
      "code": "INSUFFICIENT_PERMISSIONS",
      "message": "Vous n'avez pas les permissions nécessaires"
    }
  ]
}
```

#### Erreurs de ressource (404)
```json
{
  "success": false,
  "message": "Ressource non trouvée",
  "errors": [
    {
      "code": "RESOURCE_NOT_FOUND",
      "message": "L'utilisateur avec l'ID 'xyz' n'existe pas"
    }
  ]
}
```

### Codes d'erreur personnalisés

| Code | Description |
|------|-------------|
| `VALIDATION_ERROR` | Erreur de validation des données |
| `DUPLICATE_RESOURCE` | Ressource déjà existante |
| `RESOURCE_NOT_FOUND` | Ressource non trouvée |
| `INSUFFICIENT_PERMISSIONS` | Permissions insuffisantes |
| `EXPIRED_TOKEN` | Token expiré |
| `INVALID_CREDENTIALS` | Identifiants invalides |
| `RATE_LIMIT_EXCEEDED` | Limite de taux dépassée |
| `MAINTENANCE_MODE` | Mode maintenance actif |

## Exemples d'utilisation

### JavaScript (Fetch API)

```javascript
// Configuration de base
const API_BASE = 'https://api.example.com/v1';
const token = localStorage.getItem('authToken');

// Headers par défaut
const defaultHeaders = {
  'Content-Type': 'application/json',
  'Authorization': `Bearer ${token}`
};

// Récupérer des utilisateurs
async function getUsers(page = 1, limit = 20) {
  try {
    const response = await fetch(
      `${API_BASE}/users?page=${page}&limit=${limit}`,
      { headers: defaultHeaders }
    );

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }

    const data = await response.json();
    return data;
  } catch (error) {
    console.error('Erreur lors de la récupération des utilisateurs:', error);
    throw error;
  }
}

// Créer un utilisateur
async function createUser(userData) {
  try {
    const response = await fetch(`${API_BASE}/users`, {
      method: 'POST',
      headers: defaultHeaders,
      body: JSON.stringify(userData)
    });

    const data = await response.json();

    if (!response.ok) {
      throw new Error(data.message || 'Erreur lors de la création');
    }

    return data;
  } catch (error) {
    console.error('Erreur lors de la création de l\'utilisateur:', error);
    throw error;
  }
}
```

### Python (Requests)

```python
import requests
import json

class APIClient:
    def __init__(self, base_url, token=None):
        self.base_url = base_url
        self.session = requests.Session()
        if token:
            self.session.headers.update({
                'Authorization': f'Bearer {token}',
                'Content-Type': 'application/json'
            })

    def get_users(self, page=1, limit=20):
        """Récupérer la liste des utilisateurs"""
        params = {'page': page, 'limit': limit}
        response = self.session.get(f'{self.base_url}/users', params=params)
        response.raise_for_status()
        return response.json()

    def create_user(self, user_data):
        """Créer un nouvel utilisateur"""
        response = self.session.post(
            f'{self.base_url}/users',
            json=user_data
        )
        response.raise_for_status()
        return response.json()

# Utilisation
client = APIClient('https://api.example.com/v1', token='your-token')
users = client.get_users(page=1, limit=10)
```

### cURL

```bash
# Authentification
curl -X POST https://api.example.com/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "password123"
  }'

# Récupérer des utilisateurs
curl -X GET "https://api.example.com/v1/users?page=1&limit=10" \
  -H "Authorization: Bearer YOUR_TOKEN"

# Créer un utilisateur
curl -X POST https://api.example.com/v1/users \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "email": "newuser@example.com",
    "firstName": "John",
    "lastName": "Doe",
    "password": "securePassword123"
  }'
```

## Tests et environnement de développement

### Collection Postman

Une collection Postman est disponible pour tester l'API :
- **URL** : [Lien vers collection Postman]
- **Variables d'environnement** : Configurées pour dev/staging/prod

### Environnement de test

```bash
# Variables d'environnement pour tests
API_BASE_URL=https://api-dev.example.com/v1
API_TOKEN=test_token_1234567890

# Commandes de test
npm run test:api
npm run test:integration
```

### Mocking pour développement

```javascript
// Mock de l'API pour développement frontend
const mockAPI = {
  users: [
    { id: '1', email: 'user1@example.com', firstName: 'John', lastName: 'Doe' },
    { id: '2', email: 'user2@example.com', firstName: 'Jane', lastName: 'Smith' }
  ]
};

// Intercepter les requêtes en développement
if (process.env.NODE_ENV === 'development') {
  // Configuration MSW ou autre outil de mocking
}
```

## Support et ressources

### Contact

- **Support technique** : [support@votre-domaine.com]
- **Documentation** : [https://docs.votre-domaine.com]
- **Status page** : [https://status.votre-domaine.com]

### Ressources utiles

- **Collection Postman** : [Lien vers collection]
- **SDK JavaScript** : [Lien vers package npm]
- **SDK Python** : [Lien vers package pip]
- **Exemples de code** : [Lien vers repository GitHub]

### Changelog

#### v1.0.0 (2024-03-15)
- Version initiale de l'API
- Endpoints de base pour utilisateurs, produits, commandes
- Authentification JWT
- Rate limiting

---

**Dernière mise à jour** : [Date]
**Version de l'API** : v1.0.0
**Version de la documentation** : 1.0
