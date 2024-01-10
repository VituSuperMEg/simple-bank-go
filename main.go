package main

import "fmt"

type Bank struct {
	Account string
	Amount  int
}
type User struct {
	Nome  string
	Email string
	Bank  *Bank
}

func NewBank() *Bank {
	return &Bank{Account: "123", Amount: 0}
}
func GetAmmount(u *User) int {
	return u.Bank.Amount
}
func Deposit(v int, u *User) int {
	u.Bank.Amount += v
	return u.Bank.Amount
}

func Transaction(u *User, u2 *User, key string, value int) (string int) {

	if value < u2.Bank.Amount {
		fmt.Println("Saldo insufficient")
	}
	if key == "user1" && value <= u2.Bank.Amount {
		u.Bank.Amount += value
		u2.Bank.Amount -= value
		fmt.Printf("O valor na nova conta do usuário 1 é: %v\n", GetAmmount(u))
	}
	return 0
}

func main() {
	user := User{
		Nome:  "John",
		Email: "John@example.com",
		Bank:  NewBank(),
	}
	user2 := User{
		Nome:  "Doe",
		Email: "doe@example.com",
		Bank:  NewBank(),
	}
	// fmt.Printf("Usuário 1 é : %v\n o Usuário 2 é : %v", user, user2)
	fmt.Printf("Foi depositiado em sua conta usuario 2 : %v\n", Deposit(150, &user2))
	// fmt.Printf("Usuário : %v\ne seus dados : %s", user, user.Bank.Account)
	// fmt.Println(GetAmmount(&user))
	// fmt.Println(Deposit(110, &user))
	fmt.Printf("Extrato da Conta do Usuário 2 : %v\n", GetAmmount(&user2))
	fmt.Println(Transaction(&user, &user2, "user1", 149))
	fmt.Printf("Consulta de saldo : %v\n", GetAmmount(&user2))

}
