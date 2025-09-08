package terminal

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func SetRawMode() {
	cmd := exec.Command("stty", "-echo", "-icanon", "min", "1", "time", "0")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func RestoreMode() {
	cmd := exec.Command("stty", "sane")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func HideCursor() {
	fmt.Printf("\x1b[?25l")
}

func ShowCursor() {
	fmt.Printf("\x1b[?25h")
}

func GetSize() (int, int, error) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	output, err := cmd.Output()
	if err != nil {
		return 80, 24, err
	}
	
	dimensions := strings.TrimSpace(string(output))
	parts := strings.Fields(dimensions)
	if len(parts) != 2 {
		return 80, 24, fmt.Errorf("unexpected stty output format")
	}
	
	height, err := strconv.Atoi(parts[0])
	if err != nil {
		return 80, 24, err
	}
	
	width, err := strconv.Atoi(parts[1])
	if err != nil {
		return 80, 24, err
	}
	
	return width, height, nil
}