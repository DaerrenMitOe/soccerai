package main

import (
	"soccerai/dataset"
	"soccerai/db"
)

func main() {
	dataset.Download()
	db.LoadTest()
}
