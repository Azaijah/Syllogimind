package syllogism

type relationType string
type statementType string

const (
	moreThan relationType = "is more than"
	lessThan relationType = "is less than"
)

const (
	sameR statementType = "has the same relation as"
	diffR statementType = "has a different relation from"
)

type entity struct {
	text string
}

type premise struct {
	entity1    entity
	entity2    entity
	relation   relationship
	isInverted bool
}

type conclusion struct {
	statement1    *relationship
	statement2    *relationship
	statementType statementType
	isInverted    bool
	outcome       bool
}

type relationship struct {
	from *entity
	to   *entity
	text relationType
}

type Syllogism struct {
	premises   []premise
	conclusion conclusion
}
