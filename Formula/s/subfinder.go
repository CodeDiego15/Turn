
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// subfinderFormula represents a formula in Go.
type subfinderFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg subfinderFormula) Print() {
    fmt.Printf("Name: subfinder\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := subfinderFormula{
        Description:  "Subdomain discovery tool",
        Homepage:     "https://github.com/projectdiscovery/subfinder",
        URL:          "https://github.com/projectdiscovery/subfinder/archive/refs/tags/v2.6.6.tar.gz",
        Sha256:       "b6c890b75d2c4ad273867b4fb1206cfe620a49c0c666f1b491adf92e4a03243b",
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

    if err := pkg.Installsubfinder(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg subfinderFormula) Installsubfinder() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "v2.6.6.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "v2.6.6.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
