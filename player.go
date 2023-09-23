package player

 

type Player interface {

GetMark() string
GetName() string

}

 

type HumanPlayer struct {

Mark string
Name string

}

 

func NewHumanPlayer(mark,name string) *HumanPlayer {

return &HumanPlayer{Mark: mark, Name:name}

}

 

func (p *HumanPlayer) GetMark() string {

return p.Mark

}

func (p *HumanPlayer) GetName() string {

	return p.Name
}