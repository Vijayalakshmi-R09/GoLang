package main

import "Practice_new/employee"

func main() {
	e := employee.New("Sam", "Adolf", 30, 20)
	e.LeavesRemaining()
}
