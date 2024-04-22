# Docker Setup

All you need to do run this project with docker is to:

1. Download this folder using the postwoman.zip in this folder and extract it in your preferred location

2. Rename `.env.example` to `.env` and put your own custom values (you can just use the provided values)

3. Run `docker-compose up` in your postwoman directory while your docker engine is running

4. Go to `localhost:YOURPORT` in your browser

Check out the [startup script](https://github.com/dawitalemu4/postwoman/tree/main/.docker-setup/startup.sh) in this folder to easily start up postwoman from a shortcut on your taskbar!

> Note for mac users: If you are running into this error: `rosetta error: failed to open elf at /lib64/ld-linux-x86-64.so.2`, try to use this flag in the docker file: `FROM --platform=linux/amd64 golang:1.22.2`
