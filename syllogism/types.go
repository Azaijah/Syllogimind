package syllogism

type Syllogism struct {
	Entities []Entity
}

type Entity struct {
	Text string
}

type syllogismSystem struct {
	WordList []string `yaml:"list"`
}
