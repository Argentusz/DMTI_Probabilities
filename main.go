package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	Score      int
	Current    int
	ChoicesLog []int
}

func RandomZeroOne(p float32) int {
	rand.Seed(time.Now().UnixNano())
	random := rand.Float32()
	if random < p {
		return 1
	}
	return 0
}

func Picking(A, B *Player) {
	A.Current = RandomZeroOne(0.5)
	B.Current = RandomZeroOne(0.5)
	A.ChoicesLog = append(A.ChoicesLog, A.Current)
	B.ChoicesLog = append(B.ChoicesLog, B.Current)
}

func PickingBSmarter(A, B *Player) {
	A.Current = RandomZeroOne(0.5)
	B.Current = RandomZeroOne(0.75)
	A.ChoicesLog = append(A.ChoicesLog, A.Current)
	B.ChoicesLog = append(B.ChoicesLog, B.Current)
}

func Judge(A, B *Player) bool {
	if A.Current == B.Current {
		A.Score += 2
		B.Score -= 2
		return true
	}
	if A.Current == 0 {
		A.Score -= 1
		B.Score += 1
	} else {
		A.Score -= 3
		B.Score += 3
	}
	return false
}

func first() {
	var A, B Player
	for i := 0; i < 100; i++ {
		Picking(&A, &B)
		Judge(&A, &B)
	}
	fmt.Println("Итог:")
	fmt.Println("Игрок А: ", A.Score, "Среднее значение: ", float64(A.Score)/100)
	fmt.Println("Игрок B: ", B.Score, "Среднее значение: ", float64(B.Score)/100)
	fmt.Println("Выборы игрока А: ", A.ChoicesLog)
	fmt.Println("Выборы игрока B: ", B.ChoicesLog)
	fmt.Println("Математическое Ожидание игрока А: 0.0")
}

func second() {
	var A, B Player
	for i := 0; i < 100; i++ {
		PickingBSmarter(&A, &B)
		Judge(&A, &B)
	}
	fmt.Println("Итог:")
	fmt.Println("Игрок А: ", A.Score, "Среднее значение: ", float64(A.Score)/100)
	fmt.Println("Игрок B: ", B.Score, "Среднее значение: ", float64(B.Score)/100)
	fmt.Println("Выборы игрока А: ", A.ChoicesLog)
	fmt.Println("Выборы игрока B: ", B.ChoicesLog)
	fmt.Println("Математическое Ожидание игрока A: -0.25")
}

func main() {
	fmt.Println("Первый эксперимент:")
	first()
	fmt.Println("Второй эксперимент:")
	second()
}
