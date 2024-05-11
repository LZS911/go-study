package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {

		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			url = fmt.Sprintf("http://%s", url)
		}

		resp, err := http.Get(url)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Status code %s\n", resp.Status)

		// body, err := io.ReadAll(resp.Body)
		// resp.Body.Close()

		written, err := io.Copy(os.Stdout, resp.Body)

		fmt.Printf("\n请求获得的字节数量%d\n", written)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

	}
}
