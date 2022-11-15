package gobasics

import (
	"fmt"
	"net/http"
	"time"
)

/*
Closures or Anonymous functions are useful when you want to define a function inline without having to name it.
The state is unique to that particular function.
Essentially you can think of them like stateful functions, in the sense that they encapsulate state.

	Usage:

	1) Isolating data - To create a function that has access to data that persists even after the function exits and
			don’t want anyone else to have access to that data.
	2) Wrapping functions and creating middleware - Closures are incredibly helpful if we want to wrap our handlers with more logic.
	3) Accessing data that typically isn’t available - Write handler functions that are out of scope and not global variables

References:
1) https://rmoff.net/2020/06/29/learning-golang-some-rough-notes-s01e04-function-closures/
2) http://tleyden.github.io/blog/2016/12/20/understanding-function-closures-in-go/
3) https://www.calhoun.io/5-useful-ways-to-use-closures-in-go/
*/

// PlayWithClosure
func PlayWithClosure() {
	http.HandleFunc("/hello", timed(hello)) // Writing Middleware with closure
	http.ListenAndServe(":3000", nil)

	// Closures Will have access the data in the URL as middleware without creating global objects
	db := NewDatabase("localhost:5432")

	http.HandleFunc("/hello", dbInput(db))
	http.ListenAndServe(":3000", nil)
}

type Database struct {
	Url string
}

func NewDatabase(url string) Database {
	return Database{url}
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Hello!</h1>")
}

// This allows us to bypass the fact that http.HandleFunc() doesn’t permit us passing in
// custom variables without resorting to global variables or anything of that sort.
func dbInput(db Database) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, db.Url)
	}
}

func timed(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		f(w, r)
		end := time.Now()
		fmt.Println("The request took", end.Sub(start))
	}
}

// A SenderFunc is a function that takes no arguments and returns a boolean
// that indicates whether or not the send needs to be retried (in the case of failure)
type SenderFunc func() bool

func sendLoop(sender SenderFunc) {
	for {
		retry := sender()
		if !retry {
			return
		}
		time.Sleep(time.Second)
	}
}

func playWithClosure() {

	counter := 0         // internal state closed over and mutated by mySender function
	maxNumAttempts := 10 // internal state closed over and read by mySender function

	mySender := func() bool {

		// didn't work, any retries left?
		// only retry if we haven't exhausted attempts
		counter += 1
		return counter < maxNumAttempts
	}

	sendLoop(mySender)
}
