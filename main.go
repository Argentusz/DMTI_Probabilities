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
	Balls      []int
}

func ExpectedValue(PA, PB float32) float32 {
	return 2*((1-PA)*PB) - 3*(PA*(1-PB)) - 1*((1-PA)*PB) + 2*(PA*(1-PB))
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
	B.Current = RandomZeroOne(0.25)
	A.ChoicesLog = append(A.ChoicesLog, A.Current)
	B.ChoicesLog = append(B.ChoicesLog, B.Current)
}

func PickingABalls(A, B *Player) float32 {
	A.Current = RandomZeroOne(float32(A.Balls[0]) / float32(A.Balls[0]+A.Balls[1]))
	B.Current = RandomZeroOne(0.5)
	A.ChoicesLog = append(A.ChoicesLog, A.Current)
	B.ChoicesLog = append(B.ChoicesLog, B.Current)
	return float32(A.Balls[0]) / float32(A.Balls[0]+A.Balls[1])
}

func PickingABBalls(A, B *Player) (float32, float32) {
	A.Current = RandomZeroOne(float32(A.Balls[0]) / float32(A.Balls[0]+A.Balls[1]))
	B.Current = RandomZeroOne(float32(B.Balls[0]) / float32(B.Balls[0]+B.Balls[1]))
	A.ChoicesLog = append(A.ChoicesLog, A.Current)
	B.ChoicesLog = append(B.ChoicesLog, B.Current)
	return float32(A.Balls[0]) / float32(A.Balls[0]+A.Balls[1]), float32(B.Balls[0]) / float32(B.Balls[0]+B.Balls[1])
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

func JudgeWithABalls(A, B *Player) bool {
	if A.Current == B.Current {
		A.Score += 2
		B.Score -= 2
		A.Balls[A.Current] += 2
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

func JudgeWithABallsPunished(A, B *Player) bool {
	if A.Current == B.Current {
		A.Score += 2
		B.Score -= 2
		return true
	}
	if A.Current == 0 {
		A.Score -= 1
		B.Score += 1
		A.Balls[0] -= 1
	} else {
		A.Score -= 3
		B.Score += 3
		A.Balls[1] -= 3
	}
	return false
}

func JudgeWithABBalls(A, B *Player) bool {
	if A.Current == B.Current {
		A.Score += 2
		B.Score -= 2
		A.Balls[A.Current] += 2
		return true
	}
	if A.Current == 0 {
		A.Score -= 1
		B.Score += 1
		B.Balls[1] += 1
	} else {
		A.Score -= 3
		B.Score += 3
		B.Balls[0] += 3
	}
	return false
}

func first() {
	var A, B Player
	const PA = 0.5
	const PB = 0.5
	for i := 0; i < 100; i++ {
		Picking(&A, &B)
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
		PickingBSmarter(&A, &B)
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
		PA = PickingABalls(&A, &B)
		JudgeWithABalls(&A, &B)
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
	var PA float32
	A.Balls = append(A.Balls, 100, 100)
	for i := 0; i < 100; i++ {
		PA = PickingABalls(&A, &B)
		JudgeWithABallsPunished(&A, &B)
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
	var PA, PB float32
	A.Balls = append(A.Balls, 10, 10)
	B.Balls = append(B.Balls, 10, 10)
	for i := 0; i < 100; i++ {
		PA, PB = PickingABBalls(&A, &B)
		JudgeWithABBalls(&A, &B)
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
