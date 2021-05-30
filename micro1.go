package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func callMicro2(ctx *gin.Context) string {
	tracingHeaders := []string{
		"x-request-id",
		"x-b3-traceid",
		"x-b3-spanid",
		"x-b3-sampled",
		"x-b3-parentspanid",
		"x-b3-flags",
		"x-ot-span-context",
	}
	headersToSend := make(map[string]string)

	for _, key := range tracingHeaders {
		if val := ctx.Request.Header.Get(key); val != "" {
			headersToSend[key] = val
		}
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://micro2:8081/call", nil)
	if err != nil {
		log.Fatal(err)
	}

	for clave, valor := range headersToSend {
		req.Header.Add(clave, valor)
	}
	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
	cuerpo := " Micro1 llamando a " + string(body)

	return cuerpo

}

func setupRouter() *gin.Engine {
	r := gin.Default()
	//tracer := opentracing.GlobalTracer()
	//r.Use(ginhttp.Middleware(tracer))
	r.GET("/call", func(ctx *gin.Context) {
		ctx.String(200, callMicro2(ctx))
	})
	r.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
