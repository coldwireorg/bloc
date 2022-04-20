# Bloc

## Table of contents

- [Introduction](#introduction)
- [Security and Details](#security-and-details)
  - [Resilience](#resilience)
  - [Encryption](#encryption)
    - [User's key pair](#)
    - [User's password](#)
    - [How file are encrypted ?](#)
  - [Sharing System](#sharing-system)
    - [Public](#)
    - [Private](#)
- [Roadmap](#roadmap)
- [Quickstart](#quickstart)
  - [Requirements](#requirements)
  - [You want to contribute to the code ?](#you-want-to-contribute-to-the-code)
  - [You want to deploy bloc in production ?](#you-want-to-deploy-bloc-in-production)
- [License](#license)

## Introduction

Bloc is mainly made for people who need a safe place for their datas, like movements and activists organizing in risk zones (ARZ) or even journalists and whistleblowers needing a safe place to store their documents, proofs, etc.

## Security and Details

### Resilience

With encryption, our second security layer is the resiliency of the storage, the goal is that if a government or a company raid the servers where your instance of bloc is hosted, you just need to login with the same username and password on another instance and magic all your files are still here! This will only be archivable using the [Polar network (WIP)](https://codeberg.org/coldwire/polar/src/branch/main/dev/paper.md) and also possible to do using a self-hosted S3 infrastructure but polar is attended to be built over a big network of users and organization so it makes it harder to shutdown.

### Encryption

> The client application is responsible for encryption and the server is responsible for authentication. We use the [Liboxyd](https://codeberg.org/coldwire/liboxyd) library made in **WebAssembly** to use complex encryption functions on web browser

Algorithms used :
  - [ECIES ed25519](https://en.wikipedia.org/wiki/EdDSA)
  - Argon2id
  - XChaCha20-Poly1305
  - Blake3

#### **User's key pair**

the user's key pair is generated with the asymetric **ECIES ed25519** algorithm and is used to decrypt and encrypt the keys of the files.

It is generated in the web browser when a user registers on the Coldwire authentication service or a specific service developed by us.

The private key is stored in the database but encrypted with **XChaCha20-Poly1305** with an argon2 derivation of the password as the key.

The public key is not encrypted and can be get by anyone, it will be used for files transfer or others future features.

#### **User's password**

the user's password is hashed with **argon2id** and is used to authenticate the user on the application to get a jwt token.

#### **How file are encrypted ?**

1. A 256bits key is generated
2. The file is encrypted chunk by chunk using **XChaCha20-Poly1305** with the generated key
3. the key is encrypted using user's public key with [ECIES](https://en.wikipedia.org/wiki/Integrated_Encryption_Scheme)
4. The file is uploaded and the encrypted key stored in the metadatas.

### Sharing System

#### Private

A private sharing is the way of sharing files between users of bloc (will work betweens differents instances over polar), this is how it work:

1. Bob click on "sharing to alice"
2. Bob get alice's public key from the API
3. Bob decrypt the key of the file with its private key
4. Bob encrypt the key of the file using alice's public key
5. The file is added to alice's shared files

#### Public

For a public share, the current way is just to share the encryption key in a link, but I'm (monoko) thinking about a way to share without leaking the original key, but seems complicated with decryption on the client side.

## Roadmap

- [x] Bloc API
- [ ] Web application
  - [ ] Frontend
  - [ ] Client Side Encryption
- [ ] Storage
  - [ ] Plugins
  - [ ] Polar Storage
  - [x] S3 Storage
  - [x] File System Storage
- [ ] Database
  - [ ] Sqlite Driver
  - [x] PostgreSQL Driver
- [x] Authentication
  - [x] Oauth2 Authentication
  - [x] Local Authentication
- [ ] Native Client (Rust)
  - [ ] Using tor/lokinet/onion routed network
  - [ ] Fuse mounting

## Quickstart

### You want to contribute to the code ?

#### **With Docker**
#### Requirements

- Linux host (virtualized or bare metal)
- 500Mo of free memory
- Docker
- Docker Compose
- Git CLI

#### Launch Application

```sh
$ git clone https://codeberg.org/coldwire/bloc.git
$ cd bloc
(bloc) $ sudo systemctl start docker
(bloc) $ sudo docker-compose -f docker-compose-dev.yml up -d --build
```

#### Stop Application

```sh
(bloc) $ sudo docker-compose -f docker-compose-dev.yml down
```

#### Cleanup

```sh
(bloc) $ sudo docker-compose -f docker-compose-dev.yml down
(bloc) $ sudo docker system prune
(bloc) $ sudo docker volume prune
```

#### **Without docker**
#### Requirements

- Linux host (virtualized or bare metal)
- 500Mo of free memory
- Go
- NodeJS
- Yarn package manager
- Git CLI

```sh
# Clone repository
$ git clone https://codeberg.org/coldwire/bloc && cd bloc

# Run frontend server
(bloc) $ cd view && yarn install && yarn run dev && cd ..

# In a second terminal
(bloc) $ DEV_FRONT_URL=http://127.0.0.1:3000/ go run main.go -config config.toml # Run the backend while proxying requests to the frontend so you can dev without rebuilding the frontend everytime :)
```

### You want to deploy bloc in production ?

#### Config

An exemple config file can be found [here](./example.config.toml)

#### Building

just clone the repo and build with go <1.16

```sh
$ git clone https://codeberg.org/coldwire/bloc && cd bloc
(bloc) $ cd view
(bloc/view) $ yarn run generate
(bloc/view) $ cd ..
(bloc) $ go build main.go
```

then run it
```sh
./main.go
```

## License

You can find the license [here](LICENSE)