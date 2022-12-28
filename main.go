package main

import "time"

type Cloak struct {
	weight   string
	material string
}

// Monster содержит информацию об уроде
type Monster struct {
	body  string
	head  string
	eyes  string
	cloak Cloak
}

func NewMonster() *Monster {
	return &Monster{}
}

func (m *Monster) transform() {
	time.Sleep(time.Minute)
	m.shrinkBody()
	m.expandHead()
	m.goggle()
	m.getCloak()
}

func (m *Monster) shrinkBody() {
	m.body = "shrinked"
}

func (m *Monster) expandHead() {
	m.head = "expanded"
}

func (m *Monster) goggle() {
	m.eyes = "goggled"
}

func (m *Monster) getCloak() {
	m.cloak = Cloak{
		weight:   "heavy",
		material: "wood",
	}
}



type People interface {
	Rejoice()
}

type RoyalFamily struct {
	king     King
	queen    Queen
	princess Princess
}

type RoyalServants struct {
	Drosselmeier Drosselmeier
	MasterOfCeremonies MasterOfCeremonies
}

type Folk struct {
	CommonPeople []CommonPerson
}

type CommonPerson struct {
}

type Drosselmeier struct {
}

type MasterOfCeremonies struct {
}

type Rat struct {
	hair string
}

func (r *Rat)


func main() {

}
