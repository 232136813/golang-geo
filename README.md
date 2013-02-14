```
              ___                                                              
             /\_ \                                                             
   __     ___\//\ \      __      ___      __               __      __    ___   
 /'_ `\  / __`\\ \ \   /'__`\  /' _ `\  /'_ `\  _______  /'_ `\  /'__`\ / __`\ 
/\ \L\ \/\ \L\ \\_\ \_/\ \L\.\_/\ \/\ \/\ \L\ \/\______\/\ \L\ \/\  __//\ \L\ \
\ \____ \ \____//\____\ \__/.\_\ \_\ \_\ \____ \/______/\ \____ \ \____\ \____/
 \/___L\ \/___/ \/____/\/__/\/_/\/_/\/_/\/___L\ \        \/___L\ \/____/\/___/ 
   /\____/                                /\____/          /\____/             
   \_/__/                                 \_/__/           \_/__/              


```

# what 

A set of convience interfaces and methods that makes geo-related calculations easier for Go.

Also just an simple experiement for me to play around with in Go.

# usage

Import from github, and get geomancin'

```
import( _ "github.com/kellydunn/golang-geo")

func main() {
     db, err := geo.HandleWithSql()

     ...

     // Find all of the points of interest that are in a 15 mile radius of [42.333, 121,111]
     p := &Point{lat: 42.3333, lng: 121.111}
     res, _ := db.Within(p, 15)
}
```

# roadmap

  - Declare your mapping service / api keys and consume Geo data as needed.
  - Determine if useful to provide a database abstraction layer for convienence 