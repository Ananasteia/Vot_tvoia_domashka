package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a, b int
	var c string
	fmt.Print("Выбирете раздел :")
	fmt.Scan(&a)

	switch a {
	case 2:
		fmt.Printf("\nВыберете задание :")
		fmt.Scan(&b)
		switch b {
		case 1:
			razdel_2_odin()
		case 2:
			razdel_2_dva()
		case 3:
			razdel_2_tri()
		case 4:
			razdel_2_chetire()
		case 5:
			razdel_2_piats()
		case 6:
			razdel_2_shests()
		case 7:
			razdel_2_semerka()
		case 8:
			razdel_2_vosem()
		case 9:
			razdel_2_deviats()
		case 10:
			razdel_2_desiats()
		}
	case 3:
		fmt.Printf("\nВыберете задание :")
		fmt.Scan(&b)
		switch b {
		case 1:
			razdel_3_odin()
		case 2:
			razdel_3_dva()
		case 3:
			razdel_3_tri()
		case 4:
			razdel_3_cheire()
		case 5:
			razdel_3_piats()
		case 6:
			razdel_3_shests()
		}

	case 4:
		fmt.Printf("\nВведите, задания по какой теме Вам необходимы:\n(На русском языке с маленькой буквы, пж! Вот варианты - массивы, срезы, мапы)\n")
		fmt.Scan(&c)
		fmt.Printf("\nВыберете задание :")
		fmt.Scan(&b)
		switch c {
		case "массивы":
			switch b {
			case 1:
				razdel_4_odin()
			case 2:
				razdel_4_dva()
			case 3:
				razdel_4_tri()
			case 4:
				razdel_4_cheire()
			case 5:
				razdel_4_piats()
			case 6:
				razdel_4_shests()
			}
		case "срезы":
			switch b {
			case 1:
				razdel_4_srez_odin()
			case 2:
				razdel_4_srez_dva()
			case 3:
				razdel_4_srez_tri()
			case 4:
				razdel_4_srez_cheire()
			case 5:
				razdel_4_srez_piats()
			case 6:
				razdel_4_srez_shests()
			case 7:
				razdel_4_srez_semerka()
			}
		case "мапы":
			switch b {
			case 1:
				razdel_4_map_odin()
			case 2:
				razdel_4_map_dva()
			case 3:
				razdel_4_map_tri()
			case 4:
				razdel_4_map_chetire()
			case 5:
				razdel_4_map_piats()
			case 6:
				razdel_4_map_shests()
			}
		}

	}
}

func razdel_2_odin() {
	a := 4
	b := 4.31234
	c := "slishkom legko"
	d := true
	fmt.Print("Type int ", a, ", type flat64 ", b, ", type string ", c, ", type bool ", d, ".")
}
func razdel_2_dva() {
	a := 7
	b := 3
	fmt.Printf("Summa %v , vicheanie %v , umnozhenie %v , delenie %v.", a+b, a-b, a*b, a/b)
}
func razdel_2_tri() {
	var a, b bool
	a = 0 == 0 || true //hz cho tut hoteli, no polzovasia && i || ya umeiu uze davno, escho so vremen if else
	b = 0 != 0 && false
	fmt.Print(a, b)
}
func razdel_2_chetire() {
	a := "7"
	b, _ := strconv.Atoi(a)
	c := b + 7
	fmt.Print(c)
}
func razdel_2_piats() {
	var a int
	b := 1
	fmt.Printf("Sluchai s var %v.\nSluchai s := %v", a, b)
}
func razdel_2_shests() {
	a1 := 3.34545
	a2 := 67.65543
	a3 := a1 + a2
	fmt.Print(a3)
}
func razdel_2_semerka() {
	var a, b bool
	a = 0 != 0
	b = 0 != 1
	fmt.Printf("Znachenia sdelannie s otrisaniem %v, %v", a, b)
}
func razdel_2_vosem() {
	a := 232.42355
	b := 4
	var a1 = int(a)
	var b1 = float64(b)
	fmt.Printf("O, neeeeet, teper oni naoborot: bilo %v, %v, stalo: %v, %v!", a, b, a1, b1)
}
func razdel_2_deviats() {
	a := 21
	fmt.Printf("Pervoe znachenie %v\n", a)
	a = 2123 // tut meniatsa znachenie u peremennoi potomu chto ego meniau lichno ya
	fmt.Printf("Vtoroe znachenie %v", a)
}
func razdel_2_desiats() {
	var a string
	var b int
	fmt.Scan(&a, &b)
	a1, _ := strconv.Atoi(a)
	b2 := a1 + b
	fmt.Printf("Summa chisel %v", b2)
}

