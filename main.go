package main

import (
	"fmt"
	"math/rand"
	"time"
)


var pool = "ABCDE"


// Generate a random string of A-Z chars with len = l
func randomString(l int) string {
bytes := make([]byte, l)
for i := 0; i < l; i++ {
bytes[i] = pool[rand.Intn(len(pool))]
}
return string(bytes)
}


func main() {
rand.Seed(time.Now().UnixNano())
fmt.Println(randomString(5))
}