package main

import (
	"bufio"
	"fmt"
	"grt_test/api"
	"grt_test/safecounter"
	"grt_test/shapes"
	"grt_test/twosum"
	"grt_test/workerpool"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n--- Golang Practice Modules Menu ---")
		fmt.Println("1. Basic Concurrency: Worker Pool")
		fmt.Println("2. Thread-Safety: Safe Counter")
		fmt.Println("3. Interfaces: Shape Area System")
		fmt.Println("4. Logic & Map: Two Sum")
		fmt.Println("5. HTTP Handler: Start JSON API Server")
		fmt.Println("q. Exit")
		fmt.Print("Select an option: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			fmt.Print("Enter number of workers: ")
			wStr, _ := reader.ReadString('\n')
			w, _ := strconv.Atoi(strings.TrimSpace(wStr))
			fmt.Print("Enter number of jobs: ")
			jStr, _ := reader.ReadString('\n')
			j, _ := strconv.Atoi(strings.TrimSpace(jStr))
			fmt.Printf("Running Worker Pool with %d workers and %d jobs...\n", w, j)
			workerpool.RunWorkers(w, j)

		case "2":
			fmt.Print("Enter number of concurrent increments: ")
			nStr, _ := reader.ReadString('\n')
			n, _ := strconv.Atoi(strings.TrimSpace(nStr))
			counter := safecounter.SafeCounter{}
			var wg sync.WaitGroup
			fmt.Printf("Incrementing counter %d times concurrently...\n", n)
			for i := 0; i < n; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					counter.Inc()
				}()
			}
			wg.Wait()
			fmt.Printf("Final Counter Value: %d\n", counter.Value())

		case "3":
			fmt.Println("Choose Shape: (1) Rectangle (2) Circle")
			sChoice, _ := reader.ReadString('\n')
			if strings.TrimSpace(sChoice) == "1" {
				fmt.Print("Enter Width: ")
				wS, _ := reader.ReadString('\n')
				w, _ := strconv.ParseFloat(strings.TrimSpace(wS), 64)
				fmt.Print("Enter Height: ")
				hS, _ := reader.ReadString('\n')
				h, _ := strconv.ParseFloat(strings.TrimSpace(hS), 64)
				shapes.PrintArea(shapes.Rectangle{Width: w, Height: h})
			} else {
				fmt.Print("Enter Radius: ")
				rS, _ := reader.ReadString('\n')
				r, _ := strconv.ParseFloat(strings.TrimSpace(rS), 64)
				shapes.PrintArea(shapes.Circle{Radius: r})
			}

		case "4":
			fmt.Println("Example: nums = [2, 7, 11, 15], target = 9")
			nums := []int{2, 7, 11, 15}
			target := 9
			result := twosum.TwoSum(nums, target)
			fmt.Printf("TwoSum result for target %d: %v\n", target, result)

		case "5":
			fmt.Println("Starting API Server on :8080...")
			fmt.Println("Endpoints:")
			fmt.Println("  POST /hello (Body: {\"name\": \"...\"})")
			fmt.Println("  GET  /terminate (to shut down via HTTP)")

			go api.StartServer()

			fmt.Println("Server is running. Enter 'q' to stop server and return to menu.")
			for {
				stopInput, _ := reader.ReadString('\n')
				if strings.TrimSpace(stopInput) == "q" {
					api.ShutdownServer()
					break
				}
			}

		case "q":
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid option, try again.")
		}
	}
}
