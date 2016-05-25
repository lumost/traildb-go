package main

import (
	"fmt"

	"github.com/lumost/traildb-go"
)

type Ev struct {
	Timestamp int    `tdb:"timestamp"`
	Field1    string `tdb:"field1"`
	Field2    string `tdb:"field2"`
}

type WikiEvent struct {
	Timestamp int64  `tdb:"timestamp"`
	User      string `tdb:"user"`
	IP        string `tdb:"ip"`
	Title     string `tdb:"title"`
}

func main() {
	db, err := tdb.Open("wikipedia-history-small.tdb")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(db)
	fmt.Println(db.Version())

	var total int
	for i := 0; i < 10; i++ {
		trail, err := tdb.NewTrail(db, i)
		if err != nil {
			panic(err.Error())
		}
		// fmt.Println(trail)
		for {
			evt := trail.NextEvent()
			if evt == nil {
				trail.Close()
				break
			}
			total++
			fmt.Println(evt.ToMap())
			// Convert to struct
			we := WikiEvent{}
			fmt.Printf("struct format %+v", evt.ToStruct(we))
		}
	}
}
