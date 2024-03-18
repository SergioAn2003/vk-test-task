package main

import (
	"os"
	"vk-film-library/bimport"
	"vk-film-library/config"
	"vk-film-library/external/rest"
	"vk-film-library/internal/transaction"
	"vk-film-library/rimport"
	"vk-film-library/tools/logger"
	"vk-film-library/tools/pgdb"
	"vk-film-library/uimport"
)

const (
	module = "template"
)

var (
	version string = os.Getenv("VERSION")
)

func main() {
	log := logger.NewFileLogger(module)
	log.Infoln("version", version)

	config, err := config.NewConfig(os.Getenv("CONF_PATH"))
	if err != nil {
		log.Fatal(err)
	}

	db := pgdb.NewPostgresqlDB(config.PostgresConnectionString())
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	sm := transaction.NewSQLSessionManager(db)
	ri := rimport.NewRepositoryImports(sm)

	bi := bimport.NewEmptyBridge()
	ui := uimport.NewUsecaseImports(log, ri, bi, sm)

	bi.InitBridge(
		ui.Usecase.Actor,
	)

	server := rest.NewServer(log, ui)
	server.Run()
}
