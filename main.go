package main

import (
	"bufio"
	"fmt"
	"os"
)

type stringMap = map[string]string

func main() {
	marks := stringMap{}
	fmt.Println("Приложение для закладок.")
ConsoleApp:
	for {
		printMenu()
		command := takeCommandFromUser()
		fmt.Println("---------------------")
		flagToStop := menu(marks, command)
		if flagToStop {
			break ConsoleApp
		}
	}
	fmt.Println("Программа завершена.")
}

func printMenu() {
	fmt.Println("Меню")
	fmt.Println("1. Посмотреть закладки.")
	fmt.Println("2. Добавить закладку.")
	fmt.Println("3. Удалить закладку.")
	fmt.Println("4. Завершить.")
}

func takeCommandFromUser() string {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Введите команду: ")
		command, _ := reader.ReadString('\n')
		command = command[:len(command)-1]
		if command == "1" || command == "2" || command == "3" || command == "4" {
			return command
		}
		fmt.Println("Вы ввели неверную команду. Попробуйте ещё раз.")
	}
}

func menu(marks stringMap, command string) bool {
	switch command {
	case "1":
		printBookmarks(marks)
	case "2":
		addBookmark(marks)
	case "3":
		deleteBookmark(marks)
	case "4":
		return true
	}
	return false
}

func printBookmarks(marks stringMap) {
	fmt.Print("Результат: ")
	if len(marks) == 0 {
		fmt.Println("закладок нет.")
		return
	}
	for key, value := range marks {
		fmt.Print(key+":", value+"; ")
	}
	fmt.Println()
}

func addBookmark(marks stringMap) {
	reader := bufio.NewReader(os.Stdin)
	var key, value string
	fmt.Print("Введите название: ")
	key, _ = reader.ReadString('\n')
	key = key[:len(key)-1]
	fmt.Print("Введите адрес: ")
	value, _ = reader.ReadString('\n')
	value = value[:len(value)-1]
	marks[key] = value
}

func deleteBookmark(marks stringMap) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите название: ")
	var key string
	key, _ = reader.ReadString('\n')
	key = key[:len(key)-1]
	delete(marks, key)
}
