package main

import (
	"context"
	"log"
	"net/url"
	"os"

	"github.com/urfave/cli"
)

const (
	dpFilenameEnvVar = "DATAPOINT_FILENAME"
	clientIdEnvVar   = "MQTT_CLIENT_ID"
	mqttServerEnvVar = "MQTT_SERVER"
)

func main() {
	app := &cli.App{
		Version: "0.0.1 (alpha)",
		Usage:   "an xComfort daemon",
		Commands: []cli.Command{
			cli.Command{
				Name:    "usb",
				Aliases: []string{"u"},
				Usage:   "connect via USB",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "file, f",
						Value: os.Getenv(dpFilenameEnvVar),
						Usage: "Datapoint file exported from MRF software",
					},
					cli.IntFlag{
						Name:  "device-number, d",
						Usage: "USB device number, if more than one is available (default 0)",
					},
					cli.StringFlag{
						Name:  "client-id, i",
						Value: os.Getenv(clientIdEnvVar),
						Usage: "MQTT client id",
					},
					cli.StringFlag{
						Name:  "server, s",
						Value: os.Getenv(mqttServerEnvVar),
						Usage: "MQTT server (e.g. tcp://username:password@host:port)",
					},
				},
				Action: func(cliContext *cli.Context) error {
					ctx := context.Background()
					relay := MqttRelay{}

					if err := relay.Init(cliContext.String("file"), &relay); err != nil {
						return err
					}

					url, err := url.Parse(cliContext.String("server"))
					if err != nil {
						return err
					}

					if err := relay.Connect(ctx, cliContext.String("client-id"), url); err != nil {
						return err
					}

					return Usb(ctx, cliContext.Int("device-number"), &relay.Interface)
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
