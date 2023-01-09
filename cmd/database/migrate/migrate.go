package migrate

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/xybor/xyauth/internal/models"
)

var Command = &cobra.Command{
	Use:   "migrate",
	Short: "Auto migrate the table",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if err := models.Migrate(); err != nil {
			panic(err)
		}
		fmt.Println("Migrated all tables successfully")
	},
}
