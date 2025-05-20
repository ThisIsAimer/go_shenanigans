package main
import "fmt"

func main(){
	var hello = "hello \nworld"

	//``are used for raw strings where \n/\r doent work
	var hello_raw = `hello \n world`
	fmt.Println(hello,"\n"+hello_raw)

	//rune is like char in c
	var letter rune = 'a'
	var kanji rune = '火'
	var emoji rune = '🎃'
	var rune_str = "🎃火a"

	fmt.Println("rune number is ", letter, kanji)
	fmt.Printf("rune letter is %c, %c\n", letter, kanji)
	fmt.Printf("rune letter is %c\n", emoji)
	fmt.Println(rune_str)

}