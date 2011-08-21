package main

import "curses"
import "os"
import "fmt"

func main() {

	startGoCurses()
	defer stopGoCurses()

    curses.Init_pair(1, curses.COLOR_YELLOW, curses.COLOR_GREEN);

    smallwin, _ := curses.Stdwin.Derwin(1, 10, 10, 10)
    smallwin.Attron(curses.Color_pair(1))
    smallwin.Attron(curses.A_BOLD)

    smallwin.Addstr(0, 0, "Hello here", 0)
    smallwin.Getch()

    curses.Stdwin.Clear()
    curses.Stdwin.Refresh()

    smallwin.Resize(1, 5)
    smallwin.Mvwin(20, 20)
    smallwin.Addstr(0, 0, "There", 0)
    smallwin.Getch()

}

func startGoCurses() {
	curses.Initscr()

	if curses.Stdwin == nil {
		stopGoCurses()
		os.Exit(1)
	}

	curses.Noecho()
	curses.Curs_set(curses.CURS_HIDE)
	curses.Stdwin.Keypad(true)

	if err := curses.Start_color(); err != nil {
		fmt.Printf("%s\n", err.String())
		stopGoCurses()
		os.Exit(1)
	}

	curses.Use_default_colors()

}

func stopGoCurses() {
	curses.Endwin()
}
