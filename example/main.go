package main

import (
	"time"

	"github.com/tsny-houzz/uilive"
)

func main() {
	writer := uilive.New().Start()

	for _, f := range [][]string{{"Foo.zip", "Bar.iso"}, {"Baz.tar.gz", "Qux.img"}} {
		for i := 0; i <= 50; i++ {
			writer.Printf("Downloading %s.. (%d/%d) GB\n", f[0], i, 50)
			writer.Printf("Downloading %s.. (%d/%d) GB\n", f[1], i, 50)
			time.Sleep(time.Millisecond * 15)
		}
	}
	writer.Printf("Finished: Downloaded 150GB")

	writer = uilive.New().Start()
	writer.Printf("Waiting...")
	time.Sleep(2 * time.Second)
	writer.Printf("Done!")

	writer.Stop() // flush and stop rendering
}
