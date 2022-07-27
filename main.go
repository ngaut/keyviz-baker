package main

import (
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	nameFlag := flag.String("name", "test", "Name of this bake.")
	imagePathFlag := flag.String("image_path", "", "Path to the image to render.")
	dbFlag := flag.String("db", "", "Database connection string.")
	skipPrepareFlag := flag.Bool("skip_prepare", false, "Skip prepare steps.")
	alignSecFlag := flag.Int("align_sec", 0, "Seconds to align to start baking.")
	intervalSecFlag := flag.Uint("interval_sec", 10, "Interval in seconds to draw each Y axis.")

	flag.Parse()

	name := *nameFlag
	imagePath := *imagePathFlag
	db := *dbFlag
	skipPrepare := *skipPrepareFlag
	alignSec := *alignSecFlag
	intervalSec := *intervalSecFlag

	fmt.Println("Making baker.")
	baker, err := MakeBaker(name, imagePath, db)
	if err != nil {
		panic(err.Error())
	}
	defer baker.Close()
	fmt.Println("Baker made.")

	if !skipPrepare {
		fmt.Println("Preparing to bake.")
		err := baker.Prepare()
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("Preparation done.")
	}

	if alignSec >= 0 {
		fmt.Println(fmt.Sprintf("Aligning to %d second.", alignSec))
		baker.AlignTime(alignSec)
		fmt.Println("Aligned.")
	}

	fmt.Println(fmt.Sprintf("Start baking at interval %d seconds.", intervalSec))
	err = baker.Bake(intervalSec)
	if err != nil {
		panic(err)
	}
	fmt.Println("Baking done.")
}
