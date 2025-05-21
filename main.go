package main

import (
	handleform "3ML/handleForm"
)

func main() {
	output, err := handleform.CreateForm()
	if err != nil {
		panic(err)
	}
	println(output)
}
