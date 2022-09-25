package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/yori7/imagex"
)

var (
	out *string
	margin *float64
	bg *string
)

func main() {
	out = flag.String("o", "out.png", "output file name")
	margin = flag.Float64("m", 1.01, "mergin between images")
	bg = flag.String("b", "white", "background color")
	// strAsp []string = strings.Split(*flag.String("asp", "", "Output aspect as width:height"), ":")
	// quality = flag.Int("q", 100, "Resolution")
	flag.Parse()
	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}
	var outimg, err = imagex.Collage(flag.Args(), *bg, *margin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		flag.Usage()
		os.Exit(1)
	}
	imagex.Save(outimg, *out)
}
