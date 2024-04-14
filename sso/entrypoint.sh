#!/bin/bash
# Ожидаем, пока PostgreSQL станет полностью доступной
echo "Ожидание доступности PostgreSQL..."
until pg_isready -h db -p 5432 -U postgres; do
echo "PostgreSQL еще не готова..."
sleep 1
done
echo "PostgreSQL доступна."

# Выполнение миграций
echo "Запуск миграций..."
go run ./cmd/migrator --db-url="postgresql://postgres:root@db:5432/auth?sslmode=disable" --migrations-path=./migrations

# Если миграции прошли успешно, запускаем приложение
echo "Запуск приложения..."
exec go run ./cmd/sso --config=config/config.yaml