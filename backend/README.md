## Конфигурация сервера
Необходимо создать json файл с конфигурацией.  
Путь к нему можно задать через флаг config-path

 > configs/config.json - путь по умолчанию

### Поля конфигурационного файла:
- bindAddr -  если не указан то по дефолту ":6969") 
- logLevel - (trace, debug, info, warn, error, fatal, panic)
- dbConnStr -  "user=**user** password=**password** dbname=sberdb sslmode=disable"

```
Все поля должны быть типа string
```

### База данных
Используемая CLI для миграции - https://github.com/golang-migrate/migrate  
#### Команда для миграции:
> migrate -path migrations -database "postgres://**user**:**passwd**@localhost/sberdb?sslmode=disable" up

Позже напишу скрипт для автоматического деплоя