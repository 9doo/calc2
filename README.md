## Запуск

1. Установите [Go](https://go.dev/dl/).
2. Установите [Git](https://git-scm.com/downloads).
3. Склонируйте проект с GitHub используя командную строку:
```
    git clone https://github.com/9doo/calc2
```
4. Перейдите в папку проекта, выполните команду:
```
    cd calc2
```
5. Выполните команду:
```
    go mod tidy
```
7. Запустите сервер:
```
    go run ./cmd/calc_service/main.go
```
7. Сервис будет доступен по адресу: ```http://localhost:8080/api/v1/calculate```

### Как сменить порт (для Windiws)?
1. Для этого нужно собрать calc.exe 
2. Перейдите в папку (calc_go/cmd) проекта
3. выполните команды

```go build -o calc.exe``` 

```set "PORT=8087" & "calc.exe"``` 
В примере порт = 8087

Затем откройте новое окно «Git Bash» и снова пропишите путь, используя команду:
```
    cd calc2
```
### Примеры использования:

Успешный запрос:

```bash
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '
{
  "expression": "2*2+2"
}'
```

Ответ:

```bash
{
  "id": "..."
}
```
Затем можно увидеть, на каком этапе выполнения находится данный запрос, и его результат:
```bash
curl --location 'http://localhost:8080/api/v1/expressions'
```
Вывод:
```bash
{"expressions":[{"id":"1740240110508066400","status":"pending"}]}
```
Или получить точное значение необходимого выражения, указав его "id":
```bash
curl --location 'http://localhost:8080/api/v1/expressions/:id'
```

Ошибки при запросах:

Ошибка **404** (отсутствие выражения):
```bash
{"error":"Expression not found"}
```

Ошибка **422** (невалидное выражение):

```bash
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '
{
  "expression": "2+a"
}'
```
Ответ:

```bash
{
  "error": "Expression is not valid"
}
```

Ошибка **500** (внутренняя ошибка сервера):

```bash
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '
{
  "expression": "2/0"
}'
```
Ответ:

```bash
{
  "error": "Internal server error"
}
```

Также можно запустить тесты для evaluator:

```bash
go test ./internal/evaluator
```
В случае успешного прохождения теста будет выдан следующий результат:

```bash
ok  	calc_service/internal/evaluator	0.001s
```
