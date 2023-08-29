# AvitoTask

### Инструкция по запуску:

1. Установить все зависимости из файла go.mod.
2. В папке migrations создать папку storage, в которорую будет монтироваться в директория /var/lib/postgresql/data из контейнера docker
3. Запустить докер контейнер с postgres, сompose файл лежит в корне проекта.
4. Изменить в **config/config.yaml** путь до файла, через который будут формироваться csv отчеты о пользователях.
**filePath**: "/Users/alexander/GolandProjects/AvitoTask/logs/data.csv".
5. Скомпилировать и запустить проект из папки cmd.

### Все описания ручек описаны через swagger
