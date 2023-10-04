package main

import (
	"fmt"
)

func main() {

}

// ЦИКЛЫ 1
func razdel_5_1() {
	for i := 1; i < 11; i++ {
		fmt.Print(i, " ")
	}
}

// ЦИКЛЫ 2
func razdel_5_2() {
	for i := 2; i < 21; i += 2 {
		fmt.Print(i, " ")
	}
}

// ЦИКЛЫ 3
func razdel_5_3() {
	for i := 1; i < 11; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Print(i, " ")
	}
}

// ЦИКЛЫ 4
func razdel_5_4() {
	var s int
	for i := 1; i < 101; i++ {
		s += i
		if s > 50 {
			fmt.Printf("Сумма все ранее выведенных чисел, на которой все прерывается :%v", s)
			break
		}
		fmt.Print(i, " ")
	}
}

// ЦИКЛЫ 5
func razdel_5_5() {
	m := [3]int{1, 10, 100}
	for i, _ := range m {
		fmt.Print(i, " ")
	}
}

// ЦИКЛЫ 6
func razdel_5_6() {
	mapa := map[string]int{
		"kol":       1,
		"dvoika":    2,
		"troika":    3,
		"chetverka": 4,
		"piaterka":  5,
	}
	for i, _ := range mapa {
		fmt.Println(i)
	}
}

// ЦИКЛЫ 7
func razdel_5_7() {
	s := []int{11, 22, 33, 44}
	for i, _ := range s {
		fmt.Println(i)
	}
}

// ЦИКЛЫ 8   (НОРМ ЗАДАЧА)
func razdel_5_8() {
	s := "SKUCHNO"
	for _, e := range s {
		fmt.Println(string(e))
	}
}

//---------------------------------------------------------------------------------------------//

// ФУНКЦИИ 1
func razdel_6_1() {
	priveik_pukni_v_paketik()
}
func priveik_pukni_v_paketik() {
	fmt.Print("Приветик, пукни в петики :)")
}

// ФУНКЦИИ 2
func razdel_6_2() {
	s := summa(13, 33)
	fmt.Print(s)
}
func summa(a int, b int) int {
	return a + b
}

// ФУНКЦИИ 3
func razdel_6_3() {
	u := umnozhenie(4234, 434)
	fmt.Print(u)
}
func umnozhenie(a int, b int) int {
	return a * b
}

// ФУНКЦИИ 4
func razdel_6_4() {
	s, v := hujunksia(2342, 345)
	fmt.Print(s, v)
}
func hujunksia(a int, b int) (int, int) {
	return a + b, a - b
}

// ФУНКЦИИ 5
func razdel_6_5() {
	blin(1, 2, 3, 4, 5)
}
func blin(a ...int) {
	for _, e := range a {
		fmt.Printf("%d ", e)
	}
}

// ФУНКЦИИ 6
func razdel_6_6() {
	fmt.Print("Это шутка, что ли? Я делала так уже до, смотри там задачки выше")
}

// ФУНКЦИИ 7
const h = 45

func razdel_6_7(t int) {
	//задать константы и принимать их? Типа вызов функции будет выглядеть так razdel_6_7(h) - где h является константой указанной выше
	z := t + 6
	fmt.Print(z)
}

// ФУНКЦИИ 8!!!!!  БАЛДЕЖ ЭКСТАЗ МОЙ МОЗГ БИЛСЯ В КОНВУЛЬСИЯХ, НО СДЕЛАЛ! ЗАДАЧА НА 10000000000 ИЗ 10 ОХУЕННО, МНЕ ОЧ ПОНРАВИЛОСЬ
func razdel_6_8() {
	d := govno(5, net)
	fmt.Print(d)
}
func net(v int) int {
	return v + 2
}
func govno(x int, d func(int) int) int {
	return d(x)
}
