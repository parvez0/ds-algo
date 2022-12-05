package "ds"

type Stack struct {
    data []interface{}
}

func (s *Stack) Len() int {
    return len(s.data)
}

func (s *Stack) Pop() interface{} {
    last := s.Top()
    s.data = s.data[:s.Len()-1]
    return last
}

func (s *Stack) Top() interface{} {
    return s.data[s.Len()-1]
}

func (s *Stack) Push(e interface{}) {
    s.data = append(s.data, e)
}