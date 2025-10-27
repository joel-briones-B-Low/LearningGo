package domain

import (
	"learningGo/domain/entities"
)

func App() {
	persona := entities.Persona{Nombre: "Juan", Edad: 30}
	tipo := entities.TipoPersona{Nombre: "Estudiante"}

	persona.CumplirAnios()
	println(persona.Saludar())
	elemento := tipo.Nombre
	println("Tipo de persona:", elemento)
}
