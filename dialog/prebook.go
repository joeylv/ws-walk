package dialog

import (
	"../models"
	"github.com/lxn/walk"
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
						Text: "Name:",
					},
					LineEdit{
						Text: Bind("Name", SelRequired{}),
					},
					Label{
						Text: "Mobile:",
					},
					LineEdit{
						Text: Bind("Mobile"),
					},

					Label{
						ColumnSpan: 2,
						Text:       "Remarks:",
					},
					TextEdit{
						ColumnSpan: 2,
						MinSize:    Size{100, 50},
						Text:       Bind("Remarks"),
					},
					Label{
						Text: "Prod:",
					},
					ComboBox{
						Value:         Bind("Prod", SelRequired{}),
						BindingMember: "Id",
						DisplayMember: "Name",
						Model:         ProdList(),
					},
					Label{
						Text: "ArrivalDate:",
					},
					DateEdit{
						Date: Bind("ArrivalDate"),

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
