package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"os/exec"
	"syscall"
	"time"
)

const (
	CACHE_LINE_SIZE = 64
	NACCESS         = 128 * 1024 * 1024
)

func main() {
	_ = os.Remove("out.txt")
	f, err := os.OpenFile("out.txt", os.O_CREATE|os.O_RDWR, 0660)
	if err != nil {
		log.Fatal("openfile()に失敗しました")
	}
	defer f.Close()

	for i := 2.0; i <= 16.0; i += 0.25 {
		bufSize := int(math.Pow(2, i)) * 1024
		data, err := syscall.Mmap(-1, 0, bufSize, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_ANON|syscall.MAP_PRIVATE)
		defer syscall.Munmap(data)
		if err != nil {
			log.Fatal("mmap()に失敗しました")
		}

		fmt.Printf("バッファサイズ 2^%.2f(%d) KiB についてのデータを収集中...\n", i, bufSize/1024)
		start := time.Now()
		for i := 0; i < NACCESS/(bufSize/CACHE_LINE_SIZE); i++ {
			for j := 0; j < bufSize; j += CACHE_LINE_SIZE {
				data[j] = 0
			}
		}
		end := time.Since(start)
		f.Write([]byte(fmt.Sprintf("%f\t%f\n", i, float64(NACCESS)/float64(end.Nanoseconds()))))
	}
	command := exec.Command("./plot-cache.py")
	out, err := command.Output()
	if err != nil {
		fmt.Fprintf(os.Stderr, "コマンド実行に失敗しました: %q: %q", err, string(out))
		os.Exit(1)
	}
}
