package dialog

import (
	"../models"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"log"
)

func AddPreBook(owner walk.Form, preBook *models.PreBook) (int, error) {
	var dlg *walk.Dialog
	var db *walk.DataBinder
	var acceptPB, cancelPB *walk.PushButton

	return Dialog{
		AssignTo:      &dlg,
		Title:         Bind("预约"),
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
						Model:         Employees(),
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
