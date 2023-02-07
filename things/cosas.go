package things

//go:generate go run golang.org/x/tools/cmd/stringer -type=Cosa -trimprefix=Cosa
type Cosa int

const (
	CosaSilla Cosa = iota
	CosaMesa
)
