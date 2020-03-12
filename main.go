package utils

import (
	routes "./routes"
	log "./lib/logs"
	models "./models"
)


func main() {
	PUERTO := "5001"
	log.PrintLog("Servicio iniciado")
	models.CreateDB()
	routes.LoadRouter(PUERTO)
}
