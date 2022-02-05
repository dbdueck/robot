package lib

import (
	log "github.com/sirupsen/logrus"
	//	log "github.com/sirupsen/logrus"
)

//Convert single  String character to int  -
//PARAM: Char 'F'  --> int: 15
func HexToDecInt(b byte) int {
	log.Tracef("Hex2... %v, %v", b, int(b))
	switch {
	case b >= '0' && b <= '9':
		return int(b - '0')
	case b >= 'a' && b <= 'f':
		return int(b - 'a' + 10)
	case b >= 'A' && b <= 'F':
		return int(b - 'A' + 10)
	}
	return 0
}

//Convert a string pair
//PARM - string, eg "2F"
//RETURNS - (2, 15)
func HexStringToPair(xy string) (int, int) {

	x := HexToDecInt(xy[0])
	y := HexToDecInt(xy[1])
	return x, y
}
