package compos

import (
	. "../manager"

	"fmt"
	. "github.com/lxn/walk/declarative"
	"strconv"
)

//var tvToday, tvTom *walk.TableView
//var now = time.Now()
//var today = time.Date(now.Year(), now.Month(), now.Day()+1, 00, 00, 00, 00, time.UTC)
//var tom = time.Date(now.Year(), now.Month(), now.Day()+2, 00, 00, 00, 00, time.UTC)

//var todayModel = PreBookModel(&today)
//var tomModel = PreBookModel(&today, &tom)
//func init() {
//	fmt.Println("Init")
//}
//type preSet struct {
//	tv    *walk.TableView
//	model *ItemModel
//	title string
//	count int
//}

func Prebook(todaySet *PreSet, tomSet *PreSet) []Widget {
	//now := time.Now()
	//today := time.Date(now.Year(), now.Month(), now.Day()+1, 00, 00, 00, 00, time.UTC)
	//tom := time.Date(now.Year(), now.Month(), now.Day()+2, 00, 00, 00, 00, time.UTC)
	//preToday := PreBookModel(&today)
	//preTom := PreBookModel(&today, &tom)
	//fmt.Println(len(preToday.Items))
	return []Widget{
		Composite{
			Border:   true,
			Layout:   Grid{Columns: 2},
			MinSize:  Size{500, 370},
			MaxSize:  Size{500, 370},
			Children: tableView(todaySet),
		},
		Composite{
			Border: true,
			Layout: Grid{Columns: 2},
			//MinSize: Size{500, 770},
			//MaxSize: Size{500, 770},
			Children: tableView(tomSet),
		},
		Composite{
			//Layout:  Grid{Columns: 4, Spacing: 10},
			//MinSize: Size{1000, 100},
			Children: []Widget{
				Label{
					Text: "Mobile:",
				},
			},
		},
	}
}

func tableView(set *PreSet) []Widget {
	return []Widget{
		Label{
			Text: set.Title + strconv.Itoa(set.Count) + "人次",
			//MinSize: Size{250, 20},
		},
		TableView{
			ColumnSpan: 2,
			//MaxSize: Size{500, 420},
			AssignTo: &set.TableView,
			//CheckBoxes:       true,
			ColumnsOrderable: true,
			MultiSelection:   true,
			Columns: []TableViewColumn{
				{Title: "姓名"},
				{Title: "手机"},
				{Title: "预约时间"},
				{Title: "备注"},
			},
			Model: set.ItemModel,
			OnCurrentIndexChanged: func() {
				i := set.TableView.CurrentIndex()
				if 0 <= i {
					fmt.Printf("OnCurrentIndexChanged: %v\n", set.ItemModel.Items[i])
				}
			},
			//OnItemActivated: mw.tvItemactivated,
		},
		Label{
			Text: "XXX:",
			//MinSize: Size{250, 20},
		},
		Label{
			Text: "XXX:",
			//MinSize: Size{250, 20},
		},
	}
}
