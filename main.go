package main

import (
	"taskmanager"

	"github.com/ddkwork/app"
	"github.com/richardwilkes/unison"
)

func main() {
	app.Run("taskmanager", func(w *unison.Window) {
		w.Content().AddChild(taskmanager.New().Layout())
	})
}
