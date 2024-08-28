
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// sslsplitFormula represents a formula in Go.
type sslsplitFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg sslsplitFormula) Print() {
    fmt.Printf("Name: sslsplit\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := sslsplitFormula{
        Description:  "Man-in-the-middle attacks against SSL encrypted network connections",
        Homepage:     "https://www.roe.ch/SSLsplit",
        URL:          "https://github.com/droe/sslsplit/commit/e17de8454a65d2b9ba432856971405dfcf1e7522.patch?full_index=1",
        Sha256:       "5e62ef71e9d154f54fac451123377ad0c59d103b544386b4100ecb34da2ad2cc",
        Dependencies: []string{"check", "pkg-config", "libevent", "libnet", "libpcap", "openssl@3"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installsslsplit(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg sslsplitFormula) Installsslsplit() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "e17de8454a65d2b9ba432856971405dfcf1e7522.patch?full_index=1"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "e17de8454a65d2b9ba432856971405dfcf1e7522"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
