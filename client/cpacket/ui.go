package cpacket

import (
	"fmt"
	"strings"

	ui "github.com/gizak/termui"
)

var (
	Rows = []string{
		"[id] [%d]",
		"[服务器encrypt key] [%s]",
		"[你的加密key] [./app_private_key.pem,./app_public_key.pem]",
		"[你的朋友文件] [./user.json]",
		"typing:",
	}
)

var (
	par       *ui.Par
	ls        *ui.List
	lsContent *ui.List
)

func UIInit() {

	err := ui.Init()
	if err != nil {
		panic(err)
	}
	defer ui.Close()

	par = ui.NewPar(":PRESS Esc TO QUIT")
	par.Height = 3
	par.Width = 50
	par.TextFgColor = ui.ColorWhite
	par.BorderLabel = "Client Info"
	par.BorderFg = ui.ColorCyan
	par.X = 0
	par.Y = 0

	ls = ui.NewList()
	ls.Items = Rows
	ls.ItemFgColor = ui.ColorYellow
	ls.BorderLabel = "Infos"
	ls.Height = 7
	ls.Width = 100
	ls.Y = 4

	lsContent = ui.NewList()
	lsContent.Items = allMsg
	lsContent.ItemFgColor = ui.ColorGreen
	lsContent.BorderLabel = "Content"
	lsContent.Height = 100
	lsContent.Width = 210
	lsContent.Y = 12

	ui.Handle("<Escape>", func(ui.Event) {
		ui.StopLoop()
	})
	ui.EventHook(func(e ui.Event) {
		if e.Type == ui.KeyboardEvent {
			if e.ID == "<Enter>" {
				str := msg.Flush()
				Rows[4] = "typing:"
				SendMsg(str)
			} else {
				msg.Enque(e.ID)
				Rows[4] = fmt.Sprintf("typing:%s", strings.Join(msg, ""))
			}

			ls.Items = Rows

			reoladUI()
		}
	})

	ui.Render(ls, par, lsContent)
	ui.Loop()
}

func reoladUI() {
	if ls != nil {
		lsContent.Items = allMsg
		ui.Render(ls, par, lsContent)
	}
}

func closeUI() {
	ui.StopLoop()
}

var allMsg []string

func init() {
	allMsg = make([]string, 0)
}

func displaySend(str string) {
	msgs := append([]string{}, fmt.Sprintf("send:%s", str))
	msgs = append(msgs, allMsg...)
	allMsg = msgs
	reoladUI()
}
func displayRecv(uid uint64, str string) {
	msgs := append([]string{}, fmt.Sprintf("recv[%d]:%s", uid, str))
	msgs = append(msgs, allMsg...)
	allMsg = msgs
	reoladUI()
}
