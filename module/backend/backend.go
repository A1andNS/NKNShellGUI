//package main
//
//import (
//	"fyne.io/fyne/v2"
//	"fyne.io/fyne/v2/app"
//	"fyne.io/fyne/v2/container"
//	"fyne.io/fyne/v2/layout"
//	"fyne.io/fyne/v2/theme"
//	"fyne.io/fyne/v2/widget"
//)
//
//func main() {
//	a := app.New()
//	w := a.NewWindow("Hello")
//	w.Resize(fyne.NewSize(800, 400))
//	//主题选择
//	themes := container.NewGridWithColumns(2, widget.NewButton("Light", func() {
//		a.Settings().SetTheme(theme.LightTheme())
//	}), widget.NewButton("Dark", func() {
//		a.Settings().SetTheme(theme.DarkTheme())
//	}))
//	//右侧显示文本内容
//	rightText := widget.NewLabel("hello")
//	//右面右侧
//	right := container.NewGridWithColumns(2, rightText)
//	//命令行接收
//	cmdEntry := widget.NewEntry()
//	cmdsubmit := widget.NewButton("Execute", func() {
//		rightText.SetText(cmdEntry.Text)
//	})
//	cmdSend := container.NewVBox(cmdEntry, layout.NewSpacer(), cmdsubmit)
//	//页面左侧栏
//	menuBox := container.NewVBox(widget.NewButton("Welcome", func() {
//		rightText.SetText("welcome to the tool")
//		if len(right.Objects) == 2 {
//			right.Remove(cmdSend)
//		}
//	}), widget.NewButton("Shell", func() {
//		rightText.SetText("shell page")
//		if len(right.Objects) < 2 {
//			right.Add(cmdSend)
//		}
//	}), widget.NewButton("About", func() {
//		rightText.SetText("This is the A shell Controller based on NKNNetwork")
//		if len(right.Objects) == 2 {
//			right.Remove(cmdSend)
//		}
//	}))
//	//可以滚动的左侧栏
//	leftSroll := container.NewHScroll(menuBox)
//	//左全栏,使用空白布局填充
//	leftbar := container.NewVBox(leftSroll, layout.NewSpacer(), themes)
//	//主页面
//	maingroup := container.NewHSplit(leftbar, right)
//	maingroup.Offset = 0.2
//	w.SetContent(maingroup)
//	w.ShowAndRun()
//
//}
// Package main provides various examples of Fyne API capabilities.
package main

