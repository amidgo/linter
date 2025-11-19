package main

import (
	"context"
	"net/http"
)

func main() {
	_, _ = http.NewRequestWithContext(context.Background(), http.MethodGet, "/path", http.NoBody)
}
