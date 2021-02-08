# Scinna - Server

## What is it ?

Scinna is a self-hosted picture sharing website focusing on screenshot.

## Installing

Ready to setup your own scinna instance ? Go take at look at the [Wiki](https://github.com/scinna/server/wiki) where everything is explained.

## Our pledge

We aim to build a stable and easy to use software. Thus we pledge to never break database compatibility across last digit version (e.g. 1.2.3 has the same database structure as 1.2.4 and will not break either upgrading it or downgrading it). The second digit shows an update in the database which will go painlessly (Migration are applied when needed for upgrade, can rollback manually with a flag). Most feature will bump the second number, large features (e.g. nested folder) will bump the first digit. This is more-or-less following the (semver)[https://semver.org] standard.

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
- [x] Localization
- [x] Collections to sort your pictures

### Version 2.0

- [ ] Nested collections
- [ ] OAuth from multiple service providers (Google, Facebook, Github)
- [ ] Video sharing
- More to come

### Version 3.0

- [ ] Real folders (Like folders with sub-folders)
- [ ] Support for S3 storage (Maybe?)
- [ ] LDAP ?
- More to come

## Contributing

Since there are no contributing.md yet, please act nice when creating PR, following the same code conventions as we do.

There will be a contributing.md file later for clarification.

Oh, and we are using [SemVer](https://semver.org/) so remember to up the SCINNA_VERSION number in main.go if you have changed something related to the database, as intended. For everything else, just change the SCINNA_PATCH number
