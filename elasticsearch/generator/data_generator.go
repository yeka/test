package generator

import (
	"math/rand"
)

type Data struct {
	ID         int
	Attributes map[string]string
}

var attrColor = []string{
	"Green",
	"Black",
	"Blue",
	"Yellow",
	"Red",
	"White",
	"Orange",
}

var attrSize = []string{
	"S",
	"M",
	"L",
	"XL",
}

func Generate(size int) []Data {
	rand.Seed(1)

	d := make([]Data, size)
	for i := 0; i < size; i++ {
		d[i] = Data{
			ID: i + 1,
			Attributes: map[string]string{
				"Color": attrColor[rand.Intn(len(attrColor))],
				"Size":  attrSize[rand.Intn(len(attrSize))],
			},
		}
	}
	return d
}
