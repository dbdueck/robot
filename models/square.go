package models

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

const (
	//Double Border https://www.lookuptables.com/text/alt-codes
	TL = "╔"
	BL = "╚"
	TR = "╗"
	BR = "╝"

	//Single Walls
	T1 = "─"
	B1 = "─"
	L1 = "│"
	R1 = "│"
	//Double Walls
	T2 = "═"
	B2 = "═"
	L2 = "║" //186 Box Drawings Double Vertical
	R2 = "║"

	BoxDownDoubleHorizontalSingle = "╥" //210

	//Symbols - Sun, Moon, Star, SymbolPlanet
	SymbolPlanet = "P"
	SymbolSun    = "S"
	SymbolMoon   = "M"
	SymbolStar   = "R"
	SymbolWhirl  = "W"
	SymbolWall   = " "
	SymbolMiddle = "X"
	SymbolRobot  = "*"

	//1=bold, 4=underline, 5-blink, 7=inverse
	//https://gist.github.com/raghav4/48716264a0f426cf95e4342c21ada8e7
	ColorRed       = "\033[1;31m" // the "1;" = bold //31
	ColorRedInv    = "\033[7;91m" // the "1;" = bold
	ColorGreen     = "\033[1;32m"
	ColorGreenInv  = "\033[7;32m"
	ColorYellow    = "\033[1;33m"
	ColorYellowInv = "\033[7;33m"
	ColorBlue      = "\033[1;34m"
	ColorBlueInv   = "\033[7;34m"
	ColorPurple    = "\033[1;35m"
	ColorPurpleInv = "\033[7;35m"
	ColorCyan      = "\033[7;36m"
	ColorWhite     = "\033[37m"
	ColorBlack     = "\033[0m"
	ColorReset     = "\033[0m"
)

const ( //An Enum
	TypeEmpty = iota + 1 //because default=0
	TypeHouse
	TypeRobot
	TypeWall //This can later get overriden by a robot
	TypeMiddle
)

const ( //An Enum
	Rotate0 = iota //default=0 is OK
	RotateLeft
	RotateTwo
	RotateRight
)

const (
	RobotIdxYellow = iota
	RobotIdxRed
	RobotIdxGreen
	RobotIdxBlue
	RobotIdxWhirl
	RobotIdxWhite
)

//https://twin.sh/articles/35/how-to-add-colors-to-your-console-terminal-output-in-go
//
var ColorNames map[string]string = map[string]string{
	"yellow": ColorYellow,
	"red":    ColorRed,
	"green":  ColorGreen,
	"blue":   ColorBlue,
	"purple": ColorPurple,
	"white":  ColorWhite,

	"bgyellow": ColorYellowInv,
	"bgred":    ColorRedInv,
	"bggreen":  ColorGreenInv,
	"bgblue":   ColorBlueInv,
	"bgpurple": ColorPurpleInv,
}

var RobotColors map[int]string = map[int]string{
	RobotIdxYellow: ColorYellow,
	RobotIdxBlue:   ColorBlue,
	RobotIdxGreen:  ColorGreen,
	RobotIdxRed:    ColorRed,
	RobotIdxWhirl:  ColorPurple,
	RobotIdxWhite:  ColorWhite,
}

var RobotColorsInv map[int]string = map[int]string{
	RobotIdxYellow: ColorYellowInv,
	RobotIdxBlue:   ColorBlueInv,
	RobotIdxGreen:  ColorGreenInv,
	RobotIdxRed:    ColorRedInv,
	RobotIdxWhirl:  ColorPurpleInv,
	RobotIdxWhite:  ColorWhite,
}

//Structure of squares - robots, targets, center walls
type Square struct {
	X             int
	Y             int
	BoardSize     int
	RobotColorIdx int    //enum  Yellow=0, Red=1, Green, Blue, Whirl
	ColorName     string //or enum @deprecated
	Symbol        string //or enum  sun, moon, star, planet,  wall
	Typ           int    //enum
	//Walls
	Top    bool
	Bottom bool
	Left   bool
	Right  bool
}

// func init() {
// 	//Colors := map[string]{"yellow": colorYellow, "red": colorRed}
// }

