package serve

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/winterssy/mxget/internal/routes"
)

var (
	port int
)

var CmdServe = &cobra.Command{
	Use:   "serve",
	Short: "Run mxget as an API server.",
}

func Run(cmd *cobra.Command, args []string) {
	gin.SetMode(gin.ReleaseMode)
	app := gin.New()
	routes.Init(app)
	app.Run(fmt.Sprintf(":%d", port))
}

func init() {
	CmdServe.Flags().IntVar(&port, "port", 5050, "server listening port")
	CmdServe.Run = Run
}
