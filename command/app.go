package command

import (
	"fmt"
	"github.com/thrasher-/gocryptotrader/config"
	"github.com/thrasher-/gocryptotrader/exchanges/gdax"
	"github.com/urfave/cli"
)

// App returns the command line interface to gocryptotrader
func App() *cli.App {

	app := cli.NewApp()
	app.Name = "gocryptotrader"
	app.Usage = "Trade cryptocurrencies"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "log-level",
			EnvVar: "LOG_LEVEL",
			Value:  "DEBUG",
			Usage:  "Set log level `level`",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:   "sim",
			Usage:  "Run simulation",
			Action: startSim,
		},
	}

	return app
}

func startSim(c *cli.Context) error {
	cfg := config.GetConfig()
	cfg.LoadConfig("testdata/configtest.json")
	gdxConfig, err := cfg.GetExchangeConfig("GDAX")
	if err != nil {
		fmt.Println("failed to get exchange config: ", err.Error())
		return err
	}

	var g gdax.GDAX
	g.SetDefaults()
	g.Setup(gdxConfig)
	t, err := g.GetTicker("BTC-USD")
	fmt.Printf("Ticker: %+v", t)

	return nil
}
