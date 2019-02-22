package manager

import (
	"../models"
	"fmt"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"log"
)

func PreBooks(owner *walk.MainWindow) (int, error) {
	mw := &MyMainWindow{MainWindow: owner, model: PreBookModel()}
	var dlg *walk.Dialog
	//var db *walk.DataBinder
	//var acceptPB, cancelPB *walk.PushButton
	return Dialog{
		AssignTo: &dlg,
		Title:    "预约管理",
		MinSize:  Size{800, 600},
		Layout:   VBox{},
		Children: []Widget{
			Composite{
				Layout: HBox{MarginsZero: true},
				Children: []Widget{
					HSpacer{},
					PushButton{
						Text:      "添加",
						OnClicked: mw.NewPreBook,
					},
					PushButton{
						Text: "删除",
						OnClicked: func() {
							var items []*Item
							remove := mw.tv.SelectedIndexes()
							for i, x := range mw.model.Items {
								removeOk := false
								for _, j := range remove {
									if i == j {
										removeOk = true
									}
								}
								if !removeOk {
									items = append(items, x)
								}
							}
							mw.model.Items = items
							mw.model.PublishRowsReset()
							mw.tv.SetSelectedIndexes([]int{})
						},
					},
					PushButton{
						Text: "ExecChecked",
						OnClicked: func() {
							for _, x := range mw.model.Items {
								if x.checked {
									fmt.Printf("checked: %v\n", x)
								}
							}
							fmt.Println()
						},
					},
					PushButton{
						Text: "AddPriceChecked",
						OnClicked: func() {
							for i, x := range mw.model.Items {
								if x.checked {
									//x.Price++
									mw.model.PublishRowChanged(i)
								}
							}
						},
					},
				},
			},
			Composite{
				Layout: VBox{},
				ContextMenuItems: []MenuItem{
					Action{
						Text:        "I&nfo",
						OnTriggered: mw.tvItemactivated,
					},
					Action{
						Text: "E&xit",
						OnTriggered: func() {
							mw.Close()
						},
					},
				},
				Children: mw.tableColumn("编号", "名称", "次数", "备注"),
				//[]Widget{
				//	TableView{
				//		AssignTo:         &mw.tv,
				//		CheckBoxes:       true,
				//		ColumnsOrderable: true,
				//		MultiSelection:   true,
				//		Columns: []TableViewColumn{
				//			{Title: "编号"},
				//			{Title: "名称"},
				//			{Title: "次数"},
				//			{Title: "备注"},
				//		},
				//		Model: mw.model,
				//		OnCurrentIndexChanged: func() {
				//			i := mw.tv.CurrentIndex()
				//			if 0 <= i {
				//				fmt.Printf("OnCurrentIndexChanged: %v\n", mw.model.items[i].Name)
				//			}
				//		},
				//		OnItemActivated: mw.tvItemactivated,
				//	},
				//},
			},
		},
	}.Run(owner)
}

func AddPreBook(owner walk.Form, preBook *models.PreBook) (int, error) {
	var dlg *walk.Dialog
	var db *walk.DataBinder
	var acceptPB, cancelPB *walk.PushButton

	return Dialog{
		AssignTo:      &dlg,
		Title:         "预约",
		DefaultButton: &acceptPB,
		CancelButton:  &cancelPB,
		DataBinder: DataBinder{
			AssignTo:       &db,
			AutoSubmit:     true,
			Name:           "preBook",
			DataSource:     preBook,
			ErrorPresenter: ToolTipErrorPresenter{},
		},
		MinSize: Size{600, 300},
		Layout:  VBox{},
		Children: []Widget{
			Composite{
				Layout: Grid{Columns: 2},
				Children: []Widget{
					Label{
						Text: "姓名:",
					},
					LineEdit{
						Text: Bind("Name", SelRequired{}),
					},
					Label{
						Text: "手机:",
					},
					LineEdit{
						Text:      Bind("Mobile"),
						MaxLength: 11,
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
					Label{
						Text: "会员:",
					},
					ComboBox{
						Value:         Bind("MemId"),
						BindingMember: "Id",
						DisplayMember: "Name",
						Model:         MemberList(),
					},
					Label{
						Text: "项目:",
					},
					ComboBox{
						Value:         Bind("ProdId"),
						BindingMember: "Id",
						DisplayMember: "Name",
						Model:         ProdList(),
					},
					Label{
						Text: "员工:",
					},
					ComboBox{
						Value:         Bind("EmpId"),
						BindingMember: "Id",
						DisplayMember: "Name",
						Model:         EmployeeList(),
					},
					Label{
						Text: "预约时间:",
					},
					DateEdit{
						Date:   Bind("ArrivalDate"),
						Format: "yyyy-MM-dd HH:ss",
						//MaxDate :time.Now().AddDate(0,0,14),
						//MinDate:time.Now(),
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
