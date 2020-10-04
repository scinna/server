# Scinna - Server

## What is it ?

Scinna is a self-hosted picture sharing website focusing on screenshot.

## Installing

Scinna is a golang project so it's a simple executable.

You can put the configuration file either in the same folder (Portable mode), or in `/etc/scinna/config.json`.

Here is a sample configuration file
```json
{
  "ConfigSMTP": {
    "Enabled": true,
    "ConnectionType": "STARTTLS",
    "Hostname": "smtp.mailgun.org",
    "Port": 587,
    "Username": "",
    "Password": "",
    "Sender": ""
  },
  "ConfigDB": {
    "Hostname": "localhost",
    "Port": 5432,
    "Username": "scinna",
    "Password": "scinna",
    "Database": "scinna"
  },
  "Registration": {
    "Allowed": true,
    "Validation": "email"
  },
  "WebURL": "https://scinna.drx/",
  "WebPort": 1635,
  "MediaPath": "/tmp/medias/",
  "RealIpHeader": "X-Real-IP"
}
```


More info on the [Wiki](https://github.com/scinna/server/wiki)

## Roadmap

The first version will have the basic feature expected from a self-hosted image sharing website:

### Version 1.0

- [ ] Server management (Private, public, custom naming, ...)
- [ ] Account management
- [ ] Uploading pictures
- [ ] Picture infos (Title, Description, Date, public/private/unlisted)
- [ ] PostgreSQL

### Version 2.0

- [ ] Folders to sort your pictures
- [ ] OAuth from multiple service providers (Google, Facebook, Github)
- [ ] Localization
- More to come

### Version 3.0

- [ ] Video sharing
- [ ] Support for S3 storage (Maybe?)
- More to come

### Version 4.0

- [ ] Support for other SGBD (Maybe?, at least Sqlite)
- More to come


## Contributing

Since there are no contributing.md yet, please act nice when creating PR, following the same code conventions as we do.

There will be a contributing.md file later for clarification.

Oh, and we are using [SemVer](https://semver.org/) so remember to up the number in main.go and the DB script creation only if you have changed something there, as intended. This will maybe automatically included in a build setup when it will be created.
