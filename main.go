package main

import (
	test "github.com/Andrew-M-C/go-tools/_test"
	"github.com/Andrew-M-C/go-tools/log"
)

func main() {
	test.TestSqlToJson()
	test.TestJsonValue()
	test.TestStr()
	for i := 0; i < 200; i ++ {
		test.TestOrigJsonEffenciency()
	}
	log.Info("demo done")
	return
}
