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
		fmt.Print(" Enter the name of the menu item : ")
		fmt.Print(" ")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			menuItemName = scanner.Text()
		}
		//
		fmt.Print(" Enter the menu item icon path (.ico) : ")
		fmt.Print(" ")
		fmt.Scan(&menuItemIconAddress)
		//
		fmt.Print(` Path of the executable (i.e. "C:\Windows\System32\notepad.exe") : `)
		fmt.Print(" ")
		fmt.Scan(&menuItemJob)

		fmt.Print(" Are you sure you want the menu item to be added? : ")
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

		fmt.Print(" Do you want to add another one? : ")
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
