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

### Develop

```sh
# In a terminal
cd view && npm i && npm run dev # Ru frontend server

# In a second terminal
DEV_FRONT_URL=http://127.0.0.1:3000/ go run main.go -config config.toml # Run the backend while proxying requests to the frontend so you can dev without rebuilding the frontend everytime :)
```

### Building

just clone the repo and build with go <1.16

```sh
git clone https://codeberg.org/coldwire/bloc && cd bloc
cd view
npm run build
cd ..
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
[Documentation](API.md)