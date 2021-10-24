# LService
Простенький веб-сервис. 

# Назначение проекта

Обеспечение функционала, такого как:<br/>
+ Добавление пользователя (Имя, Фамилия, Номер телефона)<br/>
+ Вывод всех пользователей<br/>
+ Удаление пользователя по его ID<br/>
+ Поиск пользователя по номеру телефона<br/>
 
# Предварительные условия

	* Golang 1.17 
	* PostgreSQL 12.8

# Установка

1. В папке с проектами, в командной строке выполнить команду:<br/>
        <br/>
	__git clone git://github.com/F7icK/LService.git__

2. В PostgreSQL создать две базы:<br/>
        <br/>
	_User_ <br/>
	_user_test_

3. Из папки migrations выполнить миграцию up в обе базы <br/><br/>
	__migrate -path migrations -database "postgres://localhost/User?sslmode=disable" up__<br/><br/>
	__migrate -path migrations -database "postgres://localhost/user_test?sslmode=disable" up__<br/>
    
4. Выполнить таксу make в командной строке с проектом, которая соберёт бинарник.

5. Запуск. В командной строке с проектом, запускаем бинарник.<br/>

	__./LService__
	
# Примеры запросов

Добавление пользователя:<br/> [POST] <br/>
+ localhost:8080/users name=__[name]__ surname=__[surname]__ telephone=__+79991122333__<br/>
	
Вывод всех пользователей:<br/> [GET] <br/>
+ localhost:8080/users<br/>

Удаление пользователя по id:<br/> [DELETE] <br/>
+ localhost:8080/users/__[id]__<br/>

Поиск по номеру телефона:<br/> [GET] <br/>
+ localhost:8080/users?telephone=__+79991122333__<br/>

# Запуск тестов

Выполнить таску make test в командной строке с проектом, который выполнит все тесты.

# Создано с помощью

 -Golang 1.17<br/>
 -Postgres 12.8<br/>
 -Ubuntu 20.04<br/>
 
# Автор

    F7icK
