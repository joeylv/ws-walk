package manager

import (
	"../models"
	"fmt"
	"github.com/lxn/walk"
	"log"
	"sort"
	"strconv"
	"time"
)

var model = &models.Model{}

type PreSet struct {
	*walk.TableView
	*ItemModel
	Title string
	Count int
}

type MyMainWindow struct {
	*walk.MainWindow
	TodayPre *PreSet
	TomPre   *PreSet
	model    *ItemModel
	tv       *walk.TableView
}

type Item struct {
	Index       int
	Name        string
	Mobile      string
	Price       float32
	Count       int
	Remarks     string
	ArrivalDate time.Time
	checked     bool
}

type ItemModel struct {
	walk.TableModelBase
	walk.SorterBase
	sortColumn int
	sortOrder  walk.SortOrder
	table      string
	Items      []*Item
}

func (m *ItemModel) RowCount() int {
	return len(m.Items)
}

func (m *ItemModel) prebookValue(row, col int) interface{} {
	item := m.Items[row]

	switch col {
	case 0:
		return item.Name
	case 1:
		return item.Mobile
	case 2:
		return strconv.Itoa(item.ArrivalDate.Day()) + "日" + strconv.Itoa(item.ArrivalDate.Hour()) + "点" + strconv.Itoa(item.ArrivalDate.Minute()) + "分"
	case 3:
		return item.Remarks
	}
	panic("unexpected col")
}
func (m *ItemModel) Default(row, col int) interface{} {
	item := m.Items[row]

	switch col {
	case 0:
		return item.Index
	case 1:
		return item.Name
	case 2:
		switch m.table {
		case "prod":
			return item.Price
		case "prebook":
			return strconv.Itoa(item.ArrivalDate.Day()) + "日" + strconv.Itoa(item.ArrivalDate.Hour()) + "点" + strconv.Itoa(item.ArrivalDate.Minute()) + "分"
		default:
			return item.Mobile
		}
	case 3:
		return item.Remarks
	}
	panic("unexpected col")
}

func (m *ItemModel) Value(row, col int) interface{} {
	switch m.table {
	case "prebook":
		return m.prebookValue(row, col)
	default:
		return m.Default(row, col)
	}

	//switch col {
	//case 0:
	//	return item.Index
	//case 1:
	//	return item.Name
	//case 2:
	//	switch m.table {
	//	case "prod":
	//		return item.Price
	//	case "prebook":
	//		return strconv.Itoa(item.ArrivalDate.Day()) + "日" + strconv.Itoa(item.ArrivalDate.Hour()) + "点" + strconv.Itoa(item.ArrivalDate.Minute()) + "分"
	//	default:
	//		return item.Mobile
	//	}
	//case 3:
	//	return item.Remarks
	//}
	panic("unexpected col")
}

func (m *ItemModel) Checked(row int) bool {
	return m.Items[row].checked
}

func (m *ItemModel) SetChecked(row int, checked bool) error {
	m.Items[row].checked = checked
	return nil
}

func (m *ItemModel) Sort(col int, order walk.SortOrder) error {
	m.sortColumn, m.sortOrder = col, order
	sort.Stable(m)
	return m.SorterBase.Sort(col, order)
}

func (m *ItemModel) Len() int {
	return len(m.Items)
}

func (m *ItemModel) Less(i, j int) bool {
	a, b := m.Items[i], m.Items[j]

	c := func(ls bool) bool {
		if m.sortOrder == walk.SortAscending {
			return ls
		}

		return !ls
	}

	switch m.sortColumn {
	case 0:
		return c(a.Index < b.Index)
	case 1:
		return c(a.Name < b.Name)
	case 2:
		return c(a.Mobile < b.Mobile)
	case 3:
		return c(a.Remarks < b.Remarks)
	case 4:
		return c(a.Price < b.Price)
	}

	panic("unreachable")
}

func (m *ItemModel) Swap(i, j int) {
	m.Items[i], m.Items[j] = m.Items[j], m.Items[i]
}

