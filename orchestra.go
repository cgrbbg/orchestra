package main

import (
	"encoding/base64"
	"encoding/csv"
	"fmt"
	"orchestra/face"
	"orchestra/model"
	"os"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

type APIConnection struct {
	Url   string
	Token string
	Dsn   string
}

func getConnection(environment string) (APIConnection, model.IXCConnection) {
	err := godotenv.Load()
	if (err != nil) {
		fmt.Printf("Varíaveis de ambiente não carregaram")
		os.Exit(1)
	}

	var api APIConnection
	var web model.IXCConnection

	config := mysql.NewConfig()
	config.DBName = os.Getenv("db_name")
	config.User = os.Getenv("db_user")
	config.Net = os.Getenv("db_protocol")

	if environment == "desenvolvimento" {
		token := os.Getenv("dev_token")
		api.Token = base64.StdEncoding.EncodeToString([]byte(token))
		api.Url = os.Getenv("dev_url")
		web.Pass = os.Getenv("dev_password")
		web.Url = os.Getenv("dev_url")
		web.User = os.Getenv("dev_user")
		config.Addr = os.Getenv("dev_addr")
		config.Passwd = web.Pass
	} else {
		token := os.Getenv("prod_token")
		api.Token = base64.StdEncoding.EncodeToString([]byte(token))
		api.Url = os.Getenv("prod_url")
		web.Pass = os.Getenv("prod_password")
		web.Url = os.Getenv("prod_url")
		web.User = os.Getenv("prod_user")
		config.Addr = os.Getenv("prod_addr")
		config.Passwd = web.Pass
	}

	api.Dsn = config.FormatDSN()

	return api, web
}

func getCSVs(dir_name string) []string {
	var filenames []string

	files, err := os.ReadDir(dir_name)

	if err != nil {
		fmt.Println("Erro ao abrir pasta.")
		os.Exit(1)
	}

	for _, file := range files {
		filename := file.Name()

		if strings.Contains(filename, ".csv") {
			filenames = append(filenames, filename)
		}
	}

	if len(filenames) == 0 {
		fmt.Println("Não há arquivos 'csv' na pasta.")
		os.Exit(1)
	}

	return filenames
}

func getCSVData(filename string) [][]string {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Printf("Erro ao abrir arquivo '%v'.\n", filename)
		os.Exit(1)
	}

	reader := csv.NewReader(file)
	reader.TrimLeadingSpace = true
	reader.Comma = ';'
	data, err := reader.ReadAll()

	if err != nil {
		fmt.Println("Erro ao construir objeto csv.")
		os.Exit(1)
	} else if len(data) == 0 {
		fmt.Println("Arquivo csv está vazio.")
		os.Exit(1)
	} else if len(data) == 1 {
		fmt.Println("Arquivo csv está sem coluna ou sem itens.")
		os.Exit(1)
	} else if data[0][0] != "id" {
		fmt.Println("campo 'id' não encontrado na primeira posição.")
		os.Exit(1)
	}

	return data
}

func getDBConnection(dsn string) *sqlx.DB {
	db, err := sqlx.Open("mysql", dsn)

	fmt.Println("Conectando com o banco...")
	fmt.Println("--")

	if err != nil {
		fmt.Println("Erro ao conectar com o banco.")
		fmt.Printf("Erro '%v'.\n", err)
		os.Exit(1)
	}

	err = db.Ping()

	if err != nil {
		fmt.Println("Erro ao pingar com o banco.")
		fmt.Printf("Erro '%v'.\n", err)
		os.Exit(1)
	}

	fmt.Println("Banco conectado.")
	fmt.Println("--")

	return db
}

func getIDs(data [][]string) string {
	var ids string

	for i := 1; i < len(data); i++ {
		ids = ids + data[i][0] + ","
	}

	if strings.HasSuffix(ids, ",") {
		ids = ids[:len(ids)-len(",")]
	}

	fmt.Printf("ids: %v\n", ids)
	fmt.Println("--")

	return ids
}

func openLogFile(filename string) *os.File {
	logfile, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Printf("arquivo de log '%s' não abriu.\n", filename)
		os.Exit(1)
	}

	return logfile
}

func main() {
	const archive_dir = "./arquivos/"

	logfile := openLogFile("./log/log-file.txt")
	defer logfile.Close()

	filenames := getCSVs(archive_dir)

	opts := face.Surv(filenames)
	api, _ := getConnection(opts.Environment)
	data := getCSVData(archive_dir + opts.Filename)

	start := time.Now()

	ids := getIDs(data)
	db := getDBConnection(api.Dsn)

	logHeader := fmt.Sprintf("%s - %s\n", opts.Description, time.Now().Format("01/02 3:4"))
	logfile.Write([]byte(logHeader))

	if opts.Route == "cliente_contrato" {
		contratos := model.GetContratos(db, ids, data, logfile, opts.Field, opts.Value)

		if opts.Mode == "execução" {
			model.PutContratos(contratos, api.Token, api.Url)
		}
	} else if opts.Route == "bloquear_contrato" {
		blocks := model.ConstructBlockContratos(data, logfile)

		if opts.Mode == "execução" {
			model.BlockContratos(blocks, api.Token, api.Url)
		}
	} else if opts.Route == "fechar_servico" {
		servicos := model.ConstructServicos(data, logfile, opts.Message)

		if opts.Mode == "execução" {
			model.CloseServicos(servicos, api.Token, api.Url)
		}
	} else {
		fmt.Print("Rota não implementada")
		os.Exit(1)
	}

	fin := time.Since(start)

	duration := fmt.Sprintf("o programa durou '%s'.\n\n", fin)

	fmt.Print(duration)
	logfile.Write([]byte(duration))
}
