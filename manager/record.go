package manager

import (
	"../models"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"log"
)

//func CreateAnimals(db *gorm.DB) err {
//	tx := db.Begin()
//	// 注意，一旦你在一个事务中，使用tx作为数据库句柄
//
//	if err := tx.Create(&Animal{Name: "Giraffe"}).Error; err != nil {
//		tx.Rollback()
//		return err
//	}
//
//	if err := tx.Create(&Animal{Name: "Lion"}).Error; err != nil {
//		tx.Rollback()
//		return err
//	}
//
//	tx.Commit()
//	return nil
//}

func RecordModel() *ItemModel {
	records := models.Record{}.Search()
	m := &ItemModel{table: "record", Items: make([]*Item, len(records))}

	for i, j := range records {
		//fmt.Println(j.MemId)
		m.Items[i] = &Item{
			Index:   i,
			Remarks: j.Remarks,
		}
		go func() {
			mem := models.Member{}.Get(j.MemId)
			//fmt.Println(mem)
			m.Items[i].Name = mem.Name
		}()

		go func() {
			prod := models.Get(models.Prod{}, j.ProdId).Prod
			if "" != prod.Name {
				m.Items[i].Prod = prod.Name
				m.Items[i].Price = float32(prod.Price)
			}

		}()
		go func() {
			emp := models.Get(models.Employee{}, j.EmpId).Employee
			m.Items[i].EName = emp.Name
		}()
		go func() {
			card := models.Get(models.Card{}, j.CardId).Card
			if "" != card.Name {
				m.Items[i].Prod = card.Name
				m.Items[i].Price = float32(card.Price)
			}

		}()
		go func() {
			combo := models.Get(models.Combo{}, j.ComboId).Combo
			if "" != combo.Name {
				m.Items[i].Prod = combo.Name
				m.Items[i].Price = float32(combo.Price)
			}

		}()
		//fmt.Println(mem)
		//fmt.Println(prod)
		//fmt.Println(emp)
		//fmt.Println(card)
		//fmt.Println(emp.Employee.Name)
		//fmt.Println(prod.Prod.Name)
		//fmt.Println(card.Card.Name)
		//fmt.Println(models.Get(models.Card{}, j.CardId).Card.Price)

	}
	return m
}

func Records(owner *MyMainWindow) (int, error) {
	//mw := &MyMainWindow{MainWindow: owner, model: RecordModel()}
	owner.model = RecordModel()
	var dlg *walk.Dialog
	//var db *walk.DataBinder
	//var acceptPB, cancelPB *walk.PushButton
	return Dialog{
		AssignTo: &dlg,
		Title:    "消费管理",
		MinSize:  Size{800, 600},
		Layout:   VBox{},
		Children: owner.Manager("Record", "序号", "会员", "员工", "项目", "价格", "备注"),
	}.Run(owner)
}
func AddRecord(owner walk.Form, member *models.Record) (int, error) {
	var dlg *walk.Dialog
	var db *walk.DataBinder
	var acceptPB, cancelPB *walk.PushButton
	return Dialog{
		AssignTo:      &dlg,
		Title:         "添加消费记录",
		DefaultButton: &acceptPB,
		CancelButton:  &cancelPB,
		DataBinder: DataBinder{
			AssignTo:       &db,
			AutoSubmit:     true,
			Name:           "Member",
			DataSource:     member,
			ErrorPresenter: ToolTipErrorPresenter{},
		},
		MinSize: Size{400, 300},
		Layout:  VBox{},
		Children: []Widget{
			Composite{
				Layout: Grid{Columns: 2},
				Children: []Widget{
					Label{
						Text:          "会员:",
						TextAlignment: AlignFar,
					},
					ComboBox{
						Value:         Bind("MemId"),
						MinSize:       Size{50, 20},
						BindingMember: "Id",
						DisplayMember: "Name",
						Model:         MemberList(),
					},
					Label{
						Text:          "员工:",
						TextAlignment: AlignFar,
					},
					ComboBox{
						Value:         Bind("EmpId"),
						MinSize:       Size{50, 20},
						BindingMember: "Id",
						DisplayMember: "Name",
						Model:         EmployeeList(),
					},
					Label{
						Text:          "项目:",
						TextAlignment: AlignFar,
					},
					ComboBox{
						Value:         Bind("ProdId"),
						MinSize:       Size{50, 20},
						BindingMember: "Id",
						DisplayMember: "Name",
						Model:         ProdList(),
					},
					Label{
						Text:          "疗程:",
						TextAlignment: AlignFar,
					},
					ComboBox{
						Value:         Bind("ComboId"),
						MinSize:       Size{50, 20},
						BindingMember: "Id",
						DisplayMember: "Name",
						Model:         ComboList(),
					},
					Label{
						Text:          "售卡:",
						TextAlignment: AlignFar,
					},
					ComboBox{
						Value:         Bind("CardId"),
						MinSize:       Size{50, 20},
						BindingMember: "Id",
						DisplayMember: "Name",
						Model:         CardList(),
					},
					Label{
						ColumnSpan: 2,
						Text:       "备注:",
					},
					TextEdit{
						ColumnSpan: 2,
						MinSize:    Size{100, 50},
						Text:       Bind("Remarks"),
					},
				},
			},
			Composite{
				Layout: HBox{},
				Children: []Widget{
					HSpacer{},
					PushButton{
						AssignTo: &acceptPB,
						Text:     "OK",
						OnClicked: func() {
							if err := db.Submit(); err != nil {
								log.Print(err)
								walk.MsgBox(owner, "错误提示", err.Error(), walk.MsgBoxIconError)
								return
							}
							dlg.Accept()
						},
					},
					PushButton{
						AssignTo:  &cancelPB,
						Text:      "Cancel",
						OnClicked: func() { dlg.Cancel() },
					},
				},
			},
		},
	}.Run(owner)
}

