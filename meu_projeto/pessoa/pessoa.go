package pessoa

// Definindo uma struct
type Pessoa struct {
    Nome  string
    Idade int
    Cidade string
}

// Método que modifica a struct usando um ponteiro
func (p *Pessoa) FazerAniversario() {
    p.Idade++
}
