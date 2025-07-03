package models

type Persona struct {
	Id       int    `json:"id"`
	Nome     string `json:"nome"`
	Historia string `json:"historia"`
}

var Personas []Persona
