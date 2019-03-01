package main

import (
	"./compos"
	. "./manager"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"log"
	"time"
)

var isSpecialMode = walk.NewMutableCondition()

//func init() {
//	fmt.Println("Init")
//}

func main() {
	//mw := &MyMainWindow{}
	//MustRegisterCondition("isSpecialMode", isSpecialMode)
	//now := time.Now()
	//today := time.Date(now.Year(), now.Month(), now.Day()+1, 00, 00, 00, 00, time.UTC)
	//tom := time.Date(now.Year(), now.Month(), now.Day()+2, 00, 00, 00, 00, time.UTC)
	//preToday := mw.GetModel("PreBook", &today)
	//preTom := mw.GetModel("PreBook", &today, &tom)
	//&preSet{&walk.TableView{}, preToday, "今日预约", len(preToday.Items)}
	//TomPre := &PreSet{&walk.TableView{}, preTom, "明日预约", len(preTom.Items)}
	//mw.TodayPre = &PreSet{&walk.TableView{}, preToday, "今日预约", len(preToday.Items)}
	//mw.TomPre = TomPre

	MustRegisterCondition("isSpecialMode", isSpecialMode)
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day()+1, 00, 00, 00, 00, time.UTC)
	tom := time.Date(now.Year(), now.Month(), now.Day()+2, 00, 00, 00, 00, time.UTC)
	preToday := PreBookModel(&today)
	preTom := PreBookModel(&today, &tom)
	//&preSet{&walk.TableView{}, preToday, "今日预约", len(preToday.Items)}
	mw := &MyMainWindow{TodayPre: &PreSet{&walk.TableView{}, preToday, "今日预约", len(preToday.Items)}, TomPre: &PreSet{&walk.TableView{}, preTom, "明日预约", len(preTom.Items)}}

	var openAction, showAboutBoxAction *walk.Action
	//var recentMenu *walk.Menu
	var toggleSpecialModePB *walk.PushButton

	if err := (MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "無舍健康管理",
		Icon:     "img/stop.ico",
		MenuItems: []MenuItem{
			Menu{
				Text: "&会员消费",
				Items: []MenuItem{
					Action{
						AssignTo:    &openAction,
						Text:        "&Open",
						Image:       "/img/open.png",
						Enabled:     Bind("enabledCB.Checked"),
						Visible:     Bind("!openHiddenCB.Checked"),
						Shortcut:    Shortcut{walk.ModControl, walk.KeyO},
						OnTriggered: mw.OpenAction_Triggered,
					},
					Action{
						Text:        "会员管理",
						OnTriggered: mw.OpenDialog,
					},
					Action{
						Text:        "会员管理",
						OnTriggered: mw.OpenMembers,
					},
					Action{
						Text:        "预约管理",
						OnTriggered: mw.OpenPreBooks,
					},
					Action{
						Text:        "消费管理",
						OnTriggered: mw.OpenRecords,
					},
					//Menu{
					//	AssignTo: &recentMenu,
					//	Text:     "Recent",
					//},
					Separator{},
					Action{
						Text:        "E&xit",
						OnTriggered: func() { mw.Close() },
					},
				},
			},
			Menu{
				Text: "&服务疗程",
				Items: []MenuItem{
					Action{
						Text:        "项目管理",
						OnTriggered: mw.OpenProducts,
					},
					Action{
						Text:        "疗程管理",
						OnTriggered: mw.OpenCombos,
					},
					Action{
						Text:        "员工管理",
						OnTriggered: mw.OpenEmployees,
					},
				},
			},
			Menu{
				Text: "&Help",
				Items: []MenuItem{
					Action{
						Text:        "初始化",
						OnTriggered: mw.InitDataBase,
					},
					Action{
						AssignTo:    &showAboutBoxAction,
						Text:        "About",
						OnTriggered: mw.ShowAboutBoxAction_Triggered,
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
							OnTriggered: mw.NewPreBook,
						},
						Action{
							Text:        "项目",
							OnTriggered: mw.NewProd,
						},
						Action{
							Text:        "疗程",
							OnTriggered: mw.NewCombo,
						},
					},
					OnTriggered: mw.NewPreBook,
				},
				Separator{},
				Menu{
					Text:  "View",
					Image: "/img/document-properties.png",
					Items: []MenuItem{
						Action{
							Text:        "会员",
							OnTriggered: mw.NewMember,
						},
						Action{
							Text:        "员工",
							OnTriggered: mw.NewEmployee,
						},
						Action{
							Text:        "消费",
							OnTriggered: mw.NewRecord,
						},
						Action{
							Text:        "折扣卡",
							OnTriggered: mw.NewCard,
						},
					},
					OnTriggered: mw.NewMember,
				},
				Separator{},
				Action{
					Text:        "Special",
					Image:       "/img/system-shutdown.png",
					Enabled:     Bind("isSpecialMode && enabledCB.Checked"),
					OnTriggered: mw.SpecialAction_Triggered,
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
				Layout:  HBox{MarginsZero: true, SpacingZero: true},
				MinSize: Size{1000, 770},
				//MaxSize: Size{1000, 770},
				Children: compos.Prebook(mw.TodayPre, mw.TomPre),
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

	//addRecentFileActions := func(texts ...string) {
	//	for _, text := range texts {
	//		a := walk.NewAction()
	//		a.SetText(text)
	//		a.Triggered().Attach(mw.OpenAction_Triggered)
	//		recentMenu.Actions().Add(a)
	//	}
	//}
	//
	//addRecentFileActions("Foo", "Bar", "Baz")

	mw.Run()

}