const gl_constanta = 3 // dlia zadania odin, tri
func razdel_3_odin() {
	fmt.Printf("Znachenie globalnoi consanti %v", gl_constanta)
}
func razdel_3_dva() {
	const loc_consanta = 3.3
	fmt.Printf("Znachenie localnoi constanti %v", loc_consanta)
}
func razdel_3_tri() {
	const gl_constanta = 23
	fmt.Printf("Vivodit znachenie localnoi constani %v", gl_constanta)
}
func razdel_3_cheire() {
	const (
		a int = 3
		b int = 1
	)
	fmt.Printf("Summa %v, vichitanie %v, umnozhenie %v, delenie %v.", a+b, a-b, a*b, a/b)
}
func razdel_3_piats() {
	const a string = "blabla"
	// a = "blablav" - nevozmozhno ved consanti neizmeniami
	fmt.Printf("ne izmenits consantu")
}
func razdel_3_shests() {
	const (
		a int     = 2
		b string  = "2"
		c bool    = true
		d float64 = 2.2
	)
}
func razdel_4_odin() {
	massiv := [5]int{0, 1, 2, 3, 4}
	fmt.Print(massiv[0], massiv[1], massiv[2], massiv[3], massiv[4])
}
func razdel_4_dva() {
	massiv := [3]string{"nol", "odin", "dva"}
	fmt.Printf("Iznachalnii massiv %v", massiv)
	massiv[1] = "ne odin"
	fmt.Printf("izmenennie elemeni massiva: pervii element %v, vtoroi element %v, retii element %v.", massiv[0], massiv[1], massiv[2])
}
func razdel_4_tri() {
	massiv := [4]int{1, 2, 3, 4}
	fmt.Printf("dlinna massiva %v.", len(massiv))
}
func razdel_4_cheire() {
	massivInt := [2]int{1, 2}
	massivFloat := [2]float64{1.0, 2.2}
	massivString := [2]string{"odin", "dva"}
	fmt.Printf("vot eti massivi \nMassiv intov %v\nMassiv floatov %v\nMassiv stringov %v", massivInt, massivFloat, massivString)
}
func razdel_4_piats() {
	massiv := [2][2]int{{2, 5}, {5, 2}}
	fmt.Print(massiv)
}
func razdel_4_shests() {
	massiv := [3]int{1, 2, 3}
	fmt.Printf("Pervii elemen massiva %v, poslednii element massiva %v.", massiv[0], massiv[2])
}
func razdel_4_srez_odin() {
	srez := make([]int, 1)
	srez = append(srez, 1, 2)
	fmt.Print(srez)
}
func razdel_4_srez_dva() {
	srez := []string{"da", "net", "sam znaesh chei ovet"}
	srez[2] = "tot samii otvet"
	fmt.Print(srez[0], srez[1], srez[2])
}
func razdel_4_srez_tri() {
	var srez []int
	srez = append(srez, 123, 342, 34, 434, 53, 324, 234, 3, 356, 677854, 7456, 2356, 2356, 3546)
	fmt.Printf("dlinna sreza %v, emkost sreza %v.", len(srez), cap(srez))
}
func razdel_4_srez_cheire() {
	srez := make([]int, 1)
	srez = append(srez, 1, 2)
	fmt.Printf("da ya skopirovala etu zadachu iz svoei zadachi odin, a cho? podhodit zhe\n", srez)
}
func razdel_4_srez_piats() {
	var srez []int
	srez = append(srez, 123, 342, 34, 434, 53, 324, 234, 3, 356, 677854, 7456, 2356, 2356, 3546)
	fmt.Printf("i etu zadachu ya skopirovala iz svoei zadachi chetire, a cho? opiat' podhodit zhe\n", srez)
}
func razdel_4_srez_shests() {
	slise := []int{1, 34, 34534, 67, 7, 46345, 6574, 24564, 364, 446, 2354657, 235465, 234567, 23435, 3456576, 8765, 8764, 2354, 5, 67}
	srez := slise[3:12]
	fmt.Printf("Slise koroche takoi %v\nA srez slisa %v", slise, srez)
}
func razdel_4_srez_semerka() {
	slise1 := make([]int, 4)
	slise2 := []int{1, 2, 3, 4}
	_ = copy(slise1, slise2)
	fmt.Println(slise1)
	fmt.Println(slise2)
}
func razdel_4_map_odin() {
	mapi := map[string]int{
		"odin":    1,
		"dva":     2,
		"tri":     3,
		"chetire": 4,
	}
	fmt.Print(mapi)
}
func razdel_4_map_dva() {
	mapi := map[string]int{
		"odin":    1,
		"dva":     2,
		"tri":     3,
		"chetire": 5,
	}
	mapi["chetire"] = 4
	fmt.Print(mapi)
}
func razdel_4_map_tri() {
	mapi := make(map[int]int)
	mapi[1] = 1
	mapi[2] = 2
	mapi[3] = 3
	mapi[4] = 4
	fmt.Printf("Na-na-na, tut na map 1 vividet %v", mapi[1])
}
func razdel_4_map_chetire() {
	mapi := make(map[int]int)
	mapi[1] = 1
	mapi[2] = 2
	mapi[3] = 3
	mapi[4] = 4
	delete(mapi, 4)
	fmt.Print(mapi)
}
func razdel_4_map_piats() {
	mapi := map[string]int{
		"kol":       1,
		"dvoika":    2,
		"troika":    3,
		"chetverka": 4,
		"piaterka":  5,
		"shesterka": 6,
		"semerka":   7,
	}
	a := mapi["dva"]
	b := mapi["dvoika"]
	a2 := a != 0
	b2 := b != 0
	fmt.Printf("Proveiau slovo dva %v, a teper' dvoika %v\neto variant 1\nteper' variant 2\n", a2, b2)
	// variant 2
	mapa := make(map[string]struct{})
	mapa["odin"] = struct{}{}
	mapa["dva"] = struct{}{}
	mapa["tri"] = struct{}{}
	_, ok := mapa["dva"]
	fmt.Println(ok)
	_, ok = mapa["chetire"]
	fmt.Println(ok)
}
func razdel_4_map_shests() {
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
