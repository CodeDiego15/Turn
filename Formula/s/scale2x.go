
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// scale2xFormula represents a formula in Go.
type scale2xFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg scale2xFormula) Print() {
    fmt.Printf("Name: scale2x\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := scale2xFormula{
        Description:  "Real-time graphics effect",
        Homepage:     "https://www.scale2x.it/",
        URL:          "https://github.com/amadvance/scale2x/releases/download/v4.0/scale2x-4.0.tar.gz",
        Sha256:       "c8c30b93ecb6bf66a18c8004a6b56ca9ab4fc074fe4e433439f3bd49ba944005",
        Dependencies: []string{"libpng"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installscale2x(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg scale2xFormula) Installscale2x() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "scale2x-4.0.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "scale2x-4.0.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
