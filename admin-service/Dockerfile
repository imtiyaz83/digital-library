# admin-service/Dockerfile
FROM php:8.2-apache
RUN apt-get update && apt-get install -y \
    git zip unzip curl libzip-dev libonig-dev && \
    docker-php-ext-install pdo pdo_mysql zip
RUN a2enmod rewrite
WORKDIR /var/www/html
COPY --from=composer:latest /usr/bin/composer /usr/bin/composer
COPY . .
RUN composer install
RUN chown -R www-data:www-data /var/www/html && \
    [ -d /var/www/html/storage ] && chmod -R 755 /var/www/html/storage || true

EXPOSE 80
CMD ["apache2-foreground"]