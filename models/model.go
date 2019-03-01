package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"time"
)

func Search(p interface{}) *Model {
	db := GetConn()
	defer db.Close()
	model := &Model{}
	switch a := p.(type) {
	case PreBook:
		var list []PreBook
		db.Find(&list)
		//fmt.Println("Member search")
		model.PreBooks = list
		//fmt.Println("PreBook", a.Search())
	case Employee:
		var list []Employee
		db.Find(&list)
		//fmt.Println("Member search")
		model.Employees = list
		//fmt.Println("Employee", a.Search())
	case Member:
		var list []Member
		db.Find(&list)
		//fmt.Println("Member search")
		model.Members = list
		//fmt.Println("Employee", a.Search())
	case Combo:
		var list []Combo
		db.Find(&list)
		//fmt.Println("Member search")
		model.Combos = list
	case Card:
		var list []Card
		db.Find(&list)
		//fmt.Println("Member search")
		model.Cards = list
	default:
		fmt.Println("unknown type", a)
	}
	return model
}

func Get(p interface{}, id uint) *Model {
	db := GetConn()
	defer db.Close()
	model := &Model{}
	switch a := p.(type) {
	case PreBook:
		var list PreBook
		if id != 0 {
			db.First(&list, id)
		}
		//fmt.Println("Member search")
		model.PreBook = list
		//fmt.Println("PreBook", a.Search())
	case Employee:
		var list Employee
		if id != 0 {
			db.First(&list, id)
		}
		//fmt.Println("Member search")
		model.Employee = list
		//fmt.Println("Employee", a.Search())
	case Member:
		var list Member
		if id != 0 {
			db.First(&list, id)
		}
		//fmt.Println("Member search")
		model.Member = list
		//fmt.Println("Employee", a.Search())
	case Combo:
		var list Combo
		if id != 0 {
			db.First(&list, id)
		}
		//fmt.Println("Member search")
		model.Combo = list
	case Card:
		var list Card
		if id != 0 {
			db.First(&list, id)
		}
		//fmt.Println("Member search")
		model.Card = list
	default:
		fmt.Println("unknown type", a)
	}
	return model
}

type Model struct {
	Member
	Prod
	PreBook
	Combo
	Employee
	Record
	Card
	Consume

	Members   []Member
	Prods     []Prod
	PreBooks  []PreBook
	Combos    []Combo
	Employees []Employee
	Records   []Record
	Cards     []Card
	Consumes  []Consume
}

type Member struct {
	gorm.Model
	Code     string
	Name     string
	Mobile   string `gorm:"not null;unique"`
	Remarks  string
	Balance  int
	Discount float32
	//Prepaid int
	//Discount float32
	PreBooks []PreBook
	Combos   []Combo
	Records  []Record
}

type Card struct {
	gorm.Model
	Name     string
	Price    int
	Discount float32
	Remarks  string
}

type Employee struct {
	gorm.Model
	Code    string
	Name    string
	Mobile  string
	Remarks string
}

type Prod struct {
	gorm.Model
	Code    string
	Name    string
	Price   float32
	Remarks string
}

type Consume struct {
	gorm.Model
	MemId   uint
	ComboId uint
	Count   int
}

type Combo struct {
	gorm.Model
	//Prod   []Prod `gorm:"ForeignKey:ProdId"`
	ProdId uint
	Code   string
	Name   string
	Price  float32
	Count  int
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

type Record struct {
	gorm.Model
	//Name  string
	//Price float32
	//Prod    []Prod `gorm:"ForeignKey:rePId"`
	ProdId uint
	//Member  Member `gorm:"ForeignKey:recordMId"`
	MemId   uint
	EmpId   uint
	CardId  uint
	ComboId uint
	Remarks string
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
	db.AutoMigrate(&Member{}, &Employee{}, &Combo{}, &Card{})
	//db.AutoMigrate(&Combo{}, &PreBook{})
	db.AutoMigrate(&Prod{}, &Record{}, &PreBook{}, &Consume{})
	//db.Model(&Member{}).Related(&PreBook{})
}

func (m Member) Search() []Member {
	db := GetConn()
	defer db.Close()
	var list []Member
	db.Find(&list)
	return list

}

func (m Member) Get(id uint) Member {
	db := GetConn()
	defer db.Close()
	var list Member
	if id != 0 {
		db.First(&list, id)
	}
	//fmt.Println("Member search")
	return list

}
func (m Member) Save() {
	db := GetConn()
	defer db.Close()
	//fmt.Println(m)
	db.Create(&m)
}

func (m Card) Save() {
	db := GetConn()
	defer db.Close()
	//fmt.Println(m)
	db.Create(&m)
}

func (m Employee) Search(id uint) []Employee {
	db := GetConn()
	defer db.Close()
	var list []Employee
	if id == 0 {
		db.Find(&list)
	} else {
		db.First(&list, id)
	}
	//fmt.Println("Employee search")
	return list
}
func (e Employee) Save() {
	db := GetConn()
	defer db.Close()
	db.Create(&e)
}

func (p Prod) Search() []Prod {
	db := GetConn()
	defer db.Close()
	var list []Prod
	db.Find(&list) // find product with id 1
	//fmt.Println("Prod search")
	return list
}
func (p Prod) Save() {
	db := GetConn()
	defer db.Close()
	db.Create(&p)
}

func (c Consume) Search(Id ...uint) []Consume {
	db := GetConn()
	defer db.Close()
	var list []Consume
	if len(Id) > 0 {
		db.Where(&Consume{MemId: Id[0], ComboId: Id[1]}).Where("count > ? ", 0).First(&list)
	} else {
		db.Find(&list) // find product with id 1
	}
	return list
}
func (c Consume) Save(record *Record) error {
	db := GetConn()
	defer db.Close()
	tx := db.Begin()
	// 注意，一旦你在一个事务中，使用tx作为数据库句柄
	//consume[0].Save()
	if err := tx.Create(c).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(record).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (c Combo) Search() []Combo {
	db := GetConn()
	defer db.Close()
	var list []Combo
	db.Find(&list) // find product with id 1
	//fmt.Println("Combo search")
	return list
}
func (c Combo) Save() {
	db := GetConn()
	defer db.Close()
	db.Create(&c)
}

func (p PreBook) Search(time ...*time.Time) []PreBook {
	db := GetConn()
	defer db.Close()
	var list []PreBook
	if len(time) == 1 {
		//fmt.Println(time[0])
		db.Where("arrival_date < ?", time).Find(&list)
	} else if len(time) == 2 {
		//fmt.Println(time[0])
		//fmt.Println(time[1])
		db.Where("arrival_date BETWEEN ? AND ?", time[0], time[1]).Find(&list)
	} else {
		db.Find(&list) // find All
	}

	//fmt.Println("PreBook search")
	return list
}

func (p Record) Search() []Record {
	db := GetConn()
	defer db.Close()
	var list []Record
	db.Find(&list) // find product with id 1
	//fmt.Println("PreBook search")
	return list
}
func (p Record) Save() {
	db := GetConn()
	defer db.Close()
	db.Create(&p)
}

func (p PreBook) Save() {
	db := GetConn()
	defer db.Close()
	db.Create(&p)
}
