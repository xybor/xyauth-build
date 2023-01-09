package delete

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/xybor/xyauth/internal/models"
	"github.com/xybor/xyauth/internal/utils"
)

var Command = &cobra.Command{
	Use:   "delete",
	Short: "Delete the table",
	Args:  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if utils.GetConfig().GetDefault("general.environment", "dev").MustString() != "dev" {
			fmt.Println("Do not delete table if you aren't in dev environment")
			return
		}
		models.DeleteTable()
		fmt.Println("Deleted all tables successfully")
	},
}
