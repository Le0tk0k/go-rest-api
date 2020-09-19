package controllers

type Context interface {
	Param(string) string
	Bind(interface{}) error
	JSON(int, interface{}) error
}
