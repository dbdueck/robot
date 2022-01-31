package cmd

import (
	"robot/models"
	"runtime"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	section  string // Which of the 8 predefined sections to load (A-D,  1 or 2   )
	quadrant int    //One of the 4 quadrants in 16x16(I, II, III, IV), as well as 0,-1 in 8x8
)

var printQuadrantCmd = &cobra.Command{
	Use:   "printquadrant",
	Short: "Print A Ricocchet Quadrant (8x8)",
	Long: `
	
TODO: Rotation
-Example:
go run . printquadrant  -s D2  -q -1    #Loads "D2", and flips it (-1)
	
	`,
	Run: func(cmd *cobra.Command, args []string) {
		controllerPrintQuadrant()
	},
}

func init() {
	rootCmd.AddCommand(printQuadrantCmd)
	// Here you will define your flags and configuration settings.
	printQuadrantCmd.Flags().StringVarP(&section, "section", "s", "A1", "Section of the board with predefined homes")
	printQuadrantCmd.Flags().IntVarP(&quadrant, "quadrant", "q", 0, "Load into one of 6 quadrants (0=normal, -1=transposed) ")

}

func controllerPrintQuadrant() {
	log.Infof("Controller: print Quadrant, runtime=%v", runtime.GOOS)
	log.Debugf("Section=%v, Cartesian Quadrant=%v", section, quadrant)

	var quad = new(models.Quadrant)
	//sec.PrintQuadrant()
	//quad.InitAllRobots()
	//quad.LoadA1()
	//quad.LoadA2()
	quad.LoadMUX(section)
	vec := quad.LoadQuadrantVectorXform(quadrant)
	quad.XformQuadrantByVector(vec)
	quad.DebugQuadrant() //Might print something, is -v is enabled.
	quad.PrintQuadrant()
}
