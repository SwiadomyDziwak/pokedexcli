package main

func main() {
	var conf config
	conf.commands = loadCommands()
	startREPL(&conf)
}
