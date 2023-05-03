package fakedb

import (
	"encoding/json"
	"fmt"
	"github.com/mateobr0/curso.go/generics/ejercicio2/customtypes"
)


var persons string = `[{
	"Name": "Mateo",
	"Age": 21
},{
	"Name": "Esteban",
	"Age": 31
}]`

var employees string = `[{
	"Name": "Santiago",
	"Position": "Front-end"
},{
	"Name": "Pablo",
	"Position": "Back-end"
}]`

var db map[string]string = map[string]string{
	"persons":   persons,
	"employees": employees,
}

func Read[T customtypes.SlicePersonOrEmployee](key string) T{
	var resultado T
	err := json.Unmarshal([]byte(db[key]), &resultado)
	if err != nil {
		fmt.Println(err)
	}
	return resultado
}