package logs

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

var (
	DisableLineBack = false
	DisableBrush    = false
)

// SetDisableColors
func SetDisableColors(disableBrush, disableLineBack bool) {
	DisableLineBack = disableLineBack
	DisableBrush = disableBrush
}

type Brush func(format string, a ...interface{}) string

const (
	pre   = "\033["
	reset = "\033[0m"
)

func NewBrush(color string) Brush {
	return func(format string, a ...interface{}) string {
		text := format
		if len(a) > 0 {
			text = fmt.Sprintf(text, a...)
		}

		if DisableBrush {
			return text
		}
		return pre + color + "m" + text + reset
	}
}

func DefaultBrush(format string, a ...interface{}) string {
	text := format
	if len(a) > 0 {
		text = fmt.Sprintf(text, a...)
	}
	return text
}

var (
	BrushBlue    = NewBrush("1;34")
	BrushCyan    = NewBrush("1;36")
	BrushGreen   = NewBrush("1;32")
	BrushYellow  = NewBrush("1;33")
	BrushRed     = NewBrush("1;31")
	BrushMagenta = NewBrush("1;35")
)
var colors = []Brush{
	BrushBlue,    // Trace    blue
	BrushBlue,    // Debug    blue
	BrushCyan,    // Info     cyan
	BrushGreen,   // Notice   green
	BrushYellow,  // Warn     yellow
	BrushRed,     // Error    red
	BrushMagenta, // Fatal    magenta
}

type ConsoleProvider struct {
	level int
	color bool
}

func NewConsoleProvider() *ConsoleProvider {
	p := &ConsoleProvider{}
	p.color = runtime.GOOS != "windows" && IsTerminal(int(os.Stdout.Fd()))
	return p
}

func (cp *ConsoleProvider) Init() error {
	return nil
}

func (cp *ConsoleProvider) SetLevel(l int) {
	cp.level = l
}

func (cp *ConsoleProvider) WriteMsg(msg string, level int) error {
	if level < cp.level {
		return nil
	}
	if !DisableLineBack && cp.color {
		if strings.Index(msg, pre) != -1 {
			fmt.Fprint(os.Stdout, msg)
		} else {
			fmt.Fprint(os.Stdout, colors[level](msg))
		}
	} else {
		fmt.Fprint(os.Stdout, msg)
	}
	return nil
}

func (cp *ConsoleProvider) Flush() error {
	return nil
}

func (cp *ConsoleProvider) Destroy() error {
	return nil
}
