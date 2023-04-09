
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/atotto/clipboard"
	"github.com/getlantern/systray"
	"github.com/joho/godotenv"
	icon "github.com/linexjlin/systray-icons/enter-the-keyboard"
	"github.com/skratchdot/open-golang/open"
)

func getMasks() (masks []string) {
	// Define the directory path and file extension
	dir := "prompts"
	ext := ".json"

	// Use the filepath package to get a list of all files with the specified extension
	files, err := filepath.Glob(filepath.Join(dir, "*"+ext))
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		filename := filepath.Base(file)
		filenameWithoutExt := filename[:len(filename)-len(ext)]
		fmt.Println(filenameWithoutExt)
		masks = append(masks, filenameWithoutExt)
	}
	return
}

func onExit() {
	now := time.Now()
	fmt.Println("exit", now)
}

var _mClearContextSetTitle func(string)

func updateClearContextTitle(n int) {
	_mClearContextSetTitle(fmt.Sprintf(UText("Clear Context %d/%d"), n, g_userSetting.maxConext))
}

var updateHotKeyTitle func(string)

func monitorFileModification(filepath string) {
	fileInfo, err := os.Stat(filepath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	modTime := fileInfo.ModTime()
	fmt.Println("Initial modification time:", modTime)

	// Calculate the end time after 30 minutes
	endTime := time.Now().Add(30 * time.Minute)

	for time.Now().Before(endTime) {
		fileInfo, err = os.Stat(filepath)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		if modTime != fileInfo.ModTime() {
			fmt.Println("Modification time changed:", fileInfo.ModTime())
			modTime = fileInfo.ModTime()
			godotenv.Load(filepath)
		}

		time.Sleep(time.Second * 3) // Sleep for 3 seconds before checking again
	}

	fmt.Println("Monitoring completed.")
}

func onReady() {
	systray.SetTemplateIcon(icon.Data, icon.Data)

	mQuitOrig := systray.AddMenuItem(UText("Exit"), UText("Quit the whole app"))
	go func() {
		<-mQuitOrig.ClickedCh
		fmt.Println("Requesting quit")
		systray.Quit()
		fmt.Println("Finished quitting")
	}()

	mAbout := systray.AddMenuItem(UText("About"), UText("Open the project page"))
	go func() {
		for {
			<-mAbout.ClickedCh
			open.Start("https://github.com/linexjlin/inputGPT")
		}
	}()

	mSetKey := systray.AddMenuItem(UText("Set API KEY"), UText("Set the OpenAI KEY, baseurl etc.."))
	go func() {
		for {
			<-mSetKey.ClickedCh
			open.Start("env.txt")
			go monitorFileModification("env.txt")
		}
	}()

	mHotKey := systray.AddMenuItem("", UText("Click to active GPT"))
	updateHotKeyTitle = mHotKey.SetTitle

	systray.AddSeparator()

	mManager := systray.AddMenuItem(UText("Manage Prompts"), UText("Modify, Delete prompts"))
	go func() {