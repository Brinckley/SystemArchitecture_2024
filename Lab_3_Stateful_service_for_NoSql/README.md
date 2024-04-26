## Описание
!!!! Нужно заранее проверить, что у docker-compose.yaml и скриптов стоит LF в параметрах конца строки.
В данной лабороторной представлены два новых (обновленных) сервиса: `message_service` и `post_service` необходимых для взаимодействия с сообщениями и постами соответственно.   
В `docker-compose.yaml` файл была добавлена конфигурация MondoDb. Файл инициализации лежит в папке `scripts`. В нем настраивается репликация, добавляются база данных и коллекции, индексы.  

## Файлы для заполнения
Программа `insert_maker` на языке Python генерирует заданное значение постов и сообщений в формате `.json`. Сгенерированные файлы необходимо поместить в папку `scripts`. Оттуда они будут взяты из скрипта `mongo_setup.sh`, и данные оттуда будут занесены в базу. Предвратительно уже сгенерированы файлы по 80 000 сущностей сообщений и столько же постов. Id пользователей не привязаны к реальным. Это сгенерированные заглушки, которые в полностью рабочем сервисе берутся из Postgres. Заполнение осуществляется через `mongoimport`.   

## Запуск
Изначально нужно запустить контейнер: `mongo-setup`. Он нужен для запуска всех нод mongo и сервиса-healthcheck. Нужно дождаться, пока не случится выход с кодом 0. Это будет значить, что все ноды поднялись и заполнились.  
```
docker-compose up mongo-setup
```
После того, как Mongo поднялась, можно запускать остальные сервисы, которые подтянутся вслед за главным:   
```
docker-compose up user_service 
```

## Тестирование
Для скринов использовался postman. Более подробно об эндпоинтах можно посмотреть в `index.yaml`.
Добавление нового сообщения  
![add msg](https://github.com/Brinckley/SystemArchitecture_2024/blob/main/Lab_3_Stateful_service_for_NoSql/imgs/il1.jpg)

Получение сообщения по его id  
![get msg](https://github.com/Brinckley/SystemArchitecture_2024/blob/main/Lab_3_Stateful_service_for_NoSql/imgs/il2.jpg)

Получение сообщений для аккаунта с заданным id  
![get msg](https://github.com/Brinckley/SystemArchitecture_2024/blob/main/Lab_3_Stateful_service_for_NoSql/imgs/il3.jpg)

Добавление нового поста  
![add post](https://github.com/Brinckley/SystemArchitecture_2024/blob/main/Lab_3_Stateful_service_for_NoSql/imgs/il4.jpg)

Получение/удаление поста по id 
![get post](https://github.com/Brinckley/SystemArchitecture_2024/blob/main/Lab_3_Stateful_service_for_NoSql/imgs/il5.jpg) 
![delete post](https://github.com/Brinckley/SystemArchitecture_2024/blob/main/Lab_3_Stateful_service_for_NoSql/imgs/il6.jpg)

Получение постов по id аккаунта 
![get post](https://github.com/Brinckley/SystemArchitecture_2024/blob/main/Lab_3_Stateful_service_for_NoSql/imgs/il7.jpg) 
