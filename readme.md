# Scinna - Server
[![Shield](https://img.shields.io/website?down_color=%2387E7E1&down_message=Website&label=Our&up_color=%2387E7E1&up_message=Website&url=https%3A%2F%2Fscinna.app)](https://scinna.app)
[![Shield](https://discordapp.com/api/guilds/806593726859837460/widget.png?style=shield)](https://discord.gg/EYdDb72fR2)
[![Shield](https://img.shields.io/reddit/subreddit-subscribers/Scinna?color=%2387E7E1&label=Reddit%20r%2FScinna)](https://reddit.com/r/Scinna)

## What is it ?

Scinna is a self-hosted picture sharing website focusing on screenshots.

## Installing

Ready to set your own Scinna instance up ? Go take at look at the [Wiki](https://github.com/scinna/server/wiki) where everything is explained.

Versions 0.0.x are private in-dev version they should NOT be used in any way shape or form. Nothing is guaranteed on those and things WILL break.

## Our pledge

We aim to build a stable and easy to use software. Thus, we pledge to never break database compatibility across last digit version (e.g. 1.2.3 has the same database structure as 1.2.4 and will not break either upgrading it or downgrading it). The second digit shows an update in the database which will go painlessly (Migration are applied when needed for upgrade, can rollback manually with a flag). Most feature will bump the second number, large features (e.g. nested folder) will bump the first digit. This is more-or-less following the [semver](https://semver.org) standard.

## Roadmap

The first version will have the basic feature expected from a self-hosted image sharing website:

### Version 1.0

- [x] Server management through config file (Private, public, custom naming, ...)
- [x] Account management
- [x] Uploading pictures
- [x] Picture infos (Title, Description, Date, public/private/unlisted)
- [x] PostgreSQL
- [x] Database auto-initialization (--generate-db)
- [ ] Database auto-upgrading
- [x] Localization
- [x] Collections to sort your pictures
- [ ] Web UI
    - To fix: Forgotten password
- [ ] User roles (Simple user, admin, ...)
- [ ] Server admin panel
- [ ] ShareX compatibility
- [ ] URL Shortener
- [ ] Pastebin

### Version 2.0

- [ ] Support for abstract filesystems (local, S3, ftp, ...)
- [ ] OAuth from multiple service providers (Google, Facebook, Github)
- [ ] 2FA
- [ ] 2FA enforceability (Like if the admin want he can force everyone to use 2FA, like on Github) ?
- [ ] LDAP ?

### Version 3.0

- [ ] Nested collections (Like folders with sub-folders)
- [ ] Video sharing
- More to come

### Acknowledgement

Thanks to my friend @cylgom for beta-testing it and for the awesome icons he made for the web-ui!

## Contributing

Since there are no contributing.md yet, please act nice when creating PR, following the same code conventions as we do. We have a clear idea about how the software will grow in features so if you are not sure, please come by the discord to ask if the feature you want to implement is following our vision.

There will be a contributing.md file later for clarification.

Oh, and we are using [SemVer-ish](https://semver.org/) so remember to up the SCINNA_VERSION number in main.go if you have changed something related to the database, as intended. For everything else, just change the SCINNA_PATCH number. For more explanation, look at the [Our pledge](#our-pledge) section.
