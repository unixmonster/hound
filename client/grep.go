package client

import (
	"fmt"
	"os"
	"regexp"

	"github.com/unixmonster/hound/ansi"
	"github.com/unixmonster/hound/config"
)

type grepPresenter struct {
	f *os.File
}

func (p *grepPresenter) Present(
	re *regexp.Regexp,
	ctx int,
	repos map[string]*config.Repo,
	res *Response) error {

	c := ansi.NewFor(p.f)

	if _, err := fmt.Fprintf(p.f, "%s\n",
		c.Fg("// TODO(knorton): Implement", ansi.Yellow, ansi.Bold)); err != nil {
		return err
	}

	return nil
}

func NewGrepPresenter(w *os.File) Presenter {
	return &grepPresenter{w}
}
