package compos

func Prebook() []Widget {
	return []Widget{
		Composite{
			Border:  true,
			Layout:  Grid{Columns: 2},
			MinSize: Size{500, 370},
			MaxSize: Size{500, 370},
			Children: []Widget{
				//Composite{
				//	Layout:  Grid{Columns: 2, Spacing: 10},
				//	Children: []Widget{
				Label{
					Text: "今日预约",
					//MinSize: Size{250, 20},
				},
				Label{
					Text: "Count:",
					//MinSize: Size{250, 20},
				},
				TableView{
					ColumnSpan: 2,
					//MaxSize: Size{500, 420},
					//AssignTo:         &mw.tv,
					CheckBoxes:       true,
					ColumnsOrderable: true,
					MultiSelection:   true,
					Columns: []TableViewColumn{
						{Title: "编号"},
						{Title: "名称"},
						{Title: "手机"},
						{Title: "备注"},
					},
					//Model: mw.model,
					OnCurrentIndexChanged: func() {
						//i := mw.tv.CurrentIndex()
						//if 0 <= i {
						//	fmt.Printf("OnCurrentIndexChanged: %v\n", mw.model.items[i].Name)
						//}
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
				//CheckBox{
				//	Name:    "enabledCB",
				//	Text:    "Open / Special Enabled",
				//	Checked: true,
				//},
				//CheckBox{
				//	Name:    "openHiddenCB",
				//	Text:    "Open Hidden",
				//	Checked: true,
				//},
			},
			//	},
			//},
		},
		Composite{
			Border: true,
			Layout: Grid{Columns: 2},
			//MinSize: Size{500, 770},
			//MaxSize: Size{500, 770},
			Children: []Widget{
				Label{
					Text: "明日预约",
					//MinSize: Size{250, 20},
					//MaxSize: Size{250, 20},
				},
				Label{
					Text: "Count:",
					//MinSize: Size{250, 20},
				},
				TableView{
					ColumnSpan: 2,
					//MaxSize: Size{500, 420},
					//AssignTo:         &mw.tv,
					CheckBoxes:       true,
					ColumnsOrderable: true,
					MultiSelection:   true,
					Columns: []TableViewColumn{
						{Title: "编号"},
						{Title: "名称"},
						{Title: "手机"},
						{Title: "备注"},
					},
					//Model: mw.model,
					OnCurrentIndexChanged: func() {
						//i := mw.tv.CurrentIndex()
						//if 0 <= i {
						//	fmt.Printf("OnCurrentIndexChanged: %v\n", mw.model.items[i].Name)
						//}
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
			},
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
	//return &Composite{
	//	Layout:  Grid{Columns: 2, Spacing: 10},
	//	MinSize: Size{1000, 750},
	//	Children: []Widget{
	//
	//	},
	//}
}
