# LService
Простенький веб-сервис. 

# Назначение проекта

Обеспечение функционала, такого как:<br/>
+ Добавление пользователя (Имя, Фамилия, Номер телефона)_<br/>
+ Вывод всех пользователей<br/>
+ Удаление пользователя по его ID<br/>
+ Поиск пользователя по номеру телефона<br/>
 
# Предварительные условия

	* Golang 1.17 
	* PostgreSQL 12.8

# Установка

1. В папке с проектами, в командной строке выполнить команду:<br/>
	__git clone git://github.com/F7icK/LService.git__

2. В PostgreSQL создать две базы:<br/>
	_User_ <br/>
	_user_test_

3. Из папки migrations выполнить миграцию up в обе базы <br/>
	__migrate -path migrations -database "postgres://localhost/User?sslmode=disable" up__ <br/>
	__migrate -path migrations -database "postgres://localhost/user_test?sslmode=disable" up__
    
4. Выполнить таксу make в командной строке с проектом, которая соберёт бинарник.

5. Запуск

	__./LService__
	
# Примеры запросов

Добавление пользователя:<br/> [POST] <br/>
+ _localhost:8080/users name=__[name]__ surname=__[surname]__ telephone=__+79991122333___<br/>
	
Вывод всех пользователей:<br/> [GET] <br/>
+ _localhost:8080/users_<br/>

Удаление пользователя по id:<br/> [DELETE] <br/>
+ _localhost:8080/users/__[id]__<br/>

Поиск по номеру телефона:<br/> [GET] <br/>
+ _localhost:8080/users?telephone=__+79991122333___<br/>

# Запуск тестов

Выполнить таску make test в командной строке с проектом, который выполнит все тесты.

# Создано с помощью

 -Goland<br/>
 -Ubuntu 20.04<br/>
 
# Автор
	
	F7icK
