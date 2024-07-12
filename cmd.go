package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

type RuntimeArguments struct {
	EnableService        bool
	EnableCommittee      bool
	EnableStateRootCache bool
	EnableTest           bool
	EnablePprof          bool

	TestBlockHeightLimit uint
	StateRootCacheFreq   uint
	StateRootCacheNumber uint

	ConfigFilePath       string
	CommitteeIndexerName string
	CommitteeIndexerURL  string
	ProtocolName         string
	MetricAddr           string
}

func NewRuntimeArguments() *RuntimeArguments {
	return &RuntimeArguments{}
}

func (arguments *RuntimeArguments) MakeCmd() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "modular-indexer-committee",
		Short: "Activates the Nubit Committee Indexer with optional services.",
		Long: `Committee Indexer is an essential component of the Nubit Modular Indexer architecture.
This command offers multiple flags to tailor the indexer's functionality according to the user's needs.
The indexer operates on a fully user-verified execution layer for meta-protocols on Bitcoin,
leveraging Bitcoin's immutable and decentralized nature to provide a Turing-complete execution layer.
		`,
		Version: fmt.Sprintf("%v (%v)", version, gitHash),
		Run: func(cmd *cobra.Command, args []string) {
			if arguments.EnableService {
				log.Println("Service mode is enabled")
			} else {
				log.Println("Service mode is disabled")
			}
			if arguments.EnableCommittee {
				log.Println("Committee mode is enabled")
			} else {
				log.Println("Committee mode is disabled")
			}
			if arguments.EnableStateRootCache {
				log.Println("StateRoot cache is enabled")
				log.Printf("The cache will be stored per %d bitcoin blocks\n", arguments.StateRootCacheFreq)
				log.Printf("%d cache files will be stored\n", arguments.StateRootCacheNumber)
			} else {
				log.Println("StateRoot cache is disabled")
			}
			if arguments.EnableTest && arguments.TestBlockHeightLimit != 0 {
				log.Printf("Use the test mode and limit the max blockheight %d to avoid catching up to the real latest block\n", arguments.TestBlockHeightLimit)
			}

			log.Printf("The path of the config file is %s\n", arguments.ConfigFilePath)
			log.Printf("The name of the committee indexer service is %s\n", arguments.CommitteeIndexerName)
			log.Printf("The url of the committee indexer service is %s\n", arguments.CommitteeIndexerURL)
			log.Printf("The meta protocol chosen is %s\n", arguments.ProtocolName)
			log.Println("Metrics listen at:", arguments.MetricAddr)

			Execution(arguments)
		},
	}

	rootCmd.Flags().BoolVarP(&arguments.EnableService, "service", "s", false, "Enable this flag to provide API service")
	rootCmd.Flags().BoolVar(&arguments.EnableCommittee, "committee", false, "Enable this flag to provide committee service by uploading checkpoint")
	rootCmd.Flags().BoolVar(&arguments.EnableStateRootCache, "cache", true, "Enable this flag to cache State Root")
	rootCmd.Flags().BoolVarP(&arguments.EnableTest, "test", "t", false, "Enable this flag to hijack the blockheight to test the service")
	rootCmd.Flags().BoolVar(&arguments.EnablePprof, "pprof", false, "Enable the pprof HTTP handler (at `/debug/pprof/`)")

	rootCmd.Flags().UintVar(&arguments.TestBlockHeightLimit, "blockheight", 0, "When -test enabled, you can set TestBlockHeightLimit as a fixed value you want")
	rootCmd.Flags().UintVar(&arguments.StateRootCacheFreq, "cachefreq", 1000, "When -cache enabled, you can set StateRootCacheFreq as a fixed value you want")
	rootCmd.Flags().UintVar(&arguments.StateRootCacheNumber, "cachenumber", 2, "When -cache enabled, you can set StateRootCacheNumber as a fixed value you want")

	rootCmd.Flags().StringVar(&arguments.ConfigFilePath, "cfg", "config.json", "Indicate the path of config file")
	rootCmd.Flags().StringVarP(&arguments.CommitteeIndexerName, "name", "n", "", "Indicate the name of the committee indexer service")
	rootCmd.Flags().StringVarP(&arguments.CommitteeIndexerURL, "url", "u", "", "Indicate the url of the committee indexer service")
	rootCmd.Flags().StringVar(&arguments.ProtocolName, "protocol", "brc-20", "Indicate the meta protocol supported by the committee indexer")
	rootCmd.Flags().StringVar(&arguments.MetricAddr, "metrics", "0.0.0.0:8081", "Metrics listening address")
	return rootCmd
}
