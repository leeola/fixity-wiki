package fiki

import "github.com/leeola/fixity"

type Meta struct {
	Title string   `json:"title,omitempty"`
	Tags  []string `json:"tags,omitempty"`
}

type Body struct {
	Body string `json:"body,omitempty"`
	Copy string `json:"copy,omitempty"`
}

type Fiki struct {
	fixi fixity.Fixity
}

func New(fixi fixity.Fixity) *Fiki {
	return &Fiki{
		fixi: fixi,
	}
}
