# Utilisation de l'image officielle Golang
FROM golang:alpine

# Définir le répertoire de travail dans le conteneur
WORKDIR /app

# Copier les fichiers go.mod et go.sum dans le conteneur
COPY scripts/go.mod /app/scripts/

# Installer les dépendances
WORKDIR /app/scripts
RUN go mod tidy

# Copier le reste des fichiers du projet dans le conteneur
COPY . /app

# Construire l'application
RUN go build -o /app/scripts/main.go

# Définir le répertoire de travail final
WORKDIR /app

# Exposer le port sur lequel l'application s'exécute
EXPOSE 8080

# Donner les permissions d'exécution au fichier binaire
RUN chmod +x /app/scripts/main.go

# Commande pour exécuter l'application
CMD ["/app/scripts/main.go"]