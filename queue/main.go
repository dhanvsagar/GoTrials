package main

import "fmt"

type Queue struct {
	slice []int
}

func (q *Queue) Engueue(data int) {
	q.slice = append(q.slice, data)
}

func (q *Queue) Dequeue() (int, error) {
	if len(q.slice) < 0 {
		return -1, fmt.Errorf("No items in the queue")
	} else {
		data := q.slice[0]
		q.slice = q.slice[1:len(q.slice)]
		return data, nil
	}

}

func main() {
	new_q := Queue{}
	new_q.Engueue(10)
	new_q.Engueue(1)
	new_q.Engueue(12)
	fmt.Println(new_q)

	data, err := new_q.Dequeue()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(data)
	}

	fmt.Println(new_q)
}
