package main

import (
	"fmt"
	"strings"
	"time"
)

const (
	second    = time.Second
	minute    = time.Second
	crowdSize = 1000
)

// Monster содержит информацию об уроде
type Monster struct {
	body  string
	head  string
	eyes  string
	cloak Cloak
}

func NewMonster() *Monster {
	m := Monster{}
	time.Sleep(minute)
	m.shrinkBody()
	m.expandHead()
	m.goggle()
	m.getCloak()
	return &m
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

// Предметы неодушевленные

type Cloak struct {
	weight   string
	material string
}

type Nut struct {
	name      string
	unshelled bool
}

// Люди

type People interface {
	Rejoice()
}

type King struct {
}

type Queen struct {
}

type Princess struct {
	age       string
	beautiful bool
	cheeks    string
	eyes      string
	hair      string
	hasNut    bool
	monster   *Monster
}

func (p *Princess) eatNutKernel() {
	p.hasNut = false
	p.monster = nil
}

type Courtiers struct {
	trumpet string
	oboe    string
}

func (c *Courtiers) Rejoice() {
	c.trumpet = strings.ToUpper("Tarantara! tarantara!")
	c.oboe = strings.ToUpper("Ra, ra, ra, ra!")
}

type RoyalServants struct {
	Drosselmeier       Drosselmeier
	MasterOfCeremonies MasterOfCeremonies
}

type Folk struct {
	CommonPeople []CommonPerson
}

func (f *Folk) NewFolk() {
	commonPerson := CommonPerson{}
	for i := 0; i < crowdSize; i++ {
		f.CommonPeople = append(f.CommonPeople, commonPerson)
	}
}

type CommonPerson struct {
}

func (cp *CommonPerson) Rejoice() string {
	return "Hooray!"
}

type Drosselmeier struct {
	age          string
	bodyPosition string
	eyes         string
	stepsCount   int
	hasNut       bool
}

func (d *Drosselmeier) crackNut(n Nut) string {
	n.unshelled = true
	return "crack!"
}

func (d *Drosselmeier) kneel() {
	d.bodyPosition = "loyally kneeled"
	time.Sleep(second)
	d.bodyPosition = "stand up"
}

func (d *Drosselmeier) bow() {
	d.bodyPosition = "bowed politely"
	time.Sleep(second)
	d.bodyPosition = "stand up"
}

func (d *Drosselmeier) closeEyes() {
	d.eyes = "closed"
}

func (d *Drosselmeier) backAway(steps int) {
	d.stepsCount += steps
}

func (d *Drosselmeier) takeNut(m MasterOfCeremonies) {
	m.hasNut = false
	d.hasNut = true
}

func (d *Drosselmeier) giveNut(p Princess) {
	d.hasNut = false
	p.hasNut = true
}

type MasterOfCeremonies struct {
	hasNut bool
}

type Rat struct {
	hair string
}

func (r *Rat) whistle() {
}

func (r *Rat) hiss() {

}

func main() {
	nut := Nut{
		name: "Krakatuk",
	}
	drosselmeier := Drosselmeier{
		age:          "young",
		bodyPosition: "stand up",
		eyes:         "open",
	}
	masterOfCeremonies := MasterOfCeremonies{
		hasNut: true,
	}
	princess := Princess{
		age:       "young",
		beautiful: true,
		cheeks:    "like pink lilies",
		eyes:      "shiny like blue stars",
		hair:      "cute golden curls",
		monster:   NewMonster(),
	}
	folk := Folk{}
	folk.NewFolk()
	courtiers := Courtiers{}

	fmt.Println(drosselmeier)
	fmt.Println(masterOfCeremonies)
	fmt.Println(nut)
	fmt.Println(princess.monster)

	drosselmeier.bow()
	drosselmeier.takeNut(masterOfCeremonies)
	fmt.Println(drosselmeier.crackNut(nut))
	drosselmeier.kneel()
	drosselmeier.giveNut(princess)
	drosselmeier.eyes = "closed"
	drosselmeier.backAway(1)
	princess.eatNutKernel()
	drosselmeier.backAway(1)
	courtiers.Rejoice()
	fmt.Println(courtiers.trumpet)
	for _, person := range folk.CommonPeople {
		fmt.Println(person.Rejoice())
	}
	fmt.Println(courtiers.oboe)
	drosselmeier.backAway(1)

}
