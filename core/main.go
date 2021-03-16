package main

import (
	"fmt"
	"golang_practice/basic"
	"golang_practice/blockchain"
	"golang_practice/model"
	"golang_practice/utils"
	// ""
	// "hello/utils"
)

func main() {

	// fmt import 시험
	fmt.Println("what the")
	basic.Hi()
	// 디렉터리 생성과 import 예제
	utils.Helloworld()
	model.HelloModel("hong")

	/**
	Closer 실패 사례
	*/
	// tmp1 := dosomething()
	// fmt.Println(tmp1)
	// tmp2 := dosomething()
	// fmt.Println(tmp2)
	// tmp3 := dosomething()
	// fmt.Println(tmp3)
	// closerExample()

	lol()

	// reculsive
	result := reculsive(8)
	fmt.Println(result)

	// struct 개념 및 생성자 만들기
	// person := newPerson("hong", 22)
	// fmt.Println(person)
	block := blockchain.NewBlock("123", []byte{121})
	fmt.Println(block.Data)
	fmt.Println(block.Hash)
	fmt.Println(block.PrevBlockHash)
	fmt.Println(block.Timestamp)

	bc := blockchain.NewBlockchain()
	bc.AddBlock("send 1 btc")
	bc.AddBlock("send 2 btc to hong")

	for _, block := range bc.GetBlocks() {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}

func dosomething() func() int {

	i := 0
	// fmt.Println(i)
	return func() int {
		i++
		return i
	}
}

func lol() {
	fmt.Println("what")
}

func closerExample() {
	tmp := dosomething()
	fmt.Println(tmp())
	fmt.Println(tmp())
	fmt.Println(tmp())
	fmt.Println(tmp())
}

func reculsive(n int) int {
	if n == 10 {
		return 1
	}
	n++
	return n * reculsive(n)
}

// func newPerson(name string, age int) *model.Person {
// 	// implicit assignment to unexported field name in model.Person literalcompiler
// 	// person := model.Person{name, age}
// 	// 위의 문제 - 다른 디렉터리의 객체를 생성하기 위해서는 Name과 같이 대문자로 선언하여 public의 형태로 만들어 주어야함
// 	person := model.Person{}
// 	// person2 := model.Person{"hong", 22}
// 	// fmt.Println(person2)
// 	person.Age = age
// 	person.Name = name
// 	return &person
// }

// package main

// import (
// 	"github.com/andlabs/ui"
// 	_ "github.com/andlabs/ui/winmanifest"
// )

// func main() {
// 	err := ui.Main(func() {
// 		input := ui.NewEntry()
// 		button := ui.NewButton("Greet")
// 		greeting := ui.NewLabel("")
// 		box := ui.NewVerticalBox()
// 		box.Append(ui.NewLabel("Enter your name:"), false)
// 		box.Append(input, false)
// 		box.Append(button, false)
// 		box.Append(greeting, false)
// 		window := ui.NewWindow("Hello", 200, 100, false)
// 		window.SetMargined(true)
// 		window.SetChild(box)
// 		button.OnClicked(func(*ui.Button) {
// 			greeting.SetText("Hello, " + input.Text() + "!")
// 		})
// 		window.OnClosing(func(*ui.Window) bool {
// 			ui.Quit()
// 			return true
// 		})
// 		window.Show()
// 	})
// 	if err != nil {
// 		panic(err)
// 	}
// }
