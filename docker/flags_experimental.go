// +build experimental

package main

func initExperimental() {
	dockerCommands = append(dockerCommands, command{"network", "Network management"})
}
