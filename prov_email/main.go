package main

import (
	"fmt"
	"bufio" //удобное чтение ввода (строками)
	"os" //доступ к системе (stdin = клавиатура)
	"net/mail" //проверка email формата
	"strings" //работа со строками (очистка, обрезка)
)

func prov(email string) string{
	_,err := mail.ParseAddress(email)
	if err != nil {
		return "Invalid Email" //если err != nil → email неправильный
	}
	return "Valid Email"
}
func main(){
	fmt.Print("Enter email: ")
	reader := bufio.NewReader(os.Stdin)
	email,_ := reader.ReadString('\n') //читает строку до Enter

	email = strings.TrimSpace(email) // убираем \n и пробелы
	
	fmt.Println(prov(email))
}
