package main

import (
	"context"
	"log"
	"testing"
)

func Test_Lambda(t *testing.T) {
	_, err := HandleRequest(context.Background(), []byte{})
	if err != nil {
		log.Fatalln(err)
	}
}
