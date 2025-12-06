# Gator RSS feed aggregator
You will need postgres and go installed to run gator.
See the [docs](https://go.dev/doc/tutorial/compile-install) to install with `go install`

## config and command
the config file is at HOME/.gatorconfig.json 
Here is an example to get you started. Replace the username, password,and port with relevant entries.
```json
{"db_url":"postgres://postgres:username@password:port/gator?sslmode=disable","current_user_name":""}
```
```

you can run some commands:
  1. `gator addfeed "feed name" "feedURL"` //add a feed to the db
  2. `gator agg` //add posts from feeds to the db
  3. `gator browse <limit>` // print out a limited number of posts from saved posts
