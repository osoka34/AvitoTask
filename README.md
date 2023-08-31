# AvitoTask

### Инструкция по запуску:

1. Установить все зависимости из файла go.mod.
2. В папке migrations почистить папку storage или наоборот добавить, в нее будет монтироваться директория /var/lib/postgresql/data из контейнера docker.
3. Запустить докер контейнер с postgres, сompose файл лежит в корне проекта.
4. Изменить в **config/config.yaml** путь до файла, через который будут формироваться csv отчеты о пользователях.
**filePath**: "/Users/alexander/GolandProjects/AvitoTask/logs/data.csv".
5. Скомпилировать и запустить проект из папки cmd.

### Все описания ручек можно найти в swagger

#### Дополнительные задания
1. CSV файл с отчетами  

GET 0.0.0.0:11000/statistics/csv
````json
{
    "date_from":"2023-07",
    "date_to":"2023-08"
}

````

Response 200

````csv
1,Segment Y,true,2023-08-28T00:00:00Z
1,SEG 2,false,2023-08-22T00:00:00Z
1,Segment Z,true,2023-08-28T00:00:00Z
1,Segment X,true,2023-08-28T00:00:00Z
````

2. Автоматическое удаление по ttl
Сделал через горутину, которая сначала считает, кол-во записей в бд, которые можно удалить, потом удаляет. Таймаут на операции поставил небольшой, тк если ставить таймауты от одного ближайшего времени удаления до другого, то можно упустить новые записи в бд.

3. Создание сегментов, в которые автоматически добавляется определенный процент пользователей.
Реализовал через WHERE random() < probability. По сути это не совсем процент, а именно вероятность попадания в сегмент, однако при большом кол-ве пользователей вероятность как раз и будет определять процент пользователей, которые будут добавлены в сегмент.

POST 0.0.0.0:11000/segment/create_auto
````json
{
    "segment_name": "Segment J",
    "probability": 0.5
}
````

Response 200

````json
{
    "success": true
}
````


#### TODO list исправления очевидных проблем
1. Передача одних и тех же структур из хендлера в репо handler -> usecase -> repository. На уровне usecase надо делать валидацию, создавать новую структуру и передавать в репозиторий. Для тестового сойдет, но для запуска в прод надо бы переделать.
2. Добавить контексты с таймаутом для обращения в бд.
3. Отслеживать TTL через брокер сообщений, например кафка, при добавлении пользователя в сегмент с ttl записывать его туда, а потом читать, когда потребуется.
4. Если сервис очень часто спрашивает статистику, делая разные селекты в бд, то есть смысл использовать кеширование, тк SELECT это самая долгая операция. Сделав кеширование мы сократим время выдачи данных с 150-300 мс примерно до 30-50 мс, выигрыш по времени около 80%.

**Задать вопросы можно в тг: @nssqk**


