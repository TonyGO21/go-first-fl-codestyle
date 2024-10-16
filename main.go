package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func randomDamage(base int, min int, max int) int {
	return base + randint(min, max)
}

func attack(charName, charClass string) string {
	var damage int
	switch charClass {
	case "warrior":
		damage = randomDamage(5, 3, 5)
	case "mage":
		damage = randomDamage(5, 5, 10)
	case "healer":
		damage = randomDamage(5, -3, -1)
	default:
		return "неизвестный класс персонажа"
	}
	return fmt.Sprintf("%s нанес урон противнику равный %d.", charName, damage)
}

func defence(charName, charClass string) string {
	var block int
	switch charClass {
	case "warrior":
		block = randomDamage(10, 5, 10)
	case "mage":
		block = randomDamage(10, -2, 2)
	case "healer":
		block = randomDamage(10, 2, 5)
	default:
		return "неизвестный класс персонажа"
	}
	return fmt.Sprintf("%s блокировал %d урона.", charName, block)
}

func special(charName, charClass string) string {
	var skill string
	switch charClass {
	case "warrior":
		skill = fmt.Sprintf("Выносливость %d", 80+25)
	case "mage":
		skill = fmt.Sprintf("Атака %d", 5+40)
	case "healer":
		skill = fmt.Sprintf("Защита %d", 10+30)
	default:
		return "неизвестный класс персонажа"
	}
	return fmt.Sprintf("%s применил специальное умение `%s`", charName, skill)
}

func startTraining(charName, charClass string) string {
	classMessages := map[string]string{
		"warrior": "Ты Воитель - отличный боец ближнего боя.",
		"mage":    "Ты Маг - превосходный укротитель стихий.",
		"healer":  "Ты Лекарь - чародей, способный исцелять раны.",
	}

	fmt.Printf("%s, %s\n", charName, classMessages[charClass])
	fmt.Println("Потренируйся управлять своими навыками.")
	fmt.Println("Введи одну из команд: attack — чтобы атаковать противника,")
	fmt.Println("defence — чтобы блокировать атаку противника,")
	fmt.Println("special — чтобы использовать свою суперсилу.")
	fmt.Println("Если не хочешь тренироваться, введи команду skip.")

	var cmd string
	for cmd != "skip" {
		fmt.Print("Введи команду: ")
		fmt.Scanf("%s\n", &cmd)

		switch cmd {
		case "attack":
			fmt.Println(attack(charName, charClass))
		case "defence":
			fmt.Println(defence(charName, charClass))
		case "special":
			fmt.Println(special(charName, charClass))
		}
	}
	return "тренировка окончена"
}

func chooseCharClass() string {
	var charClass string
	var approveChoice string

	for approveChoice != "y" {
		fmt.Print("Введи название персонажа (Воитель — warrior, Маг — mage, Лекарь — healer): ")
		fmt.Scanf("%s\n", &charClass)

		description := map[string]string{
			"warrior": "Воитель — дерзкий воин ближнего боя. Сильный, выносливый и отважный.",
			"mage":    "Маг — находчивый воин дальнего боя. Обладает высоким интеллектом.",
			"healer":  "Лекарь — могущественный заклинатель. Черпает силы из природы, веры и духов.",
		}

		if desc, exists := description[charClass]; exists {
			fmt.Println(desc)
		} else {
			fmt.Println("Неизвестный класс персонажа.")
			continue
		}

		fmt.Print("Нажми (Y), чтобы подтвердить выбор, или любую другую кнопку, чтобы выбрать другого персонажа: ")
		fmt.Scanf("%s\n", &approveChoice)
		approveChoice = strings.ToLower(approveChoice)
	}
	return charClass
}

func main() {
	fmt.Println("Приветствую тебя, искатель приключений!")
	fmt.Print("Назови себя: ")
	var charName string
	fmt.Scanf("%s\n", &charName)
	fmt.Printf("Здравствуй, %s\n", charName)
	fmt.Println("Сейчас твоя выносливость — 80, атака — 5 и защита — 10.")
	fmt.Println("Ты можешь выбрать один из трёх путей силы: Воитель, Маг, Лекарь")

	charClass := chooseCharClass()
	fmt.Println(startTraining(charName, charClass))
}

func randint(min, max int) int {
	return rand.Intn(max-min+1) + min
}
