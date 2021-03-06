package output

import (
	"bytes"
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/hoop33/limo/model"
	"github.com/tj/go-spin"
)

var spinner *spin.Spinner

// Color is a color text output
type Color struct {
}

// Inline displays text in line
func (c *Color) Inline(s string) {
	fmt.Print(color.GreenString(s))
}

// Info displays information
func (c *Color) Info(s string) {
	color.Green(s)
}

// Error displays an error
func (c *Color) Error(s string) {
	color.Red(s)
}

// Fatal displays an error and ends the program
func (c *Color) Fatal(s string) {
	c.Error(s)
	os.Exit(1)
}

// StarLine displays a star in one line
func (c *Color) StarLine(star *model.Star) {
	var buffer bytes.Buffer

	_, err := buffer.WriteString(color.BlueString(*star.FullName))
	if err != nil {
		c.Error(err.Error())
	}

	_, err = buffer.WriteString(color.YellowString(fmt.Sprintf(" ★ :%d", star.Stargazers)))
	if err != nil {
		c.Error(err.Error())
	}

	if star.Language != nil {
		_, err = buffer.WriteString(color.GreenString(fmt.Sprintf(" %s", *star.Language)))
		if err != nil {
			c.Error(err.Error())
		}
	}

	if star.URL != nil {
		_, err = buffer.WriteString(color.RedString(fmt.Sprintf(" %s", *star.URL)))
		if err != nil {
			c.Error(err.Error())
		}
	}

	fmt.Println(buffer.String())
}

// Star displays a star
func (c *Color) Star(star *model.Star) {
	c.StarLine(star)

	if len(star.Tags) > 0 {
		var buffer bytes.Buffer
		leader := ""
		for _, tag := range star.Tags {
			_, err := buffer.WriteString(color.MagentaString(fmt.Sprintf("%s%s", leader, tag.Name)))
			if err != nil {
				c.Error(err.Error())
			}
			leader = ", "
		}
		fmt.Println(buffer.String())
	}

	if star.Description != nil && *star.Description != "" {
		color.White(*star.Description)
	}

	if star.Homepage != nil && *star.Homepage != "" {
		color.Red(fmt.Sprintf("Home page: %s", *star.Homepage))
	}

	color.Green(fmt.Sprintf("Starred on %s", star.StarredAt.Format(time.UnixDate)))
}

// Tick displays evidence that the program is working
func (c *Color) Tick() {
	if spinner == nil {
		spinner = spin.New()
		spinner.Set(spin.Spin1)
	}
	fmt.Print(color.CyanString(fmt.Sprintf("\rUpdating . . . %s ", spinner.Next())))
}

func init() {
	registerOutput(&Color{})
}
