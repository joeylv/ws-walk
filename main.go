package main

import (
	"./compos"
	"./dialog"
	"./manager"
	"./models"
	"fmt"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"log"
	"time"
)

var isSpecialMode = walk.NewMutableCondition()

type MyMainWindow struct {
	*walk.MainWindow
}

//func init() {
//	fmt.Println("Init")
//}

func main() {
	MustRegisterCondition("isSpecialMode", isSpecialMode)

	mw := new(MyMainWindow)

	var openAction, showAboutBoxAction *walk.Action
	var recentMenu *walk.Menu
	var toggleSpecialModePB *walk.PushButton

	if err := (MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "無舍健康管理",
		Icon:     "img/stop.ico",
		MenuItems: []MenuItem{
			Menu{
				Text: "&打开",
				Items: []MenuItem{
					Action{
						AssignTo:    &openAction,
						Text:        "&Open",
						Image:       "/img/open.png",
						Enabled:     Bind("enabledCB.Checked"),
						Visible:     Bind("!openHiddenCB.Checked"),
						Shortcut:    Shortcut{walk.ModControl, walk.KeyO},
						OnTriggered: mw.openAction_Triggered,
					},
					//Action{
					//	Text:        "Dialog",
					//	OnTriggered: mw.openDialog,
					//},
					//Action{
					//	Text:        "Member",
					//	OnTriggered: mw.openMember,
					//},
					Action{
						Text:        "会员",
						OnTriggered: mw.openManager,
					},
					Action{
						Text:        "预约",
						OnTriggered: mw.openReserve,
					},
					Action{
						Text:        "项目",
						OnTriggered: mw.openProducts,
					},
					Action{
						Text:        "疗程",
						OnTriggered: mw.openProducts,
					},
					Action{
						Text:        "员工",
						OnTriggered: mw.openEmployees,
					},
					Menu{
						AssignTo: &recentMenu,
						Text:     "Recent",
					},
					Separator{},
					Action{
						Text:        "E&xit",
						OnTriggered: func() { mw.Close() },
					},
				},
			},
			Menu{
				Text: "&Help",
				Items: []MenuItem{
					Action{
						Text:        "InitDataBase",
						OnTriggered: mw.InitDataBase,
					},
					Action{
						AssignTo:    &showAboutBoxAction,
						Text:        "About",
						OnTriggered: mw.showAboutBoxAction_Triggered,
					},
				},
			},
		},
		ToolBar: ToolBar{
			ButtonStyle: ToolBarButtonImageBeforeText,
			Items: []MenuItem{
				ActionRef{&openAction},
				Menu{
					Text:  "New",
					Image: "/img/document-new.png",
					Items: []MenuItem{
						Action{
							Text:        "预约",
							OnTriggered: mw.newPreBook,
						},
						Action{
							Text:        "项目",
							OnTriggered: mw.newProd,
						},
						Action{
							Text:        "疗程",
							OnTriggered: mw.newProd,
						},
					},
					OnTriggered: mw.newPreBook,
				},
				Separator{},
				Menu{
					Text:  "View",
					Image: "/img/document-properties.png",
					Items: []MenuItem{
						Action{
							Text:        "会员",
							OnTriggered: mw.newMember,
						},
						Action{
							Text:        "员工",
							OnTriggered: mw.newEmployee,
						},
						//Action{
						//	Text:        "X",
						//	OnTriggered: mw.changeViewAction_Triggered,
						//},
						//Action{
						//	Text:        "Y",
						//	OnTriggered: mw.changeViewAction_Triggered,
						//},
						//Action{
						//	Text:        "Z",
						//	OnTriggered: mw.changeViewAction_Triggered,
						//},
					},
				},
				Separator{},
				Action{
					Text:        "Special",
					Image:       "/img/system-shutdown.png",
					Enabled:     Bind("isSpecialMode && enabledCB.Checked"),
					OnTriggered: mw.specialAction_Triggered,
				},
			},
		},
		ContextMenuItems: []MenuItem{
			ActionRef{&showAboutBoxAction},
		},
		//MaxSize: Size{1000, 800},
		//MinSize: Size{1000, 800},
		Layout: VBox{},
		Children: []Widget{
			Composite{
				Border:  true,
				Layout:  HBox{},
				MinSize: Size{1000, 770},
				//MaxSize: Size{1000, 770},
				Children: compos.Prebook(),
			},
			Composite{
				Layout:  Flow{SpacingZero: true},
				Border:  true,
				MaxSize: Size{1000, 40},
				Children: []Widget{
					CheckBox{
						Name:    "enabledCB",
						Text:    "Open / Special Enabled",
						Checked: true,
					},
					CheckBox{
						Name:    "openHiddenCB",
						Text:    "Open Hidden",
						Checked: true,
					},
					PushButton{
						AssignTo: &toggleSpecialModePB,
						Text:     "Enable Special Mode",
						OnClicked: func() {
							isSpecialMode.SetSatisfied(!isSpecialMode.Satisfied())

							if isSpecialMode.Satisfied() {
								toggleSpecialModePB.SetText("Disable Special Mode")
							} else {
								toggleSpecialModePB.SetText("Enable Special Mode")
							}
						},
					},
					PushButton{
						AssignTo: &toggleSpecialModePB,
						Text:     "Enable Special Mode",
						OnClicked: func() {
							//walk.MsgBox(mw, "New", "Newing something up... or not.", walk.MsgBoxIconQuestion)
							isSpecialMode.SetSatisfied(!isSpecialMode.Satisfied())

							if isSpecialMode.Satisfied() {
								toggleSpecialModePB.SetText("Disable Special Mode")
							} else {
								toggleSpecialModePB.SetText("Enable Special Mode")
							}
						},
					},
				},
			},
		},
	}.Create()); err != nil {
		log.Fatal(err)
	}

	addRecentFileActions := func(texts ...string) {
		for _, text := range texts {
			a := walk.NewAction()
			a.SetText(text)
			a.Triggered().Attach(mw.openAction_Triggered)
			recentMenu.Actions().Add(a)
		}
	}

	addRecentFileActions("Foo", "Bar", "Baz")

	mw.Run()

}

