
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// go@1.18Formula represents a formula in Go.
type go@1.18Formula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg go@1.18Formula) Print() {
    fmt.Printf("Name: go@1.18\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := go@1.18Formula{
        Description:  "Open source programming language to build simple/reliable/efficient software",
        Homepage:     "https://go.dev/",
        URL:          "https://go.dev/dl/go1.18.10.src.tar.gz",
        Sha256:       "cc61c960b7186af57e6cc1d1f9a82beace7a93523784d78a177429d1944aa5eb",
        Dependencies: []string{"go"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installgo@1.18(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg go@1.18Formula) Installgo@1.18() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "go1.18.10.src.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "go1.18.10.src.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
