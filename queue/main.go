package main

import "fmt"

type Queue struct {
	slice []int
}

func (q *Queue) Engueue(data int) {
	q.slice = append(q.slice, data)
}

func (q *Queue) Dequeue() int {
	data := q.slice[0]
	q.slice = q.slice[1:len(q.slice)]
	return data
}

func main() {

	new_q := Queue{}
	new_q.Engueue(10)
	new_q.Engueue(1)
	new_q.Engueue(12)
	fmt.Println(new_q)

	fmt.Println(new_q.Dequeue())
}
