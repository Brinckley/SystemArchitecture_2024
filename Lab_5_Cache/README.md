## Описание
Изменения с добавлением шаблоны «сквозное чтение» и «сквозная запись» коснулись `gateway_service`, бывший `user_service`. 
Основная логика реализована вот здесь: `internal/server/router/account_handler.go#getAccount` и `internal/server/router/auth_handler.go#signUpAccount`. Кэширование происходит при чтении аккаунта по Id и при первичной регистрации (т.е. добавлении нового аккаунта). Остальное лежит в папках redis и некоторых других. Изменять другие сервисы не пришлось. TTL сущности account в кэше - 60 секунд.  

## Эндпоинты
Авторизация: 
 - /signup - Post регистрация пользователя, получение id  
 - /signin - Post вход по логину и паролю, получение токена в header "Auth-token"  

Требующие авторизации действия:
 - /account - Put/Delete запросы к аккаунту (id берется из токена)
 - /messages/msg/{message_id} - Get получить доступ может только получатель или отправитель
 - /messages/account - Get получить все сообщения для пользователя (id указано как receiver)
 - /messages - Post пользователь отправляет сообщение
 - /posts - Post сделать пост может только авторизованный пользователь
 - /posts/{post_id} - Put/Delete может сделать только авторизированный пользователь, обладатель поста  
   
Не требующие авторизации действия:
 - /account/{account_id} - Get получить аккаунт по id
 - /accounts - Get получить список имеющихся аккаунтов
 - /account/search - Get поиск по маске
 - /posts/account/{account_id} - Get посмотреть посты аккаунта
 - /posts/{post_id} - Get посмотреть пост по его id  

## Запуск
В docker-compose.yaml представлены все необходимые сервисы. 
Сначала нужно собрать mongo, для этого запускаем mongo-setup. И ждем, пока контейнер запустится и завершится с кодом 0. Это будет означать, что все реплики подняты и скрипт с инициализацией кластера сработал.
```
docker-compose up mongo-setup
```
Последний шаг - запуск gateway_service, который подтянет все остальные сервисы, а именно Redis, Postgres, account_service, message_service, post_service.  
```
docker compose up gateway_service
```
Теперь можно использовать сервис. Примеры предложены далее.  

## Тестирование
Добавление пользовтеля в базу и получение его id на этапе signup
![add msg](https://github.com/Brinckley/SystemArchitecture_2024/blob/main/Lab_5_Cache/imgs/il1.png)

Получение пользователя из базы. В первый раз он достается из СУБД, в другие разы в пределах 60 секунд - из кэша. Логика соответствует сквозному чтению.
![add msg](https://github.com/Brinckley/SystemArchitecture_2024/blob/main/Lab_5_Cache/imgs/il2.png)

