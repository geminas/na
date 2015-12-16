package controllers

import "github.com/revel/revel"

type About struct {
	*revel.Controller
}

func (c About) Events() revel.Result {
	return c.Render()
}

func (c About) Company() revel.Result {
	return c.Render()
}
func (c About) Team() revel.Result {
	return c.Render()
}
func (c About) TeamSlide() revel.Result {
	return c.Render()
}
