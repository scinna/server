# Scinna - Server

## What is it ?

Scinna is a self-hosted picture sharing website focusing on screenshot.

## Installing

Ready to setup your own scinna instance ? Go take at look at the [Wiki](https://github.com/scinna/server/wiki) where everything is explained.

## Roadmap

The first version will have the basic feature expected from a self-hosted image sharing website:

### Version 1.0

- [ ] Server management (Private, public, custom naming, ...)
- [ ] Account management
- [ ] Uploading pictures
- [ ] Picture infos (Title, Description, Date, public/private/unlisted)
- [x] PostgreSQL
- [x] Database auto-initialization (--generate-db)
- [ ] Database auto-upgrading
- [ ] Localization

### Version 2.0

- [ ] Folders to sort your pictures
- [ ] OAuth from multiple service providers (Google, Facebook, Github)
- [ ] Video sharing
- More to come

### Version 3.0

- [ ] Real folders (Like folders with sub-folders)
- [ ] Support for S3 storage (Maybe?)
- More to come

## Contributing

Since there are no contributing.md yet, please act nice when creating PR, following the same code conventions as we do.

There will be a contributing.md file later for clarification.

Oh, and we are using [SemVer](https://semver.org/) so remember to up the SCINNA_VERSION number in main.go if you have changed something related to the database, as intended. For everything else, just change the SCINNA_PATCH number
