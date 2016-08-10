package matrix

import (
	"fmt"
)

var DEBUG bool = false

type Matrix []float64
type Row []float64
type Builder []Row

/*
 * Provide `true` if you want errors to panic
 */
func SetDebug( debug bool ) {
	DEBUG = debug
}

func generateError( message string ) ( err error ) {
	if DEBUG {
		panic( message )
	} else {
		err = fmt.Errorf( message )
	}

	return
}
