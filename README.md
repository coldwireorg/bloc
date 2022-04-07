# Bloc
## Encrypted and resilient file storage

### Description and Security

Bloc is mainly made for people who need a safe place for their datas, like movements and activists organizing in risk zones (ARZ) or even journalists and whistleblowers needing a safe place to store their documents, proofs, etc.

To archieve that, every files are encrypted using ChaCha20-poly1305 on the client side, with random key generated on the clients side and these keys are encrypted using the public key of the user (ECIES with ed25519 keypair), which is also done on the client side, the server only have access to the ecrypted keys and files, which make it impossible to the server to read private datas of the users. This is made possible using [liboxyd](https://codeberg.org/coldwire/liboxyd). The actual big problem is that we are forced to use Javascript with webassembly and a web browser is not the best kind of software for security, so we are planning to build a rudimentary desktop app in rust to mount bloc with fuse with the possibility to root the trafic through tor or any other onion network (lokinet ?).

Our second security layer is the resiliency of the storage, the goal is that if a government or a company raid the servers where your instance of bloc is hosted, you just need to login with the same username and password on another instance and *magic* all your files are still here! This will only be archivable using the [Polar network (WIP)](https://codeberg.org/coldwire/polar/src/branch/main/dev/paper.md) and also possible to do using a self-hosted S3 infrastructure but polar is attended to be built over a big network of users and organization so it makes it harder to shutdown.

### About encryption

The base of the threat model: everything happen on the client side, so we don't need to trust the server, as mentioned earlier the web browser can't be 100% trusted, so a desktop client with no web tech will be written.

#### User's key pair
When a user register on one of the services of coldwire or on the auth server, an [ed25519](https://en.wikipedia.org/wiki/EdDSA) keypair is generated (on the client side as always), this pair will be useful for signing files chunks checksum on the [Polar network (WIP)](https://codeberg.org/coldwire/polar/src/branch/main/dev/paper.md) and for encrypting key files using [ECIES](https://en.wikipedia.org/wiki/Integrated_Encryption_Scheme)

The private key is stored but encrypted with ChaCha20-Poly1305 with blake3 hash of the password as the key.

The public key is not encrypted and can be get by anyone, it will be used for files transfer or others future features.

#### How files are encrypted ?
1. A 256bits key is generated
2. The file is encrypted chunk by chunk using ChaCha20-Poly1305 with the generated key
3. the key is encrypted using user's public key with [ECIES](https://en.wikipedia.org/wiki/Integrated_Encryption_Scheme)
4. The file is uploaded and the encrypted key stored in the metadatas.

#### How files are shared ?
There is two kind of sharing, public and private.

A private sharing is the way of sharing files between users of bloc (will work betweens differents instances over polar), this is how it work:

1. Bob click on "sharing to alice"
2. Bob get alice's public key from the API
3. Bob decrypt the key of the file with its private key
4. Bob encrypt the key of the file using alice's public key
5. The file is added to alice's shared files

For a public share, the current way is just to share the encryption key in a link, but I'm (monoko) thinking about a way to share without leaking the original key, but seems complicated with decryption on the client side.

#### Abouy password ?

For now, password is the only thing managed on the server side, password are hashed using argon2id.
If you think a little bit, you understand that there is a big problem, if an admin of an instance intercept the password, everything is fucked, so this will change, I just don't know how much it's safe or not to send an argon2id hash from the client. 

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
>
> Response *50X*:
> ```json
> {
>   "success": false,
>   "error": "<error message>",
> }
> ```

> ### Oauth2 url
> Description: Get oauth2 authentication url<br>
>
> Method: **GET**<br>
> Endpoint: **/api/user/auth/oauth2**<br>

> ### Oauth2 callback
> Description: Callback address to specify in the oauth2 provider<br>
>
> Method: **GET**<br>
> Endpoint: **/api/user/auth/oauth2/callback**<br>

> ### Upload a file
> Description: Upload a file<br>
>
> Method: **POST**<br>
> Endpoint: **/api/file**<br>
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

> ### Download file
> Description: Download a file from its ID<br>
>
> Method: **GET**<br>
> Endpoint: **/api/file/donwload/:id**<br>
>
> Response *50X*:
> ```json
> {
>   "success": false,
>   "error": "<error message>",
> }
> ```

> ### Toogle favorite
> Description: Set if a file is favorite or not<br>
>
> Method: **PUT**<br>
> Endpoint: **/api/file/favorite/:id**<br>
>
> Response *200*:
> ```json
> {
>   "success": true
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

> ### Delete file
> Description: Delet a file from its ID<br>
>
> Method: **DELETE**<br>
> Endpoint: **/api/file/:id**<br>
>
> Response *200*:
> ```json
> {
>   "success": true
> }
> ```

> ### Move file
> Description: Download a file from its ID<br>
>
> Method: **PUT**<br>
> Endpoint: **/api/file/:id**<br>
> Body:
> ```json
> {
>   "parent": "<new parent folder>",
> }
> ```
>
> Response *200*:
> ```json
> {
>   "success": true
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

> ### Create folder
> Description: Create a new folder<br>
>
> Method: **POST**<br>
> Endpoint: **/api/folder**<br>
> Body:
> ```json
> {
>	  "name": "<folder name>",
>	  "parent": "<parent folder>"
> }
> ```
>
> Response *200*:
> ```json
> {
>   "success": true,
>   "data": {
>     "id":"<new folder id>",
>     "name": "<new folder name>",
>     "parent": "<parent of the new folder>",
>     "owner": "<owner of the new folder>"
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

> ### Move folder
> Description: Move a folder to another one<br>
>
> Method: **PUT**<br>
> Endpoint: **/api/folder/:id**<br>
> Body:
> ```json
> {
>	  "parent": "<new parent folder>",
> }
> ```
>
> Response *200*:
> ```json
> {
>   "success": true
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

> ### Delete folder
> Description: Delete a folder and its childrens files and folders<br>
>
> Method: **PUT**<br>
> Endpoint: **/api/folder/:id**<br>
>
> Response *200*:
> ```json
> {
>   "success": true
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
