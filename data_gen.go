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

const (
	defaultHost   = "127.0.0.1"
	defaultPort   = 4000
	defaultDBName = "test"
	defaultUser   = "root"

	rowNum    = 1000000
	batchSize = 1000
)

var (
	db *gorm.DB

	host   string
	port   int
	dbname string
	user   string
	usepw  bool
	passwd string

	verbose bool
)

func init() {
	flag.StringVar(&host, "host", defaultHost, "TiDB host address")
	flag.IntVar(&port, "port", defaultPort, "TiDB host port")
	flag.StringVar(&dbname, "db", defaultDBName, "Login database")
	flag.StringVar(&user, "user", defaultUser, "Login user")
	flag.BoolVar(&usepw, "pw", false, "Login password")

	flag.BoolVar(&verbose, "V", false, "verbose")
}

func initDB() {
	if usepw {
		_, err := fmt.Scanln(&passwd)
		if err != nil {
			panic(err)
		}
	}

	connStr := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s", user, passwd, host, port, dbname)

	logLvl := logger.Error
	if verbose {
		logLvl = logger.Warn
	}

	var err error
	db, err = gorm.Open(mysql.Open(connStr), &gorm.Config{
		Logger: logger.Default.LogMode(logLvl),
	})
	if err != nil {
		panic(err)
	}
}

type Table struct{ A, B int }

var tables = []struct {
	name string
	gen  func() int
}{
	{"ndv_32", func() int { return rand.Intn(32) }},
	{"ndv_rand", func() int { return int(rand.Int31()) }},
}

func genData() error {
	data := make([]Table, rowNum)
	for _, table := range tables {
		db.Table(table.name).Migrator().DropTable(&Table{})
		db.Table(table.name).Migrator().CreateTable(&Table{})

		for i := 0; i < rowNum; i++ {
			data[i] = Table{A: rand.Intn(1000), B: table.gen()}
		}
		if err := db.Error; err != nil {
			return err
		}
		bar := progressbar.Default(int64(rowNum), table.name)
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
