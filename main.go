package main

import (
	"fmt"
	"os"

	"go.i3wm.org/i3/v4"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Need exacty one argument. Exiting.")
		os.Exit(1)
	}
	mark := os.Args[1]
	recv := i3.Subscribe(i3.WindowEventType)
	for recv.Next() {
		ev := recv.Event().(*i3.WindowEvent)
		if ev.Change != "new" {
			continue
		}

		cmd := fmt.Sprintf("[con_id=%v] mark %s", ev.Container.ID, mark)
		i3.RunCommand(cmd)
		break
	}

}
