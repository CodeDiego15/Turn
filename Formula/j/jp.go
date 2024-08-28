
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// jpFormula represents a formula in Go.
type jpFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg jpFormula) Print() {
    fmt.Printf("Name: jp\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := jpFormula{
        Description:  "Dead simple terminal plots from JSON data",
        Homepage:     "https://github.com/sgreben/jp",
        URL:          "https://github.com/sgreben/jp/archive/refs/tags/1.1.12.tar.gz",
        Sha256:       "3b35ee0680b137a2202b4ce685bbfa80ccd9dca2ec2e7ab550b2f3384b2b8f68",
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

    if err := pkg.Installjp(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg jpFormula) Installjp() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "1.1.12.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "1.1.12.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
