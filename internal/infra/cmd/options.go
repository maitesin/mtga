package cmd

type Options struct {
	Number   int    `short:"n" long:"number" description:"Number of the card in the set"`
	Set      string `short:"s" long:"set" description:"Magic the Gathering set of the card"`
	Quantity int    `short:"q" long:"quantity" default:"1" description:"Number of copies of that card"`
	Foil     bool   `short:"f" long:"foil" default:"false" description:"Card is foil"`
	Altered  bool   `short:"a" long:"altered" default:"false" description:"Card is altered"`
}
