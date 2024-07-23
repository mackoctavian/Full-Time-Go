package main

import "fmt"

type WeaponType int

const (
	Axe WeaponType = iota
	Sword
	Knife
)

func getDamage(weaponType WeaponType) int {
	switch weaponType {
	case Axe:
		return 100
	case Sword:
		return 90
	case Knife:
		return 10
	default:
		panic("Not Weapon Found")
	}
}

func main() {
	fmt.Println("Weapon Damage: ", getDamage(Axe))
}
