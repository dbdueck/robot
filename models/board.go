package models

import (
	"errors"
	"fmt"
	"robot/lib"
	"strings"

	log "github.com/sirupsen/logrus"
)

const MAXSIZE int = 17

//const SIZE int = 16 //16

type Board struct {
	Size    int
	Squares [MAXSIZE + 1][MAXSIZE + 1]*Square
}

func init() {
	log.Debug("Initializing Board")
	//Size = 8 //Set the default
}

func (b *Board) SetupBoard() *Board {

	log.Debugf("Setup Board, N:=%v", b.Size)

	//Make Each square "self-aware".  Remember - size is the actual size. But we over-size by 1
	//So:  for a SIZE=2x2, we make a board of size 3x3; Normal For loop:  0,2<SIZE
	for y := 0; y < b.Size+1; y++ {
		for x := 0; x < b.Size+1; x++ {
			b.Squares[x][y] = new(Square)
			sq := b.Squares[x][y]
			sq.X = x
			sq.Y = y
			sq.BoardSize = b.Size
			sq.Symbol = " " //fmt.Sprintf("%x", x)

			//log.Tracef("Setup, init=%v", sq.ToString())
		}
	}
	//Borders on Top and bottom of board
	for x := 0; x < b.Size; x++ {
		b.Squares[x][0].Top = true
		b.Squares[x][b.Size-1].Bottom = true
		log.Tracef("Setup, Top: %v", b.Squares[x][0].ToString())
		//b.Squares[y][SIZE].Top = true //If Bottom =true, then make the next row down true
	}
	//Log-Print Bottom
	for x := 0; x < b.Size; x++ {
		log.Tracef("Setup, Bottom: %v", b.Squares[x][b.Size-1].ToString())
	}

	// Borders on left and right side of board
	for y := 0; y < b.Size; y++ {
		b.Squares[0][y].Left = true
		b.Squares[b.Size-1][y].Right = true
		// b.Squares[y][SIZE-1].Left = true
		log.Tracef("Setup, Left: %v", b.Squares[0][y].ToString())
	}

	for y := 0; y < b.Size; y++ {
		log.Tracef("Setup, Right: %v", b.Squares[b.Size-1][y].ToString())
	}
	return b
}
func (b *Board) SetupMiddleIsland() {
	log.Debug("Setting Up Middle island")
	//Middle Island
	var N int = b.Size/2 - 1

	b.Squares[N][N].Top = true
	b.Squares[N+1][N].Top = true

	b.Squares[N][N+1].Bottom = true
	b.Squares[N+1][N+1].Bottom = true

	b.Squares[N][N].Left = true
	b.Squares[N][N+1].Left = true

	b.Squares[N+1][N].Right = true
	b.Squares[N+1][N+1].Right = true
}

//  convert all Rights to Left-1,  and Bottom to Top-1
func (b Board) Normalize() {
	log.Trace("Normalize Board")
	for y := 0; y < b.Size+1; y++ {
		for x := 0; x < b.Size+1; x++ {
			//log.Infof("x=%v, y=%v", x, y)
			sq := b.Squares[x][y]
			log.Tracef("Normalizing, %v", sq.ToString())
			if sq.Bottom { //set the top of one-higher
				b.Squares[x][y+1].Top = true
			}
			if sq.Right { //set the Left of one-to the left
				b.Squares[x+1][y].Left = true
			}

		}
	}

}
func (b Board) PrintBoard() {
	//N := SIZE + 1 //len(b.SquareArray)
	log.Tracef("Printing Board,local function, size=%v\n", b.Size)
	//Two levels:  The border (which is last), and the letter
	// for i, _ := range N {

	var (
		tbl  string
		rows string
		row0 string
		row1 string
		row2 string
		// rowdebug string
	)
	b.Normalize()
	//y is outside loop -- and x is inside loop
	for y := 0; y < b.Size; y++ {
		row0 = ""
		row1 = ""
		row2 = ""
		// rowdebug = ""
		for x := 0; x < b.Size; x++ {

			sq := b.Squares[x][y]
			log.Tracef("%v\n", sq.ToString())
			s0, s1, s2, _ := sq.GetPrintableSquare()
			row0 += s0
			row1 += s1
			row2 += s2
			//rowdebug += fmt.Sprintf("\t%s", sdebug)

		}
		//Then do the last one (right) on the row
		rows += fmt.Sprintf("%s\n%s\t%X\n", row0, row1, y /*rowdebug*/)

		if y == (b.Size-1) && row2 != "" { //Do thisonly  for the last row
			rows += fmt.Sprintf("%s\n", row2)
		}
	}
	tbl += rows
	//finally, a trace row
	rowFooter := ""
	for i := 0; i < b.Size; i++ {
		rowFooter += fmt.Sprintf(" %X", i)
	}
	rows += fmt.Sprintf("%s\n", rowFooter)

	fmt.Println(rows)
}
func (b Board) DebugBoard() string {
	var (
		rows     string
		rowdebug string
		sdebug   string
	)
	for y := 0; y < b.Size; y++ {
		rowdebug = ""
		for x := 0; x < b.Size; x++ {
			sq := b.Squares[x][y]
			sdebug = sq.ToStringBorder()
			rowdebug += fmt.Sprintf("\t%s", sdebug)
		}
		rows += fmt.Sprintf("%s\n", rowdebug)
	}
	return rows
}

