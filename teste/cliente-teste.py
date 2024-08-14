from almondegas import Test

cliente_obrigatorio_site = {
    "ativo",
    "tipo_pessoa",
    "razao",
    "contribuinte_icms",
    "endereco",
    "numero",
    "bairro",
    "cidade",
    "tipo_localidade",
    "iss_classificacao_padrao",
    "dica_imposto_retido_cliente",
}

cliente_obrigatorio = {
    "ativo",
    "tipo_pessoa",
    "razao",
    "contribuinte_icms",
    "endereco",
    "numero",
    "bairro",
    "cidade",
    "tipo_localidade",
    "iss_classificacao_padrao",
    "dica_imposto_retido_cliente",

    "cep",
    "tipo_assinante",
    "email",
    "hotsite_email",
    "hotsite_acesso",
    "senha_hotsite_md5",
    "crm_data_novo",
    "id_tipo_cliente"
}

cliente_obrigatorio_e_datas = {
    "ativo",
    "tipo_pessoa",
    "razao",
    "contribuinte_icms",
    "endereco",
    "numero",
    "bairro",
    "cidade",
    "tipo_localidade",
    "iss_classificacao_padrao",
    "dica_imposto_retido_cliente",

    "cep",
    "tipo_assinante",

    "data_nascimento",
    "data_cadastro",
    "crm_data_novo",
    "crm_data_sondagem",
    "crm_data_apresentando",
    "crm_data_negociando",
    "crm_data_vencemos",
    "crm_data_perdemos",
    "crm_data_abortamos",
    "crm_data_sem_porta_disponivel",
    "crm_data_sem_viabilidade",
    "data_nascimento_conjuge",
    "emp_data_admissao",
    "emp_data_admissao",
}

