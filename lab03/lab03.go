package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func solve(v []string) string {
	var s string = ""
	var a int
	var b int
	var err error
	if len(v) != 4 {
		return "Error!"
	}
	if a, err = strconv.Atoi(v[2]); err != nil {
		return "Error!"
	}
	if b, err = strconv.Atoi(v[3]); err != nil {
		return "Error!"
	}
	if v[1] == "add" {
		s += v[2] + " + " + v[3] + " = " + strconv.Itoa(a+b)
	} else if v[1] == "sub" {
		s += v[2] + " - " + v[3] + " = " + strconv.Itoa(a-b)
	} else if v[1] == "mul" {
		s += v[2] + " * " + v[3] + " = " + strconv.Itoa(a*b)
	} else if v[1] == "div" {
		if v[3] == "0" {
			return "Error!"
		}
		s += v[2] + " / " + v[3] + " = " + strconv.Itoa(a/b) + ", reminder = " + strconv.Itoa(a%b)
	} else {
		return "Error!"
	}
	return s
}

func Calculator(w http.ResponseWriter, r *http.Request) {
	// TODO: implement a calculator
	str := html.EscapeString(r.URL.Path)
	slice := strings.Split(str, "/")
	// fmt.Printf("%s", solve(slice))
	fmt.Fprintf(w, "%s", solve(slice))
}

func main() {
	http.HandleFunc("/", Calculator)
	log.Fatal(http.ListenAndServe(":8083", nil))
}
