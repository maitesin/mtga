package cmd

type Options struct {
	//Verbose []bool `short:"v" long:"verbose" description:"Show verbose debug information"`
	Number   int    `short:"n" long:"number" description:"Number of the card in the set"`
	Set      string `short:"s" long:"set" description:"Magic the Gathering set of the card"`
	Quantity int    `short:"q" long:"quantity" default:"1" description:"Number of copies of that card"`
}
