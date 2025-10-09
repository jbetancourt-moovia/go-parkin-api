package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func LoadEnvFile() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("No se pudo obtener el directorio de trabajo: %v", err)
	}

	envPaths := []string{
		filepath.Join(wd, ".env"),          // raíz del proyecto (cuando corres con go run)
		filepath.Join(wd, "..", ".env"),    // por si ejecutas desde /cmd
		filepath.Join(wd, "../..", ".env"), // fallback
	}

	log.Println("Buscando archivo .env en las siguientes rutas:")
	for _, path := range envPaths {
		log.Printf(" - %s", path)
	}

	var loadedPath string
	for _, path := range envPaths {
		if _, err := os.Stat(path); err == nil {
			if err := godotenv.Load(path); err == nil {
				loadedPath = path
				break
			}
		}
	}

	if loadedPath != "" {
		log.Printf("Archivo .env cargado desde: %s", loadedPath)
	} else {
		log.Println("No se encontró ningún archivo .env en las rutas esperadas")
	}
}
