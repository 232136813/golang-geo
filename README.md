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

♫ around the world ♪
```
[![Build Status](https://travis-ci.org/kellydunn/go-art.png)](https://travis-ci.org/kellydunn/golang-geo)
# what 

This library provides convenience functions for translating, geocoding, and calculating distances between geographical points.  It is inspired by ruby's `geokit` and `geokit-rails` gems, and aims to make dealing with geographical data a little bit easier.

# documentation

You can read the documentation [here](http://godoc.org/github.com/kellydunn/golang-geo).

# usage

Import from github, and get geomancin'

```
import("github.com/kellydunn/golang-geo")
```

Currently, `golang-geo` provides the following functionality:

  - Querying for points within a radius using your own SQL data tables.
  - Transposing a point for a given distance and bearing.
  - Calculating the Great Circle Distance bewteen two points.
  - Geocoding an address using Google Maps API or Open Street Maps API.
  - Reverse Geocoding a Point using the same services.

Keep in mind that you do not need to use SQL in order to perform [simple Point operations](http://godoc.org/github.com/kellydunn/golang-geo#Point) and the only function that relies on SQL is [`PointsWithinRadius`](http://godoc.org/github.com/kellydunn/golang-geo#SQLMapper.PointsWithinRadius). 

## using SQL

The project is configured to connect to a SQL database by reading a `config/geo.yml` file in the root level of your project.  If it does not exist, it will use a Default SQL configuration that will use the postgres driver as described by [lib/pq](http://github.com/lib/pq).  The Default SQL configuration will attempt to connect as a user named "postgres" and with the password "postgres" to a database named "points".  

If you want to supply a custom database conifguration, feel free to do so by using the template below:

```
development:
  driver: postgres
  openStr: user=username password=password dbname=points sslmode=disable
  table: points
  latCol: lat
  lngCol: lng
```

Or if you want to connect via MySQL, you can do that as well!

```
development:
  driver: mysql
  openStr: points/username/password
  table: points
  latCol: lat
  lngCol: lng  
```

Once you've supplied your configuration, you may connect to your database with the following line of code:

```
db, err := geo.HandleWithSQL()
```

# notes

  - `golang-geo` currently only uses metric measurements to do calculations
  - The `$GO_ENV` environment variable is used to determine which configuration group in `config.yml` is to be used.  For example, if you wanted to use the PostgreSQL configuration listed above, you could specify `GO_ENV=development` which would read `config.yml` and use the configuration under the root-level key `development`.

# roadmap
  - More Tests!
  - Redis / NOSQL Mapper
  - Bing Maps?
  - Add an abstraction layer for PostgreSQL earthdistance / PostGIS

# testing

By default, `golang-geo` will attempt to run its test suite against a PostgreSQL database.  However, you may run the tests with mocked SQL queries by specifying that you want to do so on the command line:

```
DB=mock go test
```

The `$DB` environment variable is used to specify which database you'd like to run the tests against.  You may specify `postgres`, `mysql`, or `mock`.  The [Travis CI builds](https://travis-ci.org/kellydunn/golang-geo) for this project currently runs against all of these when running the test suite.

# contributing
  - Fork
  - Create a topic branch
  - Make dem commits!
  - Write dem tests!
  - Submit Pull Request once Tests are Passing
  - do this (づ￣ ³￣)づ

Thanks! ｡◕‿◕｡