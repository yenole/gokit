package scaffold

import "sync"

type Scaffold struct {
}

func New() *Scaffold {
	return &Scaffold{}
}

func (sc *Scaffold) Setup() error {
	return nil
}

func (sc *Scaffold) Run() {
	var wg sync.WaitGroup

	wg.Wait()
}
