package manager

import (
	"../models"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"log"
)

func ProductModel() *ItemModel {
	memList := models.Prod{}.Search()
	m := &ItemModel{table: "prod", Items: make([]*Item, len(memList))}
	//m.items = make([]*Item, len(memList))
	for i, j := range memList {
		m.Items[i] = &Item{
			Index:   i,
			Name:    j.Name,
			Price:   j.Price,
			Remarks: j.Remarks,
		}
	}
	return m
}

func Products(owner *MyMainWindow) (int, error) {
	//mw := &MyMainWindow{MainWindow: owner, model: ProductModel()}
	owner.model = ProductModel()
	var dlg *walk.Dialog
	//var db *walk.DataBinder
	//var acceptPB, cancelPB *walk.PushButton
	return Dialog{
		AssignTo: &dlg,
		Title:    "项目管理",
		MinSize:  Size{800, 600},
		Layout:   VBox{},
		Children: owner.Manager("Prod", "编号", "名称", "价格", "备注"),
	}.Run(owner)
}

func AddProduct(owner walk.Form, prod *models.Prod) (int, error) {
	var dlg *walk.Dialog
	var db *walk.DataBinder
	var acceptPB, cancelPB *walk.PushButton

	return Dialog{
		AssignTo:      &dlg,
		Title:         "项目",
		FixedSize:     true,
		DefaultButton: &acceptPB,
		CancelButton:  &cancelPB,
		DataBinder: DataBinder{
			AssignTo:       &db,
			AutoSubmit:     true,
			Name:           "Prod",
			DataSource:     prod,
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
						Text: Bind("Name"),
					},
					Label{
						Text: "价格:",
					},
					NumberEdit{
						Decimals: 1,
						Value:    Bind("Price"),
					},
					Label{
						Text: "备注:",
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
