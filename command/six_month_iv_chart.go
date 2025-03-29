package command

import "github.com/spf13/cobra"

var SixMonthIVChart = &cobra.Command{
	Use: "six_month_iv_chart",
	Run: sixMonthIVChart,
}

func sixMonthIVChart(cmd *cobra.Command, args []string) {}
