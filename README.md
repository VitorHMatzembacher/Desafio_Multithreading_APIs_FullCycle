# Desafio de Concorrência com APIs - FullCycle

## Descrição

Este projeto foi desenvolvido como parte da pós-graduação FullCycle e tem como objetivo aplicar conceitos de:

Concorrência em Go (`goroutines` + `select`)
Controle de tempo com `context.WithTimeout`
Consumo de APIs públicas

O programa consulta dois serviços distintos de CEP em paralelo e exibe no terminal a resposta **mais rápida** entre eles.

## Funcionalidade

Dadas duas APIs que retornam dados de endereço com base em um CEP:

[`BrasilAPI`](https://brasilapi.com.br/api/cep/v1/{cep})
[`ViaCEP`](https://viacep.com.br/ws/{cep}/json/)

O programa faz chamadas simultâneas para ambas e exibe:
A resposta da **API mais rápida**
Os **dados do endereço** retornados
Qual **API respondeu primeiro**

### Timeout

O tempo máximo para obter a resposta é de **1 segundo**.  
Se nenhuma das APIs responder nesse tempo, será exibido um erro de timeout.

### Comando de Execucao 

go run main.go
