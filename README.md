# aram-builds
aram builds site

## Set up traefik first to make this work properly!
`docker-compose up -d --build`

## migrations
`docker-compose run backend ./backend migrate`

## add admin user
`docker-compose run backend ./backend user -u <USERNAME> -p <PASSWORD>`

The passwords are stored in the browser in plaintext (needed for basicauth) so you probably want to make a unique password

## dev environment
`docker-compose -f docker-compose.dev.yml up -d --build`

The other commands (migrations and users) are also the same, except you have to add `-f docker-compose.dev.yml` before `run`