package utils

import "sync"

// merge all error channels to a single channel
func MergeErrors(cs ...<-chan error) <-chan error {
	var wg sync.WaitGroup
	out := make(chan error, len(cs))

	// read function for each channel
	readChannel := func(c <-chan error) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))

	for _, c := range cs {
		go readChannel(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
