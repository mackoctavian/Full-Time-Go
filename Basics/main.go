package main

import "fmt"

var (
	name     = ""
	lastName = "Oct"
)

type Weapon string

func main() {
	player := Player{
		health:      10,
		name:        "Mack",
		attackPower: 10.0,
	}

	fmt.Printf("This is players: %d\n", player.getHealth())

	users := map[string]int{}
	user1 := make(map[string]int)

	users["age"] = 10
	user1["age"] = 20
	age, ok := user1["mack"]
	if ok {
		fmt.Printf("%d", age)
	} else {
		fmt.Printf("Empty")
	}

	for k, v := range users {
		fmt.Printf("The key %s and the value %d\n", k, v)
	}

	//Slices
	numbers := []int{1, 2, 3, 4, 5}
	otherNumbers := make([]int, 2)
	numbers = append(numbers, 10)
	otherNumbers = append(otherNumbers, 20)
	otherNumbers = append(otherNumbers, 30)
	otherNumbers = append(otherNumbers, 40)
	otherNumbers = append(otherNumbers, 50)

	var name Weapon = "Macqk"
	println("", name)
}

type Player struct {
	health      int
	name        string
	attackPower float64
}

func (player Player) getHealth() int {
	return player.health
}

func getWeapon(w Weapon) string {
	return string(w)
}
