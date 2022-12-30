package main

import (
	"fmt"
	"math/rand"
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

func NewMonster(transformationTime time.Duration) *Monster {
	m := Monster{}
	time.Sleep(transformationTime)
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

type King struct {
	xPosition int
}

func (k *King) rejoice() string {
	return k.jumpOnOneLeg()
}

func (k *King) jumpOnOneLeg() string {
	return "jump-jump!"
}

type Queen struct {
	xPosition int
}

func (q *Queen) rejoice() string {
	return q.faint()
}

func (q *Queen) faint() string {
	return "thud!"
}

type Princess struct {
	age       string
	beautiful bool
	cheeks    string
	eyes      string
	hair      string
	hasNut    bool
	monster   *Monster
	xPosition int
}

func (p *Princess) eatNutKernel() {
	p.hasNut = false
	p.monster = nil
}

type Courtiers struct {
	trumpet   string
	oboe      string
	xPosition int
}

func (c *Courtiers) rejoice(isLoud bool) {
	var (
		trumpetSound = "Tarantara! tarantara!"
		oboeSound    = "Ra, ra, ra, ra!"
	)
	if isLoud == true {
		trumpetSound = strings.ToUpper(trumpetSound)
		oboeSound = strings.ToUpper(oboeSound)
	}
	c.trumpet = trumpetSound
	c.oboe = oboeSound
}

type RoyalServants struct {
	Drosselmeier       Drosselmeier
	MasterOfCeremonies MasterOfCeremonies
}

type Folk struct {
	CommonPeople []CommonPerson
}

func (f *Folk) NewFolk() {
	commonPerson := CommonPerson{xPosition: 14}
	for i := 0; i < crowdSize; i++ {
		f.CommonPeople = append(f.CommonPeople, commonPerson)
	}
}

type CommonPerson struct {
	xPosition int
}

func (cp *CommonPerson) rejoice() string {
	return fmt.Sprintf("Hooray! ")
}

type Drosselmeier struct {
	age          string
	bodyPosition string
	eyes         string
	xPosition    int
	hasNut       bool
	monster      *Monster
}

func (d *Drosselmeier) crackNut(n *Nut) string {
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

func (d *Drosselmeier) backAway(steps int) string {
	d.xPosition += steps
	return "click!"
}

func (d *Drosselmeier) takeNut(m *MasterOfCeremonies) {
	m.hasNut = false
	d.hasNut = true
}

func (d *Drosselmeier) giveNut(p *Princess) {
	d.hasNut = false
	p.hasNut = true
}

func (d *Drosselmeier) stepOnRat(r *Rat) {
	r.isHurting = true
}

type MasterOfCeremonies struct {
	hasNut    bool
	xPosition int
}

type Rat struct {
	hair          string
	xPosition     int
	isUnderground bool
	isHurting     bool
}

func (r *Rat) whistleAndHiss(isLoud bool) (string, string) {
	var (
		whistle = "phheeew!"
		hiss    = "hissss!"
	)
	if isLoud == true {
		whistle = strings.ToUpper(whistle)
		hiss = strings.ToUpper(hiss)
	}
	return whistle, hiss
}

func main() {
	// объявление действующих лиц и предметов
	nut := Nut{
		name: "Krakatuk",
	}
	drosselmeier := Drosselmeier{
		age:          "young",
		bodyPosition: "stand up",
		eyes:         "open",
		xPosition:    0,
		hasNut:       false,
	}
	masterOfCeremonies := MasterOfCeremonies{
		hasNut:    true,
		xPosition: 0,
	}
	princess := Princess{
		age:       "young",
		beautiful: true,
		cheeks:    "like pink lilies",
		eyes:      "shiny like blue stars",
		hair:      "cute golden curls",
		monster:   NewMonster(0),
		hasNut:    false,
		xPosition: 0,
	}
	rat := Rat{
		hair:          "gray",
		xPosition:     7,
		isUnderground: true,
		isHurting:     false,
	}
	folk := Folk{}
	folk.NewFolk()
	courtiers := Courtiers{}

	// повествование
	drosselmeier.bow()
	drosselmeier.takeNut(&masterOfCeremonies)
	fmt.Println(drosselmeier.crackNut(&nut))
	drosselmeier.kneel()
	drosselmeier.giveNut(&princess)
	drosselmeier.eyes = "closed"
	fmt.Println(drosselmeier.backAway(1))
	princess.eatNutKernel()
	courtiers.rejoice(true)

	for drosselmeier.xPosition < 6 {
		fmt.Println(makeJoiceSound(&drosselmeier, &courtiers, &folk))
	}

	rat.isUnderground = false
	fmt.Println(rat.whistleAndHiss(true))
	drosselmeier.backAway(1)
	drosselmeier.stepOnRat(&rat)
	drosselmeier.monster = NewMonster(minute)

}

func makeJoiceSound(drosselmeier *Drosselmeier, courtiers *Courtiers, folk *Folk) string {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(4)
	var sound string
	switch n {
	case 0:
		sound = drosselmeier.backAway(1)
	case 1:
		sound = courtiers.trumpet
	case 2:
		sound = courtiers.trumpet
	case 3:
		sound = folk.CommonPeople[rand.Intn(crowdSize)].rejoice()
	}
	return sound
}
