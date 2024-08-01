package action

import "fmt"

type Defense struct {
}

func (s *Defense) TakeAction() {
	fmt.Printf("O player decidiu defender! \n")
}
