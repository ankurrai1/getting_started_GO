//Suppose we wanted to create a file, write to it, and then close when we’re done.
//Here’s how we could do that with defer.

package main
import (
    "fmt"
    "os" //Package os provides a platform-independent interface to operating system functionality.
)
// os use case example
// file, err := os.Open("file.go") For read access.
// if err != nil {
// 	log.Fatal(err)
// }


func main() {
    f := createFile("/defer.txt")
    // Immediately after getting a file object with createFile, we defer the closing of that file with closeFile.
    //  This will be executed at the end of the enclosing function (main), after writeFile has finished.
    defer closeFile(f)
    writeFile(f)
}


func createFile(p string) *os.File {  // will create a file if not exist and retuen the referance of that file
    fmt.Println("creating")
    f, err := os.Create(p)
    if err != nil {
        panic(err) // panic is used to handle if panic came, better discribed in panic.go
    }
    return f  // retuning a file referance which is created or which existed and opened
}


func writeFile(f *os.File) {
    fmt.Println("writing")
    fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
    fmt.Println("closing")
    err := f.Close()
    if err != nil {  //It’s important to check for errors when closing a file, even in a deferred function.
        fmt.Fprintf(os.Stderr, "error: %v\n", err) // Fprintf formats according to a format specifier and writes to First arg
        os.Exit(1)
    }
}
