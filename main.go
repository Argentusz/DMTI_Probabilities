package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Player struct {
	Score      int
	Current    int
	ChoicesLog []int
	Balls      []int
}

func ExpectedValue(PA, PB float32) float32 {
	return 2*((1-PA)*PB) - 3*(PA*(1-PB)) - 1*((1-PA)*PB) + 2*(PA*(1-PB))
}

func StdDeviation(expValue, res float32) float32 {
	return float32(math.Pow(math.Pow(float64(expValue), 2)+math.Pow(float64(res), 2), 0.5))
}

func Dispersion(expValue, res float32) float32 {
	return float32(math.Pow(float64(res-expValue), 2)) / 2
}

func RandomZeroOne(p float32) int {
	rand.Seed(time.Now().UnixNano())
	random := rand.Float32()
	if random < p {
		return 0
	}
	return 1
}

func Picking(A, B *Player, p1, p2 float32) {
	A.Current = RandomZeroOne(p1)
	B.Current = RandomZeroOne(p2)
	A.ChoicesLog = append(A.ChoicesLog, A.Current)
	B.ChoicesLog = append(B.ChoicesLog, B.Current)
}

func Judge(A, B *Player) int {
	if A.Current == B.Current {
		A.Score += 2
		B.Score -= 2
		return 2
	}
	if A.Current == 1 {
		A.Score -= 1
		B.Score += 1
		return -1
	} else {
		A.Score -= 3
		B.Score += 3
		return -3
	}
}

func first() {
	var A, B Player
	const PA = 0.5
	const PB = 0.5
	for i := 0; i < 100; i++ {
		Picking(&A, &B, PA, PB)
		Judge(&A, &B)
	}
	fmt.Println("Итог:")
	fmt.Println("Игрок А: ", A.Score, "Среднее значение: ", float32(A.Score)/100)
	fmt.Println("Игрок B: ", B.Score, "Среднее значение: ", float32(B.Score)/100)
	fmt.Println("Выборы игрока А: ", A.ChoicesLog)
	fmt.Println("Выборы игрока B: ", B.ChoicesLog)
	fmt.Println("Математическое Ожидание игрока А:", ExpectedValue(PA, PB))
}

func second() {
	var A, B Player
	const PA float32 = 0.5
	const PB float32 = 0.25
	for i := 0; i < 100; i++ {
		Picking(&A, &B, PA, PB)
		Judge(&A, &B)
	}
	fmt.Println("Итог:")
	fmt.Println("Игрок А: ", A.Score, "Среднее значение: ", float32(A.Score)/100)
	fmt.Println("Игрок B: ", B.Score, "Среднее значение: ", float32(B.Score)/100)
	fmt.Println("Выборы игрока А: ", A.ChoicesLog)
	fmt.Println("Выборы игрока B: ", B.ChoicesLog)
	fmt.Println("Математическое Ожидание игрока A:", ExpectedValue(PA, PB))
}

func third() {
	var A, B Player
	const PB float32 = 0.5
	var PA float32
	A.Balls = append(A.Balls, 10, 10)
	for i := 0; i < 100; i++ {
		Picking(&A, &B, PA, PB)
		if Judge(&A, &B) > 0 {
			A.Balls[A.Current] += 2
		}
		PA = float32(A.Balls[0]) / float32(A.Balls[0]+A.Balls[1])
	}
	fmt.Println("Итог:")
	fmt.Println("Игрок А: ", A.Score, "Среднее значение: ", float32(A.Score)/100)
	fmt.Println("Игрок B: ", B.Score, "Среднее значение: ", float32(B.Score)/100)
	fmt.Println("Выборы игрока А: ", A.ChoicesLog)
	fmt.Println("Выборы игрока B: ", B.ChoicesLog)
	fmt.Println("Математическое Ожидание игрока A:", ExpectedValue(PA, PB))
}
func fourth() {
	var A, B Player
	const PB float32 = 0.5
	A.Balls = append(A.Balls, 100, 100)
	var PA = float32(A.Balls[0]) / float32(A.Balls[0]+A.Balls[1])
	for i := 0; i < 100; i++ {
		Picking(&A, &B, PA, PB)
		temp := Judge(&A, &B)
		if temp < 0 {
			A.Balls[A.Current] += temp
		}
		PA = float32(A.Balls[0]) / float32(A.Balls[0]+A.Balls[1])
	}
	fmt.Println("Итог:")
	fmt.Println("Игрок А: ", A.Score, "Среднее значение: ", float32(A.Score)/100)
	fmt.Println("Игрок B: ", B.Score, "Среднее значение: ", float32(B.Score)/100)
	fmt.Println("Выборы игрока А: ", A.ChoicesLog)
	fmt.Println("Выборы игрока B: ", B.ChoicesLog)
	fmt.Println("Математическое Ожидание игрока A:", ExpectedValue(PA, PB))
}

func fifth() {
	var A, B Player
	A.Balls = append(A.Balls, 10, 10)
	B.Balls = append(B.Balls, 10, 10)
	var PA, PB = float32(A.Balls[0]) / float32(A.Balls[0]+A.Balls[1]), float32(B.Balls[0]) / float32(B.Balls[0]+B.Balls[1])
	for i := 0; i < 100; i++ {
		Picking(&A, &B, PA, PB)
		temp := Judge(&A, &B)
		if temp > 0 {
			A.Balls[A.Current] += temp
		} else {
			B.Balls[B.Current] -= temp
		}
		PA, PB = float32(A.Balls[0])/float32(A.Balls[0]+A.Balls[1]), float32(B.Balls[0])/float32(B.Balls[0]+B.Balls[1])
	}
	fmt.Println("Итог:")
	fmt.Println("Игрок А: ", A.Score, "Среднее значение: ", float32(A.Score)/100)
	fmt.Println("Игрок B: ", B.Score, "Среднее значение: ", float32(B.Score)/100)
	fmt.Println("Выборы игрока А: ", A.ChoicesLog)
	fmt.Println("Выборы игрока B: ", B.ChoicesLog)
	fmt.Println("Математическое Ожидание игрока A:", ExpectedValue(PA, PB))
}

func main() {
	fmt.Println("Первый эксперимент:")
	first()
	fmt.Println()
	fmt.Println("Второй эксперимент:")
	second()
	fmt.Println()
	fmt.Println("Третий эксперимент:")
	third()
	fmt.Println()
	fmt.Println("Четвёртый эксперимент:")
	fourth()
	fmt.Println()
	fmt.Println("Пятый эксперимент:")
	fifth()
	fmt.Println()
}