func (s *MyMainWindow) GetModel(d string, time ...*time.Time) *ItemModel {
	fmt.Println("GoGoGo!!!")
	//var memList *models.Model
	var m *ItemModel
	switch d {
	case "PreBook":
		go func() {
			memList := models.Search(models.PreBook{}, 0)
			m := &ItemModel{table: "prebook", Items: make([]*Item, len(memList.PreBooks))}
			for i, j := range memList.PreBooks {
				model := models.Search(models.Member{}, j.MemId)
				//fmt.Println(item.Members)
				//mem := models.Member{}.Search(j.MemId)
				if len(model.Members) > 0 {
					m.Items[i] = &Item{
						Index:       i,
						Name:        j.Name,
						Mobile:      model.Members[0].Mobile,
						Remarks:     j.Remarks,
						ArrivalDate: j.ArrivalDate,
					}
				}
			}
		}()
	case "Member":
		memList := models.Search(models.Member{}, 0)
		m := &ItemModel{table: "member", Items: make([]*Item, len(memList.Members))}
		//m.items = make([]*Item, len(memList))
		for i, j := range memList.Members {
			m.Items[i] = &Item{
				Index:   i,
				Name:    j.Name,
				Mobile:  j.Mobile,
				Remarks: j.Remarks,
			}
		}
	case "Combo":
		memList := models.Search(models.Combo{}, 0)
		//models.Combo{}.Search()
		m = &ItemModel{table: "combo", Items: make([]*Item, len(memList.Combos))}
		//m.items = make([]*Item, len(memList))
		//fmt.Println(memList)
		for i, j := range memList.Combos {
			//fmt.Println(i)
			//fmt.Println(j.Name)
			//m.items
			m.Items[i] = &Item{
				Index: i,
				Name:  j.Name,
				Count: j.Count,
				Price: j.Price,
				//Mobile:  j.Mobile,
				//Remarks: j.Remarks,
			}
			//fmt.Println(reflect.TypeOf(j).Elem())
		}
	case "Employee":
		memList := models.Search(models.Employee{}, 0)
		//models.Combo{}.Search()
		m = &ItemModel{table: "emp", Items: make([]*Item, len(memList.Employees))}
		//m.items = make([]*Item, len(memList))
		//fmt.Println(memList)
		for i, j := range memList.Employees {
			//fmt.Println(i)
			//fmt.Println(j.Name)
			//m.items
			m.Items[i] = &Item{
				Index:   i,
				Name:    j.Name,
				Mobile:  j.Mobile,
				Remarks: j.Remarks,
			}
			//fmt.Println(reflect.TypeOf(j).Elem())
		}
	}

	return m

	//m.items = make([]*Item, len(memList))

}

func PreBookModel(time ...*time.Time) *ItemModel {
	memList := models.PreBook{}.Search(time...)
	m := &ItemModel{table: "prebook", Items: make([]*Item, len(memList))}
	//m.items = make([]*Item, len(memList))
	for i, j := range memList {
		model := models.Search(models.Member{}, j.MemId)
		//fmt.Println(item.Members)
		//mem := models.Member{}.Search(j.MemId)
		if len(model.Members) > 0 {
			m.Items[i] = &Item{
				Index:       i,
				Name:        j.Name,
				Mobile:      model.Members[0].Mobile,
				Remarks:     j.Remarks,
				ArrivalDate: j.ArrivalDate,
			}
		}
	}
	return m
}

//
func ComboModel() *ItemModel {
	memList := models.Combo{}.Search()
	m := &ItemModel{table: "combo", Items: make([]*Item, len(memList))}
	//m.items = make([]*Item, len(memList))
	//fmt.Println(memList)
	for i, j := range memList {
		fmt.Println(i)
		fmt.Println(j.Name)
		//m.items
		m.Items[i] = &Item{
			Index: i,
			Name:  j.Name,
			Count: j.Count,
			Price: j.Price,
		}
	}

	return m
}

