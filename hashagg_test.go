package main

import (
	"flag"
	"fmt"
	"os"
	"testing"
)

var funcs = []string{"avg", "count", "group_concat", "sum"}

func BenchmarkDistinct(b *testing.B) {
	for _, fn := range funcs {
		for _, table := range tables {
			b.Run(fmt.Sprintf("table:%s,func:%s,group:%v", table.name, fn, groupNum), func(b *testing.B) {
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					_, err := db.Table(table.name).Group("a").Select(fmt.Sprintf("%s(distinct b)", fn)).Rows()
					if err != nil {
						b.Fatal(err)
					}
				}
				b.StopTimer()
			})
		}
	}
}

type aggFunc struct {
	name     string
	distinct bool
}

func (f aggFunc) String() string {
	if f.distinct {
		return f.name + "_distinct"
	}
	return f.name
}

var mix_funcs = [][]aggFunc{
	{{"avg", false}, {"count", false}, {"group_concat", false}, {"sum", false}},
	{{"avg", true}, {"count", false}, {"group_concat", false}, {"sum", false}},
	{{"avg", false}, {"count", true}, {"group_concat", false}, {"sum", false}},
	{{"avg", false}, {"count", false}, {"group_concat", true}, {"sum", false}},
	{{"avg", false}, {"count", false}, {"group_concat", false}, {"sum", true}},
	{{"avg", true}, {"count", false}, {"group_concat", false}, {"sum", true}},
	{{"avg", false}, {"count", true}, {"group_concat", false}, {"sum", true}},
	{{"avg", false}, {"count", false}, {"group_concat", true}, {"sum", true}},
	{{"avg", true}, {"count", true}, {"group_concat", false}, {"sum", true}},
	{{"avg", true}, {"count", true}, {"group_concat", true}, {"sum", true}},
}

func aggFuncsToString(fns []aggFunc) (ret []string) {
	for _, fn := range fns {
		if fn.distinct {
			ret = append(ret, fn.name+"(distinct b)")
		} else {
			ret = append(ret, fn.name+"(b)")
		}
	}
	return ret
}

func BenchmarkMix(b *testing.B) {
	for _, fns := range mix_funcs {
		for _, table := range tables {
			b.Run(fmt.Sprintf("table:%s,funcs:%s,group:%v", table.name, fns, groupNum), func(b *testing.B) {
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					_, err := db.Table(table.name).Group("a").Select(aggFuncsToString(fns)).Rows()
					if err != nil {
						b.Fatal(err)
					}
				}
				b.StopTimer()
			})
		}
	}
}

func TestMain(m *testing.M) {
	flag.Parse()
	initDB(false)
	os.Exit(m.Run())
}
