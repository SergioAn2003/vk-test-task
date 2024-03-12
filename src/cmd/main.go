package main

import (
	"grpc-test/bimport"
	"grpc-test/config"
	"grpc-test/external/grpc"
	"grpc-test/internal/transaction"
	"grpc-test/rimport"
	"grpc-test/tools/logger"
	"grpc-test/tools/pgdb"
	"grpc-test/uimport"
	"os"
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

	db := pgdb.NewPostgresqlDB(config.PostgresqlConnectionString())
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	sm := transaction.NewSQLSessionManager(db)

	ri := rimport.NewRepositoryImports(sm)

	dblog := logger.NewDBLog(module, ri)

	bi := bimport.NewEmptyBridge()
	ui := uimport.NewUsecaseImports(log, dblog, ri, bi, sm)

	bi.InitBridge(
		ui.Usecase.Info,
	)

	server := grpc.NewServer(log, ui)
	server.RunGrpcServer()
}
