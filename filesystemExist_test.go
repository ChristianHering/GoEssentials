package goessentials

import (
	"log"
)

func ExampleFileNotExist() {
	if FileNotExist("DoesNotExist.here") {
		log.Println("The file doesn't exist... Cool!")
	} else {
		log.Fatalln("This file shouldn't exist")
	}
}

func ExampleFileExists() {
	if FileExists("filesystemExist_test.go") {
		log.Println("Just making sure I'm still here")
	} else {
		log.Fatalln("We seem to have disappeared")
	}
}

func ExampleFolderNotExist() {
	if FolderNotExist("DoesNotExist") {
		log.Println("It'd be weird if that directory was a thing anyway")
	} else {
		log.Fatalln(`Did you just create a directory called "DoesNotExist"?`)
	}
}

func ExampleFolderExists() {
	if FolderExists("../GoEssentials") {
		log.Println("Just checking in on my parent")
	} else {
		log.Fatalln("Where'd Go go this time")
	}
}
