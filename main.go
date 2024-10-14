package main

import (
	"fivetran/api"
	"os"
)

func main() {
	api.Refresh("", "Warehouse", "1sIoTnItnQRuOFrL9L1rRpfw8valt_1faaz4V9-1nIrg")
	//cmd.Execute()
	os.Exit(0)
}
