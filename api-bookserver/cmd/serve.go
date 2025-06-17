/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	//"bookapi/routes"
	"github.com/SajjadSadi074/client_go/api-bookserver/Deployment"

	"github.com/spf13/cobra"
)

var (
	// Port stores port number for starting a connection
	Port     int
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "start cmd starts the server on a port",
		Long: `It starts the server on a given port number, 
				Port number will be given in the cmd`,
		Run: func(cmd *cobra.Command, args []string) {
			Deployment.BookServerDeployment()
		},
	}
)

func init() {
	serveCmd.PersistentFlags().IntVarP(&Port, "port", "p", 8081, "Port number for starting server")
	rootCmd.AddCommand(serveCmd)
}
