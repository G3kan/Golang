package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/* Dining Philosophers Problem

Aşağıdaki kısıtlamalar ve değişikliklerle birlikte yemek yiyen filozoflar problemini (Dining Philosophers Problem) uygulandı:

Toplam 5 filozof olacak ve her birinin arasında birer çubuk (chopstick) paylaşılacak.

Her filozof yalnızca 3 kez yemek yemeli (derste yaptığımız gibi sonsuz döngüde değil).

Filozoflar çubukları sırasız (en küçük numaralıyı ilk almadan) rastgele sırayla almalı.

Yemek yiyebilmek için, kendi goroutine’inde çalışan bir ev sahibinden (host) izin almaları gerekiyor.

Ev sahibi aynı anda en fazla 2 filozofun yemek yemesine izin verecek.

Her filozof 1’den 5’e kadar numaralandırılmıştır.

Bir filozof yemek yemeye başladığında (gerekli kilitleri aldıktan sonra) şu satır yazdırılır:
starting to eat <sayı>
Örneğin: starting to eat 3

Bir filozof yemeği bitirdiğinde (kilitleri bırakmadan önce) şu satır yazdırılır:
finishing eating <sayı>
Örneğin: finishing eating 3

*/

// Philosopher struct

type Philosopher struct {
	id               int
	leftChopstick    *sync.Mutex
	rightChopstick   *sync.Mutex
	timesEaten       int
	hostPermissionCh chan int
	doneEating       chan int
}

// Host struct
type Host struct {
	permitChan chan int
	doneChan   chan int
}

// NewHost initializes the host
func NewHost() *Host {
	return &Host{
		permitChan: make(chan int),
		doneChan:   make(chan int),
	}
}

// Run host logic - allow max 2 philosophers to eat
func (h *Host) Run() {
	activeEaters := 0
	queue := []int{}

	for {
		select {
		case id := <-h.permitChan:
			if activeEaters < 2 {
				activeEaters++
				h.doneChan <- id
			} else {
				queue = append(queue, id)
			}
		case <-h.doneChan:
			activeEaters--
			if len(queue) > 0 {
				nextID := queue[0]
				queue = queue[1:]
				activeEaters++
				h.doneChan <- nextID
			}
		}
	}
}

// Eat method for each philosopher
func (p *Philosopher) Eat(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		p.hostPermissionCh <- p.id // ask host permission
		<-p.doneEating             // wait for host approval

		// Pick up chopsticks in random order
		if rand.Intn(2) == 0 {
			p.leftChopstick.Lock()
			p.rightChopstick.Lock()
		} else {
			p.rightChopstick.Lock()
			p.leftChopstick.Lock()
		}

		fmt.Printf("starting to eat %d\n", p.id)
		time.Sleep(time.Millisecond * time.Duration(100+rand.Intn(200)))
		fmt.Printf("finishing eating %d\n", p.id)

		p.rightChopstick.Unlock()
		p.leftChopstick.Unlock()

		p.doneEating <- p.id // notify host eating is done
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Create 5 chopsticks
	chopsticks := make([]*sync.Mutex, 5)
	for i := range chopsticks {
		chopsticks[i] = &sync.Mutex{}
	}

	// Create host
	host := NewHost()
	go host.Run()

	var wg sync.WaitGroup
	philosophers := make([]*Philosopher, 5)

	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{
			id:               i + 1,
			leftChopstick:    chopsticks[i],
			rightChopstick:   chopsticks[(i+1)%5],
			hostPermissionCh: host.permitChan,
			doneEating:       host.doneChan,
		}
	}

	for _, p := range philosophers {
		wg.Add(1)
		go p.Eat(&wg)
	}

	wg.Wait()
}
