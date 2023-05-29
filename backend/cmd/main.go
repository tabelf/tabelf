package main

import "tabelf/backend/cmd/actions"

func main() {
	actions.InitDefaultLogger()
	actions.Execute()
}
