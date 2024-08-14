package model

import (
	"fmt"
	"os"
	"time"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

type IXCConnection struct {
	Url string
	User string
	Pass string
}

func openIXC(conn IXCConnection) selenium.WebDriver {
	caps := selenium.Capabilities{"browserName": "chrome"}
	chromeCaps := chrome.Capabilities{
		Path: "", Args: []string{ "--headless", "--no-sandbox" },
	}
	caps.AddChrome(chromeCaps)

	wd, err := selenium.NewRemote(caps, "http://localhost:8080")
	if err != nil {
		fmt.Printf("erro ao abrir selenium: %v", err) 
		os.Exit(1)
	}

	err = wd.Get(conn.Url)
	if err != nil {
		fmt.Printf("erro ao abrir ixc: %v", err) 
		os.Exit(1)
	}

	el, err := wd.FindElement("id", "email")
	el.SendKeys(conn.User)

	el, err = wd.FindElement("id", "senha")
	el.SendKeys(conn.Pass)

	el, err = wd.FindElement("id", "entrar")
	el.Click()
	time.Sleep(10)

	return wd
}

func OpenVendas(web IXCConnection) {
	wd := openIXC(web)
	defer wd.Quit()
}
