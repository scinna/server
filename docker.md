# Scinna

## What's Scinna ?

Scinna is a self-hosted picture sharing website focusing on screenshot.

## Setting it up
Pick a folder and create a config.json containing the following sample
```
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
    "__comment__": "true means open registration, false means requiring an invite code",
    "Allowed": false,
    "Validation": "open|email|admin"
  },
  "WebURL": "https://i.scinna.app/",
  "WebPort": 1635,
  "MediaPath": "/medias/",
  "RealIpHeader": "X-Real-IP"
}
```

You can then create a docker-compose file with the following content:
```
version: "3.8"
services:
    postgres:
        image: postgres:alpine
        restart: always
        volumes:
            - ./db:/var/lib/postgresql/data/pgdata
        environment:
            - POSTGRES_DB=scinna
            - POSTGRES_USER=scinna
            - POSTGRES_PASSWORD=scinna
            - PGDATA=/var/lib/postgresql/data/pgdata

    scinna:
        image: scinna/server
        restart: unless-stopped
        depends_on:
            - postgres
        volumes:
            - ./data:/medias
            - ./config.json:/config.json
        ports:
            - 1635:1635
```

Since the first version of scinna won't auto-create the database, you'll need to set it up yourself
```
$ wget https://raw.githubusercontent.com/scinna/server/master/SQL/Initialize.sql
$ docker-compose up postgres -d
$ docker ps # Find the container ID for your scinna_postgres container
$ docker copy Initialize.sql CONTAINER_ID:/Initialize.sql
$ docker-compose exec postgres psql scinna scinna -f /Initialize.sql
$ docker-compose down
$ docker-compose up -d
```

We got it running. Now let's setup our reverse proxy.
Here is a sample configuration for nginx. It uses SSL from Let's Encrypt

```
server {
        listen 80;
        server_name i.scinna.app;

        return 302 https://$server_name$request_uri;
}

server {
        listen 443 http2 ssl;
        server_name i.scinna.app;

        include snippets/ssl.conf;

        client_max_body_size 10M;

        access_log /var/log/scinna.log;
        error_log /var/log/scinna_error.log warn;

        location / {
                proxy_pass http://127.0.0.1:1635;
                proxy_set_header Host $host;
                proxy_set_header X-Real-IP $remote_addr;
                proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
                proxy_set_header X-Forwarded-Proto $scheme;
        }

}

```

Take a look at scinna logs, it will have generated an invitation code for you to create your account.
