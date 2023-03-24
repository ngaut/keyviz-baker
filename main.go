package main

import (
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	nameFlag := flag.String("name", "keyviz_baker", "Name of this bake.")
	imagePathFlag := flag.String("image_path", "", "Path to the image to render. Must be PNG format.")
	ripenessFlag := flag.Uint("ripeness", 256, "Ripeness (or brightness) to bake.")
	dbFlag := flag.String("db", "", "Database connection string. E.g. 'root:@tcp(127.0.0.1:4000)/test'.")
	skipPrepareFlag := flag.Bool("skip_prepare", false, "Skip preparation.")
	alignSecFlag := flag.Int("align_sec", 0, "Second to align to start baking.")
	intervalSecFlag := flag.Uint("interval_sec", 5, "Interval in seconds to draw each Y axis.")

	flag.Parse()

	name := *nameFlag
	imagePath := *imagePathFlag
	ripeness := *ripenessFlag
	db := *dbFlag
	skipPrepare := *skipPrepareFlag
	alignSec := *alignSecFlag
	intervalSec := *intervalSecFlag

	if len(imagePath) == 0 {
		_ = fmt.Errorf("Image path must not be empty.\n")
		return
	}
	if len(db) == 0 {
		_ = fmt.Errorf("Database connection string must not be empty.\n")
		return
	}

	fmt.Println("Making baker.")
	baker, err := MakeBaker(name, imagePath, ripeness, db)
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

	fmt.Sprintln("Aligning to", alignSec, "second.")
	baker.AlignTime(alignSec % 60)
	fmt.Println("Aligned.")

	fmt.Sprintln("Start baking at interval", intervalSec, "seconds.")
	err = baker.Bake(intervalSec)
	if err != nil {
		panic(err)
	}
	fmt.Println("Baking done.")
}
