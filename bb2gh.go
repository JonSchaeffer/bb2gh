package main

import (
	"log"
)


func main() {

	filename := parseArgs()
	config, err := ReadConfig(filename)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	migrate(config.Bitbucket.Org, config.Bitbucket.Repos, config.Github.Org)

// 	fmt.Printf("Bitbucket Org: %s\n", config.Bitbucket.Org)
// 	fmt.Println("Bitbucket Repos:")
// 	for _, repo := range config.Bitbucket.Repos {
// 		fmt.Printf("\t%s\n", repo)
// 	}
// 	fmt.Printf("Github Org: %s\n", config.Github.Org)
//
// 	fmt.Println("Hello, World!")
}
