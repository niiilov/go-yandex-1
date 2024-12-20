# Калькулятор на Go
Веб-сервис для вычисления математических выражений

## Запуск
Для запсука веб-сервиса проделайте следующие операции:
1. Установите [Go](https://go.dev/) с официального сайта

2. Переместитесь в необходимую вам директорию и скопируйте репозиторий с помощью команды
    ```
    git clone https://github.com/niiilov/go-yandex-1
    ```

3. Перейдите в директорию проекта командой
    ```
    cd go-yandex-1
    ```

4. Запустите веб-сервис командой
    ```
    go run ./cmd/main.go
    ```

5. После запуска сервис будет доступен по адресу [http://localhost:8090/api/v1/calculate](http://localhost:8090/api/v1/calculate)

## Функционал
Данный веб-сервис принимает POST запрос к эндпоинту `/api/v1/calculate` содержащий математическое выражение и возвращает результат его выполнения.

POST запрос можно отправить с помощью curl, он будет иметь следующую структуру

```
curl --location 'localhost:8090/api/v1/calculate' --header 'Content-Type: application/json' --data '{"expression": "2+2"}'
```

Где `expression` - матечатическое выражение которое необходимо решить

Ответ на запрос придет в JSON формате и будет содержаться в поле `result`

## Правила
При работе с сервисом существуют некоторые ограничения

### Ошибка 405
Возникает если на веб-сервис будет отправлен не POST запрос, например 
```
curl http://localhost:8090/api/v1/calculate
```

### Ошибка 400
Возникает если на веб-сервис будет отправлен POST запрос, body которого сформирован неккоректно
```
curl --location 'localhost:8090/api/v1/calculate' --header 'Content-Type: application/json' --data '{not exspression}'
```

### Ошибка 422
Возникает если выражение составленно неккоректно и его невозможно решить
```
curl --location 'localhost:8090/api/v1/calculate' --header 'Content-Type: application/json' --data '{"expression": "2x++56*t"}'
```
Примеры неккоректного ввода:
- 2(1+1)
- 2+(-1)
- 2-+*3
- 1+1=2

### Ошибка 500
Возникает при непридвиденной панике

## Примеры выполнения

|Запрос|Вывод|Статус|
|:-|:-|:-:|
|`curl --location 'localhost:8090/api/v1/calculate' --header 'Content-Type: application/json' --data '{"expression": "2+2"}'`|`{"result":"4"}`|200|
|`curl http://localhost:8090/api/v1/calculate`|`{"error":"Method Not Allowed"}`|405|
|`curl --location 'localhost:8090/api/v1/calculate' --header 'Content-Type: application/json' --data '{not exspression}'`|`{"error":"Bad Request"}`|400|
|`curl --location 'localhost:8090/api/v1/calculate' --header 'Content-Type: application/json' --data '{"expression": "2x++56*t"}'`|`{"error":"Unprocessable Entity"}`|422|
