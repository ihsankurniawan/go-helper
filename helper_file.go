package helper

import(
	"io"
	"os"
	"mime/multipart"
)

// Function to save multipart file into desired location
// Ex: SaveFileToDisk(file, "/tmp/temporary_file")
func SaveFileToDisk(file multipart.File, saveLoc string) error {
	defer file.Close()
	
	f, e := os.OpenFile(saveLoc, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if e != nil {
		return e
	}
	defer f.Close()
	
	io.Copy(f, file)
	return nil
}