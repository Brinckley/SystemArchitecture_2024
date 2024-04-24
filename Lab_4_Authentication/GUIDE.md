## Описание
В данной лабороторной приложение `user_service` отвечает за работу с пользователями и распределяет запросы по другим микросервисам.  
Генерация JWT токенов происходит в сервисе `account_service`. Там написана логика работы signin/signup. Signup просто добавляет пользователя в базу, а signin выдает токен, который возращается пользователю на фронт. Там его можно вставить в поле Bearer token для получения доступа к командам, выполнение которых требует авторизации.  
Сама авторизация в `user_service` проверяется в виде декоратора. Из токена вытаскивается user_id, который затем отправляется в соответствующий сервис `account_service`, `message_service` или `post_service`. 

## Эндпоинты
Требующие авторизации действия:
 - /account - Get/Put/Delete запросы к аккаунту (id берется из токена)
 - /messages/msg/{message_id} - Get получить доступ может только получатель или отправитель
 - /messages/account - Get получить все сообщения для пользователя (id указано как receiver)
 - /messages - Post пользователь отправляет сообщение
 - /posts - Post сделать пост может только авторизованный пользователь
 - /posts/{post_id} - Put/Delete может сделать только авторизированный пользователь, обладатель поста  
   
Не требующие авторизации действия:
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
Теперь поднимаем CУБД. Она будет работать в фоне.
```
docker-compose up -d postgres
```
Последний шаг - запуск user_service, который подтянет все остальные сервисы.
```
docker-compose up user_service
```
Теперь можно использовать сервис. Примеры предложены далее.
## Тестирование
Для скринов использовался postman. Более подробно об эндпоинтах можно посмотреть в `index.yaml`.

SignUp
![add msg](https://github.com/Brinckley/SystemArchitecture_2024/blob/main/Lab_4_Authentication/imgs/il1.png)

SignIn
![add msg](https://github.com/Brinckley/SystemArchitecture_2024/blob/main/Lab_4_Authentication/imgs/il2.png)

Getting accounts without auth
![add msg](https://github.com/Brinckley/SystemArchitecture_2024/blob/main/Lab_4_Authentication/imgs/il12.png)

Message sending with Auth
![add msg](https://github.com/Brinckley/SystemArchitecture_2024/blob/main/Lab_4_Authentication/imgs/il3.png)

Getting message by Id with auth + examples with wrong auth
![add msg](https://github.com/Brinckley/SystemArchitecture_2024/blob/main/Lab_4_Authentication/imgs/il4.png)
![add msg](https://github.com/Brinckley/SystemArchitecture_2024/blob/main/Lab_4_Authentication/imgs/il5.png)
![add msg](https://github.com/Brinckley/SystemArchitecture_2024/blob/main/Lab_4_Authentication/imgs/il6.png)

Getting all received messages
![add msg](https://github.com/Brinckley/SystemArchitecture_2024/blob/main/Lab_4_Authentication/imgs/il7.png)

Manipulations with posts and auth
![add msg](https://github.com/Brinckley/SystemArchitecture_2024/blob/main/Lab_4_Authentication/imgs/il8.png)
![add msg](https://github.com/Brinckley/SystemArchitecture_2024/blob/main/Lab_4_Authentication/imgs/il9.png)
![add msg](https://github.com/Brinckley/SystemArchitecture_2024/blob/main/Lab_4_Authentication/imgs/il10.png)
![add msg](https://github.com/Brinckley/SystemArchitecture_2024/blob/main/Lab_4_Authentication/imgs/il11.png)