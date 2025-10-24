# Go CRUD Starter - API REST

API REST Go professionnelle avec Gin framework, architecture propre et performances optimales.

## 🚀 Fonctionnalités

- **Go + Gin** : Performance exceptionnelle et simplicité
- **Architecture propre** : Handlers/Repository/Models séparés
- **CRUD complet** : Créer, Lire, Mettre à jour et Supprimer des utilisateurs
- **Pagination** : Support de page/limit pour les listes
- **Recherche** : Recherche d'utilisateurs par nom ou email
- **Validation** : Validation avec gin binding tags
- **Rate Limiting** : Middleware personnalisé de limitation
- **CORS** : Support des requêtes cross-origin
- **Graceful shutdown** : Arrêt propre du serveur
- **Pool de connexions** : Gestion optimisée PostgreSQL

## 📋 Prérequis

- Go 1.21+
- PostgreSQL (hébergé en ligne)

## 🔧 Installation

1. Cloner le projet
2. Télécharger les dépendances :
```bash
go mod download
```

3. Configurer les variables d'environnement :
```bash
cp .env.example .env
```

4. Modifier le fichier `.env` :
```
DATABASE_URL=postgresql://user:password@host:port/database
PORT=8000
GIN_MODE=release
```

5. Compiler l'application :
```bash
go build -o app .
```

## 🏃 Démarrage

### Développement
```bash
go run main.go
```

### Production
```bash
go build -o app .
./app
```

## 📡 API Endpoints

### Health Check
```http
GET /health
```

**Réponse** :
```json
{
  "status": "healthy",
  "database": "connected",
  "service": "go-crud-api"
}
```

### Créer un utilisateur
```http
POST /api/users
Content-Type: application/json

{
  "name": "John Doe",
  "email": "john@example.com",
  "phone": "+33612345678"
}
```

**Réponse** :
```json
{
  "message": "Utilisateur créé avec succès",
  "data": {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com",
    "phone": "+33612345678",
    "created_at": "2025-01-01T00:00:00Z",
    "updated_at": "2025-01-01T00:00:00Z"
  }
}
```

### Récupérer tous les utilisateurs (avec pagination)
```http
GET /api/users?page=1&limit=20
```

**Paramètres** :
- `page` : Numéro de page (défaut 1)
- `limit` : Nombre d'utilisateurs (défaut 20, max 100)

**Réponse** :
```json
{
  "data": [...],
  "total": 150,
  "page": 1,
  "limit": 20,
  "total_pages": 8
}
```

### Rechercher des utilisateurs
```http
GET /api/users/search?q=john
```

**Réponse** :
```json
{
  "results": [...],
  "count": 5
}
```

### Récupérer un utilisateur par ID
```http
GET /api/users/1
```

### Mettre à jour un utilisateur
```http
PUT /api/users/1
Content-Type: application/json

{
  "name": "Jane Doe",
  "email": "jane@example.com",
  "phone": "+33698765432"
}
```

### Supprimer un utilisateur
```http
DELETE /api/users/1
```

## 📦 Structure du projet

```
go-crud-starter/
├── config/
│   └── database.go          # Configuration PostgreSQL
├── handlers/
│   └── user_handler.go      # Handlers HTTP
├── middleware/
│   └── rate_limiter.go      # Rate limiting personnalisé
├── models/
│   └── user.go              # Structures de données
├── repository/
│   └── user_repository.go   # Accès aux données
├── routes/
│   └── routes.go            # Configuration des routes
├── main.go                  # Point d'entrée
├── go.mod                   # Dépendances Go
├── .env.example            # Template des variables
├── .gitignore              # Fichiers à ignorer
└── README.md               # Documentation
```

## 🛡️ Sécurité et Performance

### Performance
- **Compilation native** : Binaire optimisé et ultra-rapide
- **Concurrence** : Goroutines pour les opérations asynchrones
- **Pool de connexions** : Max 25 connexions, 5 idle
- **Index DB** : Index sur la colonne email
- **Gin framework** : Un des frameworks les plus rapides

