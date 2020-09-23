package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"

	"github.com/schollz/progressbar/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db *gorm.DB

	host   string
	port   int
	dbname string
	user   string
	usepw  bool
	passwd string
)

func init() {
	flag.StringVar(&host, "host", "", "")
	flag.StringVar(&host, "h", "127.0.0.1", "TiDB host address")
	flag.IntVar(&port, "port", 4000, "TiDB host port")
	flag.StringVar(&dbname, "db", "test", "Login database")
	flag.StringVar(&user, "user", "", "")
	flag.StringVar(&user, "u", "root", "Login user")
	flag.BoolVar(&usepw, "p", false, "Login password")
}

func initDB() {
	if usepw {
		_, err := fmt.Scanln(&passwd)
		if err != nil {
			panic(err)
		}
	}

	conn_str := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s", user, passwd, host, port, dbname)

	var err error
	db, err = gorm.Open(mysql.Open(conn_str), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		panic(err)
	}
}

const rowNum = 10000000
const batchSize = 10000
const groupNum = 1000

type Table struct{ A, B int }

var tables = []struct {
	name string
	gen  func() int
}{
	{"dense", func() int { return rand.Intn(32) }},
	{"sparse", func() int { return int(rand.Int31()) }},
}

func genData() error {
	data := make([]Table, rowNum)
	for _, table := range tables {
		for i := 0; i < rowNum; i++ {
			data[i] = Table{A: rand.Intn(groupNum), B: table.gen()}
		}

		db.Exec(fmt.Sprintf("drop table if exists %s", table.name))
		db.Exec(fmt.Sprintf("create table %s(a int, b int)", table.name))
		if err := db.Error; err != nil {
			return err
		}
		bar := progressbar.Default(rowNum, table.name)
		for front := 0; front < rowNum; front += batchSize {
			end := front + batchSize
			if end > rowNum {
				end = rowNum
			}
			batch := data[front:end]
			db = db.Table(table.name).Create(&batch)
			if err := db.Error; err != nil {
				return err
			}
			bar.Add(end - front)
		}
		bar.Finish()
	}
	return nil
}

func main() {
	flag.Parse()
	initDB()
	err := genData()
	if err != nil {
		log.Fatal(err)
	}
}
