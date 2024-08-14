package face

import (
	"fmt"
	"orchestra/model"
	"os"

	"github.com/AlecAivazis/survey/v2"
)

type Options struct {
	Description string
	Message		string
	Filename    string
	Mode        string
	Route       string
	Field       string
	Value       string
	Environment string
}

func Surv(filenames []string) Options {
	var opts Options

	exit_if_error := func(err error, label interface{}) {
		if err != nil {
			fmt.Printf("Erro ao executar questionário \"%v\".\n", label)
			os.Exit(1)
		}
	}

	sel := &survey.Select{Message: "Qual o ambiente?", Options: []string{"desenvolvimento", "produção"}}
	error := survey.AskOne(sel, &opts.Environment)
	exit_if_error(error, sel.Message)
	sel = &survey.Select{Message: "Qual o modo?", Options: []string{"visualização", "execução"}}
	error = survey.AskOne(sel, &opts.Mode)
	exit_if_error(error, sel.Message)
	sel = &survey.Select{
		Message: "Qual rota?", 
		Options: []string{"cliente_contrato", "fechar_servico", "bloquear_contrato"},
	}
	error = survey.AskOne(sel, &opts.Route)
	exit_if_error(error, sel.Message)

	if opts.Mode == "execução" {
		if opts.Route == "cliente_contrato" {
			iscampo := "não"
			sel := &survey.Select{Message: "Especificar campo e valor?", Options: []string{"sim", "não"}}
			error = survey.AskOne(sel, &iscampo)
			exit_if_error(error, sel.Message)

			if iscampo == "sim" {
				sel := &survey.Select{Message: "Qual o campo?", Options: model.ContratoFields}
				error = survey.AskOne(sel, &opts.Field)
				exit_if_error(error, sel.Message)
				prom := &survey.Input{Message: "Qual o valor?"}
				error = survey.AskOne(prom, &opts.Value)
				exit_if_error(error, prom.Message)
			}
		} else if opts.Route == "fechar_servico" {
			prom := &survey.Input{Message: "Qual a mensagem de fechamento?"}
			error = survey.AskOne(prom, &opts.Message)
			exit_if_error(error, prom.Message)
		} else if opts.Route != "bloquear_contrato" {
			fmt.Print("Rota não implementada!")
			os.Exit(1)
		}
	} else {
		if opts.Route == "fechar_servico" {
			prom := &survey.Input{Message: "Qual a mensagem de fechamento?"}
			error = survey.AskOne(prom, &opts.Message)
			exit_if_error(error, prom.Message)
		}
	}

	prom := &survey.Input{Message: "Qual a descrição?"}
	error = survey.AskOne(prom, &opts.Description)
	exit_if_error(error, prom.Message)
	sel = &survey.Select{Message: "Qual arquivo 'csv'?", Options: filenames}
	error = survey.AskOne(sel, &opts.Filename)
	exit_if_error(error, sel.Message)

	return opts
}
