package main

import (
	"flag"
	"fmt"
	"os/exec"
	"runtime"
	"time"
)

type wg struct {
	i int
	c chan struct{}
}

//type wg chan struct{}

func makewg(i int) (w wg) {
	w.i = i
	w.c = make(chan struct{}, i)
	return
}
func (w *wg) take() { w.c <- struct{}{} }
func (w *wg) put()  { <-w.c }
func (w *wg) drain() {
	for j := 0; j < w.i; j++ {
		w.take()
	}
}

func run(c *exec.Cmd) {
	fmt.Printf("running %q %q\n", c.Path, c.Args)
	o, err := c.CombinedOutput()
	fmt.Printf("%q %q -> %+v\n%s\n", c.Path, c.Args, err, string(o))
}

func optimize(w *wg, fn string) {
	w.take()
	defer w.put()

	run(exec.Command("c:/utils/optipng.exe", "-o7", "-zm1-9", fn))
	run(exec.Command("c:/utils/advpng.exe", "-z4", fn))

}

func main() {
	procs := flag.Int("c", runtime.NumCPU(), "maximum number of concurrent processes")
	flag.Parse()
	w := makewg(*procs)
	for _, v := range flag.Args() {
		go optimize(&w, v)
	}
	time.Sleep(1 * time.Second)
	w.drain()
}