func (mw *MyMainWindow) openManager() {
	var outTE *walk.TextEdit
	if cmd, err := manager.Member(mw.MainWindow); err != nil {
		log.Print(err)
	} else if cmd == walk.DlgCmdOK {
		outTE.SetText(fmt.Sprintf("%+v", "CMD"))
	}
}
func (mw *MyMainWindow) openEmployees() {
	var outTE *walk.TextEdit
	if cmd, err := manager.Employees(mw.MainWindow); err != nil {
		log.Print(err)
	} else if cmd == walk.DlgCmdOK {
		outTE.SetText(fmt.Sprintf("%+v", "CMD"))
	}
}
func (mw *MyMainWindow) openReserve() {
	var outTE *walk.TextEdit
	if cmd, err := manager.PreBooks(mw.MainWindow); err != nil {
		log.Print(err)
	} else if cmd == walk.DlgCmdOK {
		outTE.SetText(fmt.Sprintf("%+v", "CMD"))
	}
}

func (mw *MyMainWindow) openProducts() {
	var outTE *walk.TextEdit
	if cmd, err := manager.Products(mw.MainWindow); err != nil {
		log.Print(err)
	} else if cmd == walk.DlgCmdOK {
		outTE.SetText(fmt.Sprintf("%+v", "CMD"))
	}
}

func (mw *MyMainWindow) openCombos() {
	//var outTE *walk.TextEdit
	if cmd, err := manager.Combos(mw.MainWindow); err != nil {
		log.Print(err)
	} else if cmd == walk.DlgCmdOK {
		fmt.Println("OK")
		//outTE.SetText(fmt.Sprintf("%+v", "CMD"))
	}
}

func (mw *MyMainWindow) newMember() {
	//var outTE *walk.TextEdit
	member := new(models.Member)
	if cmd, err := dialog.AddMember(mw, member); err != nil {
		log.Print(err)
	} else if cmd == walk.DlgCmdOK {
		fmt.Println("OK", member)
		member.Save()
		//outTE.SetText(fmt.Sprintf("%+v", member))
	}
}

func (mw *MyMainWindow) newEmployee() {
	//var outTE *walk.TextEdit
	emp := new(models.Employee)
	if cmd, err := dialog.AddEmployee(mw, emp); err != nil {
		log.Print(err)
	} else if cmd == walk.DlgCmdOK {
		fmt.Println("OK", emp)
		emp.Save()
		//outTE.SetText(fmt.Sprintf("%+v", member))
	}
}
func (mw *MyMainWindow) newPreBook() {
	//var outTE *walk.TextEdit
	preBook := new(models.PreBook)
	//loc, _ := time.LoadLocation("Local")
	preBook.ArrivalDate = time.Now()
	if cmd, err := dialog.AddPreBook(mw, preBook); err != nil {
		log.Print(err)
	} else if cmd == walk.DlgCmdOK {
		fmt.Println("OK")
		//outTE.SetText(fmt.Sprintf("%+v", member))
	}
}
func (mw *MyMainWindow) newProd() {
	//var outTE *walk.TextEdit
	prod := new(models.Prod)
	//loc, _ := time.LoadLocation("Local")
	if cmd, err := dialog.AddProduct(mw, prod); err != nil {
		log.Print(err)
	} else if cmd == walk.DlgCmdOK {
		fmt.Println("OK")
		//outTE.SetText(fmt.Sprintf("%+v", member))
	}
}

func (mw *MyMainWindow) newCombo() {
	combo := new(models.Combo)
	if cmd, err := dialog.AddCombo(mw, combo); err != nil {
		log.Print(err)
	} else if cmd == walk.DlgCmdOK {
		fmt.Println("OK")
	}
}

func (mw *MyMainWindow) openDialog() {
	//var outTE *walk.TextEdit
	animal := new(models.Animal)
	if cmd, err := dialog.RunAnimalDialog(mw, animal); err != nil {
		log.Print(err)
	} else if cmd == walk.DlgCmdOK {
		fmt.Println("OK")
		//outTE.SetText(fmt.Sprintf("%+v", animal))
	}
}

func (mw *MyMainWindow) InitDataBase() {
	models.Migrate()
	//dbcon.Create()
}
func (mw *MyMainWindow) openAction_Triggered() {
	walk.MsgBox(mw, "Open", "Pretend to open a file...", walk.MsgBoxIconInformation)
}

func (mw *MyMainWindow) newAction_Triggered() {
	walk.MsgBox(mw, "New", "Newing something up... or not.", walk.MsgBoxIconInformation)
}

func (mw *MyMainWindow) changeViewAction_Triggered() {
	walk.MsgBox(mw, "Change View", "By now you may have guessed it. Nothing changed.", walk.MsgBoxIconInformation)
}

func (mw *MyMainWindow) showAboutBoxAction_Triggered() {
	walk.MsgBox(mw, "About", "Walk Actions Example", walk.MsgBoxIconInformation)
}

func (mw *MyMainWindow) specialAction_Triggered() {
	walk.MsgBox(mw, "Special", "Nothing to see here.", walk.MsgBoxIconInformation)
}
