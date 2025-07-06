// Concurrency :
// Concurrency is about dealing with multiple tasks at once, but not necessarily doing them simultaneously. It involves the management of multiple tasks, allowing them to make progress by interleaving execution. In a concurrent system, multiple tasks can start, run, and complete in overlapping time periods.

// Parallelism
// Parallelism is about executing multiple tasks simultaneously. It requires hardware with multiple processing units (cores). In a parallel system, multiple tasks literally run at the same time.

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// fetchURL fetches the content of the URL and sends the response body length to the results channel
func fetchURL(url string, results chan string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		results <- fmt.Sprintf("Error fetching %s: %v", url, err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		results <- fmt.Sprintf("Error reading response from %s: %v", url, err)
		return
	}

	elapsed := time.Since(start).Seconds()
	results <- fmt.Sprintf("Fetched %s in %.2f seconds. Length: %d bytes", url, elapsed, len(body))
}

func main() {
	// List of URLs to fetch
	urls := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.linkedin.com",
		"https://www.github.com",
	}

	// Channel to receive the results
	results := make(chan string)

	// Launch a goroutine for each URL
	for _, url := range urls {
		go fetchURL(url, results)
	}

	// Receive the results from all goroutines
	for range urls {
		fmt.Println(<-results)
	}
}

// Parellism

// package main

// import (
//     "fmt"
//     "runtime"
//     "time"
// )

// func task(name string) {
//     for i := 0; i < 3; i++ {
//         fmt.Println(name, "running")
//         time.Sleep(100 * time.Millisecond)
//     }
// }

// func main() {
//     runtime.GOMAXPROCS(2) // Set the number of threads to the number of CPU cores
//     go task("Task1")
//     go task("Task2")
//     time.Sleep(1 * time.Second)
// }

//Concurrency:
// Manages multiple tasks by interleaving their execution.
// Can be achieved on a single-core processor.
// Tasks appear to run simultaneously but actually do not.

// Parallelism:
// Executes multiple tasks simultaneously.
// Requires multiple cores or processors.
// Tasks truly run at the same time.
