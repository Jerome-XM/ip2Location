package ip2proxy

import (
	"fmt"
	"testing"
)

func TestIPWhoer(t *testing.T) {
	db, err := OpenDB("./bin/IP2LOCATION-LITE-DB1.IPV6.BIN")

	if err != nil {
		t.Error(err)
		return
	}
	ip := "185.158.35.255"
	all, err := db.GetAll(ip)

	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Printf("ModuleVersion: %s\n", ModuleVersion())
	fmt.Printf("PackageVersion: %s\n", db.PackageVersion())
	fmt.Printf("DatabaseVersion: %s\n", db.DatabaseVersion())

	fmt.Printf("ModuleVersion: %s\n", ModuleVersion())
	fmt.Printf("PackageVersion: %s\n", db.PackageVersion())
	fmt.Printf("DatabaseVersion: %s\n", db.DatabaseVersion())
	fmt.Printf("CountryShort: %s\n", all["CountryShort"])
	fmt.Printf("CountryLong: %s\n", all["CountryLong"])

	db.Close()
}
