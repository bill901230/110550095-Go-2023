package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// TODO: Create a struct to hold the data sent to the template
type PageData struct {
	Expression string
	Result     string
}

func Calculator(w http.ResponseWriter, r *http.Request) {
	// TODO: Finish this function
	key1 := "op"
	key2 := "num1"
	key3 := "num2"
	var op string = r.URL.Query().Get(key1)
	var num1 string = r.URL.Query().Get(key2)
	var num2 string = r.URL.Query().Get(key3)
	var Error int = 0
	var expr string
	var res string
	var filename string
	n1, err := strconv.Atoi(num1)
	if err != nil {
		Error = 1
	}
	n2, err := strconv.Atoi(num2)
	if err != nil {
		Error = 1
	}
	gcd := GCD(n1, n2)
	if op == "add" {
		expr = num1 + " + " + num2
		res = strconv.Itoa(n1 + n2)
	} else if op == "sub" {
		expr = num1 + " - " + num2
		res = strconv.Itoa(n1 - n2)
	} else if op == "mul" {
		expr = num1 + " * " + num2
		res = strconv.Itoa(n1 * n2)
	} else if op == "div" {
		if n2 != 0 {
			expr = num1 + " / " + num2
			res = strconv.Itoa(n1 / n2)
		} else {
			Error = 1
		}
	} else if op == "gcd" {
		expr = "GCD(" + num1 + ", " + num2 + ")"
		res = strconv.Itoa(gcd)
	} else if op == "lcm" {
		if gcd != 0 {
			expr = "LCM(" + num1 + ", " + num2 + ")"
			res = strconv.Itoa(n1 * n2 / gcd)
		} else {
			Error = 1
		}
	} else {
		Error = 1
	}
	data := PageData{
		Expression: expr,
		Result:     res,
	}
	if Error == 0 {
		filename = "index.html"
	} else {
		filename = "error.html"
	}
	tmpl, err := template.ParseFiles(filename)
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	http.HandleFunc("/", Calculator)
	log.Fatal(http.ListenAndServe(":8084", nil))
}
