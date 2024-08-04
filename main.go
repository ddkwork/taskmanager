package main

import (
	"taskmanager"

	"github.com/ddkwork/unison"
	"github.com/ddkwork/unison/app"
)

func main() {
	app.Run("taskmanager", func(w *unison.Window) {
		w.Content().AddChild(taskmanager.New().Layout())
	})
}
