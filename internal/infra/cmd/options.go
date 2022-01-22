package cmd

type Options struct {
	Add    Add    `command:"add"`
	Export Export `command:"export"`
	Update Update `command:"update"`
}
