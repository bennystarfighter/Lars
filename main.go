package main

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"strings"

	"github.com/marcusolsson/tui-go"
	"github.com/spf13/viper"
)

func main() {
	var bots []string
	viper.SetConfigName("config")
	if runtime.GOOS == "linux" {
		viper.AddConfigPath("$HOME/.config/lars-tui")
	}
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	history := tui.NewVBox()

	historyScroll := tui.NewScrollArea(history)
	historyScroll.SetAutoscrollToBottom(true)

	historyBox := tui.NewVBox(historyScroll)
	historyBox.SetBorder(true)
	input := tui.NewEntry()
	input.SetFocused(true)
	input.SetSizePolicy(tui.Expanding, tui.Maximum)

	inputBox := tui.NewHBox(input)
	inputBox.SetBorder(true)
	inputBox.SetSizePolicy(tui.Expanding, tui.Maximum)

	chat := tui.NewVBox(historyBox, inputBox)
	chat.SetSizePolicy(tui.Expanding, tui.Expanding)

	bots = viper.AllKeys()

	root := tui.NewHBox(chat)

	ui, err := tui.New(root)
	if err != nil {
		log.Fatal(err)
	}

	ui.SetKeybinding("Esc", func() { ui.Quit() })
	for _, element := range bots {
		path := viper.GetString(element)
		go launchbot(path, element, history, ui)
	}
	history.Append(tui.NewHBox(
		tui.NewLabel("Starting bots!"),
	))
	input.OnSubmit(func(e *tui.Entry) {
		in := e.Text()
		words := strings.Split(strings.ToLower(in), ",")
		if words[0] == "launch" {
			go launchbot(words[2], words[1], history, ui)
			input.SetText("")
		} else if words[0] == "exit" {
			ui.Quit()
		} else {
			input.SetText("")
		}
	})
	if err := ui.Run(); err != nil {
		log.Fatal(err)
	}
}

func launchbot(path string, id string, out *tui.Box, ui tui.UI) {
	//c := color.New(color.BgGreen, color.FgBlack).Add(color.Underline)
	out.Append(tui.NewHBox(
		tui.NewLabel("Bot " + id + " started from " + path),
	))
	cmd := exec.Command(path)
	output, err := cmd.StdoutPipe()
	if err != nil {
		out.Append(tui.NewHBox(
			tui.NewLabel(err.Error()),
		))
		return
	}
	reader := bufio.NewReader(output)
	err = cmd.Start()
	if err != nil {
		out.Append(tui.NewHBox(
			tui.NewLabel(err.Error()),
		))
		return
	}

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			out.Append(tui.NewHBox(
				tui.NewLabel(err.Error()),
			))
			return
		}
		out.Append(tui.NewHBox(
			tui.NewLabel(id + " : " + line),
		))
		ui.Repaint()
	}
}
