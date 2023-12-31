package main

import (
	"sync"

	"log"
	"time"
)

type Worker interface {
	Task()
}

type Pool struct {
	work chan Worker
	wg   sync.WaitGroup
}

func New(maxGoroutines int) *Pool {
	p := Pool{
		work: make(chan Worker),
	}

	p.wg.Add(maxGoroutines)
	for i := 0; i < maxGoroutines; i++ {
		go func() {
			for w := range p.work {
				w.Task()
			}
			p.wg.Done()
		}()
	}

	return &p
}

func (p *Pool) Run(w Worker) {
	p.work <- w
}

func (p *Pool) Shutdown() {
	close(p.work)
	p.wg.Wait()
}

// 测试
var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
}

type namePrinter struct {
	name string
}

func (n *namePrinter) Task() {
	time.Sleep(time.Second)
	log.Println(n.name)
}

func main() {
	p := New(2)
	for i := 0; i < 1; i++ {
		for _, name := range names {
			np := namePrinter{name}
			p.Run(&np)
			log.Println(1)
		}
	}
	p.Shutdown()
}
