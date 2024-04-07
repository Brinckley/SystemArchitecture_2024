## Старт
Все контейнеры стартуют вместе, но, благодаря зависимостям (depends on), старт имеет очередность.   
data_loader_http - программа на языке python, наполняющая базу случайно сгенерированными и связанными между собой сущностями через написанный API. Данное наполнение подходит для тестирование проекта, так как в случае репликации или хэширования пароля, данные будут корректно обрабатываться.  
insert_maker - программа на языке python для записи напрямую в базу. Сначала генерируется файл в формате json, затем из него данные читаются и отправляются в БД. Запускается вместе со всеми контейнерами с командой docker-compose up.


## Получить все аккаунты
![Get all accounts](https://github.com/Brinckley/SystemArchitecture_2024/blob/main/Lab_2_Stateful_service_for_RDBMS/imgs/il1.jpg)
## Добавить новый аккаунт
![Add accounts](https://github.com/Brinckley/SystemArchitecture_2024/blob/main/Lab_2_Stateful_service_for_RDBMS/imgs/il2.jpg)
## Получить аккаунт по номеру
![Get account](https://github.com/Brinckley/SystemArchitecture_2024/blob/main/Lab_2_Stateful_service_for_RDBMS/imgs/il3.jpg)
![No account](https://github.com/Brinckley/SystemArchitecture_2024/blob/main/Lab_2_Stateful_service_for_RDBMS/imgs/il4.jpg)
## Обновить аккаунт по номеру
![Get account](https://github.com/Brinckley/SystemArchitecture_2024/blob/main/Lab_2_Stateful_service_for_RDBMS/imgs/il5.jpg)
![No account](https://github.com/Brinckley/SystemArchitecture_2024/blob/main/Lab_2_Stateful_service_for_RDBMS/imgs/il6.jpg)
## Поиск по маске
![Mask search](https://github.com/Brinckley/SystemArchitecture_2024/blob/main/Lab_2_Stateful_service_for_RDBMS/imgs/il7.jpg)