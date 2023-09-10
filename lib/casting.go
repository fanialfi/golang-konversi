package lib

import "fmt"

func KonversiWithCasting() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("panic error occured with message \"%v\"\n", r)
			// fmt.Println(r)
		}
	}()

	// type assertions pada interface kosong (interface{} / any)
	data := map[string]any{
		"nama":   "Fani Alfirdaus",
		"grade":  2,
		"height": 160,
		"isMale": true,
		"hobies": []string{"eating", "sleeping"},
	}
	fmt.Printf("%T\t%v\n", data["nama"], data["nama"])
	fmt.Printf("%T\t%v\n", data["grade"], data["grade"])
	fmt.Printf("%T\t%v\n", data["height"], data["height"])
	fmt.Printf("%T\t%v\n", data["isMale"], data["isMale"])
	fmt.Printf("%T\t%v\n", data["hobies"], data["hobies"])
	fmt.Println()

	for _, value := range data {
		switch value.(type) {
		case string:
			fmt.Printf("%T\t%s\n", value.(string), value.(string))
		case int:
			fmt.Printf("%T\t%d\n", value.(int), value.(int))
		case bool:
			fmt.Printf("%T\t%t\n", value.(bool), value.(bool))
		case []string:
			fmt.Printf("%T\t%v\n", value.([]string), value.([]string))
		default:
			fmt.Println("value", value)
		}
	}

	fmt.Println()
	// casting string <--> byte
	text := "😊 hello"

	a := []byte(text)
	a = append(a, 252, 169)
	for index, elm := range a {
		fmt.Printf("index ke-%d\t%d\t%s\n", index, elm, string(elm))
	}
	fmt.Println(string(a))

	fmt.Println()

	// tipe data rune bisa mengatas karakter ASCII non standar (seperti karakter emoji)
	b := []rune(text)
	b = append(b, 252, 169)
	for index, elm := range b {
		fmt.Printf("index ke-%d\t%d\t%s\n", index, elm, string(elm))
	}
	fmt.Println(string(b))
}
