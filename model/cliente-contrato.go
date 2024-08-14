package model

import (
	"bytes"
	"crypto/tls"
	"reflect"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"github.com/jmoiron/sqlx"
)

type Dict map[string]interface{}

// type contrato struct {
// 	Id string `json:"id"`
// 	Tipo string `json:"tipo"`
// 	Id_cliente string `json:"id_cliente"`
// 	Id_vd_contrato string `json:"id_vd_contrato"`
// 	Contrato string `json:"contrato"`
// 	Id_tipo_contrato string `json:"id_tipo_contrato"`
// 	Id_modelo string `json:"id_modelo"`
// 	Id_filial string `json:"id_filial"`
// 	Data string `json:"data"`
// 	Id_tipo_documento string `json:"id_tipo_documento"`
// 	Id_carteira_cobranca string `json:"id_carteira_cobranca"`
// 	Id_vendedor string `json:"id_vendedor"`
// 	Cc_previsao string `json:"cc_previsao"`
// 	Tipo_cobranca string `json:"tipo_cobranca"`
// 	Renovacao_automatica string `json:"renovacao_automatica"`
// 	Base_geracao_tipo_doc string `json:"base_geracao_tipo_doc"`
// 	Bloqueio_automatico string `json:"bloqueio_automatico"`
// 	Aviso_atraso string `json:"aviso_atraso"`
// 	Endereco_padrao_cliente string `json:"endereco_padrao_cliente"`
//
// 	Data_ativacao string `json:"data_ativacao"`
// 	Data_renovacao string `json:"data_renovacao"`
//
// }

// func (c *contrato) scan(rows *sql.Rows) error {
// 	return rows.Scan(
// 		&c.Id, &c.Tipo, &c.Id_cliente, &c.Id_vd_contrato, &c.Contrato,
// 		&c.Id_tipo_contrato, &c.Id_modelo, &c.Id_filial, &c.Data,
// 		&c.Id_tipo_documento, &c.Id_carteira_cobranca, &c.Id_vendedor, &c.Cc_previsao,
// 		&c.Tipo_cobranca, &c.Renovacao_automatica, &c.Base_geracao_tipo_doc,
// 		&c.Bloqueio_automatico, &c.Aviso_atraso, &c.Endereco_padrao_cliente,
// 		&c.Data_ativacao, &c.Data_renovacao,
// 	)
// }

// func (c *contrato) log(logfile *os.File) {
// 	cliente_contrato_formated := fmt.Sprintf("%v\n", c)
// 	_, error := logfile.Write([]byte(cliente_contrato_formated))
//
// 	if error != nil {
// 		fmt.Println("falha ao escrever para arquivo 'log-file.txt'.")
// 		os.Exit(1)
// 	}
// }

// func (c *contrato) toMap() Dict {
// 	cliente_map := Dict{}
//
// 	cliente_reflect := reflect.ValueOf(c)
//
// 	if cliente_reflect.Kind() == reflect.Ptr {
// 		cliente_reflect = cliente_reflect.Elem()
// 	}
//
// 	cliente_type := cliente_reflect.Type()
//
// 	for i := 0; i < cliente_reflect.NumField(); i++ {
// 		cliente_field := cliente_type.Field(i)
//
// 		cliente_field_value := cliente_field.Tag.Get("json")
//
// 		if cliente_field_value != "" {
// 			cliente_map[cliente_field_value] = cliente_reflect.Field(i).Interface()
// 		}
// 	}
//
// 	return cliente_map
// }

