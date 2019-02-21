package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"time"
)

func (m Model) Save(d string) {
	db := GetConn()
	defer db.Close()
	switch d {
	case "PreBook":
		db.Create(&m.PreBook)
		//fmt.Println("PreBook", a.Search())
	case "Employee":
		db.Create(&m.Employee)
		//fmt.Println("Employee", a.Search())
	case "Member":
		db.Create(&m.Member)
	case "Prod":
		db.Create(&m.Prod)
		//fmt.Println("Employee", a.Search())
	case "Combo":
		db.Create(&m.Combo)
	default:
		fmt.Println("unknown type", m)

	}

}

func Search(p interface{}, id uint) *Model {
	db := GetConn()
	defer db.Close()
	model := &Model{}
	switch a := p.(type) {
	case PreBook:
		//fmt.Println("PreBook", a.Search())
	case Employee:
		//fmt.Println("Employee", a.Search())
	case Member:
		var product []Member
		if id == 0 {
			db.Find(&product)
		} else {
			db.First(&product, id)
		}
		//fmt.Println("Member search")
		model.Members = product
		//fmt.Println("Employee", a.Search())
	default:
		fmt.Println("unknown type", a)
	}
	return model
}

type Model struct {
	*Member
	*Prod
	*PreBook
	*Combo
	*Employee

	Members   []Member
	Prods     []Prod
	PreBooks  []PreBook
	Combos    []Combo
	Employees []Employee
}

type Member struct {
	gorm.Model
	Code     string
	Name     string
	Mobile   string `gorm:"not null;unique"`
	Remarks  string
	Discount float32
	PreBooks []PreBook
}

func (m Member) Search(id uint) []Member {
	db := GetConn()
	defer db.Close()
	var product []Member
	if id == 0 {
		db.Find(&product)
	} else {
		db.First(&product, id)
	}
	//fmt.Println("Member search")
	return product

}
func (m Member) Save() {
	db := GetConn()
	defer db.Close()
	//fmt.Println(m)
	db.Create(&m)
}

type Employee struct {
	gorm.Model
	Code    string
	Name    string
	Mobile  string
	Remarks string
}

func (m Employee) Search(id uint) []Employee {
	db := GetConn()
	defer db.Close()
	var product []Employee
	if id == 0 {
		db.Find(&product)
	} else {
		db.First(&product, id)
	}
	//fmt.Println("Employee search")
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
	Price   float32
	Remarks string
}

func (p Prod) Search() []Prod {
	db := GetConn()
	defer db.Close()
	var product []Prod
	db.Find(&product) // find product with id 1
	//fmt.Println("Prod search")
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
	Price  float32
	Count  int
}

func (c Combo) Search() []Combo {
	db := GetConn()
	defer db.Close()
	var product []Combo
	db.Find(&product) // find product with id 1
	//fmt.Println("Combo search")
	return product
}
func (c Combo) Save() {
	db := GetConn()
	defer db.Close()
	db.Create(&c)
}

type PreBook struct {
	gorm.Model
	Name   string
	Mobile string
	//Member      Member `gorm:"ForeignKey:MemId"`
	MemId uint
	//Employee    Employee `gorm:"ForeignKey:EmpId"`
	EmpId uint
	//Prod        []Prod `gorm:"ForeignKey:prePId"`
	ProdId      uint
	ArrivalDate time.Time
	Remarks     string
}

func (p PreBook) Search(time ...*time.Time) []PreBook {
	db := GetConn()
	defer db.Close()
	var product []PreBook
	if len(time) == 1 {
		//fmt.Println(time[0])
		db.Where("arrival_date < ?", time).Find(&product)
	} else if len(time) == 2 {
		//fmt.Println(time[0])
		//fmt.Println(time[1])
		db.Where("arrival_date BETWEEN ? AND ?", time[0], time[1]).Find(&product)
	} else {
		db.Find(&product) // find All
	}

	//fmt.Println("PreBook search")
	return product
}
func (p PreBook) Save() {
	db := GetConn()
	defer db.Close()
	db.Create(&p)
}

type Record struct {
	gorm.Model
	Name    string
	Price   float32
	Prod    []Prod `gorm:"ForeignKey:rePId"`
	ProdId  uint
	Member  Member `gorm:"ForeignKey:recordMId"`
	MemId   uint
	Remarks string
}

func (p Record) Search() []Record {
	db := GetConn()
	defer db.Close()
	var product []Record
	db.Find(&product) // find product with id 1
	//fmt.Println("PreBook search")
	return product
}
func (p Record) Save() {
	db := GetConn()
	defer db.Close()
	db.Create(&p)
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
	db.Model(&Member{}).Related(&PreBook{})
}
