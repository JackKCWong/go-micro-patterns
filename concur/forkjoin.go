package concur

import "sync"

type ForkJoinGroup struct {
	wg sync.WaitGroup
}

func (g *ForkJoinGroup) Fork(f func()) {
	g.wg.Add(1)
	go func() {
		defer g.wg.Done()
		f()
	}()
}

func (g *ForkJoinGroup) Join(whenDone func()) <-chan bool {
	ch := make(chan bool, 1)
	go func() {
		g.wg.Wait()
		if whenDone != nil {
			whenDone()
		}
		ch <- true
		close(ch)
	}()

	return ch
}