var ContratoFields = []string {
	"tipo",
	"id_cliente",
	"id_vd_contrato",
	"descricao_aux_plano_venda",
	"contrato",
	"id_tipo_contrato",
	"id_modelo",
	"assinatura_digital",
	"integracao_assinatura_digital",
	"liberacao_bloqueio_manual",
	"id_filial",
	"indicacao_contrato_id",
	"data_assinatura",
	"isentar_contrato",
	"data_ativacao",
	"data",
	"data_renovacao",
	"pago_ate_data",
	"status",
	"status_internet",
	"status_velocidade",
	"tipo_produtos_plano",
	"motivo_inclusao",
	"id_indexador_reajuste",
	"url_assinatura_digital",
	"token_assinatura_digital",
	"id_tipo_documento",
	"id_carteira_cobranca",
	"id_vendedor",
	"moeda",
	"comissao",
	"cc_previsao",
	"tipo_cobranca",
	"renovacao_automatica",
	"gerar_finan_assin_digital_contrato",
	"base_geracao_tipo_doc",
	"id_contrato_principal",
	"num_parcelas_atraso",
	"nf_info_adicionais",
	"credit_card_recorrente_bandeira_cartao",
	"credit_card_recorrente_token",
	"ids_contratos_recorrencia",
	"tipo_doc_opc",
	"tipo_doc_opc2",
	"tipo_doc_opc3",
	"tipo_doc_opc4",
	"id_tipo_doc_ativ",
	"id_produto_ativ",
	"taxa_instalacao",
	"id_cond_pag_ativ",
	"id_responsavel",
	"id_vendedor_ativ",
	"ativacao_numero_parcelas",
	"ativacao_vencimentos",
	"ativacao_valor_parcela",
	"fidelidade",
	"data_expiracao",
	"desconto_fidelidade",
	"id_instalador",
	"taxa_improdutiva",
	"tipo_condicao_pag",
	"bloqueio_automatico",
	"nao_bloquear_ate",
	"aviso_atraso",
	"nao_avisar_ate",
	"desbloqueio_confianca",
	"desbloqueio_confianca_ativo",
	"restricao_auto_desbloqueio",
	"motivo_restricao_auto_desbloq",
	"obs",
	"nao_susp_parc_ate",
	"liberacao_suspensao_parcial",
	"utilizando_auto_libera_susp_parc",
	"restricao_auto_libera_susp_parcial",
	"motivo_restri_auto_libera_parc",
	"contrato_suspenso",
	"data_inicial_suspensao",
	"data_final_suspensao",
	"data_acesso_desativado",
	"data_cancelamento",
	"id_responsavel_cancelamento",
	"motivo_cancelamento",
	"motivo_adicional",
	"concorrente_mot_adicional",
	"obs_cancelamento",
	"data_negativacao",
	"id_responsavel_negativacao",
	"protocolo_negativacao",
	"id_motivo_negativacao",
	"obs_negativacao",
	"data_desistencia",
	"id_responsavel_desistencia",
	"motivo_desistencia",
	"obs_desistencia",
	"obs_contrato",
	"alerta_contrato",
	"imp_realizado",
	"imp_inicial",
	"imp_carteira",
	"imp_importacao",
	"imp_treinamento",
	"imp_rede",
	"imp_bkp",
	"imp_obs",
	"imp_final",
	"imp_status",
	"imp_motivo",
	"dt_ult_bloq_auto",
	"dt_ult_finan_atraso",
	"dt_ult_des_bloq_conf",
	"dt_ult_liberacao_susp_parc",
	"dt_ult_ativacao",
	"dt_utl_negativacao",
	"dt_ult_desiste",
	"data_cadastro_sistema",
	"ultima_atualizacao",
	"data_retomada_contrato",
	"endereco_padrao_cliente",
	"id_condominio",
	"bloco",
	"apartamento",
	"cep",
	"endereco",
	"numero",
	"bairro",
	"cidade",
	"complemento",
	"referencia",
	"latitude",
	"longitude",
	"tipo_localidade",
	"avalista_1",
	"avalista_2",
	"testemunha_assinatura_digital",
	"document_photo",
	"selfie_photo",
}

