FROM nginx:latest

WORKDIR /var/www/html

RUN apt-get update \
    && apt-get install -y nginx-module-njs

COPY ./oauth2.js /etc/nginx/conf.d/oauth2.js
