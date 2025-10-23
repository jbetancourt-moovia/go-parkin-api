package helpers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

func FetchExternalData[T any](method string, url string, payload any) (*T, error) {
	// Convertir a JSON
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	// Crear el cliente HTTP y hacer la solicitud
	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("error al comunicar con el servicio externo")
	}

	// Decodificar la respuesta en el tipo genérico
	var result T
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
