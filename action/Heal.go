package action

import "fmt"

type Heal struct {
}

func (s *Heal) TakeAction() {
	fmt.Printf("O player decidiu se curar!")
}
