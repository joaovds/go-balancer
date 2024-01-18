package flags

import "flag"

type Flags struct {
  Verbose *bool
  Help *bool
}

func SetupFlags() *Flags {
  flags := Flags{
    Verbose: flag.Bool("verbose", false, "verbose output, will print logs"),
    Help: flag.Bool("help", false, "print help"),
  }

  flag.Parse()

  return &flags
}

