package tasks

import (
	"coke/app"

	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(app.New())
}
