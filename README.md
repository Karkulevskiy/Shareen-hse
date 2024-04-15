Зависимости для запуска (go get [имя зависимости]):
- github.com/spf13/viper
- github.com/google/uuid
- github.com/gin-gonic/gin
- github.com/swaggo/swag/cmd/swag
- github.com/swaggo/gin-swagger
- github.com/swaggo/files
- github.com/lib/pq
Также: go install github.com/swaggo/swag/cmd/swag@latest
В папке dbSchema находится SQL файл с запросами на создание БД. Нужно ручками в постгресе создать бд, скопировав код из конфига. Позже добавлю автомиграции. 

Swagger генерируется автоматически по документации к api, поэтому стоит перед тестированием контроллеров прописать:
swag init -g .\internal\server\httpServer.go -o .\internal\docs 

Для запуска проекта:
go run ./cmd/main.go

Потом добавлю Makefile, чтобы ничего такого не прописывать
