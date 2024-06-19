// A generated module for Go01 functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"
)

type Go01 struct{}

func (m *Go01) TestAll(ctx context.Context, source *Directory) (string, error) {
	result, err := m.Build(ctx, source)
	if err != nil {
		return "", err
	}
	
	return result, nil
}

// Returns a container that echoes whatever string argument is provided
func (m *Go01) Test(ctx context.Context, source *Directory) *Container {
	result := m.BuildEnv(source).
		WithExec([]string{"go", "test", "./...", "-v"}).
		WithExec([]string{"go", "mod", "verify"})
	return result
}

func (m *Go01) Lint(ctx context.Context, source *Directory) *Container {
	return m.Test(ctx, source).
		WithExec([]string{"go", "install", "github.com/golangci/golangci-lint/cmd/golangci-lint@v1.59.1"}).
		WithExec([]string{"pwd"}).
		WithExec([]string{"golangci-lint", "run", "./src", "./modules/...", "--issues-exit-code=1"})
}

func (m *Go01) Build(ctx context.Context, source *Directory) (string, error) {
	return m.Lint(ctx, source).
		WithExec([]string{"go", "build", "-v", "./..."}).
		Stdout(ctx)
}

// Build a ready-to-use development environment
func (m *Go01) BuildEnv(source *Directory) *Container {
	return dag.Container().
		From("golang:1.22.1").
		WithDirectory("/src", source).
		WithWorkdir("/src")

}
