package main

import (
	"fmt"
	"net/http"

	"github.com/Faizal-Asep/loadBalancer/lb"
)

func main() {
	servers := []lb.Server{
		lb.NewSimpleServer("http://127.0.0.1:8000"),
		lb.NewSimpleServer("http://127.0.0.1:8001"),
		lb.NewSimpleServer("http://127.0.0.1:8002"),
	}

	balancer := lb.NewLoadBalancer("8080", servers)
	handleRedirect := func(rw http.ResponseWriter, req *http.Request) {
		balancer.ServeProxy(rw, req)
	}

	// register a proxy handler to handle all requests
	http.HandleFunc("/", handleRedirect)

	fmt.Printf("serving requests at 'localhost:%s'\n", balancer.Port)
	http.ListenAndServe(":"+balancer.Port, nil)
}
