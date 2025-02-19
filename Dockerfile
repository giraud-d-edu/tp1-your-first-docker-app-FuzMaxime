# Utilisation de l'image officielle Nginx
FROM nginx:alpine

# Copie des fichiers du site
COPY src/ /usr/share/nginx/html/

# Configuration du port d'écoute
EXPOSE 80

# Le conteneur démarre automatiquement Nginx