package main

import (
	"context"
	"fmt"
	"math"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func sub() {
	fmt.Println("sub() is running")
	time.Sleep(time.Second)
	fmt.Println("sub() is finished")
}

func main() {
	// fmt.Println("start sub()")
	// go sub()
	// time.Sleep(2 * time.Second)

	// fmt.Println("start anonymous sub()")
	// go func() {
	// 	fmt.Println("anonymous sub() is running")
	// 	time.Sleep(time.Second)
	// 	fmt.Println("anonymous sub() is finished")
	// }()

	// time.Sleep(2 * time.Second)

	fmt.Println("start sub()")
	done := make(chan bool)
	go func() {
		fmt.Println("sub() is finished")
		// こんな感じにアクセスするのはやっぱり気持ち悪いなぁとは思う
		// あー、contextとかいうので関数渡しにするのか
		done <- true
	}()
	// チャンネル開きっぱなしで止めると　all goroutines are asleep とか出て止まるの賢い
	<-done
	fmt.Println("all tasks are finished")

	pn := primeNumber()
	for n:= range pn {
		fmt.Println(n)
	}

	receiver := make(chan int)
	exiter := make(chan bool)

	go func() {
		receiver <- 10
	}()

	go func() {
		time.Sleep(time.Second)
		exiter <- true
	}()

	select {
	case data := <- receiver:
		fmt.Printf("data comes! %v\n", data)
	case <- exiter:
		fmt.Println("finished!")
	}

	fmt.Println("start context sub()")
	
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		fmt.Println("context sub() is finished!")
		time.Sleep(time.Second)
		cancel()
	}()
	<-ctx.Done()
	fmt.Println("all context tasks are finished")

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT)
	fmt.Println("Waiting SIGINT (CTRL + C)")
	<-signals
	fmt.Println("SIGINT arrived")

	q4_1(5)
}

func q4_1(waitSecond int64) {
	wait := time.After(time.Duration(waitSecond * int64(time.Second)))
	<-wait
}

func primeNumber() chan int {
	result := make(chan int)
	go func() {
		result <- 2
		for i:= 3; i < 100000; i += 2 {
			l := int(math.Sqrt(float64(i)))
			found := false
			for j := 3; j < l + 1; j += 2 {
				if i % j == 0 {
					found = true
					break;
				}
			}
			if !found {
				result <- i
			}
		}
		close(result)
	}()
	return result
}