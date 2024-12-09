package concurrency

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			// Send statement
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	// One at a time now to unpack
	for i := 0; i < len(urls); i++ {
		// Receive expression
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}