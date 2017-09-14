package main

import (
	"log"
	"net/http"

	vproxy "github.com/vamproxy/core/httpproxy"
)

func main() {
	proxy, err := vproxy.DefaultProxy("", "8888", &customVampire{})

	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	proxy.Start()
}

type customVampire struct {
}

func (v *customVampire) OnIncomingRequest(reqDump []byte, req *http.Request) {
	log.Printf("incoming request: \n%s\n\n", string(reqDump))
}

func (v *customVampire) OnOutgoingRequest(reqDump []byte, req *http.Request) {
	log.Printf("outgoing request: \n%s\n\n", string(reqDump))
}

func (v *customVampire) OnIncomingResponse(resDump []byte, res *http.Response) {
	log.Printf("incoming response: \n%s\n\n", string(resDump))
}
