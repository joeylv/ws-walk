package compos

import (
	. "../manager"
	"fmt"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"strconv"
	"time"
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
type preSet struct {
	tv    *walk.TableView
	model *ItemModel
	title string
	count int
}

func Prebook() []Widget {
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day()+1, 00, 00, 00, 00, time.UTC)
	tom := time.Date(now.Year(), now.Month(), now.Day()+2, 00, 00, 00, 00, time.UTC)
	preToday := PreBookModel(&today)
	preTom := PreBookModel(&today, &tom)
	fmt.Println(len(preToday.Items))
	return []Widget{
		Composite{
			Border:  true,
			Layout:  Grid{Columns: 2},
			MinSize: Size{500, 370},
			MaxSize: Size{500, 370},

			Children: tableView(&preSet{&walk.TableView{}, preToday, "今日预约", len(preToday.Items)}),
		},
		Composite{
			Border: true,
			Layout: Grid{Columns: 2},
			//MinSize: Size{500, 770},
			//MaxSize: Size{500, 770},
			Children: tableView(&preSet{&walk.TableView{}, preTom, "明日预约", len(preTom.Items)}),
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

func tableView(set *preSet) []Widget {
	return []Widget{
		Label{
			Text: set.title + strconv.Itoa(set.count) + "人次",
			//MinSize: Size{250, 20},
		},
		TableView{
			ColumnSpan: 2,
			//MaxSize: Size{500, 420},
			AssignTo: &set.tv,
			//CheckBoxes:       true,
			ColumnsOrderable: true,
			MultiSelection:   true,
			Columns: []TableViewColumn{
				{Title: "姓名"},
				{Title: "手机"},
				{Title: "预约时间"},
				{Title: "备注"},
			},
			Model: set.model,
			OnCurrentIndexChanged: func() {
				i := set.tv.CurrentIndex()
				if 0 <= i {
					//fmt.Printf("OnCurrentIndexChanged: %v\n", )
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
