package manager

import (
	"../dialog"
	"../models"
	"fmt"
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

//type ProductMainWindow struct {
//	*walk.MainWindow
//	model *ItemModel
//	tv    *walk.TableView
//}

func Products(owner *walk.MainWindow) (int, error) {
	mw := &MWindow{MainWindow: owner, model: ProductModel()}
	var dlg *walk.Dialog
	//var db *walk.DataBinder
	//var acceptPB, cancelPB *walk.PushButton
	return Dialog{
		AssignTo: &dlg,
		Title:    "项目管理",
		MinSize:  Size{800, 600},
		Layout:   VBox{},
		Children: []Widget{
			Composite{
				Layout: HBox{MarginsZero: true},
				Children: []Widget{
					HSpacer{},
					PushButton{
						Text:      "添加",
						OnClicked: mw.openProduct,
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
				Children: mw.tableColumn("编号", "名称", "价格", "备注"),
				//[]Widget{
				//	TableView{
				//		AssignTo:         &mw.tv,
				//		CheckBoxes:       true,
				//		ColumnsOrderable: true,
				//		MultiSelection:   true,
				//		Columns: []TableViewColumn{
				//			{Title: "编号"},
				//			{Title: "名称"},
				//			{Title: "手机"},
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

func (mw *MWindow) openProduct() {
	prod := new(models.Prod)
	if cmd, err := dialog.AddProduct(mw, prod); err != nil {
		log.Print(err)
	} else if cmd == walk.DlgCmdOK {
		fmt.Println("DlgCmdOK")
		prod.Save()
		mw.model.Items = append(mw.model.Items, &Item{
			Index:   mw.model.Len(),
			Name:    prod.Name,
			Price:   prod.Price,
			Remarks: prod.Remarks,
		})
		mw.model.PublishRowsReset()
	}
}