// const contratoQuery = ` 
// 	select
// 		id, tipo, id_cliente, id_vd_contrato, contrato, id_tipo_contrato,
// 		id_modelo, id_filial, DATE_FORMAT(data, "%%d/%%m/%%Y") as data, 
// 		id_tipo_documento, id_carteira_cobranca, id_vendedor, cc_previsao, 
// 		tipo_cobranca, renovacao_automatica, base_geracao_tipo_doc, 
// 		bloqueio_automatico, aviso_atraso, endereco_padrao_cliente,
// 		DATE_FORMAT(data_ativacao, "%%d/%%m/%%Y") as data_ativacao,
// 		DATE_FORMAT(data_renovacao, "%%d/%%m/%%Y") as data_renovacao
// 	from cliente_contrato where id in (%s)
// `

const contratoQuery = `
	select
		id,
		tipo, 
		id_cliente, 
		id_vd_contrato, 
		id_tipo_contrato,
		id_modelo, 
		id_filial, 
		DATE_FORMAT(data, "%%d/%%m/%%Y") as data, 
		case when data_assinatura is NULL then "0" else data_assinatura end data_assinatura,
		id_tipo_documento, 
		id_carteira_cobranca, 
		id_vendedor, 
		cc_previsao, 
		tipo_cobranca, 
		renovacao_automatica, 
		base_geracao_tipo_doc, 
		bloqueio_automatico, 
		aviso_atraso, 
		endereco_padrao_cliente,
		DATE_FORMAT(data_ativacao, "%%d/%%m/%%Y") as data_ativacao,
		DATE_FORMAT(data_renovacao, "%%d/%%m/%%Y") as data_renovacao,
		descricao_aux_plano_venda,
		contrato,
		assinatura_digital,
		integracao_assinatura_digital,
		liberacao_bloqueio_manual,
		indicacao_contrato_id,
		DATE_FORMAT(data_assinatura, "%%d/%%m/%%Y") as data_assinatura,
		isentar_contrato,
		DATE_FORMAT(pago_ate_data, "%%d/%%m/%%Y") as pago_ate_data,
		status,
		status_internet,
		status_velocidade,
		tipo_produtos_plano,
		motivo_inclusao,
		id_indexador_reajuste,
		url_assinatura_digital,
		token_assinatura_digital,
		moeda,
		comissao,
		gerar_finan_assin_digital_contrato,
		id_contrato_principal,
		num_parcelas_atraso,
		nf_info_adicionais,
		credit_card_recorrente_bandeira_cartao,
		credit_card_recorrente_token,
		ids_contratos_recorrencia,
		tipo_doc_opc,
		tipo_doc_opc2,
		tipo_doc_opc3,
		tipo_doc_opc4,
		id_tipo_doc_ativ,
		id_produto_ativ,
		taxa_instalacao,
		id_cond_pag_ativ,
		id_responsavel,
		id_vendedor_ativ,
		ativacao_numero_parcelas,
		ativacao_vencimentos,
		ativacao_valor_parcela,
		fidelidade,
		DATE_FORMAT(data_expiracao, "%%d/%%m/%%Y") as data_expiracao,
		desconto_fidelidade,
		id_instalador,
		taxa_improdutiva,
		DATE_FORMAT(nao_bloquear_ate, "%%d/%%m/%%Y") as nao_bloquear_ate,
		DATE_FORMAT(nao_avisar_ate, "%%d/%%m/%%Y") as nao_avisar_ate,
		desbloqueio_confianca,
		desbloqueio_confianca_ativo,
		restricao_auto_desbloqueio,
		motivo_restricao_auto_desbloq,
		obs,
		DATE_FORMAT(nao_susp_parc_ate, "%%d/%%m/%%Y") as nao_susp_parc_ate,
		liberacao_suspensao_parcial,
		utilizando_auto_libera_susp_parc,
		restricao_auto_libera_susp_parcial,
		motivo_restri_auto_libera_parc,
		contrato_suspenso,
		DATE_FORMAT(data_inicial_suspensao, "%%d/%%m/%%Y") as data_inicial_suspensao,
		DATE_FORMAT(data_final_suspensao, "%%d/%%m/%%Y") as data_final_suspensao,
		DATE_FORMAT(data_acesso_desativado, "%%d/%%m/%%Y") as data_acesso_desativado,
		DATE_FORMAT(data_cancelamento, "%%d/%%m/%%Y") as data_cancelamento,
		id_responsavel_cancelamento,
		motivo_cancelamento,
		motivo_adicional,
		concorrente_mot_adicional,
		obs_cancelamento,
		DATE_FORMAT(data_negativacao, "%%d/%%m/%%Y") as data_negativacao,
		id_responsavel_negativacao,
		protocolo_negativacao,
		id_motivo_negativacao,
		obs_negativacao,
		DATE_FORMAT(data_desistencia, "%%d/%%m/%%Y") as data_desistencia,
		id_responsavel_desistencia,
		motivo_desistencia,
		obs_desistencia,
		obs_contrato,
		alerta_contrato,
		imp_realizado,
		DATE_FORMAT(imp_inicial, "%%d/%%m/%%Y") as imp_inicial,
		imp_carteira,
		imp_importacao,
		imp_treinamento,
		imp_rede,
		imp_bkp,
		imp_obs,
		DATE_FORMAT(imp_final, "%%d/%%m/%%Y") as imp_final,
		imp_status,
		imp_motivo,
		DATE_FORMAT(dt_ult_bloq_auto, "%%d/%%m/%%Y") as dt_ult_bloq_auto,
		DATE_FORMAT(dt_ult_bloq_manual, "%%d/%%m/%%Y") as dt_ult_bloq_manual,
		DATE_FORMAT(dt_ult_bloq_manual, "%%d/%%m/%%Y") as dt_ult_bloq_manual,
		DATE_FORMAT(dt_ult_finan_atraso, "%%d/%%m/%%Y") as dt_ult_finan_atraso,
		DATE_FORMAT(dt_ult_des_bloq_conf, "%%d/%%m/%%Y") as dt_ult_des_bloq_conf,
		DATE_FORMAT(dt_ult_liberacao_susp_parc, "%%d/%%m/%%Y") as dt_ult_liberacao_susp_parc,
		DATE_FORMAT(dt_ult_ativacao, "%%d/%%m/%%Y") as dt_ult_ativacao,
		DATE_FORMAT(dt_utl_negativacao, "%%d/%%m/%%Y") as dt_utl_negativacao,
		DATE_FORMAT(dt_ult_desiste, "%%d/%%m/%%Y") as dt_ult_desiste,
		DATE_FORMAT(data_cadastro_sistema, "%%d/%%m/%%Y") as data_cadastro_sistema,
		ultima_atualizacao,
		DATE_FORMAT(data_retomada_contrato, "%%d/%%m/%%Y") as data_retomada_contrato,
		id_condominio,
		bloco,
		apartamento,
		cep,
		endereco,
		numero,
		bairro,
		cidade,
		complemento,
		referencia,
		latitude,
		longitude,
		tipo_localidade,
		avalista_1,
		avalista_2,
		testemunha_assinatura_digital,
		document_photo,
		selfie_photo
	from cliente_contrato where id in (%s)
`

