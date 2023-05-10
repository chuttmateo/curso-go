package main

import (
	"errors"
	"fmt"
	"log"
)

type raizError struct {
	lat  string
	long string
	err  error
}

func (re raizError) Error() string {
	return fmt.Sprintf("error matemático: %v %v %v", re.lat, re.long, re.err)
}

func main() {

	_, err := raiz(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func raiz(f float64) (float64, error) {
	re := raizError{
		lat: "50.2289 N",
		long: "99.4656 W",
		err: errors.New("error my friend"),
	}
	if f < 0 {
		return 0, re
		// Escribe tu código aquí
	}
	return 42, nil
}
