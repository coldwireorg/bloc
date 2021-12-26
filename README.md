<div align="center">
  <br>
  <img width="500" src="./imgs/logo-full.svg" alt="BLOC">
  <br>
  <br>
</div>

<p align="center">
  <i>The most resilient and safe place to store your files</i><br>
</p>

<br>

## Description

Bloc is a project of federated cloud storage with strong encryption.<br>
The goal is to have a stable, simple to use and secure place to store our files and to not even worry about the loss of an instance of the network.

## Screenshots

<img src="./imgs/screen-login.png">
<img src="./imgs/screen-app.png">

## Available languages (in dev)
- English
- Swedish
- French

## Features
*because we are currently in early developpement, the federation system is not implemented yet, the protocol will be written in Rust in another repo*

### Available features
- User creation / authentification (only using username/password system to avoid collecting metadatas)
- File uploading / downloading
- File deleting
- Favorite system (add the files to a list of favorites)

### Planned features
- Private file sharing (the file encryption key is re-encrypted using the receiver public key)
- Public file sharing (sharing a file using a link, this one expose the file's encryption key!)
- Making possible to scale bloc on nomad/kubernetes clusters 
- Moderation / Admin panel for configuring
- Authentification from others servers of the federated network
- Sharing files over the federated network
- Automated backup over the network (encrypted files are sliced into many little parts and sent to a swarm of servers)

## Installation / Selft-Hosting

*note: before running the backend, setup a postgresql server, you can see the tables in `database/tables.sql`*

### Simple way: with docker-compose
```sh
git clone --recurse-submodules https://github.com/coldwireorg/bloc.git
cd bloc
```

Then edit the enviroment variables in the `docker-compose.yml` file.
You will find more detail about environment variables below.

```sh
# after configuration, lets's run bloc!
docker-compose up -d
```

#### Web server / load balancing
It's up to you to setup your own web server with nginx, apache, træfik, etc.

nginx exemple:
```conf
upstream bloc-frontend {
  server 127.0.0.1:3000;
}

upstream bloc-backend {
  server 127.0.0.1:3001;
}

server {
  server_name bloc.coldwire.org;
  location / {
    proxy_pass http://bloc-fontend;
  }

  location /api {
    proxy_pass http://bloc-backend;
  }

  listen 80;
}
```

#### Environment variables:
##### Backend
*theses are the envrionment variables for the backend (in `bloc-api` in the `docker-compose.yml` file)*
```sh
SERVER_DOMAIN=dev.bloc.coldwire.org # the domain name to use for the cookies
SERVER_HOST=0.0.0.0 # the address to bind the backend on
SERVER_PORT=3000 # the port of the backend
SERVER_HTTPS=false # set to true if the your server use SSL (IT MUST BE ACTIVATED IN PRODUCTION)

TOR_ADDRESS=7at4g6c1312grosobafach5goaacab3ab5a2fcvoocvfu474jynbgzyd.onion # tor address
DB_URL=postgresql://postgres:12345@bloc-db:5432/bloc # address of the database

STORAGE_DIR=/opt/bloc/files # path to your file server/storage folder where all the files will be stored
STORAGE_QUOTA=4096 # Total file size limit for users in Mb
```

##### Fontend
*theses are the envrionment variables for the frontend (in `bloc-front` in the `docker-compose.yml` file)*
```sh
API_BASE=http://bloc.coldwire.org/api # address of the backend server, if you want to change it you must re-build the container by excuting "docker-compose build" command
```

## A note about security
/!\ **SSL MUST BE ACTIVATED IN PRODUCTION** 

### Our encryption protocol:
1. All the users have their own Private / Public keypair generated using secp256k1
2. Each files have their own encryption key, theses keys are encrypted usign the user public key.
3. The files are encrypted using Xchacha20-poly1305.
4. The private key is encrypted with Xchacha20-poly1305 and the encryption key is a 32bytes argon2 derivation of the user's password.

Currently all the opperations are done on the server side, we are considering the possibility to build a WASM library in Rust to makes all the operations on the user side, which would makes easier to trust the servers but would be impossible for users who are disabling javascript (with TOR for exemple)

## License

This project is licensed under the [NPOSL 3.0](https://opensource.org/licenses/NPOSL-3.0) License.<br>
The logo and all associated visuals are under the [CC BY-NC-ND 3.0](https://creativecommons.org/licenses/by-nc-nd/3.0/) License.