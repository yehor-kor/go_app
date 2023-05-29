package main

import (
	"net"
	"encoding/gob"
)

func calculateL20() int {
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

func main0() {
	// Прийняття з'єднання від головного комп'ютера (host)
	hostAddress := "192.168.1.202:5000"
	ln, err := net.Listen("tcp", hostAddress)
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	conn, err := ln.Accept()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Відправка результату L2 головному комп'ютеру
	result := calculateL20()
	encoder := gob.NewEncoder(conn)
	err = encoder.Encode(result)
	if err != nil {
		panic(err)
	}
}
