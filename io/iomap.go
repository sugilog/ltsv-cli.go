package io

import (
	coreio "io"
)

type IOMap struct {
	Out coreio.Writer
	Err coreio.Writer
	In  coreio.Reader
}
