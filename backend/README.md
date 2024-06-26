## Ключевые изменения по сравнению с API V2
- эндпоинт с сотрудниками был изменен с /users/employee/ на /employee/ т.к users был лишним
- эндпоинт для токенов был изменён с /employee/auth/token/login/ на /token/login/
- изменены многие поля сущностей (в частности везде где было first_name, second_name...)
теперь **fullname** одним полем
- Все POST запросы должны возвращать созданную вами СУЩность
- В конце **всех** ссылок теперь надо писать / для единообразия. Даже если после будут параметры.  
Например:
> sbermeeting.tk/api/v3/employees/?client_id=1&direction_id=1
- Те значения, которые могут быть NULL помечены **NULL***
---
## Авторизация

Получение токена
- /token/login/  
POST запрос
~~~
- username
- password
~~~
Возвращает:
- employee - сотрудник, которому принадлежит токен
- auth_token - сам токен


## Сотрудники
~~~
- fullname
- client_list - список клиентов, с которыми работает сотрудники (см. сотрудники)
- role (см. роли)
- direction (см. направления) - NULL*
~~~
### Эндпоинты:  

- /employees/ 

GET - получить список пользователей  
Query параметры(?):
  - direction_id - получить пользователей по направлению
  - client_id - получить пользователей, которые работают с данным клиентом  

POST - создать пользователя
~~~
- fullname
- username
- password
~~~

- employees/current/ 

GET - Получить текущего пользователя

PUT - Изменить текущего пользователя
~~~
- fullname
- direction (см. направления)
- role (см. роли)
- client_list - массив клиентов (см. клиентов)

Хотя вы можете добавить роль в список полей, сама роль в этом эндпоинте не изменится по соображениям безопасности.
~~~


- /employees/**id**/  

GET - Получить пользователя по id


## Направления 
Направления используются для сотрудников.  
Например - ММБ, Эквайринг.
~~~
- id
- name
~~~

### Эндпоинты:  
Получить направление по id:
- employees/directions/**id**/  

Получить список направлений:
- /employees/directions/  
POST - добавить новое направление
~~~
   - name
~~~ 

 




## Ошибки
> Читайте json с ошибкой, часто он довольно полезен!
~~~
Если запрос отработал с ошибкой, то возвращается json с двумя полями:

    code - код ошибки
    error - текст ошибки
~~~

### Ошибки, возвращаемые с **API**:

- 404 - указанной страницы не существует.  
  
    Возможные варианты:
    - не был добавлен / в конце ссылки
    - ошибка в синтаксисе ссылки  

- 400 - сервер не смог прочитать/распарсить запрос, или же понял но отказывается выполнять по каким-то причинам

  Возможные варианты:
    - ошибка синтаксиса в заголовке / json'e
    - нарушены ограничения модели, например встрече в поле "время начала" передали строку
    - есть шанс что ошибка на стороне сервера, но вероятность крайне мала
    - те же варианты что и в 404 :)

- 401 - ошибка авторизации.

  Возможные варианты:
    - попытка получить доступ к ресурсу, который требует авторизации
    - неправильный логин/пароль при получении токена

- 500 - внутренняя ошибка сервера.  
    Сервер насрал в штаны.  
    Возможно по вашей, а возможно и по моей вине.
    
- 405 - метод не разрешён
  Например у эндпоинта разрешён только GET, а вы пытаетесь сделать POST
## Конфигурация сервера
Необходимо создать json файл с конфигурацией.  
Путь к нему можно задать через флаг config-path

 > configs/config.json - путь по умолчанию

### Поля конфигурационного файла:
- bindAddr -  если не указан то по дефолту ":8080") 
- logLevel - (trace, debug, info, warn, error, fatal, panic)
- dbConnStr -  "user=**user** password=**password** dbname=sberdb sslmode=disable"
- salt - соль, которая нужна для генерации токенов

```
Все поля должны быть типа string
```

### База данных
Используемая CLI для миграции - https://github.com/golang-migrate/migrate  
#### Команда для миграции:
> migrate -path migrations -database "postgres://**user**:**passwd**@localhost/sberdb?sslmode=disable" up

#### Подключение к бд через консоль:
> docker exec -it **container-name** psql -U postgres

Позже напишу скрипт для автоматического деплоя