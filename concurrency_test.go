package notarealpackage

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

// var rpcQ []rpc.RPC
var intQ = []int{}
var activeBool = false

var oop = []int{}
var op = 0

var mutQ sync.Mutex
var mutB sync.Mutex

var g sync.WaitGroup

func Test_conc(t *testing.T) {
	g.Add(2)
	go putterLoop(100)
	go putterLoop(200)
	g.Wait()
	//fmt.Println("Final order of operations: ", oop)

	// Separate operations
	op1 := []int{}
	op2 := []int{}
	for index, i := range oop {
		if i >= 100 && i < 200 {
			op1 = append(op1, oop[index])
		} else {
			op2 = append(op2, oop[index])
		}
	}

	//fmt.Println("Separated:\n", op1, "\n", op2)

	var before = -1
	for _, i := range op1 {
		if before > i {
			t.Error("Failed order of operations in op1")
		}

		before = i
	}

	for _, i := range op2 {
		if before > i {
			t.Error("Failed order of operations in op2")
		}

		before = i
	}
}

func putterLoop(prefix int) {
	for i := prefix; i < 100+prefix; i++ {
		if rand.Intn(5) > 3 {
			time.Sleep(time.Second)
		}

		AddToQueue(i)
	}
	time.Sleep(time.Second)
	g.Done()
}

func AddToQueue(added int) {
	now := time.Now()
	mutQ.Lock()
	intQ = append(intQ, added)
	//fmt.Println("q: ", intQ)
	mutQ.Unlock()

	if waitTime := time.Since(now).Microseconds(); waitTime != 0 {
		//fmt.Println("!!! delay !!! microseconds waiting for add to q: ", waitTime)
	}
	go initHandleLoop()
}

func initHandleLoop() {
	mutB.Lock()
	if activeBool {
		mutB.Unlock()
		return
	} else {
		activeBool = true
		mutB.Unlock()
		handleLoop()
	}
}

func handleLoop() {
	for {
		mutB.Lock()
		mutQ.Lock()
		if len(intQ) > 0 {
			qHead := intQ[0]
			intQ = intQ[1:]
			oop = append(oop, qHead)
			mutB.Unlock()
			mutQ.Unlock()
		} else {
			activeBool = false
			mutB.Unlock()
			mutQ.Unlock()
			break
		}
	}
}
