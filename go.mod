//This file serves as your go module dependency guide
//It's like composer.json in PHP
//It's like package.json in JS

//This line will enable you to locally import your packages.
//The repo name should ideally be the folder name
module github.com/karenirenecano/go-mysql-export-to-excel

go 1.14

require (
	github.com/360EntSecGroup-Skylar/excelize v1.4.1
	github.com/go-sql-driver/mysql v1.5.0
)
