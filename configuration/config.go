package configuration

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Load loads the configuration from the file
func Load() (Configuration, error) {

	dsn, exists := os.LookupEnv("POSTGRES_DSN")
	if !exists {
		panic("No database url found! (POSTGRES_DSN)")
	}

	port, exists := os.LookupEnv("WEB_PORT")
	if !exists {
		panic("No listening port found! (WEB_PORT)")
	}

	headerIPField, exists := os.LookupEnv("HEADER_IP_FIELD")
	if !exists {
		fmt.Println("The header for the IP field is not set (HEADER_IP_FIELD). If you are using a reverse-proxy please be sure to set it according to its configuration.\nTo disable this message, add the environment variable with an empty value.")
	}

	registrationAllowed, exists := os.LookupEnv("REGISTRATION_ALLOWED")
	var registrationAllowedInt int
	if !exists {
		fmt.Println("Registration is allowed by default. You can't hide this message or turn it off by filling the \"REGISTRATION_ALLOWED\" environment variable with either NO|INVITE|YES.")
		registrationAllowedInt = 0
	} else {
		switch strings.ToLower(registrationAllowed) {
		case "no":
			registrationAllowedInt = 2
			break
		case "invite":
			registrationAllowedInt = 1
			break
		case "yes":
			registrationAllowedInt = 0
			break

		default:
			fmt.Println("The registration setting can't be parsed. Setting it to private for safety")
			registrationAllowedInt = 2
			break
		}
	}

	websiteURL, exists := os.LookupEnv("WEB_URL")
	if !exists {
		panic("Can't find website URL (WEB_URL). This is required to make the link in the validation email and the forgotten password email")
	}

	picturepath, exists := os.LookupEnv("PICTURE_PATH")
	if !exists {
		panic("No picture folder found! (PICTURE_PATH)")
	}

	ratelimiting, exists := os.LookupEnv("RATE_LIMITING")
	if !exists {
		panic("Please set a rate limiting for the API (RATE_LIMITING). X request max per 5 minutes")
	}

	ratelimitingInt, err := strconv.Atoi(ratelimiting)
	if err != nil {
		panic(err)
	}

	return Configuration{
		PostgresDSN:         dsn,
		PicturePath:         picturepath,
		HeaderIPField:       headerIPField,
		RegistrationAllowed: registrationAllowedInt,
		WebURL:              websiteURL,
		WebPort:             port,
		RateLimiting:        ratelimitingInt,
	}, nil
}

// Configuration represents the current config for the server
type Configuration struct {
	PostgresDSN         string
	PicturePath         string
	HeaderIPField       string
	RegistrationAllowed int
	WebURL              string
	WebPort             string
	RateLimiting        int
}
