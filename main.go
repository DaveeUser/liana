package liana

import (
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"liana/memory"

	"github.com/gin-gonic/gin"
)

var (
	RandomKeyError = errors.New("error generate key hash")
)

// The RandomSHA256Key function generates a 256-bit SHA-256
// random key and returns its hexadecimal representation in string format.

func RandomSHA256Key() (string, error) {
	var key [32]byte
	n, err := rand.Read(key[:])
	if n != len(key) || err != nil {
		return "", RandomKeyError
	}

	hasher := sha256.New()
	_, err = hasher.Write(key[:])
	if err != nil {
		return "", RandomKeyError
	}

	result := hasher.Sum(nil)
	return fmt.Sprintf("%x", result[:10]), nil
}

//star of the services

func main() {

	var limit int
	var port int

	flag.IntVar(&limit, "limit", 1, "limit flag of the memory")
	flag.IntVar(&port, "port", 8080, "port of the services")

	flag.Parse()

	mem := memory.NewMemory(limit)

	r := gin.Default()

	// Code block shows how you can define a "/send" route in Gin to handle HTTP POST requests. In this case, it will read the
	// data sent in the request body, trim whitespace, and finally generate a random SHA256 key and input an input element to a mem instance.

	r.POST("/send", func(c *gin.Context) {

		data, err := ioutil.ReadAll(c.Request.Body)

		myData := strings.TrimSpace(string(data))

		key, err := RandomSHA256Key()
		if err != nil {
			log.Fatal(err)
		}

		mem.Load(key, myData)

		c.String(http.StatusOK, key)
	})

	// This block of code defines a new GET route for your Gin web application. Any request directed to `/:id`
	// will be served by this anonymous function. Inside that function, it retrieves the value associated with `:id`, then calls the `Get()`
	// function of the `mem` instance. Finally, it returns a 200 status message and the result of `Get()`.

	r.GET("/:id", func(c *gin.Context) {

		id := c.GetString("id")

		data := mem.Get(id)

		c.String(http.StatusOK, data)
	})

	http.ListenAndServe(":"+strconv.Itoa(port), r)
}
