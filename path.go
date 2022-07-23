package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

///visits the path of folders

func visit(pathes string, f os.FileInfo, err error) error {
	fmt.Printf("Visited path are: %s\n", pathes)
	return nil
}
func main() {
	flag.Parse()
	root := flag.Arg(0)
	err := filepath.Walk(root, visit)
	fmt.Printf("filepath.Walk() returned %v\n", err)
}

/////output

/* S C:\Users\faiks\OneDrive\Masaüstü\week4> go run "c:\Users\faiks\OneDrive\Masaüstü\week4\path.go" C:\Users\faiks\OneDrive\Masaüstü\week4
Visited: C:\Users\faiks\OneDrive\Masaüstü\week4
Visited: C:\Users\faiks\OneDrive\Masaüstü\week4\calculator.go
Visited: C:\Users\faiks\OneDrive\Masaüstü\week4\conversation.go
Visited: C:\Users\faiks\OneDrive\Masaüstü\week4\math
Visited: C:\Users\faiks\OneDrive\Masaüstü\week4\math\math.go
Visited: C:\Users\faiks\OneDrive\Masaüstü\week4\path.go
*/
