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