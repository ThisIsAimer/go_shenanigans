package main

import "fmt"

func main(){
	var a int = 77
	b := int32(a)
	c:= float64(b)

	d := "1"
	// e := int(d) not possible

	f := 3.14
	g := int(f)

	h := "i love my brother!"

	i := []byte(h)

	j := []byte{255,32,111,72}

	h = string(j)


	fmt.Println("",a,b,c,d,g,"\n",i,"\n",h)

}