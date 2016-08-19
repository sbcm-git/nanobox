package dev

import (
	"github.com/nanobox-io/nanobox/processors"
)

// Up ...
type Up struct {
}

//
func (up Up) Run() error {

	// run a nanobox start
	processorStart := processors.Start{}
	if err := processorStart.Run(); err != nil {
		return err
	}

	// run a nanobox build
	processorBuild := &processors.Build{}
	if err := processorBuild.Run(); err != nil {
		return err
	}

	// run a dev start
	devStart := &Start{Env: processorBuild.Env}
	if err := devStart.Run(); err != nil {
		return err
	}

	// run a dev deploy
	devDeploy := Deploy{Env: processorBuild.Env, App: devStart.App}
	if err := devDeploy.Run(); err != nil {
		return err
	}

	return nil
}