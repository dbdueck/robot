package models

import (
	log "github.com/sirupsen/logrus"
)

/*
II   |  I
----------
III |  IV

*/
var (
	//Transform from computer format (0,0 = top left) -> to
	Xform0 = [3][3]int{{0, 1, 0}, {0, 0, 1}, {Rotate0, 0, 0}}   // Do Nothing
	Xform1 = [3][3]int{{0, 0, 1}, {0, 1, 0}, {RotateTwo, 0, 0}} //Transpose To Quadrant I

	//                    x           y            borders
	XformQ1 = [3][3]int{{8, 0, 1}, {7, -1, 0}, {RotateLeft, 0, 0}} //OK
	//These still need to be calculated
	XformQ2 = [3][3]int{{7, -1, 0}, {7, 0, -1}, {RotateTwo, 0, 0}}
	XformQ3 = [3][3]int{{7, 0, -1}, {8, 1, 0}, {RotateRight, 0, 0}} //OK
	XformQ4 = [3][3]int{{8, 1, 0}, {8, 0, 1}, {Rotate0, 0, 0}}      //OK
	//Q2:= [...][...]int{{1,2,3},{4,5,6},}
	//Borders: Rotate by 0, 1, 2,
)

type Quadrant struct {
	RobotRed        Square
	RobotYellow     Square
	RobotBlue       Square
	RobotGreen      Square
	RobotSwirl      Square
	FenceHorz       Square
	FenceVert       Square
	Middle          Square
	Items           []*Square //This is a pointer to the list
	transposeVector [3][3]int
}

//TODO: Does this one run?  Only when this gets loaded
func (quad *Quadrant) init() {
	//This only needs to be set up once
	quad.Middle = Square{X: 0, Y: 0, Symbol: SymbolMiddle, ColorName: "white", Typ: TypeMiddle, Left: true, Right: true, Top: true, Bottom: true}
	//This needs to get set each time you create a new one
	quad.Items = []*Square{
		&quad.RobotRed, &quad.RobotGreen,
	}
	log.Tracef("init:AllRobots, %v", quad.Items)
}

func (quad *Quadrant) InitAllRobots() {
	log.Trace("InitAllRobots")
	//Always add the "Middle" to each block.
	//Should check about adding the Whirl -- don't want to add it each time
	quad.Items = []*Square{
		&quad.RobotYellow, &quad.RobotBlue, &quad.RobotRed, &quad.RobotGreen, &quad.RobotSwirl,
		&quad.FenceHorz, &quad.FenceVert, &quad.Middle,
	}
}

func (quad *Quadrant) LoadA1() {
	log.Trace("LoadA1#")
	quad.RobotYellow = Square{X: 0, Y: 2, Symbol: SymbolPlanet, RobotColorIdx: RobotIdxYellow, ColorName: "yellow", Typ: TypeHouse, Right: true, Bottom: true}
	quad.RobotBlue = Square{X: 5, Y: 3, Symbol: SymbolStar, RobotColorIdx: RobotIdxBlue, ColorName: "blue", Typ: TypeHouse, Left: true, Top: true}
	quad.RobotGreen = Square{X: 1, Y: 5, Symbol: SymbolMoon, RobotColorIdx: RobotIdxGreen, ColorName: "green", Typ: TypeHouse, Left: true, Bottom: true}
	quad.RobotRed = Square{X: 6, Y: 6, Symbol: SymbolSun, RobotColorIdx: RobotIdxRed, ColorName: "red", Typ: TypeHouse, Top: true, Right: true}
	quad.FenceHorz = Square{X: 7, Y: 2, Top: true, Typ: TypeWall, Symbol: SymbolWall}
	quad.FenceVert = Square{X: 4, Y: 7, Left: true, Typ: TypeWall, Symbol: SymbolWall}
	// log.Debugf("LoadA1: Red=%v", quad.RobotRed)
	// log.Debugf("LoadA1: Green=%v", quad.RobotGreen)
	quad.InitAllRobots() //This even inits an empty swirl
	log.Tracef("All Robots, %v", quad.Items)
}

