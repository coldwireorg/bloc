# Bloc API

⚠️ Since the encryption is done on the client side, if you are sending a file without encrypting it yourself, it not be encrypted byt the server !! When developping a client, you MUST inclust file encryption to your software!!

### Result
The API alway return 4 fields: 
```json
{
	"status": "<STATUS_OF_THE_RESPONSE>",
	"message": "<just a message to describe the result>",
	"content": "<can be a string, an array, object or just null, depend of the kind of reponse>",
	"error": "<most of the time empty but can containe more informations about an error>"
}
```

### Auth **/api/auth**
> ### **POST** /register
> To register a new user to the instance
> ```json
> {
>   "username": "<username>",
>   "password": "<password">,
>   "private_key": "<private_key>",
>   "public_key": "<public_key>"
> }
> ```

> ### **POST** /login
> To login the instance
> ```json
> {
>   "username": "<username>",
>   "password": "<password">
> }
> ```

> ### **POST** /oauth/callback
> Url to set as callback url for openid providers

> ### **GET** /logout
> To logout of the instance

> ### **GET** /oauth
> To get the authorization url


### Files **/api/files** (require a valid JWT token in a cookie)
> ### **POST** /
> Upload a file
> Form inputs:
>> *parent* -> parent folder of the file to upload<br>
>> *key* -> Encryption key of the file encrypted with user's public key

> ### **DELETE** /:id
> Delete a file with it ID

> ### **GET** /list/:folder
> List files in a folder

> ### **GET** /downlaoder/:id
> Download a file with its id

> ### **PUT** /favorite/:id
> Toogle if a file is favorite or not
> ```json
> {
>   "favorite": "<true | false>"
> }
> ```

### Folders **/api/folders** (require a valid JWT token in a cookie)
> ### **POST** /
> Create a new folder
> ```json
> {
>   "name": "<name of the folder>",
>   "parent": "<ID of the parent folder>"
> }
> ```

> ### **PUT** /:id
> Move folder to another one
> ```json
> {
>   "parent": "<ID of the new parent folder>"
> }
> ```

> ### **DELETE** /:id
> Delete a folder with id ID and all its childrens (files and folders)

### Shares **/api/shares** (require a valid JWT token in a cookie)
> ### **POST** /
> Create a new folder
> ```json
> {
>   "is_file": "<true | false>",
>   "to_share": "<ID of the file to share>",
>   "share_to": "<username of the person to share the file to>",
>   "key": "<key of the file encrypted with user's public key>"
> }
> ```

> ### **DELETE** /:id
> Revoke a share by with its id

> ### **GET** /
> List all shared a user have access to