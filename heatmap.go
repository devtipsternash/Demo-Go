package main

import (
	"encoding/csv"
	"fmt"
	"os"

	ex "github.com/markus-wa/demoinfocs-golang/v5/examples"
	demoinfocs "github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs"
	events "github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs/events"
)

func main() {
	// Abrir demo
	f, err := os.Open(ex.DemoPathFromArgs())
	checkError(err)
	defer f.Close()

	parser := demoinfocs.NewParser(f)
	defer parser.Close()

	// Criar CSV
	csvFile, err := os.Create("demo_data.csv")
	checkError(err)
	defer csvFile.Close()
	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	// Cabeçalho
	writer.Write([]string{"Shooter", "X", "Y", "Z", "Weapon"})

	// Registrar evento WeaponFire
	parser.RegisterEventHandler(func(e events.WeaponFire) {
		pos := e.Shooter.Position()
		writer.Write([]string{
			e.Shooter.Name,
			fmt.Sprintf("%.2f", pos.X),
			fmt.Sprintf("%.2f", pos.Y),
			fmt.Sprintf("%.2f", pos.Z),
			e.Weapon.String(),
		})
	})

	// Parse demo
	err = parser.ParseToEnd()
	checkError(err)

	fmt.Println("CSV gerado com sucesso: demo_data.csv")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
