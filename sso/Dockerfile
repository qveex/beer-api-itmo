# Используем образ golang для сборки приложения
FROM golang:1.22.2

# Обновляем список пакетов и устанавливаем netcat
RUN apt-get update && apt-get install -y netcat-openbsd postgresql-client

# Копируем все файлы внутрь контейнера
COPY . /sso

# Задаем рабочую директорию
WORKDIR /sso

# Запускаем go mod tidy
RUN go mod tidy

## Собираем приложение
#RUN go build -o sso ./cmd/sso

# Добавляем и делаем исполняемым entrypoint скрипт
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

# Устанавливаем скрипт в качестве entrypoint
ENTRYPOINT ["/entrypoint.sh"]

# Команда по умолчанию - запуск приложения
CMD ["./sso", "--config=config/config.yaml"]