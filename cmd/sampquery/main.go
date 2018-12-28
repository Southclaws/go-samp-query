package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/Southclaws/go-samp-query"
)

func main() {
	var decode = flag.Bool("decode", false, "attempt to decode badly encoded characters")
	flag.Parse()

	if len(flag.Args()) == 0 {
		fmt.Println("Usage: sampquery [-decode] <address>")
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	server, err := sampquery.GetServerInfo(ctx, flag.Arg(0), *decode)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err = json.NewEncoder(os.Stdout).Encode(server); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}