//
func MemberModel() *ItemModel {
	memList := models.Member{}.Search(0)
	m := &ItemModel{table: "member", Items: make([]*Item, len(memList))}
	//m.items = make([]*Item, len(memList))
	for i, j := range memList {
		m.Items[i] = &Item{
			Index:   i,
			Name:    j.Name,
			Mobile:  j.Mobile,
			Remarks: j.Remarks,
		}
	}
	return m
}

func (mw *MyMainWindow) OpenMembers() {
	var outTE *walk.TextEdit
	if cmd, err := Member(mw.MainWindow); err != nil {
		log.Print(err)
	} else if cmd == walk.DlgCmdOK {
		outTE.SetText(fmt.Sprintf("%+v", "CMD"))
	}
}
func (mw *MyMainWindow) OpenEmployees() {
	var outTE *walk.TextEdit
	if cmd, err := Employees(mw.MainWindow); err != nil {
		log.Print(err)
	} else if cmd == walk.DlgCmdOK {
		outTE.SetText(fmt.Sprintf("%+v", "CMD"))
	}
}
func (mw *MyMainWindow) OpenPreBooks() {
	var outTE *walk.TextEdit
	if cmd, err := PreBooks(mw.MainWindow); err != nil {
		log.Print(err)
	} else if cmd == walk.DlgCmdOK {
		outTE.SetText(fmt.Sprintf("%+v", "CMD"))
	}
}

func (mw *MyMainWindow) OpenProducts() {
	var outTE *walk.TextEdit
	if cmd, err := Products(mw.MainWindow); err != nil {
		log.Print(err)
	} else if cmd == walk.DlgCmdOK {
		outTE.SetText(fmt.Sprintf("%+v", "CMD"))
	}
}

func (mw *MyMainWindow) OpenCombos() {
	var outTE *walk.TextEdit
	if cmd, err := Combos(mw.MainWindow); err != nil {
		log.Print(err)
	} else if cmd == walk.DlgCmdOK {
		fmt.Println("OK")
		outTE.SetText(fmt.Sprintf("%+v", "CMD"))
	}
}

func (mw *MyMainWindow) OpenRecords() {
	var outTE *walk.TextEdit
	if cmd, err := Records(mw.MainWindow); err != nil {
		log.Print(err)
	} else if cmd == walk.DlgCmdOK {
		fmt.Println("OK")
		outTE.SetText(fmt.Sprintf("%+v", "CMD"))
	}
}

func (mw *MyMainWindow) NewMember() {
	member := &models.Member{}
	if cmd, err := AddMember(mw, member); err != nil {
		log.Print(err)
	} else if cmd == walk.DlgCmdOK {
		//fmt.Println("OK", member)
		//models.Save(member)
		model.Member = member
		model.Save("member")
		//member.Save()
		//outTE.SetText(fmt.Sprintf("%+v", member))
	}
}

func (mw *MyMainWindow) NewEmployee() {
	//var outTE *walk.TextEdit
	emp := &models.Employee{}
	if cmd, err := AddEmployee(mw, emp); err != nil {
		log.Print(err)
	} else if cmd == walk.DlgCmdOK {
		//fmt.Println("OK", emp)
		model.Employee = emp
		model.Save("Employee")
		//emp.Save()
		//outTE.SetText(fmt.Sprintf("%+v", member))
	}
}

