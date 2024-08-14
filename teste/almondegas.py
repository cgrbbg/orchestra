import base64
from typing import Dict
import requests
import json
import urllib3

urllib3.disable_warnings(urllib3.exceptions.InsecureRequestWarning)
base_token = "868:b33c00f955056de28940fd9ade79798c54064939a8be9a1b3d09989aebd3ed5f".encode('utf-8')

class Test:
    token = 'Basic {}'.format(base64.b64encode(base_token).decode('utf-8'))
    error = False
    params = { 'qtype': '', 'query': '', 'oper': '==', 'page': '1' }
    headersl = { 'ixcsoft': 'listar', 'Content-Type': 'application/json', 'Authorization': token }
    headersd = { 'ixcsoft': '', 'Content-Type': 'application/json', 'Authorization': token }
    fetched = []

    def __init__(self, rota: str, id: str) -> None:
        url = "https://172.31.255.155/webservice/v1/"

        self.id = id
        self.rota = rota
        self.params["query"] = id
        self.url = url+rota

        if not id.isnumeric(): 
            print("você fez cringe")
            exit(1)
            
        if rota == "cliente":
            self.params["qtype"] = "cliente.id"
        else:
            print("você fez cringe")
            exit(1)

        self.fetched.append(self.fetch())

        with open(self.rota+"_teste.txt", mode="w") as file:
            file.write(json.dumps(self.fetched[0]))
            file.write("\n")
            file.close()

    def normalize(self, registro: Dict):
        for k in registro.keys():
            if registro[k] in ["0", "0000-00-00"]: 
                registro[k] = ""
            elif registro[k].count('-') == 2:
                dates = registro[k].split('-')
                registro[k] = dates[2]+"/"+dates[1]+"/"+dates[0]

        return registro

    def fetch(self) -> Dict:
        response = requests.post(
            self.url, 
            data=json.dumps(self.params), 
            headers=self.headersl, 
            verify=False)
        response = response.json()
        registro = response['registros'][0]

        return self.normalize(registro)

    def update(self, data: Dict[str, str]):
        response = requests.put(
            self.url+"/"+self.id, 
            data=json.dumps(data), 
            headers=self.headersd, 
            verify=False)
        print(response.text)
        print("--")

        if response.text.find("error") != -1:
            self.error = True
        else:
            self.error = False
            self.fetched.append(self.fetch())


    def filter(self, cols:set[str]) -> Dict[str, str]:
        dic: Dict[str, str] = {}

        for col in cols:
            if col in self.fetched[0].keys():
                dic[col] = self.fetched[0][col]

        return dic

    def isequal(self):
        if len(self.fetched) < 2:
            print("estúpido")
            print("--")
            return None

        if self.error:
            print("estado em erro")
            print("")
            return None

        errs = 0
        for keys in self.fetched[0].keys():
            prev_val = self.fetched[0][keys]
            new_val = self.fetched[len(self.fetched)-1][keys]
            stimmt =  prev_val == new_val
            if not stimmt:
                print("{} incorreto - veio {} e {}".format(keys, prev_val, new_val))
                errs = errs + 1

        if errs == 0: 
            print("todos os campos vieram corretos")
        print("")

    def undo(self):
        self.update(self.fetched[0])