func (quad *Quadrant) LoadA2() {
	log.Trace("LoadA2#")
	quad.RobotYellow = Square{X: 6, Y: 1, Symbol: SymbolPlanet, RobotColorIdx: RobotIdxYellow, ColorName: "yellow", Typ: TypeHouse, Right: true, Bottom: true}
	quad.RobotBlue = Square{X: 2, Y: 5, Symbol: SymbolStar, RobotColorIdx: RobotIdxBlue, ColorName: "blue", Typ: TypeHouse, Left: true, Top: true}
	quad.RobotGreen = Square{X: 5, Y: 3, Symbol: SymbolMoon, RobotColorIdx: RobotIdxGreen, ColorName: "green", Typ: TypeHouse, Left: true, Bottom: true}
	quad.RobotRed = Square{X: 0, Y: 2, Symbol: SymbolSun, RobotColorIdx: RobotIdxRed, ColorName: "red", Typ: TypeHouse, Top: true, Right: true}
	quad.FenceHorz = Square{X: 7, Y: 3, Top: true, Typ: TypeWall, Symbol: SymbolWall}
	quad.FenceVert = Square{X: 4, Y: 7, Left: true, Typ: TypeWall, Symbol: SymbolWall}
	quad.InitAllRobots() //This even inits an empty swirl
	log.Tracef("All Robots, %v", quad.Items)
}

func (quad *Quadrant) LoadB1() {
	log.Trace("LoadB1#")
	quad.RobotYellow = Square{X: 1, Y: 3, Symbol: SymbolSun, RobotColorIdx: RobotIdxYellow, ColorName: "yellow", Typ: TypeHouse, Right: true, Bottom: true}
	quad.RobotBlue = Square{X: 2, Y: 1, Symbol: SymbolMoon, RobotColorIdx: RobotIdxBlue, ColorName: "blue", Typ: TypeHouse, Left: true, Bottom: true}
	quad.RobotGreen = Square{X: 6, Y: 4, Symbol: SymbolPlanet, RobotColorIdx: RobotIdxGreen, ColorName: "green", Typ: TypeHouse, Left: true, Bottom: true}
	quad.RobotRed = Square{X: 5, Y: 6, Symbol: SymbolStar, RobotColorIdx: RobotIdxRed, ColorName: "red", Typ: TypeHouse, Top: true, Right: true}
	quad.RobotSwirl = Square{X: 4, Y: 0, Symbol: SymbolWhirl, RobotColorIdx: RobotIdxWhirl, ColorName: "purple", Typ: TypeHouse, Top: true, Left: true}

	quad.FenceHorz = Square{X: 7, Y: 3, Top: true, Typ: TypeWall, Symbol: SymbolWall}
	quad.FenceVert = Square{X: 2, Y: 7, Left: true, Typ: TypeWall, Symbol: SymbolWall}
	// log.Debugf("LoadA1: Red=%v", quad.RobotRed)
	// log.Debugf("LoadA1: Green=%v", quad.RobotGreen)
	quad.InitAllRobots() //This even inits an empty swirl
	log.Tracef("All Robots, %v", quad.Items)
}

func (quad *Quadrant) LoadB2() {
	log.Trace("LoadB2#")
	quad.RobotYellow = Square{X: 6, Y: 4, Symbol: SymbolSun, RobotColorIdx: RobotIdxYellow, ColorName: "yellow", Typ: TypeHouse, Right: true, Bottom: true}
	quad.RobotBlue = Square{X: 1, Y: 6, Symbol: SymbolMoon, RobotColorIdx: RobotIdxBlue, ColorName: "blue", Typ: TypeHouse, Left: true, Top: true}
	quad.RobotGreen = Square{X: 2, Y: 3, Symbol: SymbolPlanet, RobotColorIdx: RobotIdxGreen, ColorName: "green", Typ: TypeHouse, Left: true, Bottom: true}
	quad.RobotRed = Square{X: 5, Y: 2, Symbol: SymbolStar, RobotColorIdx: RobotIdxRed, ColorName: "red", Typ: TypeHouse, Top: true, Right: true}
	quad.RobotSwirl = Square{X: 0, Y: 2, Symbol: SymbolWhirl, RobotColorIdx: RobotIdxWhirl, ColorName: "purple", Typ: TypeHouse, Top: true, Left: true}

	quad.FenceHorz = Square{X: 7, Y: 1, Top: true, Typ: TypeWall, Symbol: SymbolWall}
	quad.FenceVert = Square{X: 4, Y: 7, Left: true, Typ: TypeWall, Symbol: SymbolWall}
	quad.InitAllRobots() //This even inits an empty swirl
	log.Tracef("All Robots, %v", quad.Items)
}

