package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"

	demoinfocs "github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs"
	events "github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs/events"
)

func main() {
	demoPath := `C:\Users\aldre\demo-parser\esports-world-cup-2025-vitality-vs-falcons-bo3-8ZTMZQ0BkOa0azICXTbCYv\vitality-vs-falcons-m1-inferno-p1.dem`
	outputFile := `C:\Users\aldre\demo-parser\resultado_detalhado.csv`

	// Cria arquivo CSV
	file, err := os.Create(outputFile)
	if err != nil {
		log.Panic("Erro ao criar arquivo CSV: ", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Cabeçalho CSV
	writer.Write([]string{
		"Round", "Time", "Killer", "KillerTeam", "KillerX", "KillerY",
		"Weapon", "Headshot", "Wallbang", "Victim", "VictimTeam", "VictimX", "VictimY",
	})

	// Evento de kill
	onKill := func(kill events.Kill) {
		hs := "No"
		if kill.IsHeadshot {
			hs = "Yes"
		}

		wb := "No"
		if kill.PenetratedObjects > 0 {
			wb = "Yes"
		}

		// Round e tempo
		round := strconv.Itoa(kill.Round)
		time := strconv.FormatFloat(kill.Time.Seconds(), 'f', 2, 64)

		// Posições e times
		killerX := strconv.FormatFloat(kill.Killer.Position().X, 'f', 2, 64)
		killerY := strconv.FormatFloat(kill.Killer.Position().Y, 'f', 2, 64)
		victimX := strconv.FormatFloat(kill.Victim.Position().X, 'f', 2, 64)
		victimY := strconv.FormatFloat(kill.Victim.Position().Y, 'f', 2, 64)

		killerTeam := kill.Killer.Team().String()
		victimTeam := kill.Victim.Team().String()

		writer.Write([]string{
			round,
			time,
			kill.Killer.String(),
			killerTeam,
			killerX,
			killerY,
			kill.Weapon.String(),
			hs,
			wb,
			kill.Victim.String(),
			victimTeam,
			victimX,
			victimY,
		})
	}

	// Parse da demo
	err = demoinfocs.ParseFile(demoPath, func(p demoinfocs.Parser) error {
		p.RegisterEventHandler(onKill)
		return nil
	})

	if err != nil {
		log.Panic("Falha ao parsear demo: ", err)
	}

	log.Println("CSV detalhado gerado com sucesso em:", outputFile)
}
