# Менеджер паролей GophKeeper

GophKeeper представляет собой клиент-серверную систему, 
позволяющую пользователю надёжно и безопасно хранить логины, 
пароли, бинарные данные и прочую приватную информацию.

## Сервер

Сервер реализует следующую бизнес-логику:
- регистрация, аутентификация и авторизация пользователей;
- хранение приватных данных;
- синхронизация данных между несколькими авторизованными клиентами одного владельца;
- передача приватных данных владельцу по запросу.

### Взаимодействие с сервером

Регистрация:
```
curl -k -d '{"email":"user@mail.com", "password":"password"}' \
    -H 'Content-Type: application/json' \
    http://localhost:8080/api/user/register
```
Авторизация:
```
curl -k -d '{"email":"user@mail.com", "password":"password"}' \
    -H 'Content-Type: application/json' \
    http://localhost:8080/api/user/login
```
Статус:
```
curl -H 'Content-Type: application/json' \
    -H "Authorization:TOKEN" \
    http://localhost:8080/api/data/status
```
Пулл:
```
 curl -H 'Content-Type: application/json' \
    -H "Authorization:TOKEN" \
    http://localhost:8080/api/data/pull?id=UUID
```
Пуш:
```
curl -X POST \
    -H "Content-Type: application/json" \
    -H "Authorization:TOKEN" \
    -d '{"id":"UUID", "category":0, "data":"0KLQtdC60YHRgiDRgdC+0L7QsdGJ0LXQvdC40Y8=", "version": 0}' \
    http://localhost:8080/api/data/push
```

### Принцип работы
Задача сервера - сохранять данные, которые ему прислал клиент. Для этого выбран postgres. Адрес хранения - таблица `data`.

| Поле     | Тип      | Комментарий                                           | 
|----------|----------|-------------------------------------------------------| 
| id       | uuid     | идентификатор                                         | 
| user_id  | uuid     | идентификатор пользователя                            | 
| category | smallint | тип данных: логин-пароль, текст, бинарные, банковские | 
| data     | text     | данные                                                | 
| version  | integer  | версия для сравнения с сервером                       | 

Сервер не может расшифровать полученные данные, но может сравнить их с теми, что хранятся у него. Если разницы в данных нет, то никаких изменений в таблице не происходит. В противном случае сервер следит, чтобы версия отправленных данных совпадала с версией на сервере.

Версионирование данных происходит через поле `version`. Например, если клиент отправит "устаревшие" данные с `version` = 1, а на сервере данные обновились до `version` = 2, то клиенту откажут в сохранении. В этом случае ему вначале придется подтянуть к себе изменения с сервера. 

## Клиент

Клиент реализует следующую бизнес-логику:
- аутентификация и авторизация пользователей на удалённом сервере;
- доступ к приватным данным по запросу.

Сборка клиента с флагами:
```shell
go build -ldflags "-X main.Version=1.0.0 -X main.BuildDate=2024-04-18"
```

### Поддерживаемые команды
```
Авторизация
./client register -u user@email.com -p 12345678
./client login -u user@email.com -p 12345678

Добавление записей
./client create credentials -l administrator -p password
./client create text -t "My private text"
./client create binary -b "SGVsbG8sIHdvcmxkIQ=="
./client create card -n 1234123412341234 -o "Name Surname" -c 000

Список записей 
./client list

Обновление записей
./client update credentials -i 44e221ed-ceee-4cca-b4bc-431512f7b8ee -l administrator -p password
./client update text -i 4ad45bb6-feaa-494d-9c00-163bc8a36bc1 -t "New text"
./client update binary -i eb9a7864-d81e-4599-b9a2-0dba6e4e4b0c -b "111SGVsbG8sIHdvcmxkIQ=="
./client update card -i b1253183-f7c0-4020-9ab2-cc2fb3b798ac -n 4321432143214321 -o "Ivan Ivanov" -c 123

Удаление записей
./client remove -i eb9a7864-d81e-4599-b9a2-0dba6e4e4b0c

Синхронизация с сервером
./client status
./client push -i 44e221ed-ceee-4cca-b4bc-431512f7b8ee
./client pull -i 44e221ed-ceee-4cca-b4bc-431512f7b8ee
```

## Хранение данных
Для хранения данных на стороне клиента выбран sqlite (файл "data.db"). Все записи там хранятся в таблице `data`.

| Поле     | Тип      | Комментарий                                           |
|----------|----------|-------------------------------------------------------|
| id       | uuid     | идентификатор                                         |
| category | smallint | тип данных: логин-пароль, текст, бинарные, банковские |
| data     | text     | данные                                                |
| version  | integer  | версия для сравнения с сервером                       |

Данные в этой таблице хранятся в открытом (незашифрованном) виде. Однако перед их отправкой на сервер, они проходят шифровку ключом пользователя, который он, предположительно, носит с собой. В проекте этот ключ для упрощения хранится в `cmd/client/secret.txt`.

### Принцип пользования
С клиентским приложением можно работать "офлайн": добавлять новые записи, редактировать их и удалять. Но для того, чтобы взаимодействовать с сервером, пользователю нужно пройти аутентификацию. После этого можно вызывать команды `status`, `push` и `pull`. 