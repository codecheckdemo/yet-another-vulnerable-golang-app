package main

import (
	"log"
    "os"
    "os/exec"
    "crypto/rand"
	"github.com/BurntSushi/toml"
	"github.com/hashicorp/consul/logger"
	"github.com/rs/cors"
)

type Foo struct {
	Bar string
}

func main() {
	log.Printf("This application does nothing...")
	var foo Foo
	if _, err := toml.Decode("Bar = 'Qux'", &foo); err != nil {
		log.Printf("An error has occurred.")
	}
	log.Printf(foo.Bar)
	_ = cors.Options{}
	_ = logger.LevelFilter()
    good, _ := rand.Read(nil)
	println(good)
	run := "sleep" + os.Getenv("SOMETHING")
	cmd := exec.Command(run, "5")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
}