import (
	"fyne.io/fyne/v2/dialog"
	"log"
	"net/url"

	"NKNshellgo/module/tutorials"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/cmd/fyne_settings/settings"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

const preferenceCurrentTutorial = "currentTutorial"

var topWindow fyne.Window

func main() {
	//新建APP并且设置ID
	a := app.NewWithID("io.nknShell.nn")
	//设置app图标
	a.SetIcon(tutorials.ResourceLOGOPng)
	//生命周期管理
	logLifecycle(a)
	//新建一个窗口，斌且传递给topWindows
	w := a.NewWindow("NN")
	topWindow = w
	//设置主菜单
	w.SetMainMenu(makeMenu(a, w))
	//退出此川口就退出了程序
	w.SetMaster()
	//content布局是max
	content := container.NewMax()
	//文本栏title和intro
	//title := widget.NewLabel("Component name")
	//intro := widget.NewLabel("An introduction would probably go\nhere, as well as a")
	//intro文字自动换行
	//intro.Wrapping = fyne.TextWrapWord

	setTutorial := func(t tutorials.Tutorial) {
		//title.SetText(t.Title)
		//intro.SetText(t.Intro)
		// content容器内的布局
		content.Objects = []fyne.CanvasObject{t.View(w)}
		//刷新content容器
		content.Refresh()
	}

	tutorial := container.NewBorder(
		nil, nil, nil, nil, content)
	split := container.NewHSplit(makeNav(setTutorial, true), tutorial)
	split.Offset = 0.2
	w.SetContent(split)
	w.Resize(fyne.NewSize(840, 460))
	w.ShowAndRun()
}

//记录生命周期，开始和结束，前台和后台
func logLifecycle(a fyne.App) {
	a.Lifecycle().SetOnStarted(func() {
		log.Println("Lifecycle: Started")
	})
	a.Lifecycle().SetOnStopped(func() {
		log.Println("Lifecycle: Stopped")
	})
	a.Lifecycle().SetOnEnteredForeground(func() {
		log.Println("Lifecycle: Entered Foreground")
	})
	a.Lifecycle().SetOnExitedForeground(func() {
		log.Println("Lifecycle: Exited Foreground")
	})
}

//创建菜单
func makeMenu(a fyne.App, w fyne.Window) *fyne.MainMenu {
	//新建其他项目
	mail := fyne.NewMenuItem("Mail", func() {
		mailInfo := dialog.NewInformation("EMAIL", "lilongxinalan@gmail.com", w)
		mailInfo.Show()
	})
	//设置菜单项目
	settingsItem := fyne.NewMenuItem("Settings", func() {
		w := a.NewWindow("Settings")
		w.SetContent(settings.NewSettings().LoadAppearanceScreen(w))
		w.Resize(fyne.NewSize(480, 480))
		w.Show()
	})
	//help菜单项目
	helpMenu := fyne.NewMenu("Help",
		fyne.NewMenuItem("Documentation", func() {
			u, _ := url.Parse("https://github.com/A1andNS")
			_ = a.OpenURL(u)
		}),
		fyne.NewMenuItem("Author", func() {
			u, _ := url.Parse("https://github.com/A1andNS")
			_ = a.OpenURL(u)
		}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("About", func() { //关于页面，新弹窗并且居中
			w1 := a.NewWindow("About")
			w1.Resize(fyne.Size{Width: 200, Height: 200})
			w1.CenterOnScreen()
			aboutPage := container.NewVBox(layout.NewSpacer(), widget.NewLabelWithStyle("This is a Fujian Agriculture and Forest University Undergraduate Graduation Project", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}), widget.NewLabelWithStyle("NN v1.0 rc", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}), widget.NewLabelWithStyle("Copyright© 2022 A1andNS. All Rights Reserved.", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}), layout.NewSpacer())
			w1.SetContent(aboutPage)
			w1.Show()
		}))

	// a quit item will be appended to our first (File) menu
	file := fyne.NewMenu("File", mail, fyne.NewMenuItemSeparator(), settingsItem)
	return fyne.NewMainMenu(
		file,
		helpMenu,
	)
}

func makeNav(setTutorial func(tutorial tutorials.Tutorial), loadPrevious bool) fyne.CanvasObject {
	//获取当前app
	a := fyne.CurrentApp()

	tree := &widget.Tree{
		ChildUIDs: func(uid string) []string {
			return tutorials.TutorialIndex[uid]
		},
		IsBranch: func(uid string) bool {
			children, ok := tutorials.TutorialIndex[uid]

			return ok && len(children) > 0
		},
		CreateNode: func(branch bool) fyne.CanvasObject {
			return widget.NewLabel("Collection Widgets")
		},
		UpdateNode: func(uid string, branch bool, obj fyne.CanvasObject) {
			t, ok := tutorials.Tutorials[uid]
			if !ok {
				fyne.LogError("Missing tutorial panel: "+uid, nil)
				return
			}
			obj.(*widget.Label).SetText(t.Title)
		},
		OnSelected: func(uid string) {
			if t, ok := tutorials.Tutorials[uid]; ok {
				a.Preferences().SetString(preferenceCurrentTutorial, uid)
				setTutorial(t)
			}
		},
	}

	if loadPrevious {
		currentPref := a.Preferences().StringWithFallback(preferenceCurrentTutorial, "welcome")
		tree.Select(currentPref)
	}

	themes := fyne.NewContainerWithLayout(layout.NewGridLayout(2),
		widget.NewButton("Dark", func() {
			a.Settings().SetTheme(theme.DarkTheme())
		}),
		widget.NewButton("Light", func() {
			a.Settings().SetTheme(theme.LightTheme())
		}),
	)

	return container.NewBorder(nil, themes, nil, nil, tree)
}

func shortcutFocused(s fyne.Shortcut, w fyne.Window) {
	if focused, ok := w.Canvas().Focused().(fyne.Shortcutable); ok {
		focused.TypedShortcut(s)
	}
}
