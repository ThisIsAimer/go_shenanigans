package main

import (
	"fmt"
	"net/url"
)

func main() {
	// [scheme://][userinfo@][host][:port][/path][?quary][#fragment]
	//scheme : http/https/ftp
	//userinfo contains username and pass (optional)
	//host - domain name or ip address
	//port - port no. (op)
	//path to the resource of the server
	//quary- key value pairs
	//fragment - specify a location within the resource (op)

	var rawUrl = "https://example.com/path?quary=param#fragment"
	var parsedUrl ,err = url.Parse(rawUrl)
	if err !=nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Sceme:",parsedUrl.Scheme)
	fmt.Println("Host:",parsedUrl.Host)
	fmt.Println("Port:",parsedUrl.Port())
	fmt.Println("Path:",parsedUrl.Path)
	fmt.Println("Quary:",parsedUrl.Query())
	fmt.Println("Raw Quary:",parsedUrl.RawQuery)
	fmt.Println("Fragment:",parsedUrl.Fragment)

	fmt.Println("-------------------------------------------")

	var rawUrl1 = "https://example.com/path?name=john&age=30"
	Url1,err := url.Parse(rawUrl1)
	if err !=nil {
		fmt.Println(err)
		return
	}
	quaryUrl1 := Url1.Query()

	fmt.Println("quary url1:",quaryUrl1)
	fmt.Println("name:", quaryUrl1.Get("name"))
	fmt.Println("age:", quaryUrl1.Get("age"))



	customUrl := *&url.URL{Scheme: "https",Host: "myWife.com",Path: "/path"}
	customQuary := customUrl.Query()
	customQuary.Add("name","Akriti")
	customQuary.Add("age","21")
	customUrl.RawQuery = customQuary.Encode()

	fmt.Println("custom Url:", customUrl.String())

	//customed quary with just values
	values := url.Values{}
	values.Add("bro","code")
	values.Add("language","go")

	var encodedQuary = values.Encode()

	var customUrl1 = "https://broCode.com/path"

	cusQuary := customUrl1 + "?" + encodedQuary

	fmt.Println("custom quary1:",cusQuary)


}