# AvitoTask

### Инструкция по запуску:

1. Установить все зависимости из файла go.mod.
2. В папке migrations почистить папку storage, в которорую будет монтироваться в директория /var/lib/postgresql/data из контейнера docker
3. Запустить докер контейнер с postgres, сompose файл лежит в корне проекта.
4. Изменить в **config/config.yaml** путь до файла, через который будут формироваться csv отчеты о пользователях.
**filePath**: "/Users/alexander/GolandProjects/AvitoTask/logs/data.csv".
5. Скомпилировать и запустить проект из папки cmd.

### Все описания ручек можно найти в swagger

#### TODO list исправления очевидных проблем
1. Передача одних и тех же структур из хендлера в репо handler -> usecase -> repository. На уровне usecase надо делать валидацию, создавать новую структуру и передавать в репозиторий. Для тестового сойдет, но для запуска в прод надо бы переделать.
2. Отслеживать TTL через брокер сообщений, например кафка, при создании пользователя с ttl записывать его туда, а потом читать, когда потребуется.
3. Если сервис очень часто спрашивает статистику, делая разные селекты в бд, то есть смысл использовать кеширование, тк SELECT это самая долгая операция. Сделав кеширование мы сократим время выдачи данных с 150-300 мс примерно до 30-50 мс, выигрыш по времени около 80%.

**Задать вопросы можно в тг: @nssqk**


