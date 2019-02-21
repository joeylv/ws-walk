package dialog

import (
	"../models"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"log"
)

func init() {
	//fmt.Println("Member Init")
	//member :=Member{}
}
func AddEmployee(owner walk.Form, member *models.Employee) (int, error) {
	var dlg *walk.Dialog
	var db *walk.DataBinder
	var acceptPB, cancelPB *walk.PushButton
	return Dialog{
		AssignTo:      &dlg,
		Title:         "添加员工",
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
						Text: "工号:",
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

							//if db.Dirty() {
							//	fmt.Println(db.Dirty())
							//	fmt.Println(member.Name)
							//	fmt.Println(db.DataSource())

							if err := db.Submit(); err != nil {
								log.Print(err)
								walk.MsgBox(owner, "错误提示", err.Error(), walk.MsgBoxIconError)
								return
							}
							dlg.Accept()
							//} else {
							//fmt.Println(db.Dirty())
							//fmt.Println(member.Name)
							//fmt.Println(member.Mobile)
							//fmt.Println(member.Code)
							//fmt.Println(member.Remarks)
							//fmt.Println(db.DataSource())
							//if err := db.Submit(); err != nil {
							//	log.Print(err)
							//	walk.MsgBox(owner, "错误提示", err.Error(), walk.MsgBoxIconError)
							//	return
							//}
							//dbcon.Create("member", member.Name, member.Mobile, member.Code, member.Remarks)
							//	walk.MsgBox(owner, "错误提示", "...", walk.MsgBoxIconError)
							//}

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