func AddCard(owner walk.Form, prepaid *models.Card) (int, error) {
	var dlg *walk.Dialog
	var db *walk.DataBinder
	var acceptPB, cancelPB *walk.PushButton
	return Dialog{
		AssignTo:      &dlg,
		Title:         "添加折扣卡",
		DefaultButton: &acceptPB,
		CancelButton:  &cancelPB,
		DataBinder: DataBinder{
			AssignTo:       &db,
			AutoSubmit:     true,
			Name:           "Prepaid",
			DataSource:     prepaid,
			ErrorPresenter: ToolTipErrorPresenter{},
		},
		MinSize: Size{600, 300},
		Layout:  VBox{},
		Children: []Widget{
			Composite{
				Layout: Grid{Columns: 2},
				Children: []Widget{
					Label{
						Text: "名称:",
					},
					LineEdit{
						Text:    Bind("Name"),
						MinSize: Size{50, 20},
					},
					Label{
						Text: "价格:",
					},
					NumberEdit{
						Value:   Bind("Price"),
						MinSize: Size{50, 20},
					},
					Label{
						Text: "折扣:",
					},
					//ComboBox{
					//	Value:         Bind("Discount", SelRequired{}),
					//	BindingMember: "Value",
					//	DisplayMember: "Name",
					//	Model:         Discounts(),
					//},
					NumberEdit{
						Value:    Bind("Discount", Range{0.30, 0.95}),
						MinSize:  Size{50, 20},
						Decimals: 2,
					},
					Label{
						ColumnSpan: 2,
						Text:       "备注:",
					},
					TextEdit{
						ColumnSpan: 2,
						MinSize:    Size{100, 50},
						Text:       Bind("Remarks"),
					},
				},
			},
			Composite{
				Layout: HBox{},
				Children: []Widget{
					HSpacer{},
					PushButton{
						AssignTo: &acceptPB,
						Text:     "OK",
						OnClicked: func() {
							if err := db.Submit(); err != nil {
								log.Print(err)
								walk.MsgBox(owner, "错误提示", err.Error(), walk.MsgBoxIconError)
								return
							}
							dlg.Accept()
						},
					},
					PushButton{
						AssignTo:  &cancelPB,
						Text:      "Cancel",
						OnClicked: func() { dlg.Cancel() },
					},
				},
			},
		},
	}.Run(owner)
}
