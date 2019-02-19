package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"time"
)

type Item struct {
	Member
	Prod
	PreBook
	Combo
	Employee
}

type Member struct {
	gorm.Model
	Code     string
	Name     string
	Mobile   string `gorm:"not null;unique"`
	Remarks  string
	Discount float32
}

func (m Member) Search() []Member {
	db := GetConn()
	defer db.Close()
	var product []Member
	db.First(&product, 1) // find product with id 1
	fmt.Println("Member search")
	return product

}
func (m Member) Save() {
	db := GetConn()
	defer db.Close()
	fmt.Println(m)
	db.Create(&m)
}

type Employee struct {
	gorm.Model
	Code    string
	Name    string
	Mobile  string
	Remarks string
}

func (m Employee) Search() []Employee {
	db := GetConn()
	defer db.Close()
	var product []Employee
	db.Find(&product) // find product with id 1
	fmt.Println("Employee search")
	return product
}
func (e Employee) Save() {
	db := GetConn()
	defer db.Close()
	db.Create(&e)
}

type Prod struct {
	gorm.Model
	Code    string
	Name    string
	Price   int
	Remarks string
}

func (p Prod) Search() []Prod {
	db := GetConn()
	defer db.Close()
	var product []Prod
	db.Find(&product) // find product with id 1
	fmt.Println("Prod search")
	return product
}
func (p Prod) Save() {
	db := GetConn()
	defer db.Close()
	db.Create(&p)
}

type Combo struct {
	Prod   []Prod `gorm:"ForeignKey:ProdId"`
	ProdId uint
	Code   string
	Name   string
	Price  int
	Count  int
}

func (c Combo) Search() []Combo {
	db := GetConn()
	defer db.Close()
	var product []Combo
	db.Find(&product) // find product with id 1
	fmt.Println("Combo search")
	return product
}
func (c Combo) Save() {
	db := GetConn()
	defer db.Close()
	db.Create(&c)
}

type PreBook struct {
	gorm.Model
	Name        string
	Mobile      string
	Member      Member `gorm:"ForeignKey:MemId"`
	MemId       uint
	Employee    Employee `gorm:"ForeignKey:EmpId"`
	EmpId       uint
	Prod        []Prod `gorm:"ForeignKey:prePId"`
	ProdId      uint
	ArrivalDate time.Time
	Remarks     string
}

func (p PreBook) Search() []PreBook {
	db := GetConn()
	defer db.Close()
	var product []PreBook
	db.Find(&product) // find product with id 1
	fmt.Println("PreBook search")
	return product
}
func (p PreBook) Save() {
	db := GetConn()
	defer db.Close()
	db.Create(&p)
}

type Record struct {
	gorm.Model
	Name   string
	Price  int
	Prod   []Prod `gorm:"ForeignKey:rePId"`
	ProdId uint
	Member Member `gorm:"ForeignKey:recordMId"`
	MemId  uint
}
type Animal struct {
	Name          string
	ArrivalDate   time.Time
	SpeciesId     int
	Speed         int
	Sex           Sex
	Weight        float64
	PreferredFood string
	Domesticated  bool
	Remarks       string
	Patience      time.Duration
}

func (a *Animal) PatienceField() *DurationField {
	return &DurationField{&a.Patience}
}

type DurationField struct {
	p *time.Duration
}

func (*DurationField) CanSet() bool       { return true }
func (f *DurationField) Get() interface{} { return f.p.String() }
func (f *DurationField) Set(v interface{}) error {
	x, err := time.ParseDuration(v.(string))
	if err == nil {
		*f.p = x
	}
	return err
}
func (f *DurationField) Zero() interface{} { return "" }

type Sex byte

func GetConn() *gorm.DB {
	db, err := gorm.Open("sqlite3", "gorm.db")
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func Migrate() {
	db := GetConn()
	defer db.Close()
	fmt.Println("AutoMigrate")
	// Migrate the schema
	db.AutoMigrate(&Member{}, &Employee{}, &Combo{})
	//db.AutoMigrate(&Combo{}, &PreBook{})
	db.AutoMigrate(&Prod{}, &Record{}, &PreBook{})
}
