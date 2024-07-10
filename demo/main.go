package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	apiKey        = "YOUR_IBM_CLOUD_API_KEY" // Substitua com sua API Key
	authEndpoint  = "https://iam.cloud.ibm.com/identity/token"
	quantumAPIURL = "https://cloud.ibm.com/apidocs/quantum-computing" // URL base para a API
)

func main() {
	// Autenticação e obtenção do token de acesso
	token, err := getAuthToken()
	if err != nil {
		log.Fatalf("Erro ao obter token: %v", err)
	}

	// Fazendo uma requisição à API de computação quântica da IBM
	err = makeQuantumAPIRequest(token)
	if err != nil {
		log.Fatalf("Erro ao fazer requisição à API: %v", err)
	}
}

func getAuthToken() (string, error) {
	data := "grant_type=urn:ibm:params:oauth:grant-type:apikey&apikey=" + apiKey
	req, err := http.NewRequest("POST", authEndpoint, bytes.NewBufferString(data))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}

	token, ok := result["access_token"].(string)
	if !ok {
		return "", fmt.Errorf("token de acesso não encontrado na resposta")
	}

	return token, nil
}

func makeQuantumAPIRequest(token string) error {
	// Exemplo de requisição para a API
	req, err := http.NewRequest("GET", quantumAPIURL+"/your-endpoint", nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("Resposta da API: %s\n", string(body))
	return nil
}