//Get the 4 values for the square -- border .  Top and Left
//  a b   	/=
//  c d   	|val
//PARAM: Row1=border
//PARAM: Row2=the value
//PARAM: Row3= border, only for last row. usually blank (so don't print it)
func (sq Square) GetPrintableSquare() (string, string, string, string) {

	// 00	10	20
	// 01	11	21
	// 02	12	22

	//First Column
	c00 := T1 //"."
	c01 := L1 //"."
	c02 := "" //".\t" /
	//Middle Column
	c10 := T1 //"."
	c11 := SymbolWall

	c12 := "" //".\t"
	//Last column
	c20 := ""
	c21 := ""
	c22 := ""

	c11 = sq.getSymbolAndColor()

	if sq.Top {
		c00 = T2
		c10 = T2
	}
	if sq.Left {
		//c00 = BoxDownDoubleHorizontalSingle //╥
		c00 = L2
		c01 = L2
	}
	if sq.Left && sq.Top {
		c00 = TL
	}

	//3 possibilities: right border.. 1/mostly ignore these.
	if sq.Right { //Right border - only print this in the last column
		// c0 = R2                //Should check for top right corner
		// c12 = R2                //Override what would normally be a blank
		if sq.X == (sq.BoardSize - 1) { //2/Last one on the right, then we'll add the right side border
			if sq.Y == 0 { //Top Right
				c20 = TR
				c21 = R2
			} else if sq.Y == (sq.BoardSize - 1) { //Bottom Right
				//Bottom Right side -- add an extra row
				c20 = R2
				c21 = R2
				c22 = BR
			} else { //Middle Right
				c20 = R2
				c21 = R2
			}

		}

	}

	if sq.Bottom {
		if sq.Y == (sq.BoardSize - 1) { //OK. So this is the last row.So lets add extra stuff
			if sq.X == 0 { //Bottom Left
				c02 = BL
				c12 = B2
			} else if sq.X == (sq.BoardSize - 1) { //Bottom Right
				c02 = B2
				c12 = B2
				c22 = BR
			} else { //Bottom Middle
				c02 = B2
				c12 = B2
			}
		}
	}

	row0 := fmt.Sprintf("%s%s%s", c00, c10, c20)
	row1 := fmt.Sprintf("%s%s%s", c01, c11, c21)
	row2 := fmt.Sprintf("%s%s%s", c02, c12, c22)
	rowdebug := sq.ToStringBorder()
	return row0, row1, row2, rowdebug
}

func (sq Square) getSymbolAndColor() string {
	var c11 string = SymbolWall
	//Want to override a blank square
	if sq.Typ == TypeHouse {
		//c11 = fmt.Sprintf("%s%s%s", ColorNames["bg"+sq.ColorName], sq.Symbol, ColorReset)
		c11 = fmt.Sprintf("%s%s%s", RobotColorsInv[sq.RobotColorIdx], sq.Symbol, ColorReset)
		log.Infof("calcSymbolColor, Type=House, Symbol=%s, SQ=%v", sq.Symbol, sq)
	} else if sq.Typ == TypeRobot {
		c11 = fmt.Sprintf("%s%s%s", RobotColors[sq.RobotColorIdx], SymbolRobot, ColorReset)
		log.Infof("calcSymbolColor, Type=Robot, Symbol=%s, SQ=%v", sq.Symbol, sq)
	} else if sq.Typ == TypeMiddle {
		c11 = SymbolMiddle
	} else if sq.Typ == TypeWall {
		c11 = SymbolWall //Empty box, I guess, unless they want something else Say a Robot gets placed here later
	} else { // For null; or the possibly empty swirl (which isn't really defined.)
		c11 = SymbolWall
	}
	return c11
}

func (sq Square) ToString() string {
	return fmt.Sprintf("%v,%v, %v", sq.X, sq.Y, sq.ToStringBorder())
}

func (sq Square) ToStringBorder() string {
	return fmt.Sprintf("%v%v%v%v", bool01(sq.Left), bool01(sq.Top), bool01(sq.Bottom), bool01(sq.Right))
}

func bool01(b bool) string {
	if b {
		return "1"
	} else {
		return "0"
	}

}

//Rotate the borders +1, -1, or by 2 (+2 or -2)
func (sq *Square) RotateBorder(rotation int) {

	OLD := sq.ToStringBorder()
	switch rotation {

	case RotateRight: //	PI/2
		temp := sq.Top
		sq.Top = sq.Left
		sq.Left = sq.Bottom
		sq.Bottom = sq.Right
		sq.Right = temp
	case RotateTwo: //	PI
		temp := sq.Top
		sq.Top = sq.Bottom
		sq.Bottom = temp

		temp = sq.Left
		sq.Left = sq.Right
		sq.Right = temp

	case RotateLeft: //  3PI/2Same as +3
		temp := sq.Top
		sq.Top = sq.Right
		sq.Right = sq.Bottom
		sq.Bottom = sq.Left
		sq.Left = temp

	default:
		//Do Nothing

	}
	log.Tracef("RotateBorders:%v->%v", OLD, sq.ToStringBorder())

}
