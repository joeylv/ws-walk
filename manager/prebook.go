package manager

import (
	"../dialog"
	"../models"
	"fmt"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"log"
	"sort"
	"strconv"
	"time"
)

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

func PreBooks(owner *walk.MainWindow) (int, error) {
	mw := &MWindow{MainWindow: owner, model: PreBookModel(nil)}
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
						OnClicked: mw.openPreBook,
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

func (mw *MWindow) openPreBook() {
	preBook := &models.PreBook{ArrivalDate: time.Now()}
	if cmd, err := dialog.AddPreBook(mw, preBook); err != nil {
		log.Print(err)
	} else if cmd == walk.DlgCmdOK {
		fmt.Println("DlgCmdOK")
		preBook.Save()
		mw.model.Items = append(mw.model.Items, &Item{
			Index:   mw.model.Len(),
			Name:    preBook.Name,
			Remarks: preBook.Remarks,
		})
		mw.model.PublishRowsReset()
	}
}
