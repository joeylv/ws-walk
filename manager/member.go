package manager

import (
	"../models"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"log"
)

func Member(owner *MyMainWindow) (int, error) {
	//mw := &MyMainWindow{MainWindow: owner, model: MemberModel()}
	owner.model = MemberModel()
	//mw.model = mw.GetModel("Member")
	var dlg *walk.Dialog
	//var db *walk.DataBinder
	//var acceptPB, cancelPB *walk.PushButton
	return Dialog{
		AssignTo: &dlg,
		Title:    "会员管理",
		MinSize:  Size{800, 600},
		Layout:   VBox{},
		Children: owner.Manager("Member", "编号", "名称", "手机", "备注"),
	}.Run(owner)
}

func AddMember(owner walk.Form, member *models.Member) (int, error) {
	var dlg *walk.Dialog
	var db *walk.DataBinder
	var acceptPB, cancelPB *walk.PushButton
	return Dialog{
		AssignTo:      &dlg,
		Title:         "添加会员",
		DefaultButton: &acceptPB,
		CancelButton:  &cancelPB,
		DataBinder: DataBinder{
			AssignTo:       &db,
			AutoSubmit:     true,
			Name:           "Member",
			DataSource:     member,
			ErrorPresenter: ToolTipErrorPresenter{},
		},
		MinSize: Size{600, 300},
		Layout:  VBox{},
		Children: []Widget{
			Composite{
				Layout: Grid{Columns: 2},
				Children: []Widget{
					Label{
						Text: "姓名:",
					},
					LineEdit{
						Text: Bind("Name"),
					},
					Label{
						Text: "手机:",
					},
					LineEdit{
						Text:      Bind("Mobile"),
						MaxLength: 11,
					},
					Label{
						Text: "编号:",
					},
					LineEdit{
						Text: Bind("Code"),
					},
					Label{
						ColumnSpan: 2,
						Text:       "备注:",
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
								walk.MsgBox(owner, "错误提示", err.Error(), walk.MsgBoxIconError)
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
