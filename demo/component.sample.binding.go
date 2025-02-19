package main

import (
	"strconv"

	"github.com/yuriizinets/go-ssc"
)

type ComponentSampleBinding struct {
	FirstValue  string
	SecondValue string
	Result      string
}

func (c *ComponentSampleBinding) Init(p ssc.Page) {
	c.FirstValue = "5"
	c.SecondValue = "5"
	c.Result = "10"
}

func (c *ComponentSampleBinding) Actions() ssc.ActionsMap {
	return ssc.ActionsMap{
		"Calculate": func(args ...interface{}) {
			fv, err := strconv.Atoi(c.FirstValue)
			if err != nil {
				c.Result = "Can't calculate"
			}
			sv, err := strconv.Atoi(c.SecondValue)
			if err != nil {
				c.Result = "Can't calculate"
			}
			c.Result = strconv.Itoa(fv + sv)
		},
	}
}
