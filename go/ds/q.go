package "ds"

type Queue struct {
    data []interface{}
}

func(q *Queue) Len() int {
    return len(q.data)
}

func(q *Queue) Front() interface{} {
    last := q.Peek()
    q.data = q.data[1:]
    return last
}

func(q *Queue) Peek() interface{} {
    return q.data[0]
}

func(q *Queue) Push(e interface{}) {
    q.data = append(q.data, e)
}