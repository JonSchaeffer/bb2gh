package main

import (
	"flag"
	"log"
	"os"
	"fmt"
	"io/ioutil"
)

func parseArgs() string {
	var filename string
	var generateConfig bool
	flag.StringVar(&filename, "f", "", "Path to YAML config file")
	flag.StringVar(&filename, "file", "", "Path to YAML config file")
	flag.BoolVar(&generateConfig, "gc", false, "Generate a sample YAML config file")
	flag.BoolVar(&generateConfig, "generate-config", false, "Generate a sample YAML config file")

	flag.Usage = func() {
		fmt.Println(`
___.  ___.    ________         .__     
\_ |__\_ |__  \_____  \   ____ |  |__  
 | __ \| __ \  /  ____/  / ___\|  |  \ 
 | \_\ \ \_\ \/       \ / /_/  >   Y  \
 |___  /___  /\_______ \\___  /|___|  /
     \/    \/         \/_____/      \/ 
`)
		fmt.Println("bb2gh (bitbucket to github) migrates repositories from Bitbucket to GitHub.")
		fmt.Println("The 'gh' command line tool needs to be installed and authenticated. You will also need ssh key read/write access to Bitbucket and GitHub")
		fmt.Println("bb2gh reads a YAML configuration file that specifies the Bitbucket and GitHub organizations and the repositories to migrate.")
		fmt.Println("Options:")
		flag.PrintDefaults()
	}

	flag.Parse()
	
	if generateConfig {
		generateSampleConfig()
		os.Exit(0)
	}

	if filename == "" {
		log.Fatal("Please provide a YAML file with -f or --file flag")
	}


	return filename
}

func generateSampleConfig() {
	sampleConfig := `bitbucket:
  org: my_bitbucket_org
  repos:
    - repo1
    - repo2
    - repo3
github:
  org: myorg_github_org
`
	err := ioutil.WriteFile("sample_config.yaml", []byte(sampleConfig), 0644)
	if err != nil {
		log.Fatalf("Failed to write sample config file: %v", err)
	}
	fmt.Println("Generated sample_config.yaml")
}
