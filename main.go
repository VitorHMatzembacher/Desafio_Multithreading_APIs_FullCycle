package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Resultado struct {
	API      string
	Conteudo map[string]interface{}
	Err      error
}

func main() {
	cep := "01153000"
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	resposta := make(chan Resultado, 2)

	go buscarNaBrasilAPI(ctx, cep, resposta)
	go buscarNoViaCEP(ctx, cep, resposta)

	select {
	case resultado := <-resposta:
		if resultado.Err != nil {
			fmt.Println("Erro:", resultado.Err)
			return
		}
		fmt.Printf("Resposta da API %s:\n", resultado.API)
		for k, v := range resultado.Conteudo {
			fmt.Printf("%s: %v\n", k, v)
		}
	case <-ctx.Done():
		fmt.Println("Erro: Timeout após 1 segundo")
	}
}

func buscarNaBrasilAPI(ctx context.Context, cep string, resposta chan Resultado) {
	url := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)
	enviarRequest(ctx, url, "brasilapi", resposta)
}

func buscarNoViaCEP(ctx context.Context, cep string, resposta chan Resultado) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
	enviarRequest(ctx, url, "viacep", resposta)
}

func enviarRequest(ctx context.Context, url, api string, resposta chan Resultado) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		resposta <- Resultado{API: api, Err: err}
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		resposta <- Resultado{API: api, Err: err}
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		resposta <- Resultado{API: api, Err: err}
		return
	}

	var conteudo map[string]interface{}
	err = json.Unmarshal(body, &conteudo)
	if err != nil {
		resposta <- Resultado{API: api, Err: err}
		return
	}

	resposta <- Resultado{
		API:      api,
		Conteudo: conteudo,
		Err:      nil,
	}
}
