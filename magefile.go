//+build mage

package main

import (
	"errors"
	"fmt"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"
)

const (
	packageName = "github.com/NodeFactoryIo/hactar-deamon"
)

var goexe = "go"

func init() {
	if exe := os.Getenv("GOEXE"); exe != "" {
		goexe = exe
	}

	// We want to use Go 1.11 modules even if the source lives inside GOPATH.
	os.Setenv("GO111MODULE", "on")
}

// mage build
func Build() error {
	return sh.Run(goexe, "build", "./...")
}

// mage buildall
func BuildAll() error {
	mg.Deps(Test)

	platforms := [...]string{"linux", "darwin"}
	archs := [...]string{"386", "amd64"}

	for _, os := range platforms {
		for _, a := range archs {
			build(os, a)
		}
	}

	return nil
}

func build(os string, arch string) error {
	env := map[string]string{
		"GOOS":   os,
		"GOARCH": arch,
	}
	timeStamp := time.Now().Format("2006-01-02")
	fileName := fmt.Sprintf("./builds/%s/hactar-%s-%s-%s", os, os[:3], arch, timeStamp)

	return sh.RunWith(env, goexe, "build", "-o", fileName)
}

// mage install
func Install() error {
	return sh.Run(goexe, "install", "./...")
}

// mage test
func Test() error {
	mg.Deps(Lint, Vet)

	fmt.Println("Run tests:")
	env := map[string]string{"GOFLAGS": testGoFlags()}
	return runCmd(env, goexe, "test", "./...")
}

func testGoFlags() string {
	// long tests can be skipped when running on CI
	if isCI() {
		return "-test.short"
	}
	return ""
}

//  Run go vet linter
func Vet() error {
	if err := runCmd(nil, goexe, "vet", "./..."); err != nil {
		return fmt.Errorf("Error running go vet: %v", err)
	} else {
		fmt.Println("Linter vet finished")
	}
	return nil
}

// Run gofmt linter
func Fmt() error {
	if !isGoLatest() {
		return nil
	}
	pkgs, err := hactarPackages()
	if err != nil {
		return err
	}
	failed := false
	first := true
	for _, pkg := range pkgs {
		files, err := filepath.Glob(filepath.Join(pkg, "*.go"))
		if err != nil {
			return nil
		}
		for _, f := range files {
			// gofmt doesn't exit with non-zero when it finds unformatted code
			// so we have to explicitly look for output, and if we find any, we
			// should fail this target.
			s, err := sh.Output("gofmt", "-l", f)
			if err != nil {
				fmt.Printf("ERROR: running gofmt on %q: %v\n", f, err)
				failed = true
			}
			if s != "" {
				if first {
					fmt.Println("The following files are not gofmt'ed:")
					first = false
				}
				failed = true
				fmt.Println(s)
			}
		}
	}
	if failed {
		return errors.New("improperly formatted go files")
	}
	return nil
}

func Lint() error {
	fmt.Println("Linter golint")
	pkgs, err := hactarPackages()
	if err != nil {
		return err
	}
	failed := false
	for _, pkg := range pkgs {
		// We don't actually want to fail this target if we find golint errors,
		// so we don't pass -set_exit_status, but we still print out any failures.
		if _, err := sh.Exec(nil, os.Stderr, nil, "golint", pkg); err != nil {
			fmt.Printf("ERROR: running go lint on %q: %v\n", pkg, err)
			failed = true
		} else {
			fmt.Println("\u2713 ", pkg)
		}
	}
	if failed {
		return errors.New("errors running golint")
	}
	return nil
}

// Util functions

var (
	pkgPrefixLen = len("github.com/NodeFactoryIo/hactar-deamon")
	pkgs         []string
	pkgsInit     sync.Once
)

func hactarPackages() ([]string, error) {
	var err error
	pkgsInit.Do(func() {
		var s string
		s, err = sh.Output(goexe, "list", "./...")
		if err != nil {
			return
		}
		pkgs = strings.Split(s, "\n")
		for i := range pkgs {
			pkgs[i] = "." + pkgs[i][pkgPrefixLen:]
		}
	})
	return pkgs, err
}

func isCI() bool {
	return os.Getenv("CI") != ""
}

func isGoLatest() bool {
	return strings.Contains(runtime.Version(), "1.13")
}

func runCmd(env map[string]string, cmd string, args ...string) error {
	if mg.Verbose() {
		return sh.RunWith(env, cmd, args...)
	}

	output, err := sh.OutputWith(env, cmd, args...)
	if err != nil {
		fmt.Fprint(os.Stderr, output)
	} else {
		fmt.Fprint(os.Stdout, output)
	}

	return err
}
