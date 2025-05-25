package main

import (
	"bufio"
	"fmt"
	"html/template"
	"os"
	"strings"
)

func main() {

	var programme int = 1
	//makes template or throws error if template wrong
	if programme == 0 {
		var tmpl, err = template.New("example").Parse("welcome {{.name}}! how are you doing there are {{.number}} people preasent.\n")
		//var tmpl = template.Must(template.New("example").Parse("welcome {{.name}}! how are you doing there are {{.number}} people preasent.\n"))
		if err != nil {
			panic(err)
		}

		data := map[string]interface{}{
			"name": "bro",
			"number": 100,
			//"name": 1,
		}

		//prints and throws error if something goes wrong
		err = tmpl.Execute(os.Stdout, data)

		if err != nil {
			panic(err)
		}
	} else{
		//---------------------------------------------------------------------

		fmt.Println("please enter your name:")
		buffer := bufio.NewReader(os.Stdin)
		name, _ := buffer.ReadString('\n')

		//trimspace gets rid of \n
		name = strings.TrimSpace(name)

		madeTemplates := map[string]string{
			"name": "hello {{.name}}, welcome\n",
			"notification" : "{{.name}} you have a notification: {{.notification}}\n",
			"error" : "Error: {{.error}}\n",
		}

		parsedTemplates := make(map[string]*template.Template)

		for key, value := range madeTemplates{
			parsedTemplates[key] = template.Must(template.New(key).Parse(value))
		}

		for{
			println("-----------------------------------------------")
			fmt.Println("choose the following: ")
			fmt.Println("1. join")
			fmt.Println("2. notification")
			fmt.Println("3. error")
			fmt.Println("4. exit")

			//we could have used bufio here too
			var input string
			fmt.Scanln(&input)

			var tmpl *template.Template
			var mappings map[string]interface{}

			switch input{

			case "1":
					tmpl = parsedTemplates["name"]
					mappings = map[string]interface{}{
						"name" : name,
					}
					
				case "2":
					fmt.Println("enter new notification:")
					buffer := bufio.NewReader(os.Stdin)
					notification, _ := buffer.ReadString('\n')
					notification = strings.TrimSpace(notification)

					tmpl = parsedTemplates["notification"]

					mappings = map[string]any{
						"name" : name,
						"notification" : notification,
					}
					
				case "3":
					fmt.Println("enter new error:")
					buffer := bufio.NewReader(os.Stdin)
					myErr , _ := buffer.ReadString('\n')
					myErr = strings.TrimSpace(myErr)

					tmpl = parsedTemplates["error"]

					mappings = map[string]any{
						"name": name,
						"error": myErr,
					}
				case "4":
					fmt.Println("process ending...")
					return
				default:
					fmt.Println("smth went wrong... please input again!")
					continue
			}
			var err = tmpl.Execute(os.Stdout,mappings)
			if err != nil{
				fmt.Println("error executing template")
			}
			

		}

	}

}

