package data

type ClientSettings interface {
	Name() (string, error)
}
