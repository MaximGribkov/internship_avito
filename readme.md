# Тестовое задание на  позицию стажера бекэнд разработчика.

Микросервис для динамического сегментирования пользователей

Реализованные задачи:
* Создание пользователя.
* Создание сегмента.


Немного команд примененных в ходе освоения docker.

Запуск docker-compose:
`docker-compose up app`

Остановка контейнера: 
`docker stop <name>`

Запуск миграций: 
`migrate -path ./schema -database 'postgres://avito:avito@localhost:5433/avito?sslmode=disable' up`

Вход в контейнер:
`docker exec -it 3eae8f23ed61 /bin/bash`

Открытие списка запущенных контейнеров:
`docker ps`

Сборка образа:
`docker build -t docker-avito .`

Подключение к терминалу Postgres:
`psql -U avito`

Примеры запросов:

(POST) user/create Создание пользователя
![img.png](img/img.png)

(POST) segments/create Создание сегмента
![img.png](img/img2.png)

(DELETE) segments/delete Удаление сегмента 
![img.png](img/img3.png)