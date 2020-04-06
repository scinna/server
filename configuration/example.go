package configuration

import "fmt"

// PrintExample shows an example for the configuration file
func PrintExample() {

	fmt.Println(`
# This is the connection string for the postgresql database
PostgresDSN: "postgres://username:password@host[:port]/database"

# This feeds the generator that create unique IDs for your medias
IdAlphabet: "0123456789abcdefghijklmnopqrstuvwxyzABCEFGHIJKLMNOPQRSTUVWXYZ"
IdSize: 10

# This is the URL used when sending activation mail or forgotten mail for example
WebURL: https://scinna.local

# This is the port the server is listening to 
WebPort: 1635

# This is where the medias will be stored
MediaPath: "/tmp/scinna/medias"

# /!\ REALLY IMPORTANT /!\
# This field must be set according to what you reverse proxy fills
# Any misconfiguration can lead to the client being able to make the server belive any IP it wants you to belive
# This will determinate the user IP (Prevent ddos, bruteforce, etc...)
# If you are not using reverse-proxy, leave this empty
HeaderIPField: X-Real-IP

SMTPSender: noreply@scinna.local
SMTPHost: localhost:1025
SMTPUser:
SMTPPass:

# This will let or not users register on this Scinna instance.
# Accepted values are NO, YES or INVITE
RegistrationAllowed: yes

# This is the amount of request per 5 minutes an IP is allowed to do
RateLimiting: 100`)

}
