package main

import (
	"log"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type Gui struct {
	App      fyne.App
	Client   Client
	Messages []Message
}

func NewGui(app fyne.App, client Client) *Gui {
	return &Gui{App: app, Client: client, Messages: []Message{}}
}

func (g Gui) Run() error {
	g.showLogin()
	g.App.Run()

	return nil
}

func (g Gui) refresh(grid *fyne.Container) {
	grid.RemoveAll()
	var err error
	g.Messages, err = g.Client.getmessage()

	if err != nil {
		log.Println(err)
	}
	for _, msg := range g.Messages {
		grid.Add(layout.NewSpacer())

		label1 := widget.NewLabel(msg.Username + ": ")
		value1 := widget.NewLabel(msg.Content)

		grid2 := container.New(layout.NewFormLayout(), label1, value1)
		grid.Add(grid2)
	}
}

func (g Gui) addForm(grid *fyne.Container) *widget.Form {
	textArea := widget.NewMultiLineEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{},
		OnSubmit: func() {
			msg := Message{Username: g.Client.Username, Content: textArea.Text}

			err := g.Client.sendmessage(msg)
			if err != nil {
				log.Println(err)
			}
			grid.Add(layout.NewSpacer())

			label1 := widget.NewLabel(g.Client.Username + ": ")
			value1 := widget.NewLabel(textArea.Text)

			grid2 := container.New(layout.NewFormLayout(), label1, value1)
			grid.Add(grid2)

		},
	}

	form.Append("Nachricht", textArea)
	return form
}

func (g Gui) showLogin() {
	myWindow := g.App.NewWindow("Fam Chat Login")

	username := widget.NewEntry()
	server := widget.NewEntry()
	server.Text = "http://localhost:110"

	form := &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{Text: "Username", Widget: username}, {Text: "Server", Widget: server}},
		OnSubmit: func() { // optional, handle form submission
			g.Client.Uri = server.Text
			g.Client.Username = username.Text
			g.showApp()
			myWindow.Close()
		},
	}

	// we can also append items
	// form.Append("Text", textArea)

	myWindow.SetContent(form)
	myWindow.Resize(fyne.NewSize(600, 40))
	myWindow.Show()
}

func (g Gui) showApp() {
	myWindow := g.App.NewWindow("Fam Chat")

	grid := container.New(layout.NewVBoxLayout())
	form := g.addForm(grid)

	ticker := time.NewTicker(5 * time.Second)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				g.refresh(grid)
			}
		}
	}()
	g.refresh(grid)

	myWindow.SetContent(container.New(layout.NewVBoxLayout(), grid, layout.NewSpacer(), form))
	myWindow.Resize(fyne.NewSize(1024, 768))
	myWindow.Show()
}
