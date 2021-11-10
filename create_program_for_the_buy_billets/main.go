package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"

	"github.com/kakoitouser/learn/create_program_for_the_buy_billets/helpers"
	"github.com/kakoitouser/learn/create_program_for_the_buy_billets/models"
	"github.com/pterm/pterm"
)

func createBilletsList(count int) []models.Billet {
	billetsArr := make([]models.Billet, 0)
	i := 0
	for i < count {
		rndRocket := rand.Intn(5)
		rndSpaceLine := rand.Intn(3)
		rndTriptype := rand.Intn(2)
		billet, err := models.NewBillet(models.Spacelines[rndSpaceLine], models.TripTypes[rndTriptype], models.Rockets[rndRocket])
		if err != nil {
			continue
		}
		billetsArr = append(billetsArr, billet)
		i++
	}
	return billetsArr
}

func displayBillets(billetsArr []models.Billet) {

	data := pterm.TableData{{"Spaceline", "Days", "Tripe type", "Price"}}
	for _, val := range billetsArr {
		var newBil []string = []string{string(val.GetSpaceLineName()), strconv.Itoa(int(val.GetDays())), string(val.GetTripType()), "$ " + strconv.Itoa(int(val.GetPrice()))}
		data = append(data, newBil)
	}
	data = append(data, []string{"some", "some", "some", "$ " + helpers.Get_format_money(990009)})
	pterm.DefaultTable.WithHasHeader().WithData(data).Render()
}

func main() {
	billetsArr := createBilletsList(10)
	displayBillets(billetsArr)

	var inpt string
	var user models.User
	check := true
	for check {
		fmt.Println(" 1-register  2-login")
		fmt.Print("Введите:")
		fmt.Scan(&inpt)

		switch inpt {
		case "1":
			check = false
		case "2":
			var login, password string
			fmt.Print("login: ")
			fmt.Scan(&login)
			fmt.Print("password: ")
			fmt.Scan(&password)
			user, err := models.GetUser(login, password)
			if err != nil {
				log.Fatal(err)
			}
			check = false
			fmt.Println(user.GetName(), " Добро пожаловать")
		default:
			fmt.Print(" не правильный ввод ")
		}
	}

	check = true
	for check {
		displayBillets(billetsArr)
		fmt.Println(" 1-buy billet  2-my billets")
		fmt.Print("Введите:")
		fmt.Scan(&inpt)

		switch inpt {
		case "1":
			var bilNum int
			fmt.Print("номер биллета: ")
			fmt.Scan(&bilNum)

			if bilNum > 0 && bilNum < len(billetsArr) {
				err := user.BuyBillet(billetsArr[bilNum])
				if err != nil {
					log.Fatal(err)
				}
			}
			check = true
		case "2":
			fmt.Println("Ваши билеты")
			displayBillets(user.GetListBillets())
		default:
			fmt.Print(" не правильный ввод ")
		}

	}
}