func (quad *Quadrant) LoadC1() {
	log.Trace("LoadC1#")
	quad.RobotYellow = Square{X: 1, Y: 4, Symbol: SymbolStar, RobotColorIdx: RobotIdxYellow, ColorName: "yellow", Typ: TypeHouse, Left: true, Top: true}
	quad.RobotBlue = Square{X: 4, Y: 1, Symbol: SymbolPlanet, RobotColorIdx: RobotIdxBlue, ColorName: "blue", Typ: TypeHouse, Right: true, Top: true}
	quad.RobotGreen = Square{X: 6, Y: 5, Symbol: SymbolSun, RobotColorIdx: RobotIdxGreen, ColorName: "green", Typ: TypeHouse, Left: true, Bottom: true}
	quad.RobotRed = Square{X: 3, Y: 6, Symbol: SymbolMoon, RobotColorIdx: RobotIdxRed, ColorName: "red", Typ: TypeHouse, Bottom: true, Right: true}
	quad.FenceHorz = Square{X: 7, Y: 2, Top: true, Typ: TypeWall, Symbol: SymbolWall}
	quad.FenceVert = Square{X: 6, Y: 7, Left: true, Typ: TypeWall, Symbol: SymbolWall}
	// log.Debugf("LoadA1: Red=%v", quad.RobotRed)
	// log.Debugf("LoadA1: Green=%v", quad.RobotGreen)
	quad.InitAllRobots() //This even inits an empty swirl
	log.Tracef("All Robots, %v", quad.Items)
}

func (quad *Quadrant) LoadC2() {
	log.Trace("LoadC2#")
	quad.RobotYellow = Square{X: 1, Y: 6, Symbol: SymbolStar, RobotColorIdx: RobotIdxYellow, ColorName: "yellow", Typ: TypeHouse, Left: true, Top: true}
	quad.RobotBlue = Square{X: 1, Y: 2, Symbol: SymbolPlanet, RobotColorIdx: RobotIdxBlue, ColorName: "blue", Typ: TypeHouse, Left: true, Bottom: true}
	quad.RobotGreen = Square{X: 6, Y: 5, Symbol: SymbolSun, RobotColorIdx: RobotIdxGreen, ColorName: "green", Typ: TypeHouse, Right: true, Bottom: true}
	quad.RobotRed = Square{X: 4, Y: 1, Symbol: SymbolMoon, RobotColorIdx: RobotIdxRed, ColorName: "red", Typ: TypeHouse, Top: true, Right: true}
	quad.FenceHorz = Square{X: 7, Y: 2, Top: true, Typ: TypeWall, Symbol: SymbolWall}
	quad.FenceVert = Square{X: 3, Y: 7, Left: true, Typ: TypeWall, Symbol: SymbolWall}
	quad.InitAllRobots() //This even inits an empty swirl
	log.Tracef("All Robots, %v", quad.Items)
}

//-------------------------------------------------------------------------------------
func (quad *Quadrant) LoadD1() {
	log.Trace("LoadD1#")
	quad.RobotYellow = Square{X: 1, Y: 3, Symbol: SymbolMoon, RobotColorIdx: RobotIdxYellow, ColorName: "yellow", Typ: TypeHouse, Right: true, Bottom: true}
	quad.RobotBlue = Square{X: 5, Y: 1, Symbol: SymbolSun, RobotColorIdx: RobotIdxBlue, ColorName: "blue", Typ: TypeHouse, Left: true, Bottom: true}
	quad.RobotGreen = Square{X: 2, Y: 6, Symbol: SymbolStar, RobotColorIdx: RobotIdxGreen, ColorName: "green", Typ: TypeHouse, Left: true, Top: true}
	quad.RobotRed = Square{X: 6, Y: 5, Symbol: SymbolPlanet, RobotColorIdx: RobotIdxRed, ColorName: "red", Typ: TypeHouse, Top: true, Right: true}
	quad.FenceHorz = Square{X: 7, Y: 4, Top: true, Typ: TypeWall, Symbol: SymbolWall}
	quad.FenceVert = Square{X: 4, Y: 7, Left: true, Typ: TypeWall, Symbol: SymbolWall}
	// log.Debugf("LoadA1: Red=%v", quad.RobotRed)
	// log.Debugf("LoadA1: Green=%v", quad.RobotGreen)
	quad.InitAllRobots() //This even inits an empty swirl
	log.Tracef("All Robots, %v", quad.Items)
}