func (mw *MyMainWindow) NewPreBook() {
	//var outTE *walk.TextEdit
	preBook := &models.PreBook{ArrivalDate: time.Now()}
	//loc, _ := time.LoadLocation("Local")
	if cmd, err := AddPreBook(mw, preBook); err != nil {
		log.Print(err)
	} else if cmd == walk.DlgCmdOK {
		//fmt.Println("OK")
		model.PreBook = preBook
		model.Save("PreBook")

		//new := models.Search(models.Member{}, preBook.MemId)
		//fmt.Println(item.Members)

		//mem := models.Member{}.Search(preBook.MemId)
		//if len(model.Members) > 0 {
		//	m.Items[i] = &Item{
		//		Index:       i,
		//		Name:        j.Name,
		//		Mobile:      mem.Members[0].Mobile,
		//		Remarks:     j.Remarks,
		//		ArrivalDate: j.ArrivalDate,
		//	}
		//}
		mw.TodayPre.Items = append(mw.TodayPre.Items, &Item{
			Index:       mw.TodayPre.Len(),
			Name:        "test",
			Mobile:      model.Members[0].Mobile,
			Remarks:     "test",
			ArrivalDate: time.Now(),
		})
		mw.model.PublishRowsReset()
		//preBook.Save()
		//outTE.SetText(fmt.Sprintf("%+v", member))
	}
}
func (mw *MyMainWindow) NewProd() {
	//var outTE *walk.TextEdit
	prod := &models.Prod{}
	//loc, _ := time.LoadLocation("Local")
	if cmd, err := AddProduct(mw, prod); err != nil {
		log.Print(err)
	} else if cmd == walk.DlgCmdOK {
		//fmt.Println("OK")
		model.Prod = prod
		model.Save("Prod")
		//prod.Save()
		//outTE.SetText(fmt.Sprintf("%+v", member))
	}
}

func (mw *MyMainWindow) NewCombo() {
	combo := &models.Combo{}
	if cmd, err := AddCombo(mw, combo); err != nil {
		log.Print(err)
	} else if cmd == walk.DlgCmdOK {
		fmt.Println("DlgCmdOK")
		combo.Save()
		mw.model.Items = append(mw.model.Items, &Item{
			Index: mw.model.Len(),
			Name:  combo.Name,
			Count: combo.Count,
			Price: combo.Price,
		})
		mw.model.PublishRowsReset()
		//outTE.SetText(fmt.Sprintf("%+v", member))
	}

}

func (mw *MyMainWindow) NewRecord() {
	record := new(models.Record)
	if cmd, err := AddRecord(mw, record); err != nil {
		log.Print(err)
	} else if cmd == walk.DlgCmdOK {
		fmt.Println("DlgCmdOK")
		record.Save()
		mw.model.Items = append(mw.model.Items, &Item{
			Index: mw.model.Len(),
			//Name:    record.Name,
			//Price:   record.Price,
			Remarks: record.Remarks,
		})
		mw.model.PublishRowsReset()
	}
}

func (mw *MyMainWindow) openDialog() {
	//var outTE *walk.TextEdit
	animal := new(models.Animal)
	if cmd, err := RunAnimalDialog(mw, animal); err != nil {
		log.Print(err)
	} else if cmd == walk.DlgCmdOK {
		fmt.Println("OK")
		//outTE.SetText(fmt.Sprintf("%+v", animal))
	}
}

func (mw *MyMainWindow) InitDataBase() {
	models.Migrate()
	//dbcon.Create()
}
func (mw *MyMainWindow) OpenAction_Triggered() {
	walk.MsgBox(mw, "Open", "Pretend to open a file...", walk.MsgBoxIconInformation)
}

func (mw *MyMainWindow) NewAction_Triggered() {
	walk.MsgBox(mw, "New", "Newing something up... or not.", walk.MsgBoxIconInformation)
}

func (mw *MyMainWindow) ChangeViewAction_Triggered() {
	walk.MsgBox(mw, "Change View", "By now you may have guessed it. Nothing changed.", walk.MsgBoxIconInformation)
}

func (mw *MyMainWindow) ShowAboutBoxAction_Triggered() {
	walk.MsgBox(mw, "About", "Walk Actions Example", walk.MsgBoxIconInformation)
}

func (mw *MyMainWindow) SpecialAction_Triggered() {
	walk.MsgBox(mw, "Special", "Nothing to see here.", walk.MsgBoxIconInformation)
}
