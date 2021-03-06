// +build ignore

package main

import (
	"context"
	"log"
	"net"

	"github.com/goinsane/accepter"
)

func main() {
	a := &accepter.Accepter{
		Handler: accepter.HandlerFunc(func(ctx context.Context, conn net.Conn) {
			for {
				var b [32 * 1024]byte
				n, err := conn.Read(b[:])
				if err != nil {
					break
				}
				m, err := conn.Write(b[:n])
				if err != nil {
					break
				}
				if m < n {
					break
				}
			}
		}),
	}
	log.Fatal(a.ListenAndServe("tcp", ":1234"))
}
