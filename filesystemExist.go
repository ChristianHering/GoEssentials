package goessentials

import "os"

//FileNotExist checks to see if a file
//exists and returns false if it exists
func FileNotExist(filename string) bool {
	return !FileExists(filename)
}

//FileExists checks to see if a file
//exists and returns true if it exists
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

//FolderNotExist checks to see if a dir
//exists and returns false if it exists
func FolderNotExist(path string) bool {
	return !FolderExists(path)
}

//FolderExists checks to see if a dir
//exists and returns true if it exists
func FolderExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}

	return info.IsDir()
}
