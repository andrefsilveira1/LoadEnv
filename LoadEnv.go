package LoadEnv

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

func LoadEnv(env string) string {
	envFile, err := os.Open(".env")
	if err != nil {
		log.Fatalln("Local variables couldn't be loaded", err)
	}

	file := bufio.NewReader(envFile)
	op := make(chan string)
	result := make(chan string)
	wg := new(sync.WaitGroup)

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go ReadEnv(env, wg, result, op)
	}

	go func() {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			op <- scanner.Text()
			fmt.Println("Inside:", scanner.Text())
		}
		close(op)
	}()

	go func() {
		wg.Wait()
		close(result)
	}()

	iterator := ""
	for i := range result {
		iterator += i
	}
	return strings.Trim(iterator, "\"")
}

func ReadEnv(env string, wg *sync.WaitGroup, result chan<- string, operators <-chan string) {
	defer wg.Done()
	for i := range operators {
		res := strings.Split(i, "=")
		if env == res[0] {
			result <- res[1]
		}
	}
}
