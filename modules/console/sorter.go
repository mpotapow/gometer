package console

import "gometer/modules/console/contracts"

// CommandByName ...
type CommandByName []contracts.Command

func (a CommandByName) Len() int           { return len(a) }
func (a CommandByName) Less(i, j int) bool { return a[i].GetName() < a[j].GetName() }
func (a CommandByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// OptionByName ...
type OptionByName []contracts.Option

func (a OptionByName) Len() int           { return len(a) }
func (a OptionByName) Less(i, j int) bool { return a[i].GetName() < a[j].GetName() }
func (a OptionByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
