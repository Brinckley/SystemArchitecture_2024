## Описание
В данной лабороторной представлены два новых (обновленных) сервиса: `message_service` и `post_service` необходимых для взаимодействия с сообщениями и постами соответственно.   
В `docker-compose.yaml` файл была добавлена конфигурация MondoDb. Файл инициализации лежит в папке `scripts`. В нем настраивается репликация, добавляются база данных и коллекции, индексы.  

## Запуск
Изначально нужно запустить два контейнера: `mongo-setup`. Он нужен для запуска всех нод mongo и сервиса-healthcheck.  
```
docker-compose mongo-setup up -d
```
После того, как Mongo поднялась, можно запускать сервисы. В следующей последовательности:   
```
docker-compose user_service up -d
```
Этими командами поднялись сервисы, отвечающие за сообщения и посты, а также главный сервис для распределения запросов от пользователя.  
В данной версии отсутствует взаимодействие с Postgres, но данные функции можно добавить, подняв базу и `account_service`. Для этого нужно раскомментировать строчки в зависимостях depends_on в `user_service` и `account_service`. Тогда можно заново запустить  user_service следующей командой и поднимется еще и СУБД.  
```
docker-compose user_service up -d
```

## Тестирование
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