// func GetContratosFields() []string {
// 	contrato := &contrato{}
// 	cliente_reflect := reflect.ValueOf(contrato).Elem()
// 	contrato_len := cliente_reflect.NumField()
// 	cliente_type := cliente_reflect.Type()
//
// 	colunas := make([]string, contrato_len)
//
// 	for i := 0; i < contrato_len; i++ {
// 		cliente_field := cliente_type.Field(i)
//
// 		cliente_field_value := cliente_field.Tag.Get("json")
// 		colunas[i] = cliente_field_value
// 	}
//
// 	return colunas
// }

func GetContratos(
	db *sqlx.DB, ids string, data [][]string, logfile *os.File, field string, value string,
) []Dict {
	var contratos []Dict

	quer := fmt.Sprintf(contratoQuery, ids)
	rows, error := db.Queryx(quer)
	// defer rows.Close()

	if error != nil {
		fmt.Println("Erro ao executar query.")
		fmt.Println(error)
		os.Exit(1)
	}

	index := 1 
	for rows.Next() {
		// contrato := contrato{}
		// error := contrato.scan(rows)

		contrato_map := make(Dict)
		error := rows.MapScan(contrato_map)
		// contrato_map, error := rows.SliceScan()

		if error != nil {
			fmt.Println("Erro ao escanear cliente_contrato.")
			fmt.Println(error)
			os.Exit(1)
		}

		for k, v := range contrato_map {
			t := reflect.TypeOf(v)
			if t != nil {
				switch t.Kind() {
				case reflect.Slice:
					contrato_map[k] = fmt.Sprintf("%s", v)

				default:
					// do nothing
				}
			}
		}

		// contrato.log(logfile)
		// contrato_map := contrato.toMap()

		cliente_contrato_formated := fmt.Sprintf("%v\n", contrato_map)
		_, error = logfile.Write([]byte(cliente_contrato_formated))

		if error != nil {
			fmt.Println("falha ao escrever para arquivo 'log-file.txt'.")
			os.Exit(1)
		}

		row_len := len(data[0])
		for i := 0; i < row_len; i++ {
			contrato_map[data[0][i]] = data[index][i]
			//fmt.Printf("%s %s", data[0][i], data[index][i])
		}

		if field != "" && value != "" {
			_, ok := contrato_map[field]

			if !ok {
				fmt.Printf("'%v' não existe!", field)
				os.Exit(1)
			}

			contrato_map[field] = value
		} else {
			fmt.Printf("contrato: %+v\n", contrato_map)
		}

		contratos = append(contratos, contrato_map)
		index++
	}

	err := rows.Err()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	return contratos
}

