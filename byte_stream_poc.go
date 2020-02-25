package main

import (
	"fmt"
	"strconv"
	"time"
)

type Source struct {
	downstream chan []byte
	fn         func() //produces multiple data into downstream
}

func (s *Source) mapto(fn func([] byte) []byte) *Mappable {
	return &Mappable{
		upstream:   s.downstream,
		downstream: make(chan [] byte),
		fn:         fn,
	}
}

type Mappable struct {
	upstream   chan []byte
	downstream chan []byte
	fn         func([] byte) []byte
}

func (m *Mappable) mapto(fn func([] byte) []byte) *Mappable {
	n := &Mappable{
		upstream:   m.downstream,
		downstream: make(chan [] byte),
		fn:         fn,
	}
	go n.start()
	return n
}

func (m *Mappable) start() {
	m.downstream <- m.fn(<-m.upstream)
}

func (m *Mappable) Subscribe(success func([]byte)) {
	b := <-m.downstream
	success(b)
}

func square(b []byte) []byte {
	s := string(b)
	i, _ := strconv.Atoi(s)
	n := i * i
	return []byte(strconv.Itoa(n))
}
func double(b []byte) []byte {
	s := string(b)
	i, _ := strconv.Atoi(s)
	n := 2 * i
	return []byte(strconv.Itoa(n))
}

func main() {
	d := make(chan [] byte)
	until100 := Source{
		downstream: d,
		fn: func() {
			go func() {
				for i := 1; i <= 100; i++ {
					time.Sleep(100 * time.Millisecond)
					d <- []byte(strconv.Itoa(i))
				}
			}()
		},
	}

	until100.
		mapto(square).
		mapto(double).
		Subscribe(
			func(b []byte) {
				fmt.Println(string(b))
			},
		)
}
