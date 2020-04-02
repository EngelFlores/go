package main

import (
	"bufio"
	"os"
	"strconv"
	"encoding/json"
)

type Relatorio struct {
	Client_id            string `json:"client_id "`
	Employee_name string `json:"employee_name"`
	Description   string `json:"description"`
}

func criarRelatorio(firstParam string, secondParam string, field string, id int, name string, relatorio []Relatorio) []Relatorio {
	var description string
	var convertedId = strconv.Itoa(id)
	if firstParam != secondParam {
		description = "The " + field + " is not the same in the client database and our database"
		relatorio = append(relatorio, Relatorio{
			Client_id :            convertedId,
			Employee_name: name,
			Description:   description,
		})
	}
	return relatorio
}

func read(clienteData []Data, ourData []Data) {
	var relatorio []Relatorio

	for _, dadoCliente := range clienteData {
		for _, dadoNosso := range ourData {
			if dadoNosso.Id == dadoCliente.Id && dadoNosso.Employee_name == dadoCliente.Employee_name {
				relatorio = criarRelatorio(dadoNosso.Medical_plan, dadoCliente.Medical_plan, "medical plan", dadoCliente.Id, dadoCliente.Employee_name, relatorio)
				relatorio = criarRelatorio(dadoNosso.Dental_plan, dadoCliente.Dental_plan, "dental plan", dadoCliente.Id, dadoCliente.Employee_name, relatorio)
				relatorio = criarRelatorio(dadoNosso.Language, dadoCliente.Language, "language", dadoCliente.Id, dadoCliente.Employee_name, relatorio)
				relatorio = criarRelatorio(dadoNosso.Claimant_name, dadoCliente.Claimant_name, "claimant name", dadoCliente.Id, dadoCliente.Employee_name, relatorio)
				relatorio = criarRelatorio(dadoNosso.Relationship_type, dadoCliente.Relationship_type, "relationship type", dadoCliente.Id, dadoCliente.Employee_name, relatorio)
				relatorio = criarRelatorio(dadoNosso.Gender, dadoCliente.Gender, "gender", dadoCliente.Id, dadoCliente.Employee_name, relatorio)
				relatorio = criarRelatorio(dadoNosso.Effective_date, dadoCliente.Effective_date, "effective date", dadoCliente.Id, dadoCliente.Employee_name, relatorio)
				relatorio = criarRelatorio(dadoNosso.Termination_date, dadoCliente.Termination_date, "termination date", dadoCliente.Id, dadoCliente.Employee_name, relatorio)
			}
		}
		writeFile(relatorio)

	}

}

func writeFile(relatorio []Relatorio) {
	pwd, _ := os.Getwd()
	resultado, _ := json.Marshal(relatorio)

	file, _ := os.Create(pwd + "/temp-file/relatorio.json")
	writer := bufio.NewWriter(file)
	writer.WriteString(string(resultado))
	writer.Flush()

}
