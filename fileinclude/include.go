package fileinclude

import (
	"embed"
	"fmt"
	"log"
	"os"
)

func init() {
	log.Println(fmt.Sprintf("%s", "init fileinclude"))
}

func Include(d *embed.FS) {
	log.Println(d)

	if _f, e := os.ReadDir("files"); e != nil {
		log.Fatal(e)
	} else {
		for _, d := range _f {
			log.Println(d)
		}
	}

	includeDir(d)
	includeFile(d)
	includeFileOpen(d)
}

func includeDir(d *embed.FS) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error ::: ", err)
		}
	}()
	if dirs, err := d.ReadDir("files/"); err != nil {
		log.Println("includeDir ::: ", err)
	} else {
		for _, d := range dirs {
			log.Println(d)
		}
	}
}

func includeFile(d *embed.FS) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error ::: ", err)
		}
	}()
	if f, err := d.ReadFile("files/dep-develop.yml"); err != nil {
		log.Println("includeFile ::: ", err)
	} else {
		log.Println(string(f))
	}
}

func includeFileOpen(d *embed.FS) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error ::: ", err)
		}
	}()

	if f, err := d.Open("files/dep-develop.yml"); err != nil {
		log.Println("includeFileOpen ::: ", err)
	} else {
		log.Println(f.Stat())
	}
}
