package main

import (
	"flag"
	"fmt"
)

var (
	certFile string
	keyFile  string
)

func init() {
	flag.StringVar(&certFile, "cert", "test.crt", "TLS Cert File")
	flag.StringVar(&keyFile, "key", "test.key", "TLS Key File")
	flag.Parse()
}

func main() {
	var block = make(chan struct{})
	go mainGRPC("localhost:4443")
	go mainREST("localhost:4444")
	<-block
}

// Validatable interface to describe what is validatable
type Validatable interface {
	Validate() error
}

// validationErrors - custom error type that can take many errors
type validationErrors []error

// Error - implementation of error interface
func (ve validationErrors) Error() string {
	var errStr string
	for _, v := range ve {
		errStr += fmt.Sprintf("%s\n", v.Error())
	}
	return errStr
}

func validate(v Validatable) error {
	return v.Validate()
}
