package main

import ( 
    "fmt"
    "io"
    "net/http"
    "os"
)

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.

func downloadFile(filePath string, url string) error {
    
    // Create the file
    out, err := os.Create(filePath)
    
    if err != nil {
        return err
    }
    defer out.Close()

    // Get the data
    resp, err := http.Get(url)
    if err != nil {
        return err
    }

    // Write the body to a file
    _, err = io.Copy(out, resp.Body)
    if err != nil {
        return err
    }
    
    return nil
}

func main() {
    
    fmt.Printf("Downloading file...\n")

    fileUrl := "https://golangcode.com/images/avatar.jpg"

    err := downloadFile("avatar.jpg", fileUrl)
    if err != nil {
        panic(err)
    }

}

