package player

 

type Player interface {

GetMark() string
GetName() string

}

//cannot use detail1 (variable of type *"contactapp/contactdetails".ContactDetails) as *"contactapp/contactDetails".ContactDetails value in argument to appendcompilerIncompatibleAssign
//var detail1 *contactdetails.ContactDetails

 

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
