// cmd/iamsad/main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	imgcat "github.com/martinlindhe/imgcat/lib"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "I am sad"
	app.Usage = "Shows you cute things when you're feeling down."

	dog := Dog{}

	breedFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "breed",
			Value: "hound",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "breeds",
			Usage: "List the available dog breeds",
			Action: func(c *cli.Context) error {
				breeds := dog.Breeds()

				fmt.Println(breeds)
				return nil
			},
		},
		{
			Name:  "subbreeds",
			Usage: "Given a breed, list the available sub breeds",
			Flags: breedFlags,
			Action: func(c *cli.Context) error {
				subbreeds := dog.SubBreeds("hound")

				fmt.Println(subbreeds)
				return nil
			},
		},
		{
			Name:  "photo",
			Usage: "Show a random photo of a dog!",
			Action: func(c *cli.Context) error {
				photoPath := dog.RandomPhoto()

				res, err := http.Get(photoPath)
				if err != nil {
					log.Fatal(err)
				}

				imgcat.Cat(res.Body, os.Stdout)
				return nil
			},
		},
	}

	// start our application
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
