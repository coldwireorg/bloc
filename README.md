# Bloc
## Encrypted and resilient file storage

### Description and Security

Bloc is mainly made for people who need a safe place for their datas, like movements and activists organizing in risk zones (ARZ) or even journalists and whistleblowers needing a safe place to store their documents, proofs, etc.

To archieve that, every files are encrypted using ChaCha20-poly1305 on the client side, with random key generated on the clients side and these keys are encrypted using the public key of the user (ECIES with ed25519 keypair), which is also done on the client side, the server only have access to the ecrypted keys and files, which make it impossible to the server to read private datas of the users. This is made possible using [liboxyd](https://codeberg.org/coldwire/liboxyd). The actual big problem is that we are forced to use Javascript with webassembly and a web browser is not the best kind of software for security, so we are planning to build a rudimentary desktop app in rust to mount bloc with fuse with the possibility to root the trafic through tor or any other onion network (lokinet ?).

Our second security layer is the resiliency of the storage, the goal is that if a government or a company raid the servers where your instance of bloc is hosted, you just need to login with the same username and password on another instance and *magic* all your files are still here! This will only be archivable using the [Polar network (WIP)](https://codeberg.org/coldwire/polar/src/branch/main/dev/paper.md) and also possible to do using a self-hosted S3 infrastructure but polar is attended to be built over a big network of users and organization so it makes it harder to shutdown.

### Config

An exemple config file can be found [here](./exemple.config.toml)

### Building

just clone the repo and build with go <1.16

```sh
git clone https://codeberg.org/coldwire/bloc && cd bloc
go build main.go
```

then run it
```sh
./main.go
```

### Deploying

#### Docker

// TODO

#### Nomad
Coldwire' internal infrastructure is using nomad, so for those who are using it too, you're lucky, here is a sample config: 
*no yet*

// TODO

### API

⚠️ Since the encryption is done on the client side, if you are sending a file without encrypting it yourself, it not be encrypted byt the server !! When developping a client, you MUST inclust file encryption to your software!!

> ### Register a new user
> Description: If oauth2 is not enable, users need to register theu account directly on bloc<br>
>
> Method: **POST**<br>
> Endpoint: **/api/user/auth/register**<br>
> Body:
> ```json
> {
>   "username": "<username>",
>   "password": "<password>",
>   "auth_mode": "LOCAL",
>   "private_key": "<private key>",
>   "public_key": "<private key>",
> }
> ```
>
> Response *200*:
> ```json
> {
>   "success": true,
>   "data": {
>     "token": "<jwt token>"
>   }
> }
> ```
>
> Response *50X*:
> ```json
> {
>   "success": false,
>   "error": "<error message>",
> }
> ```

> ### Login user
> Description: If oauth2 is not enable, users need to login directly on bloc<br>
>
> Method: **POST**<br>
> Endpoint: **/api/user/auth/login**<br>
> Body:
> ```json
> {
>   "username": "<username>",
>   "password": "<password>",
>   "private_key": "<private key>",
>   "public_key": "<private key>",
> }
> ```
>
> Response *200*:
> ```json
> {
>   "success": true,
>   "data": {
>     "token": "<jwt token>"
>   }
> }
> ```
>
> Response *50X*:
> ```json
> {
>   "success": false,
>   "error": "<error message>",
> }
> ```

> ### Logout user
> Description: If users want to logout they just need to visit this address<br>
>
> Method: **GET**<br>
> Endpoint: **/api/user/auth/logout**<br>

> ### Oauth2 callback
> Description: Callback address to specify in the oauth2 provider<br>
>
> Method: **GET**<br>
> Endpoint: **/api/user/auth/oauth2/callback**<br>

> ### Upload a file
> Description: Upload a file<br>
>
> Method: **POST**<br>
> Endpoint: **/api/user/auth/login**<br>
> Form values:
>   - *file* : the file to upload
>   - *parent* : Parent directory
>   - *key* : Encrypted encryption key of the file
>
> Response *200*:
> ```json
> {
>   "success": true,
>   "data": {
>     "id": "<file id>"
>   }
> }
> ```
>
> Response *50X*:
> ```json
> {
>   "success": false,
>   "error": "<error message>",
> }
> ```