package manager

import (
	"../dialog"
	"../models"
	"fmt"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"log"
)

func MemberModel() *ItemModel {
	memList := models.Member{}.Search()

	m := new(ItemModel)
	m.items = make([]*Item, len(memList))
	for i, j := range memList {
		fmt.Println(i)
		fmt.Println(j.Name)
		//m.items
		m.items[i] = &Item{
			Index:   i,
			Name:    j.Name,
			Mobile:  j.Mobile,
			Remarks: j.Remarks,
		}
	}
	return m
}

type MWindow struct {
	*walk.MainWindow
	model *ItemModel
	tv    *walk.TableView
}

func Member(owner *walk.MainWindow) (int, error) {
	mw := &MWindow{MainWindow: owner, model: MemberModel()}
	var dlg *walk.Dialog
	//var db *walk.DataBinder
	//var acceptPB, cancelPB *walk.PushButton
	return Dialog{
		AssignTo: &dlg,
		Title:    "会员管理",
		MinSize:  Size{800, 600},
		Layout:   VBox{},
		Children: []Widget{
			Composite{
				Layout: HBox{MarginsZero: true},
				Children: []Widget{
					HSpacer{},
					PushButton{
						Text:      "添加",
						OnClicked: mw.openMember,
						//func() {
						//	//mw.model.items = append(mw.model.items, &Condom{
						//	//	Index: mw.model.Len() + 1,
						//	//	Name:  "第六感",
						//	//	Price: mw.model.Len() * 5,
						//	//})
						//	mw.model.PublishRowsReset()
						//	mw.tv.SetSelectedIndexes([]int{})
						//},
					},
					PushButton{
						Text: "删除",
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
					PushButton{
						Text: "ExecChecked",
						OnClicked: func() {
							for _, x := range mw.model.items {
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
							for i, x := range mw.model.items {
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
				Children: mw.tableColumn(),
			},
		},
	}.Run(owner)
}

func (mw *MWindow) tvItemactivated() {
	msg := ``
	for _, i := range mw.tv.SelectedIndexes() {
		msg = msg + "\n" + mw.model.items[i].Name
	}
	walk.MsgBox(mw, "title", msg, walk.MsgBoxIconInformation)
}

func (mw *MWindow) openMember() {
	//walk.MsgBox(*mw, "title", "sss", walk.MsgBoxIconInformation)
	//var outTE *walk.TextEdit
	member := new(models.Member)
	if cmd, err := dialog.AddMember(mw, member); err != nil {
		log.Print(err)
	} else if cmd == walk.DlgCmdOK {
		fmt.Println("DlgCmdOK")
		member.Save()
		mw.model.items = append(mw.model.items, &Item{
			Index:   mw.model.Len(),
			Name:    member.Name,
			Mobile:  member.Mobile,
			Remarks: member.Remarks,
		})
		mw.model.PublishRowsReset()
		//outTE.SetText(fmt.Sprintf("%+v", member))
	}
}

func (mw MWindow) tableColumn() []Widget {
	return []Widget{
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
	}
}
