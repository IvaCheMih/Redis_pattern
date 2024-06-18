package models

type Good struct {
	Id         int    `json:"id"`
	SomeString string `json:"some_string"`
	SomeInt    int    `json:"some_int"`
	SomeBool   bool   `json:"some_bool"`
}
