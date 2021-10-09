package tournament

import (
	"fmt"
	"io"
	"io/ioutil"
	"sort"
	"strings"
)

type Team struct {
	name           string
	mp, w, d, l, p int
}

func NewTeam(name string) *Team {
	return &Team{
		name: name,
	}
}

func (t *Team) Win() {
	t.mp++
	t.w++
	t.p += 3
}

func (t *Team) Lose() {
	t.mp++
	t.l++
}

func (t *Team) Draw() {
	t.mp++
	t.d++
	t.p++
}

func (t Team) String() string {
	formattedName := t.name + strings.Repeat(" ", 31-len(t.name))
	return fmt.Sprintf("%s|  %d |  %d |  %d |  %d |  %d", formattedName, t.mp, t.w, t.d, t.l, t.p)
}

type Table struct {
	standings []*Team
	teams     map[string]*Team
}

func NewTournament() *Table {
	return &Table{
		standings: []*Team{},
		teams:     make(map[string]*Team),
	}
}

func (t *Table) AddTeam(name string) *Team {
	team := NewTeam(name)
	t.standings = append(t.standings, team)
	t.teams[name] = team
	return team
}
func (t *Table) GetTeam(team string) *Team {
	if _, ok := t.teams[team]; !ok {
		return t.AddTeam(team)
	}
	return t.teams[team]
}

func (Table) Match(t1, t2 *Team, result string) error {
	switch result {
	case "win":
		t1.Win()
		t2.Lose()
		return nil
	case "loss":
		t1.Lose()
		t2.Win()
		return nil
	case "draw":
		t1.Draw()
		t2.Draw()
		return nil
	}
	return fmt.Errorf("invalid result")
}

func (t Table) Len() int {
	return len(t.standings)
}

func (t Table) Less(i, j int) bool {
	if t.standings[i].p == t.standings[j].p {
		return t.standings[i].name < t.standings[j].name
	}
	return t.standings[i].p > t.standings[j].p
}

func (t Table) Swap(i, j int) {
	t.standings[i], t.standings[j] = t.standings[j], t.standings[i]
}

func (t Table) String() string {
	s := "Team                           | MP |  W |  D |  L |  P\n"
	for _, v := range t.standings {
		s += v.String() + "\n"
	}
	return s
}

func Tally(reader io.Reader, buffer io.Writer) error {
	bytesArray, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}
	tmt := NewTournament()
	text := strings.TrimSpace(string(bytesArray))
	matches := strings.Split(text, "\n")
	for _, m := range matches {
		m = strings.TrimSpace(m)
		if m == "" || string(m[0]) == "#" {
			continue
		}
		match := strings.Split(m, ";")
		if len(match) != 3 {
			return fmt.Errorf("invalid input %s", m)
		}
		t1, t2, res := match[0], match[1], match[2]
		team1 := tmt.GetTeam(t1)
		team2 := tmt.GetTeam(t2)
		err := tmt.Match(team1, team2, res)
		if err != nil {
			return err
		}
	}
	sort.Sort(tmt)
	buffer.Write([]byte(tmt.String()))
	return nil
}
