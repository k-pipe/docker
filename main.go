package memtestpackage

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	size, _ := strconv.Atoi(GetEnv("MEM_SIZE", "1"))
	fmt.Println("Allocating ", size, " MB ram...")
	allocateMem(size)
	time.Sleep(time.Duration(30) * time.Minute)
	fmt.Println("Done")
}

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

func allocateMem(sizeInMB int) {
	size := sizeInMB * 1024 * 1024
	mem := make([]byte, size)
	for i := 0; i < size; i++ {
		mem[i] = 1
	}
}
