package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/urfave/cli/v2"
)

func main() {
	// To inspect whether domain is down or up
	// type the command 👉️ "go run . -domain domain-name(e.g medium.com)"
	app := cli.App{
		Name:  "WebHEALTH-INSPECT",
		Usage: "A CLI tool that checks whether the given domain is down.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "domain",
				Aliases:  []string{"d"},
				Usage:    "To inspect Domain name.",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "port",
				Aliases:  []string{"p"},
				Usage:    "Port number to check.",
				Required: false,
			},
		},
		Action: func(c *cli.Context) error {
			port := c.String("port")
			if c.String("port") == "" {
				port = "80"
			}
			status := Inspect(c.String("domain"), port)

			const col = 20
			fmt.Println("Inspecting...")
			for i := 0; i < col; i++ {
				fmt.Print("=")
				time.Sleep(200 * time.Millisecond)
			}
			fmt.Print(">")
			fmt.Println(" Done!")
			fmt.Println(status)
			domain_name := c.String("domain")

			if strings.Contains(status, "UP") {
				fmt.Println("Opening the website...")
				website := "https://" + domain_name + "/"
				exec.Command("xdg-open", website).Run()
				fmt.Println(website) //https://medium.com/
			} else {
				fmt.Println("Health of website is DOWN")
			}
			return nil
		},
	}

	app.Run(os.Args)

}
