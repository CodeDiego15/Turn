package main

import (
	"fmt"
	"os"
	"os/exec"
)

// Definición de la estructura Formula
type Formula struct {
	Name        string
	Description string
	Homepage    string
	URL         string
	Sha256      string
	License     string
	Install     func() error
	Test        func() error
}

func (f *Formula) InstallPackage() error {
	fmt.Printf("Installing %s...\n", f.Name)
	if err := f.Install(); err != nil {
		return fmt.Errorf("installation failed: %v", err)
	}
	return nil
}

func (f *Formula) TestPackage() error {
	fmt.Printf("Testing %s...\n", f.Name)
	if err := f.Test(); err != nil {
		return fmt.Errorf("testing failed: %v", err)
	}
	return nil
}

var python = &Formula{
	Name:        "Python",
	Description: "Python Programming Language",
	Homepage:    "https://www.python.org/",
	URL:         "https://www.python.org/ftp/python/3.12.0/python-3.12.0-linux-x86_64.tar.xz",
	Sha256:      "b1f38eb43d4e3a0e0a8e6a7b05a5cb56f32b676c9c7a5a7c61d21c15d72c3d3e",
	License:     "Python-2.0",
	Install: func() error {
		fmt.Println("Downloading Python...")
		cmd := exec.Command("curl", "-LO", "https://www.python.org/ftp/python/3.12.0/python-3.12.0-linux-x86_64.tar.xz")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Extracting Python...")
		cmd = exec.Command("tar", "-xJf", "python-3.12.0-linux-x86_64.tar.xz")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Installing Python...")
		cmd = exec.Command("cd", "python-3.12.0", "&&", "./configure", "&&", "make", "&&", "make", "install")
		if err := cmd.Run(); err != nil {
			return err
		}

		return nil
	},
	Test: func() error {
		fmt.Println("Testing Python...")
		cmd := exec.Command("python3", "--version")
		output, err := cmd.CombinedOutput()
		if err != nil {
			return err
		}
		fmt.Println(string(output))
		return nil
	},
}

func main() {
	if err := python.InstallPackage(); err != nil {
		fmt.Println("Installation failed:", err)
		os.Exit(1)
	}

	if err := python.TestPackage(); err != nil {
		fmt.Println("Testing failed:", err)
		os.Exit(1)
	}

	fmt.Println("Python installed and tested successfully!")
}