cliente_defaults = {
    'ativo': 'S',
    'id_tipo_cliente': '',
    'tipo_cliente_scm': '01',
    'pais': 'Brasil',
    'tipo_pessoa': 'F',
    'razao': '',
    'fantasia': '',
    'cnpj_cpf': '',
    'ie_identidade': '',
    'contribuinte_icms': 'N',
    'rg_orgao_emissor': '',
    'nacionalidade': 'Brasileiro',
    'cidade_naturalidade': '',
    'estado_nascimento': '',
    'data_nascimento': '',
    'Sexo': '',
    'profissao': '',
    'estado_civil': '',
    'inscricao_municipal': '',
    'isuf': '',
    'tipo_assinante': '3',
    'filial_id': '',
    'filtra_filial': 'S',
    'idx': '',
    'data_cadastro': '',
    'ativo_serasa': '',
    'convert_cliente_forn': 'N',
    'atualizar_cadastro_galaxPay': 'N',
    'grau_satisfacao': '',
    'id_condominio': '',
    'bloco': '',
    'apartamento': '',
    'cep': '',
    'cif': '',
    'endereco': '',
    'numero': '',
    'complemento': '',
    'bairro': '',
    'cidade': '',
    'referencia': '',
    'uf': '1',
    'moradia': '',
    'tipo_localidade': 'U',
    'latitude': '',
    'longitude': '',
    'cep_cob': '',
    'endereco_cob': '',
    'numero_cob': '',
    'bairro_cob': '',
    'cidade_cob': '',
    'complemento_cob': '',
    'referencia_cob': '',
    'uf_cob': '',
    'fone': '',
    'telefone_comercial': '',
    'ramal': '',
    'id_operadora_celular': '',
    'telefone_celular': '',
    'whatsapp': '',
    'email': '',
    'email_opa': '',
    'contato': '',
    'website': '',
    'skype': '',
    'facebook': '',
    'hotsite_email': '',
    'senha': '',
    'acesso_automatico_central': 'P',
    'alterar_senha_primeiro_acesso': 'P',
    'senha_hotsite_md5': 'N',
    'hotsite_acesso': '2',
    'crm': 'N',
    'id_candato_tipo': '',
    'id_campanha': '',
    'id_concorrente': '',
    'id_perfil': '',
    'responsavel': '',
    'indicado_por': '',
    'cadastrado_via_viabilidade': 'N',
    'status_prospeccao': 'C',
    'crm_data_novo': '',
    'crm_data_sondagem': '',
    'crm_data_apresentando': '',
    'crm_data_negociando': '',
    'crm_data_vencemos': '',
    'crm_data_perdemos': '',
    'crm_data_abortamos': '',
    'crm_data_sem_porta_disponivel': '',
    'crm_data_sem_viabilidade': '',
    'pipe_id_organizacao': '',
    'foto_cartao': '',
    'participa_cobranca': 'S',
    'num_dias_cob': '',
    'participa_pre_cobranca': 'S',
    'cob_envia_email': 'S',
    'cob_envia_sms': 'S',
    'id_conta': '',
    'cond_pagamento': '',
    'id_vendedor': '',
    'tabela_preco': '',
    'deb_automatico': '',
    'deb_agencia': '',
    'deb_conta': '',
    'codigo_operacao': '',
    'tipo_pessoa_titular_conta': 'F',
    'cnpj_cpf_titular_conta': '',
    'ultima_atualizacao': 'CURRENT_TIMESTAMP',
    'nome_pai': '',
    'cpf_pai': '',
    'identidade_pai': '',
    'nascimento_pai': '',
    'nome_mae': '',
    'cpf_mae': '',
    'identidade_mae': '',
    'nascimento_mae': '',
    'quantidade_dependentes': '',
    'nome_conjuge': '',
    'fone_conjuge': '',
    'cpf_conjuge': '',
    'rg_conjuge': '',
    'data_nascimento_conjuge': '',
    'nome_contador': '',
    'telefone_contador': '',
    'orgao_publico': 'N',
    'im': '',
    'nome_representante_1': '',
    'cpf_representante_1': '',
    'identidade_representante_1': '',
    'nome_representante_2': '',
    'cpf_representante_2': '',
    'identidade_representante_2': '',
    'emp_empresa': '',
    'emp_cnpj': '',
    'emp_cep': '',
    'emp_endereco': '',
    'emp_cidade': '',
    'emp_fone': '',
    'emp_cargo': '',
    'emp_remuneracao': '',
    'emp_data_admissao': '',
    'iss_classificacao_padrao': '99',
    'pis_retem': 'S',
    'cofins_retem': 'S',
    'csll_retem': 'S',
    'irrf_retem': 'S',
    'desconto_irrf_valor_inferior': 'N',
    'cli_desconta_iss_retido_total': 'S',
    'dica_imposto_retido_cliente': '',
    'ref_com_empresa1': '',
    'ref_com_fone1': '',
    'ref_com_empresa2': '',
    'ref_com_fone2': '',
    'ref_pes_nome1': '',
    'ref_pes_fone1': '',
    'ref_pes_nome2': '',
    'ref_pes_fone2': '',
    'obs': '',
    'alerta': ''
}

if __name__ == "__main__":
    teste = Test("cliente", "4")

    print("testando repor normalizado")
    print("--")
    teste.update(teste.fetched[0])
    teste.isequal()

    print("testando campos obrigatórios conforme documentação")
    print("--")
    obj = teste.filter(cliente_obrigatorio_site)
    print(obj)
    teste.update(obj)
    teste.isequal()

    print("testando campos obrigatórios")
    print("--")
    obj = teste.filter(cliente_obrigatorio)
    print(obj)
    teste.update(obj)
    teste.isequal()

    print("testando campos obrigatórios com datas")
    print("--")
    obj = teste.filter(cliente_obrigatorio_e_datas)
    print(obj)
    teste.update(obj)
    teste.isequal()

    print("testando pôr valores padrão")
    print("--")
    teste.update(cliente_defaults)
    teste.isequal()

    print("repôndo valores prévios")
    print("--")
    teste.undo()
