# docker-demo

## Intsalling docker 

    Ubuntu: https://docs.docker.com/engine/install/ubuntu/
    Windows: https://hub.docker.com/editions/community/docker-ce-desktop-windows
    MacOS: https://hub.docker.com/editions/community/docker-ce-desktop-mac

## Containers

A Linux container is a set of one or more processes that are isolated from the rest of the system. 

Docker provides (in simple terms...):

    Process isolation
    Filesystem isolation (union filesystem)
    Network isolation
    
Docker is not
    A VM
    A chroot jail (per se)
    Secure all by itself

## Docker basics

### `docker pull`

    $ docker pull ubuntu:latest
    $ docker pull ubuntu:18.04

### `docker run`

    $ docker run busybox echo "hello world"
    $ docker run -it ubuntu /bin/bash
    $ docker run -it -v /tmp:/mnt ubuntu:18.04 /bin/bash
    $ docker run -it -v /tmp:/mnt -p 8080:8080 ubuntu /bin/bash

#### Useful command line params

    -v        mount volume
    -p        open a port
    -e        environment vars
    --rm      remove the docker container once it finishes
    -d        daemonize the container
    --restart the restart policy
    -i -t     run it interactively

### `docker ps`

Shows running containers. Add `-a` to see stopped containers.

### `docker rm`

Removes stopped containers. Add `-f` to "force" remove a container.

Try this for cleanup:

    $ docker rm $(docker ps -a -q -f status=exited)

Equivalent to:

    $ docker container prune

### `docker images`

Shows docker images downloaded

### `docker rmi`

Removes downloaded docker images

### `docker build`

    $ cd /path/to/repo/repo-example
    $ docker build --tag repo-exmaple:latest .
    $ docker run -p 9090:8080 repo-example:latest
    ...
    $ docker rm $(docker ps -a -q -f status=exited)

Some variations

    $ docker run --name example -p 9090:8080 -d repo-example:latest
    $ docker run -p 9090:8080 --rm repo-example:latest

## Activity

### Local pgadmin4

    $ docker run --name pgadmin4 -p 10443:443 -e "PGADMIN_DEFAULT_EMAIL=you@example.com" -e "PGADMIN_DEFAULT_PASSWORD=very_complex_password" -v /path/to/.pgadmin:/var/lib pgadmin -d dpage/pgadmin4:4.24

### Modded Minecraft server

https://github.com/itzg/docker-minecraft-server


#### Vanilla

    $ docker run -d -p 25565:25565 -e EULA=TRUE --name mc itzg/minecraft-server

#### ToroQuest

    $ docker run -d -v /path/to/toro:/data -e VERSION=1.12.2 -e TYPE=FORGE -p 25565:25565 -e EULA=TRUE --name mc-toroquest itzg/minecraft-server

#### Checking startup

    $ docker logs -f mc

## Power stuff:

### `docker volume`

    $ docker volume help

### Backup/Restore Volumes

#### Backup

    $ docker run -v $VOLUME_FOLDER_PATH --name $CONTAINER_NAME ubuntu /bin/bash
    $ docker run --rm --volumes-from $CONTAINER_NAME -v $(pwd):/backup ubuntu tar cvf /backup/backup.tar $VOLUME_FOLDER_PATH

#### Restore

    $ docker run -v $VOLUME_FOLDER_PATH --name $NEW_CONTAINER_NAME ubuntu /bin/bash
    $ docker run --rm --volumes-from $NEW_CONTAINER_NAME -v $(pwd):/backup ubuntu bash -c "cd $VOLUME_FOLDER_PATH && tar xvf /backup/backup.tar --strip 1"

### `docker network`

    $ docker network help

## docker-compose

A `docker-compose` is a yaml file that allows you to create mutli-container deployments.

Example:

    $ cd /path/to/docker-mailserver
    $ docker-compose up -d
