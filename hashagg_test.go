package main

import (
	"flag"
	"fmt"
	"os"
	"testing"

	"gorm.io/hints"
)

var concurrency = []int{1, 4, 8, 15, 30}
var funcs = []string{"avg(distinct b)", "count(distinct b)", "sum(distinct b)"}

func BenchmarkSingleGroup(b *testing.B) {
	for _, table := range tables {
		for _, fn := range funcs {
			for _, con := range concurrency {
				err := db.Exec("set global tidb_executor_concurrency = ?;", con).Error
				if err != nil {
					b.Fatal(err)
				}
				b.Run(fmt.Sprintf("(table:%s,group_num:1,func:%s,concurrency:%v)", table.name, fn, con), func(b *testing.B) {
					b.ResetTimer()
					for i := 0; i < b.N; i++ {
						_, err := db.Clauses(hints.New("HASH_AGG()")).Table(table.name).Select(fn).Rows()
						if err != nil {
							b.Fatal(err)
						}
					}
					b.StopTimer()
				})
			}
		}
	}
}

func Benchmark1000Groups(b *testing.B) {
	for _, table := range tables {
		for _, fn := range funcs {
			for _, con := range concurrency {
				err := db.Exec("set global tidb_executor_concurrency = ?;", con).Error
				if err != nil {
					b.Fatal(err)
				}
				b.Run(fmt.Sprintf("(table:%s,group_num:1000,func:%s,concurrency:%v)", table.name, fn, con), func(b *testing.B) {
					b.ResetTimer()
					for i := 0; i < b.N; i++ {
						_, err := db.Clauses(hints.New("HASH_AGG()")).Table(table.name).Group("a").Select(fn).Rows()
						if err != nil {
							b.Fatal(err)
						}
					}
					b.StopTimer()
				})
			}
		}
	}
}

func TestMain(m *testing.M) {
	flag.Parse()
	initDB()
	os.Exit(m.Run())
}
