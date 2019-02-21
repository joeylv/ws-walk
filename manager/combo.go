package manager

import (
	"../dialog"
	"../models"
	"fmt"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"log"
)

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

func Combos(owner *walk.MainWindow) (int, error) {
	mw := &MWindow{MainWindow: owner, model: ComboModel()}
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
						OnClicked: mw.openCombo,
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

func (mw *MWindow) openCombo() {
	//walk.MsgBox(*mw, "title", "sss", walk.MsgBoxIconInformation)
	//var outTE *walk.TextEdit
	combo := new(models.Combo)
	if cmd, err := dialog.AddCombo(mw, combo); err != nil {
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
