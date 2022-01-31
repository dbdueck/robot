package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	cfgFile        string
	flagCntVerbose int
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "robot",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.robot.yaml)")
	rootCmd.PersistentFlags().CountVarP(&flagCntVerbose, "verbose", "v", "Logging Verbose Level; -vv for even more")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".robot" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".robot")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
	initLogger()

}

func initLogger() {

	if flagCntVerbose >= 4 {
		log.SetLevel(log.TraceLevel)
		log.SetReportCaller(true) //Slower, but more helpful - 20-40%
	} else if flagCntVerbose >= 3 {
		log.SetLevel(log.DebugLevel)
		log.SetReportCaller(true) //Slower, but more helpful - 20-40%
	} else if flagCntVerbose >= 2 {
		log.SetLevel(log.TraceLevel)
	} else if flagCntVerbose >= 1 {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
	log.SetOutput(os.Stdout)

	//https://stackoverflow.com/questions/48971780/how-to-change-the-format-of-log-output-in-logrus
	//log.SetFormatter(&log.TextFormatter{})

	// log.SetFormatter(&easy.Formatter{
	// 	TimestampFormat: "2006-01-02 15:04:05",
	// 	LogFormat:       "[%lvl%]: %time% - %msg%\n",
	// })

	log.SetFormatter(&log.TextFormatter{
		TimestampFormat:        "2006-0102 15:04:05",
		DisableLevelTruncation: false, //Doesn't seem to work, though
		PadLevelText:           true,
	})

	log.Trace("Flag Cnt:=", flagCntVerbose)
	//logrus.TextFormatter.PadLevelText = true

}
