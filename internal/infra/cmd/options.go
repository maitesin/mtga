package cmd

type Options struct {
	Number    int    `short:"n" long:"number" description:"Number of the card in the set"`
	Set       string `short:"s" long:"set" description:"Magic the Gathering set of the card"`
	Quantity  int    `short:"q" long:"quantity" default:"1" description:"Number of copies of that card"`
	Foil      bool   `short:"f" long:"foil" description:"Foil card"`
	Altered   bool   `short:"a" long:"altered" description:"Altered card"`
	Signed    bool   `short:"i" long:"signed" description:"Signed card"`
	Condition string `short:"c" long:"condition" default:"nm" description:"Card's condition'"`
	Language  string `short:"l" long:"language" default:"en" description:"Card's language'"`
}
