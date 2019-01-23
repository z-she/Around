package main

import (
	"Around/utils"
	"context"
	"fmt"
	"strconv"

	"cloud.google.com/go/bigtable"
)

func saveToBT(p *Post, id string) {
	ctx := context.Background()
	btClient, err := bigtable.NewClient(ctx, utils.PROJECT_ID, utils.BT_INSTANCE) //, option.WithCredentialsFile(CREDENTIAL))
	if err != nil {
		panic(err)
	}
	tbl := btClient.Open("post")
	mut := bigtable.NewMutation()
	t := bigtable.Now()
	mut.Set("post", "user", t, []byte(p.User))
	mut.Set("post", "message", t, []byte(p.Message))
	mut.Set("location", "lat", t, []byte(strconv.FormatFloat(p.Location.Lat, 'f', -1, 64)))
	mut.Set("location", "lon", t, []byte(strconv.FormatFloat(p.Location.Lon, 'f', -1, 64)))

	err = tbl.Apply(ctx, id, mut)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Post is saved to BigTable: %s\n", p.Message)
	return
}
