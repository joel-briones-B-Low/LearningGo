package entities

type Persona struct {
	Nombre string
	Edad   int
}

func (p Persona) Saludar() string {
	return "Hola, mi nombre es " + p.Nombre
}

func (p *Persona) CumplirAnios() {
	p.Edad++
}
