package action

import "fmt"

type Attack struct {
}

func (s *Attack) TakeAction() {
	fmt.Printf("O player decidiu atacar!")
}