func PutContratos(contratos []Dict, token string, url string) {
	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	client := &http.Client{Transport: tr}
	auth := fmt.Sprintf("Basic %v", token)

	for _, contrato := range contratos {
		_, ok := contrato["id"]

		if !ok {
			fmt.Println("campo 'id' não encontrado.")
			os.Exit(1)
		}

		id := contrato["id"].(string)

		url := url+"cliente_contrato/"+id
		contrato_json, error := json.Marshal(contrato)

		if error != nil {
			fmt.Println("A criação do seriamento json falhou.")
			os.Exit(1)
		}

		fmt.Printf("contrato_json: %s\n", string(contrato_json))
		req, error := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(contrato_json))

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

			bodyBytes := []byte{}
			_, err := resp.Body.Read(bodyBytes)
			if err != nil {
				fmt.Println("A leitura do corpo falhou.")
				fmt.Printf("error: %v\n", error)
				os.Exit(1)
			}

			fmt.Printf("corpo: '%s'.\n", string(bodyBytes))
			os.Exit(1)
		}

		defer resp.Body.Close()

		resp_body, error := io.ReadAll(resp.Body)

		if error != nil {
			fmt.Println("A leitura da resposta falhou.")
			os.Exit(1)
		}

		fmt.Printf("resposta: %s\n", string(resp_body))
		fmt.Printf("--")
	}
}

func ConstructBlockContratos(ids [][]string, logfile *os.File) []Dict {
	var blocks []Dict

	for i, id := range ids {
		if i == 0 {
			continue;
		}

		block := Dict{"id_contrato": id[0]}

		block_formated := fmt.Sprintf("%v\n", block)

		_, error := logfile.Write([]byte(block_formated))
		if (error != nil) {
			fmt.Println("A escrita para log falhou.")
			os.Exit(1)
		}

		blocks = append(blocks, block)
		fmt.Printf("block: %v\n", block)
	}

	fmt.Print("--\n")

	return blocks
}

func BlockContratos(blocks []Dict, token string, url string) {
	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	client := &http.Client{Transport: tr}
	auth := fmt.Sprintf("Basic %v", token)

	for _, block := range blocks {
		url := url+"cliente_contrato_15300/"
		contrato_json, error := json.Marshal(block)

		if error != nil {
			fmt.Println("A criação do seriamento json falhou.")
			os.Exit(1)
		}

		fmt.Printf("contrato_json: %s\n", string(contrato_json))
		req, error := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(contrato_json))

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
		fmt.Printf("--")
	}

}
