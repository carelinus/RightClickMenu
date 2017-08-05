package main

import (
	"log"
	"golang.org/x/sys/windows/registry"
	"fmt"
	"strings"
	"bufio"
	"os"
)

var finish bool

func main() {
	finish = true
	var menuItemName string
	var menuItemIconAddress string
	var menuItemJob string

	for finish {
		//
		fmt.Print(" Menü item adını giriniz : ")
		fmt.Print(" ")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			menuItemName = scanner.Text()
		}
		//
		fmt.Print(" Menü item ikon yolunu giriniz (.ico) : ")
		fmt.Print(" ")
		fmt.Scan(&menuItemIconAddress)
		//
		fmt.Print(` Yapacağı işi giriniz (örn. "C:\Windows\System32\notepad.exe") : `)
		fmt.Print(" ")
		fmt.Scan(&menuItemJob)

		fmt.Print(" Menü eklenecek emin misiniz? : ")
		complete := Ask4confirm()
		if complete {
		} else {
			break
		}

		//CreateKey
		_, _, e := registry.CreateKey(registry.CLASSES_ROOT, `Directory\Background\shell\`+menuItemName+`\command`, registry.ALL_ACCESS)
		nilControl(e)
		//CreateJob
		createJob, _, e := registry.CreateKey(registry.CLASSES_ROOT, `Directory\Background\shell\`+menuItemName+`\command`, registry.ALL_ACCESS)
		nilControl(e)
		createJob.SetStringValue("", menuItemJob)
		//SetIcon
		a, _, e := registry.CreateKey(registry.CLASSES_ROOT, `Directory\Background\shell\`+menuItemName, registry.ALL_ACCESS)
		nilControl(e)
		a.SetStringValue("Icon", menuItemIconAddress)

		fmt.Print(" Yeni bir tane eklemek ister misiniz? : ")
		isConfirmed := Ask4confirm()
		if isConfirmed {
			finish = true
		} else {
			finish = false
		}
	}

}

func Ask4confirm() bool {
	var s string

	fmt.Printf("(Y/n): ")
	_, err := fmt.Scan(&s)
	if err != nil {
		panic(err)
	}

	s = strings.TrimSpace(s)
	s = strings.ToLower(s)

	if s == "y" || s == "yes" {
		return true
	}
	return false
}

func nilControl(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