func (quad *Quadrant) LoadD2() {
	log.Trace("LoadD2#")
	quad.RobotYellow = Square{X: 5, Y: 6, Symbol: SymbolMoon, RobotColorIdx: RobotIdxYellow, ColorName: "yellow", Typ: TypeHouse, Right: true, Bottom: true}
	quad.RobotBlue = Square{X: 1, Y: 4, Symbol: SymbolSun, RobotColorIdx: RobotIdxBlue, ColorName: "blue", Typ: TypeHouse, Right: true, Top: true}
	quad.RobotGreen = Square{X: 6, Y: 1, Symbol: SymbolStar, RobotColorIdx: RobotIdxGreen, ColorName: "green", Typ: TypeHouse, Left: true, Top: true}
	quad.RobotRed = Square{X: 3, Y: 2, Symbol: SymbolPlanet, RobotColorIdx: RobotIdxRed, ColorName: "red", Typ: TypeHouse, Bottom: true, Left: true}
	quad.FenceHorz = Square{X: 7, Y: 3, Top: true, Typ: TypeWall, Symbol: SymbolWall}
	quad.FenceVert = Square{X: 3, Y: 7, Left: true, Typ: TypeWall, Symbol: SymbolWall}
	quad.InitAllRobots() //This even inits an empty swirl
	log.Tracef("All Robots, %v", quad.Items)
}

func (quad *Quadrant) LoadZ0() {
	log.Trace("LoadZ0#")
	quad.RobotYellow = Square{X: 0, Y: 0, Symbol: SymbolMoon, RobotColorIdx: RobotIdxYellow, ColorName: "yellow", Typ: TypeHouse, Right: true, Bottom: true}
	quad.RobotBlue = Square{X: 0, Y: 0, Symbol: SymbolSun, RobotColorIdx: RobotIdxBlue, ColorName: "blue", Typ: TypeHouse, Right: true, Top: true}
	quad.RobotGreen = Square{X: 0, Y: 0, Symbol: SymbolStar, RobotColorIdx: RobotIdxGreen, ColorName: "green", Typ: TypeHouse, Left: true, Top: true}
	quad.RobotRed = Square{X: 0, Y: 0, Symbol: SymbolPlanet, RobotColorIdx: RobotIdxRed, ColorName: "red", Typ: TypeHouse, Bottom: true, Left: true}
	quad.FenceHorz = Square{X: 3, Y: 3, Top: true, Typ: TypeWall, Symbol: SymbolWall}
	quad.FenceVert = Square{X: 3, Y: 7, Left: true, Typ: TypeWall, Symbol: SymbolWall}
	quad.InitAllRobots() //This even inits an empty swirl
	log.Tracef("All Robots, %v", quad.Items)
}

func (quad *Quadrant) LoadZ1() {
	log.Trace("LoadZ1#")
	quad.RobotYellow = Square{X: 2, Y: 6, Symbol: SymbolMoon, RobotColorIdx: RobotIdxYellow, ColorName: "yellow", Typ: TypeHouse, Right: true, Bottom: true}
	quad.RobotBlue = Square{X: 0, Y: 7, Symbol: SymbolSun, RobotColorIdx: RobotIdxBlue, ColorName: "blue", Typ: TypeHouse, Right: true, Top: true}
	quad.RobotGreen = Square{X: 0, Y: 0, Symbol: SymbolStar, RobotColorIdx: RobotIdxGreen, ColorName: "green", Typ: TypeHouse, Left: true, Top: true}
	quad.RobotRed = Square{X: 0, Y: 0, Symbol: SymbolPlanet, RobotColorIdx: RobotIdxRed, ColorName: "red", Typ: TypeHouse, Bottom: true, Left: true}
	quad.FenceHorz = Square{X: 7, Y: 3, Top: true, Typ: TypeWall, Symbol: SymbolWall}
	quad.FenceVert = Square{X: 3, Y: 7, Left: true, Typ: TypeWall, Symbol: SymbolWall}
	quad.InitAllRobots() //This even inits an empty swirl
	log.Tracef("All Robots, %v", quad.Items)
}

//--------------------------------------------------------------------------------------

//This just prints the quadrant in its literal form ComputerRQ4 quadrant
func (quad *Quadrant) PrintQuadrant() {

	log.Debug("PrintQuadrant#")
	var board8x8 = Board{Size: 8}
	board8x8.SetupBoard()

	for index, aRobot := range quad.Items {
		sq := board8x8.Squares[aRobot.X][aRobot.Y]
		log.Tracef("Square=%v", sq.ToStringBorder())
		log.Tracef("Robot[%v]=%v\n", index, aRobot)
		//red = sec.RobotRed

		sq.Top = sq.Top || aRobot.Top
		sq.Bottom = sq.Bottom || aRobot.Bottom
		sq.Left = sq.Left || aRobot.Left
		sq.Right = sq.Right || aRobot.Right
		sq.Symbol = aRobot.Symbol
		sq.ColorName = aRobot.ColorName
		sq.Typ = aRobot.Typ
		log.Tracef("SquareComplete=%v", sq.ToStringBorder())
	}
	board8x8.PrintBoard()
}

