### Sync from Deezer to Spotify

Small toy project to discover Golang language and also to migrate a full playlist from Deezer to Spotify.

### How to get Deezer Code

Visit this page using a browser to get Deezer Code:

`https://connect.deezer.com/oauth/auth.php?app_id=<APP_ID>redirect_uri=<REDIRECT_URI>&perms=basic_access,email`

### How to get Deezer Access token

`https://connect.deezer.com/oauth/access_token.php?app_id=<APP_ID>&secret=<SECRET>&code=<CODE>`

### How to run this script

`go run .`