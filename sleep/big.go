package main

import (
	"fmt"
	"math/big"
	"net/http"
	"time"
)

func factorial(n int64) *big.Int {
	result := big.NewInt(1)
	for i := int64(1); i <= n; i++ {
		result.Mul(result, big.NewInt(i))
	}
	return result
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	result := factorial(100000)
	elapsed := time.Since(start)
	fmt.Fprintf(w, "Elapsed time: %s", elapsed)
	_ = result // To avoid compiler warning
}

func main() {
	http.HandleFunc("/", myHandler)
	http.ListenAndServe(":8080", nil)
}
