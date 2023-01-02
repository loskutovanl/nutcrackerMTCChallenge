package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

const (
	second    = time.Second
	minute    = time.Minute
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

type Music struct {
	trumpet string
	oboe    string
}

func (m *Music) on(isLoud bool) {
	var (
		trumpetSound = "Tarantara! tarantara!"
		oboeSound    = "Ra, ra, ra, ra!"
	)
	if isLoud == true {
		trumpetSound = strings.ToUpper(trumpetSound)
		oboeSound = strings.ToUpper(oboeSound)
	}
	m.trumpet = trumpetSound
	m.oboe = oboeSound
}

// действующие лица

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
	xPosition   int
	isConscious bool
}

func (q *Queen) rejoice() string {
	return q.faint()
}

func (q *Queen) faint() string {
	q.isConscious = false
	return "thud!"
}

type Princess struct {
	age       string
	beautiful bool
	cheeks    string
	eyes      string
	hair      string
	nut       *Nut
	xPosition int
	monster   *Monster
}

func (p *Princess) eatNutKernel() {
	p.nut = nil
	p.monster = nil
}

type Courtiers struct {
	xPosition int
}

func (c *Courtiers) rejoice() string {
	return c.jumpOnOneLeg()
}

func (c *Courtiers) jumpOnOneLeg() string {
	return "jump-jump!"
}

type Folk struct {
	CommonPeople []CommonPerson
	xPosition    int
}

func (f *Folk) NewFolk() {
	commonPerson := CommonPerson{}
	for i := 0; i < crowdSize; i++ {
		f.CommonPeople = append(f.CommonPeople, commonPerson)
	}
}

func (f *Folk) rushToPrincess(steps int) {
	f.xPosition -= steps
}

type CommonPerson struct {
}

func (cp *CommonPerson) rejoice() string {
	return fmt.Sprintf("Hooray! ")
}

type Drosselmeier struct {
	age          string
	bodyPosition string
	eyes         string
	xPosition    int
	nut          *Nut
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

func (d *Drosselmeier) takeNut(m *MasterOfCeremonies, n *Nut) {
	m.nut = nil
	d.nut = n
}

func (d *Drosselmeier) giveNut(p *Princess, n *Nut) {
	d.nut = nil
	p.nut = n
}

func (d *Drosselmeier) stepOnRat(r *Rat) {
	r.isHurting = true
}

type MasterOfCeremonies struct {
	nut       *Nut
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
	king := King{
		xPosition: 0,
	}
	queen := Queen{
		xPosition:   0,
		isConscious: true,
	}
	princess := Princess{
		monster:   NewMonster(0),
		nut:       nil,
		xPosition: 0,
	}

	drosselmeier := Drosselmeier{
		age:          "young",
		bodyPosition: "stand up",
		eyes:         "open",
		xPosition:    0,
		nut:          nil,
	}
	nut := Nut{
		name:      "Krakatuk",
		unshelled: false,
	}
	masterOfCeremonies := MasterOfCeremonies{
		nut:       &nut,
		xPosition: 0,
	}

	drosselmeier.bow()
	drosselmeier.takeNut(&masterOfCeremonies, &nut)
	fmt.Println(drosselmeier.crackNut(&nut))
	drosselmeier.kneel()
	drosselmeier.giveNut(&princess, &nut)
	drosselmeier.eyes = "closed"
	fmt.Println(drosselmeier.backAway(1))

	princess.eatNutKernel()
	princess.beautiful = true
	princess.age = "young"
	princess.cheeks = "like pink lilies"
	princess.eyes = "shiny like blue stars"
	princess.hair = "cute golden curls"

	music := Music{
		trumpet: "",
		oboe:    "",
	}
	music.on(true)
	folk := Folk{xPosition: 12}
	folk.NewFolk()
	courtiers := Courtiers{}

	var wg sync.WaitGroup
	stepsLeftByDrosselmeier := 6
	stepWhenQueenFaints := rand.Intn(stepsLeftByDrosselmeier)
	for step := 1; step < stepsLeftByDrosselmeier; step++ {
		go func() {
			wg.Add(1)
			defer wg.Done()
			fmt.Println(drosselmeier.backAway(1))
		}()

		go func() {
			wg.Add(1)
			defer wg.Done()
			fmt.Println(music.trumpet)
			fmt.Println(music.oboe)
		}()

		go func() {
			wg.Add(1)
			defer wg.Done()
			fmt.Println(folk.CommonPeople[rand.Intn(crowdSize)].rejoice())
			folk.rushToPrincess(2)
		}()

		go func() {
			wg.Add(1)
			defer wg.Done()
			fmt.Println(king.rejoice())
			fmt.Println(courtiers.rejoice())
		}()

		if step == stepWhenQueenFaints {
			go func() {
				wg.Add(1)
				defer wg.Done()
				fmt.Println(queen.rejoice())
			}()
		}
	}

	wg.Wait()

	rat := Rat{
		hair:          "gray",
		xPosition:     7,
		isUnderground: true,
		isHurting:     false,
	}

	rat.isUnderground = false
	fmt.Println(rat.whistleAndHiss(true))
	drosselmeier.backAway(1)
	drosselmeier.stepOnRat(&rat)
	drosselmeier.monster = NewMonster(minute)
}
