package manager

import (
	"../dialog"
	"../models"
	"fmt"
	"github.com/lxn/walk"
	"log"
)

//
//type Product struct {
//	Index   int
//	Name    string
//	Price   int
//	Remarks string
//	checked bool
//}
//
//type ItemModel struct {
//	walk.TableModelBase
//	walk.SorterBase
//	sortColumn int
//	sortOrder  walk.SortOrder
//	items      []*Item
//}
//
//func (m *ItemModel) RowCount() int {
//	return len(m.items)
//}
//
//func (m *ItemModel) Value(row, col int) interface{} {
//	item := m.items[row]
//	fmt.Println(m.items)
//	switch col {
//	case 0:
//		return item.Index
//	case 1:
//		return item.Name
//	case 2:
//		return item.Price
//	case 3:
//		return item.Remarks
//	}
//	panic("unexpected col")
//}
//
//func (m *ItemModel) Checked(row int) bool {
//	return m.items[row].checked
//}
//
//func (m *ItemModel) SetChecked(row int, checked bool) error {
//	m.items[row].checked = checked
//	return nil
//}
//
//func (m *ItemModel) Sort(col int, order walk.SortOrder) error {
//	m.sortColumn, m.sortOrder = col, order
//
//	sort.Stable(m)
//
//	return m.SorterBase.Sort(col, order)
//}
//
//func (m *ItemModel) Len() int {
//	return len(m.items)
//}
//
//func (m *ItemModel) Less(i, j int) bool {
//	fmt.Println(m)
//	a, b := m.items[i], m.items[j]
//
//	c := func(ls bool) bool {
//		if m.sortOrder == walk.SortAscending {
//			return ls
//		}
//
//		return !ls
//	}
//
//	switch m.sortColumn {
//	case 0:
//		return c(a.Index < b.Index)
//	case 1:
//		return c(a.Name < b.Name)
//	case 2:
//		return c(a.Price < b.Price)
//	case 3:
//		return c(a.Remarks < b.Remarks)
//	}
//
//	panic("unreachable")
//}
//
//func (m *ItemModel) Swap(i, j int) {
//	m.items[i], m.items[j] = m.items[j], m.items[i]
//}
//
func ProductModel() *ItemModel {
	memList := models.Prod{}.Search()
	m := new(ItemModel)
	m.items = make([]*Item, len(memList))
	for i, j := range memList {
		m.items[i] = &Item{
			Index:   i,
			Name:    j.Name,
			Price:   j.Price,
			Remarks: j.Remarks,
		}
	}
	return m
}

type ProductMainWindow struct {
	*walk.MainWindow
	model *ItemModel
	tv    *walk.TableView
}

func Products(owner *walk.MainWindow) (int, error) {
	mw := &ProductMainWindow{MainWindow: owner, model: ProductModel()}
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
						Text: "Delete",
						OnClicked: func() {
							var items []*Item
							remove := mw.tv.SelectedIndexes()
							for i, x := range mw.model.items {
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
							mw.model.items = items
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
				Children: []Widget{
					TableView{
						AssignTo:         &mw.tv,
						CheckBoxes:       true,
						ColumnsOrderable: true,
						MultiSelection:   true,
						Columns: []TableViewColumn{
							{Title: "编号"},
							{Title: "名称"},
							{Title: "手机"},
							{Title: "备注"},
						},
						Model: mw.model,
						OnCurrentIndexChanged: func() {
							i := mw.tv.CurrentIndex()
							if 0 <= i {
								fmt.Printf("OnCurrentIndexChanged: %v\n", mw.model.items[i].Name)
							}
						},
						OnItemActivated: mw.tvItemactivated,
					},
				},
			},
		},
	}.Run(owner)
}

func (mw *ProductMainWindow) tvItemactivated() {
	msg := ``
	for _, i := range mw.tv.SelectedIndexes() {
		msg = msg + "\n" + mw.model.items[i].Name
	}
	walk.MsgBox(mw, "title", msg, walk.MsgBoxIconInformation)
}

func (mw *ProductMainWindow) openProduct() {
	prod := new(models.Prod)
	if cmd, err := dialog.AddProduct(mw, prod); err != nil {
		log.Print(err)
	} else if cmd == walk.DlgCmdOK {
		fmt.Println("DlgCmdOK")
		prod.Save()
		mw.model.items = append(mw.model.items, &Item{
			Index:   mw.model.Len(),
			Name:    prod.Name,
			Price:   prod.Price,
			Remarks: prod.Remarks,
		})
		//mw.model.PublishRowsReset()
	}
}
