FROM golang:1.23.2 AS builder

RUN mkdir /app

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum для кэширования зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем остальные файлы приложения
COPY . .
# Собираем приложение
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin ./src/main.go

# Используем минимальный образ для запуска
FROM alpine:latest

# Устанавливаем необходимые библиотеки
RUN apk --no-cache add ca-certificates

# Копируем собранное приложение из предыдущего этапа
# Копируем скомпилированное приложение из этапа сборки
COPY --from=builder /app/bin /app/bin

COPY schema.sql /docker-entrypoint-initdb.d/

# Копируем файл env из директории src
COPY .env .env 

# Указываем переменную окружения для порта
ENV PORT=8080

# Открываем порт
EXPOSE 8080

# Запускаем приложение
CMD ["/app/bin"]
