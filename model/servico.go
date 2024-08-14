/*
$dados = array(
    'id_chamado' => '', //insira o ID da O.S
    'data_inicio' => '02/12/2019 10:30:00', //Data e Hora de Inicio da finalização
    'data_final' => '04/12/2019 14:16:00', //Data e Hora de final da finalização

    'mensagem' => '',//Insira mensagem ao campo "Mensagem" //"Finalizada por ajuste da qualidade"
    'gera_comissao' => '',//"N" para Não "S" para Sim //"N"
    'finaliza_processo' => 'S',//"S" para finalizar o processo //"S"
    'status' => 'F',//Status "F" para finalizar //"F"
    'id_tecnico' => ''//ID do técnico Responsável); //809 //1197
*/

package model

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func ConstructServicos(ids [][]string, logfile *os.File, message string) []Dict {
	var servicos []Dict

	date := time.Now().Format("2006-01-02 15:04:05")

	for i, id := range ids {
		if i == 0 {
			continue
		}

		servico := Dict{
			"id_chamado": id[0],
			"data_inicio": date,
			"data_final": date,
			"mensagem": message,
			"gera_comissao": "N",
			"finaliza_processo": "S",
			"status": "F",
			"id_tecnico": "1197",//809
		}

		servicos = append(servicos, servico)
		fmt.Printf("servico: %v\n", servico)
	}

	fmt.Print("--\n")
	return servicos
}

func CloseServicos(servicos []Dict, token string, url string) {
	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	client := &http.Client{Transport: tr}
	auth := fmt.Sprintf("Basic %v", token)

	for _, servico := range servicos {
		url := url + "su_oss_chamado_fechar"
		servico_json, error := json.Marshal(servico)

		if error != nil {
			fmt.Println("A criação do seriamento json falhou.")
			os.Exit(1)
		}

		fmt.Printf("servico_json: %s\n", string(servico_json))
		req, error := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(servico_json))

		if error != nil {
			fmt.Println("A criação da requisição falhou.")
			os.Exit(1)
		}

		req.Header.Set("Content-type", "application/json")
		req.Header.Set("ixcsoft", "")
		req.Header.Set("Authorization", auth)

		resp, error := client.Do(req)

		if error != nil {
			fmt.Println("A requisição falhou.")
			fmt.Printf("error: %v\n", error)
			os.Exit(1)
		}

		if resp.StatusCode != http.StatusOK {
			fmt.Printf("A requisição falhou, status '%v'\n.", resp.StatusCode)
			os.Exit(1)
		}

		defer resp.Body.Close()

		resp_body, error := io.ReadAll(resp.Body)

		if error != nil {
			fmt.Println("A leitura da resposta falhou.")
			os.Exit(1)
		}

		fmt.Printf("resposta: %s\n", string(resp_body))
	}
}