func (quad *Quadrant) DebugQuadrant() {
	log.Trace("DebugQuadrant#")

	for index, element := range quad.Items {
		log.Tracef("Robot[%v]=%v\n", index, element)
	}

}

//Load a particular Gameboard/QuarterSection, given as A1, A2, B1, B2, .. D2
//CMD LINE --section or -s
//TODO: Make Data Driven
func (quad *Quadrant) LoadMUX(section string) {

	switch section {
	case "A1":
		quad.LoadA1()
	case "A2":
		quad.LoadA2()
	case "B1":
		quad.LoadB1()
	case "B2":
		quad.LoadB2()
	case "C1":
		quad.LoadC1()
	case "C2":
		quad.LoadC2()
	case "D1":
		quad.LoadD1()
	case "D2":
		quad.LoadD2()
	case "Z0":
		quad.LoadZ0()
	case "Z1":
		quad.LoadZ1()

	default:
		log.Debugf("Invalid Section, %v", section)
	}
	log.Debugf("LoadMUX, section=%v", section)
}

//Load/Determine the transposition vector to decide which quadrant this goes in.
//RETURNS: Transposing Vector
func (quad *Quadrant) LoadQuadrantVectorXform(cartesianquadrant int) [3][3]int {
	//var transposeVector [2][3]int
	log.Debugf("LoadQuadrantVectorXform# Transpose section into one of 5 quadrants, q=%v", cartesianquadrant)
	switch cartesianquadrant {
	case -1: //Simple Transpose, starting at 0,0 TopLeft
		// log.Debug("Transpose ..")
		quad.transposeVector = Xform1
	case 0: //Simple Transpose, starting at 0,0 TopLeft
		// log.Debug("Transpose ..")
		quad.transposeVector = Xform0
	case 1: //Quadrant 1
		quad.transposeVector = XformQ1
	case 2:
		quad.transposeVector = XformQ2
	case 3:
		quad.transposeVector = XformQ3
	case 4:
		quad.transposeVector = XformQ4

	default:
		quad.transposeVector = Xform1
	}
	log.Debugf("Transpose, Quadrant=%v, Vector=%v", cartesianquadrant, quad.transposeVector)
	return quad.transposeVector
}

//Do the Xform
//USES: The transposeVector, as determined elsewhere
//USES: The 4or5 robot houses, edges, etc
func (quad *Quadrant) XformQuadrantByVector(transposeVector [3][3]int) {
	log.Debugf("XformQuadrantByVector# %v", transposeVector)
	vec_x := transposeVector[0]
	xpos := vec_x[0]
	xdx := vec_x[1]
	xdy := vec_x[2]

	vec_y := transposeVector[1]
	ypos := vec_y[0]
	ydx := vec_y[1]
	ydy := vec_y[2]

	vecRotate := transposeVector[2][0] //Only 1 item in the last rotation item
	for index, item := range quad.Items {
		//sq := board8x8.Squares[aRobot.X][aRobot.Y]
		//log.Tracef("Square=%v", sq.ToString2())
		OLD := item.ToString()
		// log.Debugf("Robot[%v]=%v\n", index, item.ToString())
		//red = sec.RobotRed
		//[offset   x   y   ]
		//X= offset
		x0 := item.X
		y0 := item.Y
		item.X = (xpos) + (x0 * xdx) + (y0 * xdy)
		item.Y = (ypos) + (x0 * ydx) + (y0 * ydy)
		log.Tracef("MathX[%v]: %v + (%v * %v) + (%v * %v ) = %v + %v + %v = %v",
			x0, xpos, x0, xdx, y0, xdy, (xpos), x0*xdx, y0*xdy, item.X)
		log.Tracef("MathY[%v]: %v + (%v * %v) + (%v * %v ) = %v + %v + %v = %v",
			y0, ypos, x0, ydx, y0, ydy, (xpos), x0*ydx, y0*ydy, item.Y)
		item.RotateBorder(vecRotate)
		// sq.Top = sq.Top || item.Top
		// sq.Bottom = sq.Bottom || item.Bottom
		// sq.Left = sq.Left || item.Left
		// sq.Right = sq.Right || item.Right
		// sq.Symbol = item.Symbol
		// sq.ColorName = item.ColorName
		// sq.Typ = item.Typ
		log.Tracef("RotateBot[%v] OLD=%v, NEW=%v ", index, OLD, item.ToString())
	}

}
