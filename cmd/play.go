package cmd

import (

	//	log "github.com/sirupsen/logrus"

	"fmt"
	"robot/models"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	boardSize            int
	flagSkipMiddleIsland bool
	parmBoardSections    string
	parmRobotCoords      string
	boardSectionsArry    = []string{"A1", "B1", "C1", "D1"}
	boardRobots          []models.Square
)

var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Ricochet Robots",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:`,
	Run: func(cmd *cobra.Command, args []string) {
		controllerPlay()
	},
}

func init() {
	rootCmd.AddCommand(playCmd)

	playCmd.Flags().BoolVarP(&flagSkipMiddleIsland, "skip-middle", "m", false, "Skip Middle Island")
	playCmd.Flags().IntVarP(&boardSize, "size", "N", 16, "Size of board, eg 4 for a 4x4")
	playCmd.Flags().StringVarP(&parmBoardSections, "sections", "s", "", "Sections (A1,B1,C1,D1")
	playCmd.Flags().StringVarP(&parmRobotCoords, "robots", "r", "", "Robots, Hex (Y,B,G,R,W")

}

//var board models.Board

func controllerPlay() {
	var board = new(models.Board)
	board.Size = boardSize
	board.SetupBoard()
	//fmt.Print(board.DebugBoard())
	if !flagSkipMiddleIsland { //Note: This is backwards. Wanted the skip/default to be yes
		board.SetupMiddleIsland()
	}

	if parmBoardSections != "" {

		boardSectionsArry = parseParmBoardSections(parmBoardSections)
		for idx, sect := range boardSectionsArry {
			var quad = new(models.Quadrant)
			//idx: starts at 0, but Quadrants start at "1"
			log.Infof("Sections/Quadrants: %v,%v", idx, sect)
			quad.LoadMUX(sect)
			vec := quad.LoadQuadrantVectorXform(idx + 1)
			quad.XformQuadrantByVector(vec)

			board.LoadQuadrant(quad)
		}
	}

	if parmRobotCoords != "" {
		board.LoadRobots(parmRobotCoords)
	}

	// fmt.Print(board.DebugBoard())
	log.Debug("Printing Board")

	//log.Debugf("Board1=%v\n", board)
	// log.Debugf("square1,1=%v\n", sq)

	board.PrintBoard()
	if flagCntVerbose >= 2 {
		fmt.Print(board.DebugBoard())
	}

}

func parseParmBoardSections(sections string) []string {
	log.Debugf("parseParmBoardSections, s=%s", sections)
	s := strings.Split(sections, ",")
	for idx, sect := range s {
		log.Tracef("%v,%v", idx, sect)
	}
	return s
}

