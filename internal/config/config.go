package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"regexp"
	"smartNotify/internal/telegram"
	"strings"
)

func ProcessArgs() {
	args := strings.Join(os.Args[1:], " ")
	re := regexp.MustCompile(`-\w+(?:=[^\s-]+(?:-[^-\s]+)*)?`)
	splitArgs := re.FindAllString(args, -1)

	if len(splitArgs) == 0 {
		log.Fatalf("Software has no args. Use -h")
	}

	for _, arg := range splitArgs {
		arg = arg[1:]

		if arg == "h" {
			fmt.Println("HELP!")
		} else if strings.HasPrefix(arg, "config=") {
			filename := strings.SplitN(arg, "=", 2)[1]

			if filename == "" {
				log.Fatalf("Filename not defined")
			}

			file, err := os.ReadFile(filename)
			if err != nil {
				log.Fatalf("error reading YAML file: %v", err)
			}

			var config Config
			err = yaml.Unmarshal(file, &config)
			if err != nil {
				log.Fatalf("error parsing YAML: %v", err)
			}

			message := config.Users[0].Notifications[0].Message
			username := config.Users[0].Name
			telegramId := config.Users[0].Contact.TelegramID

			telegram.Send(telegramId, fmt.Sprintf("%s, %s", username, message))
		} else {
			log.Fatalf("Unrecognized argument: %s", arg)
		}
	}
}
