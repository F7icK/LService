# LService
Простенький веб-сервис.

# Предварительные условия

	* Golang 1.17 
	* PostgreSQL 12.8

# Установка

1. В папке с проектами, в командной строке выполнить команду:<br/>
    __git clone git://github.com/F7icK/LService.git__

2. В PostgreSQL создать две базы:<br/>
User <br/>
user_test

3. Из папки migrations выполнить миграцию up в обе базы <br/>
    __migrate -path migrations -database "postgres://localhost/User?sslmode=disable" up__ <br/>
    __migrate -path migrations -database "postgres://localhost/user_test?sslmode=disable" up__
    
4. Выполнить таксу make в командной строке с проектом, которая соберёт бинарник.

5. Запуск

		__./LService__

# Запуск тестов

Выполнить таску make test в командной строке с проектом, который выполнит все тесты.

# Создано с помощью

 -Goland<br/>
 -Ubuntu 20.04<br/>
 
# Автор
	
	F7icK
