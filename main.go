package main

import (
	handleform "3ML/handleForm"
)

func main() {
	err := handleform.CreateForm()
	if err != nil {
		panic(err)
	}
	handleform.AddOptions()
}
