package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jaypipes/ghw"
	"github.com/pippolo84/device-discovery/internal/config"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("config file required")
	}

	cfg, err := config.FromFile(os.Args[1])
	if err != nil {
		log.Fatalf("configuration: %v", err)
	}

	pci, err := ghw.PCI()
	if err != nil {
		log.Fatalf("error getting PCI info: %v", err)
	}

	for _, device := range pci.ListDevices() {
		for _, feature := range cfg.Features {
			for _, matchOn := range feature.MatchOn {
				if !matchOn.PCIID.IsValid() {
					continue
				}

				if matchOn.PCIID.Vendor != device.Vendor.ID {
					continue
				}

				if matchOn.PCIID.Device != device.Product.ID {
					continue
				}

				log.Printf("device: %q\n", device)
				fmt.Printf("%s=%s\n", feature.Name, feature.Value)
			}
		}
	}
}
