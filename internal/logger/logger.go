package logger

import (
  "fmt"
  "github.com/charmbracelet/log"
  "github.com/charmbracelet/lipgloss"
)

func StyleLogger() {
  log.ErrorLevelStyle = lipgloss.NewStyle().
    SetString("ERROR").
    Padding(0, 1, 0, 1).
    Background(lipgloss.AdaptiveColor{
        Light: "203",
        Dark:  "204",
    }).
    Foreground(lipgloss.Color("0"))
  
  // TODO: Style log fatal
  /*
  log.FatalLevelStyle = lipgloss.NewStyle().
    SetString("FATAL").
    Padding(0, 1, 0, 1).
    Background(lipgloss.AdaptiveColor{
        Light: "203",
        Dark:  "204",
    }).
    Foreground(lipgloss.Color("0"))
  */
}





// Syntax error
type SyntaxError struct {
	Line int
	Col  int
  Msg  string
}

func (e SyntaxError) Error() string {
	return fmt.Sprintf("%d:%d: %s", e.Line, e.Col, e.Msg)
}
