
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// wget2Formula represents a formula in Go.
type wget2Formula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg wget2Formula) Print() {
    fmt.Printf("Name: wget2\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := wget2Formula{
        Description:  "Successor of GNU Wget, a file and recursive website downloader",
        Homepage:     "https://gitlab.com/gnuwget/wget2",
        URL:          "https://ftp.gnu.org/gnu/wget/wget2-2.1.0.tar.gz",
        Sha256:       "e88b3e5451829851b61086d3a3962ed71e14ec41264bcefdaf474f4782c1f427",
        Dependencies: []string{"doxygen", "gnu-sed", "graphviz", "lzip", "pandoc", "pkg-config", "texinfo", "brotli", "gettext", "gnutls", "gpgme", "libassuan", "libgpg-error", "libidn2", "libmicrohttpd", "libnghttp2", "libpsl", "libtasn1", "lzlib", "nettle", "p11-kit", "pcre2", "xz", "zstd"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installwget2(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg wget2Formula) Installwget2() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "wget2-2.1.0.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "wget2-2.1.0.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