//Load a section into the appropriate quadrant
//PARAM - section - the board layout
//PARAM - Which Quadrant (I, II, III, IV)
func (b *Board) LoadQuadrant(section *Quadrant) {

	//Quadrant is already placed into its correct position
	//Convert the section into it's correct position
	//section.transposeVector()

	for index, sq := range section.Items {
		log.Trace("At index", index, "sq is", sq)
		b.CopySquare(sq)
	}

}

//TODO: Move this to the square class
func (b *Board) CopySquare(aRobot *Square) {

	sq := b.Squares[aRobot.X][aRobot.Y]
	log.Tracef("Square=%v", sq.ToStringBorder())
	log.Tracef("Robot=%v\n", aRobot)
	//red = sec.RobotRed

	sq.RobotColorIdx = aRobot.RobotColorIdx
	sq.Top = sq.Top || aRobot.Top
	sq.Bottom = sq.Bottom || aRobot.Bottom
	sq.Left = sq.Left || aRobot.Left
	sq.Right = sq.Right || aRobot.Right
	sq.Symbol = aRobot.Symbol
	sq.ColorName = aRobot.ColorName
	sq.Typ = aRobot.Typ
	log.Tracef("SquareComplete=%v", sq.ToStringBorder())

}

//PARAM x,y = Coord's (0-15,0-15)
//PARAM colorIdx = enum for the color.  Y=0, R=1, G,B,W
func (b *Board) setRobot(x int, y int, colorIdx int) {
	sq := b.Squares[x][y]
	sq.Typ = EnumTypeRobot
	sq.Symbol = "!"
	sq.RobotColorIdx = colorIdx
	log.Infof("Setting Robot[%v,%v]=%v", x, y, colorIdx)
}

func (b *Board) LoadRobots(robotstring string) {
	log.Infof("LoadRobots, %v", robotstring)
	robotList := strings.Split(robotstring, ",")
	//Quadrant is already placed into its correct position
	//Convert the section into it's correct position
	//section.transposeVector()

	for index, xy := range robotList {
		//Xy="6F" as a string.
		x, y := lib.HexStringToPair(xy)
		log.Tracef("LoadRobots: At index=%v,xy=%s, x,y=%v,%v", index, xy, x, y)

		b.setRobot(x, y, index)
	}

}

// //Parse the List of Robots  eg. F3, 5E, G4, 12,13
// func parseListOfRobotCoordsAsString(robotstring string) []model.Squares {
// 	log.Infof("parseParmRobots, s=%s", robotstring)
// 	s := strings.Split(robotstring, ",")
// 	for idx, robot := range s {
// 		log.Infof("%v,%v", idx, robot)
// 	}
// 	return s
// }

//Solve the board,
//PARAM:  Solve for this robot:  "Red, Star"
func (b *Board) Solve(parmSolve string) {
	log.Infof("Board::Solve, %v", parmSolve)

	strArry := strings.SplitN(parmSolve, ",", 2)
	solveColor, solveShape := strArry[0], strArry[1]
	log.Infof("Board::Solve, %v, %v, %v", parmSolve, solveColor, solveShape)
	solveColorEnum := getRobotEnumFromColor(solveColor)
	//solveShape is a 1 char value
	//Remove this one eventually - it's not in the right place
	log.Infof("Solving for:  Color, %v, c=%v, Shape=%v", solveColorEnum, solveColor, solveShape)

	robotSq, err := b.FindSquareByTypeColorShape(EnumTypeHouse, solveColorEnum, solveShape)
	if err != nil {
		log.Fatalf("%v", err)

	} else {

		log.Infof("House we are looking for is at: %v,%v", robotSq.X, robotSq.Y)
	}
	//

}

//PARM: Color as a string. Only care about 1st char
//RETURN RobotIdxYellow = iota,	RobotIdxRed,	RobotIdxGreen,	RobotIdxBlue,	RobotIdxWhirl,	RobotIdxWhite
func getRobotEnumFromColor(c string) int {
	c0 := strings.ToUpper(string(c[0]))
	idx := strings.Index(EnumColorAbbrev, c0)
	log.Tracef("getEnum:: Abbrev=%v, c0=%s, idx=%v", EnumColorAbbrev, c0, idx)
	return idx
}

//Get the square with the Robot in it
//PARAM: Typ:  Enum
//PARAM: color: RobotColorIdx   (enum)
//PARAM: shape - "Sun, Moon, Stars, Planet"
func (b *Board) FindSquareByTypeColorShape(findTyp SquareType, findColor int, findShape string) (*Square, error) {
	log.Infof("Searching for shape: Typ=%v, Color=%v (%v), Shape=%v",
		findTyp, findColor, ColorNames[findColor], findShape)

	for y := 0; y < b.Size; y++ {
		for x := 0; x < b.Size; x++ {
			sq := b.Squares[x][y]
			if sq.Symbol != SymbolBlank {
				log.Infof("SQ:=T=%v, C=%v, S=%v, [%v]", sq.Typ, sq.RobotColorIdx, sq.Symbol, sq)
			}
			if sq.Typ == findTyp && sq.RobotColorIdx == findColor && sq.Symbol == findShape {
				return sq, nil
			}
		}
	}
	return nil, errors.New("shape not found, Color")
}
