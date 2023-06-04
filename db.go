// Открытьб openPanel. Запустить сервер. открыть phpAdmin. логин root, пароль пустой.

package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Users struct {
	ID   int
	Name string `json:"name"`
	Age  uint16 `json:"age"`
}

func main() {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/golangDB")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("Подключено к MySQL")

	// Добавление данных
	// insert, err := db.Query("INSERT INTO `users` (`name`, `age`) VALUES ('Bob', 35)")
	// if err != nil {
	// panic(err)
	// }
	// defer insert.Close()

	// Выборка данных
	for {
		var sqlcommand string
		fmt.Print("Please input the SQL command: ")
		fmt.Scan(&sqlcommand)
		_, err := db.Exec(sqlcommand)
		if err != nil {
			fmt.Println("паника здесь")
			panic(err)
		}

		res, err := db.Query("SELECT * FROM `users`")
		if err != nil {
			panic(err)
		}

		for res.Next() {
			var user Users
			err = res.Scan(&user.ID, &user.Name, &user.Age)
			if err != nil {
				panic(err)
			}
			fmt.Println(fmt.Sprintf("User %d: %s with age %d", user.ID, user.Name, user.Age))
		}

		defer res.Close()
	}
}
