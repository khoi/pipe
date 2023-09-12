package main

import (
	"context"
	"errors"

	"github.com/khoi/pipe/manifest"
)

// App struct
type App struct {
	ctx       context.Context
	manifests map[string]*manifest.Manifest
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.manifests = make(map[string]*manifest.Manifest)
	for _, m := range a.ListManifests() {
		a.manifests[m.ID] = m
	}
}

func (a *App) ListManifests() []*manifest.Manifest {
	return manifest.ListManifests(fixtures)
}

func (a *App) RunManifest(id string, input *string) (string, error) {
	m, ok := a.manifests[id]
	if !ok {
		return "", errors.New("manifest not found")
	}
	return m.Execute(a.ctx, input)
}
