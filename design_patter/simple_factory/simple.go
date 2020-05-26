package simple_factory

import "fmt"

type Api interface {
	Say(name string) string
}

func NewApi() Api {
	return &HiApi{}
}

type HiApi struct{}

func (h *HiApi) Say(name string) string {
	return fmt.Sprintf("hi,my name is %s", name)
}
