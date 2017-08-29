package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s {TIMESTAMP}\n", os.Args[0])
		return
	}

	end, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	duration := time.Unix(int64(end), 0).Sub(time.Now())
	if duration.Seconds() < 0 {
		fmt.Println("BLASTOFF!!!!")
		return
	}

	fmt.Printf("%d:%02d:%02d\n", int(duration.Hours()), int(duration.Minutes())%60, int(duration.Seconds())%60)
}
