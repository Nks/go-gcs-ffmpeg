package models

type Mediafile struct {
	metadata    Metadata
	inputFile   string
	outputFile  string
	preCommands []string
	commands    []string
}

/*** SETTERS ***/
func (m *Mediafile) SetMetadata(v Metadata) {
	m.metadata = v
}

func (m *Mediafile) AddCommand(v string, before bool) {
	if before == true {
		m.commands = append([]string{v}, m.commands...)
	} else {
		m.commands = append(m.commands, v)
	}

}

func (m *Mediafile) AddInputFile(v string) {
	m.inputFile = v
}

func (m *Mediafile) AddOutputFile(v string) {
	m.outputFile = v
}

/*** GETTERS ***/
func (m *Mediafile) Metadata() Metadata {
	return m.metadata
}

func (m *Mediafile) Commands() []string {
	return m.commands
}

func (m *Mediafile) ToString() []string {
	cmd := append(m.preCommands, "-i", m.inputFile)
	cmd = append(cmd, m.commands...)

	return cmd
}
