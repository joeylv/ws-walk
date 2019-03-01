package manager

import (
	"fmt"
	. "github.com/lxn/walk/declarative"
)

func (mw MyMainWindow) Manager(model string, args ...string) []Widget {
	return []Widget{
		Composite{
			Layout:   HBox{MarginsZero: true},
			Children: mw.Actions(model),
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
			Children: mw.tableColumn(args...),
		},
	}
}

func (mw MyMainWindow) Actions(model string) []Widget {
	return []Widget{
		HSpacer{},
		PushButton{
			Text: "添加",
			OnClicked: func() {
				switch model {
				case "Member":
					mw.NewMember()
				case "Record":
					mw.NewRecord()
				case "Employee":
					mw.NewEmployee()
				case "Prod":
					mw.NewProd()
				case "Combo":
					mw.NewCombo()
				case "PreBook":
					mw.NewPreBook()
				default:
					fmt.Println("::::::::未知类型 ")
				}
			},
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
	}
}

func (mw MyMainWindow) tableColumn(column ...string) []Widget {
	var tableViewColumn []TableViewColumn
	for _, title := range column {
		//fmt.Println(s)
		//fmt.Println(j)
		tableViewColumn = append(tableViewColumn, TableViewColumn{Title: title})
	}
	//fmt.Println(tableViewColumn)
	return []Widget{
		TableView{
			AssignTo: &mw.tv,
			//CheckBoxes:       true,
			ColumnsOrderable: true,
			//MultiSelection:   true,
			Columns: tableViewColumn,
			Model:   mw.model,
			OnCurrentIndexChanged: func() {
				i := mw.tv.CurrentIndex()
				if 0 <= i {
					fmt.Printf("OnCurrentIndexChanged: %v\n", mw.model.Items[i].Name)
				}
			},
			OnItemActivated: mw.tvItemactivated,
		},
	}
}
