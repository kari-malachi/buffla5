package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/kari-malachi/buffla5/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
