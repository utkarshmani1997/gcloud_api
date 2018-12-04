package main

import (
	"fmt"
	"log"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/compute/v1"
)

func main() {
	ctx := context.Background()

	c, err := google.DefaultClient(ctx, compute.CloudPlatformScope)
	if err != nil {
		log.Fatal(err)
	}

	computeService, err := compute.New(c)
	if err != nil {
		log.Fatal(err)
	}

	// Project ID for this request.
	project := "my-project" // TODO: Update placeholder value.

	// The name of the zone for this request.
	zone := "my-zone" // TODO: Update placeholder value.

	rb := &compute.Disk{
		// TODO: Add desired fields of the request body.
	}

	resp, err := computeService.Disks.Insert(project, zone, rb).Context(ctx).Do()
	if err != nil {
		log.Fatal(err)
	}

	// TODO: Change code below to process the `resp` object:
	fmt.Printf("%#v\n", resp)
}
