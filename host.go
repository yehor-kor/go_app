package main

import (
	"fmt"
	"net"
)

type Matrix struct {
	// Опис структури матриці
	Rows int
	Columns int
	Elements [][]int
}

type Result struct {
	L2 int
}

func calculateL1() int {
	// Код для обчислення найбільшого власного числа L1 за допомогою методу Лавєр'є-Ньютона
	matrix := [][]int{
		{-100, 50, 10, 25, 30},
		{50, -60, 5, 10, 25},
		{10, 5, -50, 20, 50},
		{25, 10, 20, -40, 10},
		{30, 25, 50, 10, -10},
	}

	maxValue := matrix[0][0] // Припускаємо, що перший елемент є найбільшим

	// Перебираємо всі елементи матриці і оновлюємо maxValue, якщо знаходимо більше значення
	for _, row := range matrix {
		for _, val := range row {
			if val > maxValue {
				maxValue = val
			}
		}
	}
	
	return maxValue // Значення L1 (приклад)
}

func calculateL2() int {
	// Код для обчислення найбільшого власного числа L2 за допомогою методу Лавєр'є-Ньютона
	matrix := [][]int{
		{700, 20, 50, 95, 10},
		{50, -30, 55, 10, 205},
		{10, 6, -50, 200, 500},
		{50, 10, 320, 440, 1100},
		{35, 55, 20, 140, 10},
	}

	maxValue := matrix[0][0] // Припускаємо, що перший елемент є найбільшим

	// Перебираємо всі елементи матриці і оновлюємо maxValue, якщо знаходимо більше значення
	for _, row := range matrix {
		for _, val := range row {
			if val > maxValue {
				maxValue = val
			}
		}
	}
	
	return maxValue // Значення L2
}

func main() {
	// Задача 1: Обчислення L1
	L1 := calculateL1() // Обчислення L1

	// З'єднання з slave-комп'ютером (slave-ОК)
	slaveAddress := "10.0.2.15:5000"
	conn, err := net.Dial("tcp", slaveAddress)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Задача 2: Обчислення L2
	L2 := calculateL2() // Обчислення L2

	// Задача 3: Обчислення L3
	var L3 int
	if L1 > L2 {
		L3 = L1 * L1 * L2
	} else {
		L3 = L1 * L2 * L2
	}

	// Виведення результатів на консоль
	fmt.Println("L1:", L1)
	fmt.Println("L2:", L2)
	fmt.Println("L3:", L3)
}
