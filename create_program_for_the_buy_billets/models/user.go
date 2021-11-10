package models

import "errors"

type User struct {
	name     string
	login    string
	password string
	balance  uint32
	billets  []Billet
}

func newUser(name, login, password string) (*User, error) {
	if len(name) > 0 && len(login) > 0 && len(password) > 0 {
		//TODO: check
		return &User{name: name, login: login, password: password}, nil
	}
	return &User{}, errors.New("bad parameters")
}
func (u *User) SetName(name string) {
	u.name = name
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) SetLogin(login string) {
	u.login = login
}

func (u *User) GetLogin() string {
	return u.login
}

func (u *User) SetPassword(password string) {
	u.password = password
}

func (u *User) GetPassword() string {
	return u.password
}

func (u *User) AddMoneyToBalance(sum uint32) {
	u.balance += sum
}

func (u *User) GetBalance() uint32 {
	return u.balance
}

func (u *User) BuyBillet(b Billet) error {
	if u.GetBalance() < b.GetPrice() {
		return errors.New("Не хватает денег")
	}
	u.billets = append(u.billets, b)
	u.balance -= b.GetPrice()
	return nil
}
func (u *User) GetListBillets() []Billet {
	return u.billets
}

// datas

var Users = []User{
	User{name: "Person1", login: "person1", password: "123", balance: 1000000000},
	User{name: "Person2", login: "person2", password: "123", balance: 2000000000},
	User{name: "Person3", login: "person3", password: "123", balance: 3000000000},
	User{name: "Person4", login: "person4", password: "123", balance: 4000000000},
	User{name: "Person5", login: "person5", password: "123", balance: 4000000000},
	User{name: "Person6", login: "person6", password: "123", balance: 4200000000},
}

func GetUser(login, password string) (*User, error) {
	for k := range Users {
		if Users[k].GetLogin() == login && Users[k].GetPassword() == password {
			return &Users[k], nil
		}
	}
	return nil, errors.New(" Не удалось найти пользователя ")
}
