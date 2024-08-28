
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// slirp4netnsFormula represents a formula in Go.
type slirp4netnsFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg slirp4netnsFormula) Print() {
    fmt.Printf("Name: slirp4netns\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := slirp4netnsFormula{
        Description:  "User-mode networking for unprivileged network namespaces",
        Homepage:     "https://github.com/rootless-containers/slirp4netns",
        URL:          "https://raw.githubusercontent.com/rootless-containers/slirp4netns/v1.2.1/tests/test-slirp4netns-api-socket.sh",
        Sha256:       "075f43c98d9a848ab5966d515174b3c996deec8c290873d92e200dc6ceae1500",
        Dependencies: []string{"autoconf", "automake", "pkg-config", "bash", "jq", "glib", "libcap", "libseccomp", "libslirp"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installslirp4netns(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg slirp4netnsFormula) Installslirp4netns() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "test-slirp4netns-api-socket.sh"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "test-slirp4netns-api-socket"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
