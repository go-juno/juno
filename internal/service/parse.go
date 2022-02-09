package service

type Filed struct {
	Name  string
	Type  string
	Field []*Filed
}

type Package struct {
	Name string
}
