package main

import (
	"html/template"
	"os"
)

func main() {
	//makes template or throws error if template wrong
	var tmpl, err = template.New("example").Parse("welcome {{.name}}! how are you doing")
	if err != nil {
		panic(err)
	}

	data := map[string]interface{}{
		"name": "bro",
		//"name": 1,
	}

	//prints and throws error if something goes wrong
	err = tmpl.Execute(os.Stdout, data)

	if err != nil {
		panic(err)
	}

}