### Sécurité
- **CORS** : Configuration cross-origin
- **Rate Limiting** : 100 requêtes/minute par IP
- **Validation** : Validation stricte avec binding tags
- **Prepared statements** : Protection contre les injections SQL
- **Error handling** : Gestion propre des erreurs

### Validation des données
- **Nom** : 1-255 caractères, requis
- **Email** : Format email valide, unique, requis
- **Téléphone** : Max 50 caractères, optionnel
- **ID** : Entier positif
- **Pagination** : page ≥ 1, limit 1-100

## 🎯 Gestion des erreurs

L'API retourne des réponses JSON structurées :

- **400** : Erreur de validation
- **404** : Ressource non trouvée
- **409** : Conflit (email déjà existant)
- **429** : Trop de requêtes (rate limit)
- **500** : Erreur serveur

**Exemple d'erreur** :
```json
{
  "error": "Un utilisateur avec cet email existe déjà"
}
```

## 📝 Commandes utiles

```bash
go mod download          # Télécharger les dépendances
go build -o app .        # Compiler l'application
go run main.go           # Lancer en mode développement
go test ./...            # Lancer les tests
go fmt ./...             # Formater le code
```

## 🌐 Déploiement

Configuration pour votre plateforme :
- **Build command** : `go mod download && go build -o app .`
- **Start command** : `./app`

Variables d'environnement requises :
- `DATABASE_URL` : URL de connexion PostgreSQL
- `PORT` : Port d'écoute (optionnel, défaut 8000)
- `GIN_MODE` : Mode Gin (release/debug)

## 📝 Best Practices

- **Architecture propre** : Séparation claire des responsabilités
- **Repository pattern** : Isolation de la logique d'accès aux données
- **Validation** : Binding tags pour validation automatique
- **Error handling** : Gestion cohérente des erreurs
- **Concurrence** : Utilisation des goroutines
- **Graceful shutdown** : Arrêt propre avec signaux
- **Pool de connexions** : Réutilisation des connexions DB
- **Rate limiting** : Protection contre les abus
- **CORS** : Support des applications front-end

## 🔍 Avantages Go

### Performance
- **Compilation native** : Binaire ultra-rapide
- **Concurrence** : Goroutines légères et efficaces
- **Faible empreinte mémoire** : Consommation minimale
- **Démarrage instantané** : Pas de warm-up
- **Garbage collector** : Optimisé pour faible latence

### Développement
- **Simplicité** : Syntaxe claire et concise
- **Typage statique** : Détection des erreurs à la compilation
- **Tooling** : Outils intégrés (fmt, test, doc)
- **Déploiement** : Un seul binaire, pas de dépendances
- **Cross-compilation** : Compilation pour toutes les plateformes

### Écosystème
- **Gin** : Framework web ultra-rapide
- **Standard library** : Riche et complète
- **Communauté** : Active et croissante
- **Production-ready** : Utilisé par Google, Uber, Docker

## 🆚 Différences avec les autres projets

Go offre des avantages uniques :
- **Performance** : 10-100x plus rapide que Python/Node.js
- **Binaire unique** : Pas de runtime à installer
- **Concurrence native** : Goroutines vs threads/async
- **Faible latence** : Idéal pour les APIs haute performance
- **Déploiement simple** : Un fichier, pas de dépendances
- **Consommation mémoire** : Très faible vs autres langages
- **Compilation rapide** : Feedback instantané

## 🔧 Développement

### Ajouter une nouvelle route
1. Créer la structure dans `models/`
2. Ajouter les méthodes dans `repository/`
3. Créer le handler dans `handlers/`
4. Définir la route dans `routes/routes.go`

### Tester l'API
Utilisez le fichier `test_api.http` avec l'extension REST Client de VS Code.

### Hot reload en développement
```bash
go install github.com/cosmtrek/air@latest
air
```

## 🚀 Optimisations

### Build optimisé pour production
```bash
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-w -s" -o app .
```

Flags :
- `-w` : Supprime les infos de debug
- `-s` : Supprime la table des symboles
- `CGO_ENABLED=0` : Binaire statique

### Docker (optionnel)
```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -o app .

FROM alpine:latest
COPY --from=builder /app/app /app
CMD ["/app"]
```
