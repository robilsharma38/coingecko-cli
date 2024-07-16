package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/robilsharma38/coingecko-cli/client"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

var coinList = &cobra.Command{
	Use:   "Coin List",
	Short: "Getting Coin List",
	Run: func(cmd *cobra.Command, args []string) {
		coins, err := client.GetCoinList()
		if err != nil {
			fmt.Println(err)
		}
		printCoinsTable(coins)
	},
}

func printCoinsTable(coins []client.CoinList) {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	tbl := table.New("ID", "Name", "Symbol")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for _, c := range coins {
		tbl.AddRow(c.Id, c.Name, c.Symbol)
	}
	tbl.Print()
}

func init() {
	rootCmd.AddCommand(coinList)
}
