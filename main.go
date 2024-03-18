package main

import (
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	// Define pink color
	pink := ui.ColorMagenta

	// Pet info text
	petInfo := "Name:      Kubi\nMood:      Neutral\nHealth:    Good\nAge:       1 day\nSize:      Tiny"

	// Pet info widget
	petInfoWidget := widgets.NewParagraph()
	petInfoWidget.Text = petInfo
	petInfoWidget.SetRect(0, 0, 30, 10) // Adjust size accordingly
	petInfoWidget.TextStyle.Fg = pink
	petInfoWidget.BorderStyle.Fg = pink

	// Pet ASCII art
	petArtNeutral := `
            .^._.^.
            | . . |
           (  ---  )
           .'     '.
           |/     \|
            \ /-\ /
             V   V`

	// Calculate the starting x position to centralize the pet art
	startX := (33-len(".^._.^."))/2 + 31 // Center the art within the given width

	// Pet art widget
	petArtWidget := widgets.NewParagraph()
	petArtWidget.Text = petArtNeutral
	petArtWidget.SetRect(31, 0, 64, 10) // Adjust size and position accordingly
	petArtWidget.TextStyle.Fg = pink
	petArtWidget.BorderStyle.Fg = pink
	petArtWidget.PaddingLeft = startX
	petArtWidget.PaddingTop = 1

	// Kubepet label widget
	kubepetLabel := widgets.NewParagraph()
	kubepetLabel.TextStyle.Fg = pink
	kubepetLabel.BorderStyle.Fg = pink
	// Center the label between the top and bottom widgets
	labelHeight := 3
	kubepetLabel.SetRect(0, 10, 64, 10+labelHeight)

	// Description widget below the pet info and art
	descriptionWidget := widgets.NewParagraph()
	descriptionWidget.Text = "Kubi will live in your Kubernetes cluster and grow based on your cluster size.\nThe health of your Kubi depends on your cluster status.\nNo one knows what decides its mood.\nWait too long and it might leave you.\nYou can feed Kubi by deploying food."
	descriptionWidget.SetRect(0, 10+labelHeight, 64, 20) // Adjust size and position accordingly
	descriptionWidget.TextStyle.Fg = pink
	descriptionWidget.BorderStyle.Fg = pink

	ui.Render(petInfoWidget, petArtWidget, kubepetLabel, descriptionWidget) // Render the widgets

	// Event loop
	for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			break
		}
	}
}
