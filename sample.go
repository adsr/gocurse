package main

import "curses"
import "os"
import "fmt"

func main() {
	x := 10;
	y := 10;
	startGoCurses();
	defer stopGoCurses();

	curses.Init_pair(1, curses.COLOR_RED, curses.COLOR_BLACK);

	loop(x, y);
}

func startGoCurses() {
	curses.Initscr();

	if curses.Stdwin == nil {
		stopGoCurses();
		os.Exit(1);
	}

	curses.Noecho();
	curses.Curs_set(curses.CURS_HIDE);
	curses.Stdwin.Keypad(true);

	if err := curses.Start_color(); err != nil {
		fmt.Printf("%s\n", err.String());
		stopGoCurses();
		os.Exit(1);
	}
}

func stopGoCurses() {
	curses.Endwin();
}

func loop(x, y int) {
	quit := false

	for {
		curses.Stdwin.Addstr(0, 0, "Hello,\nworld!", 0);
		curses.Stdwin.Addstr(3, 0, "use the cursor keys to move around", 0);
		curses.Stdwin.Addstr(4, 0, "press the 'q' key to quit", 0);
		inp := curses.Stdwin.Getch();
		quit = false

		switch inp {
		case 'q':
			quit = true
		case curses.KEY_LEFT:
			x = x - 1;
		case curses.KEY_RIGHT:
			x = x + 1;
		case curses.KEY_UP:
			y = y - 1;
		case curses.KEY_DOWN:
			y = y + 1;
		}

		if quit {
			break
		}

		curses.Stdwin.Clear();
		curses.Stdwin.Addch(y, x, '@', curses.Color_pair(1));
		curses.Stdwin.Refresh();
	}
}
