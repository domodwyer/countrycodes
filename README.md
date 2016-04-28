[![Build Status](https://travis-ci.org/domodwyer/countrycodes.svg?branch=master)](https://travis-ci.org/domodwyer/countrycodes)

#Country Codes
Quick library file for translating __ISO 3166-1 ALPHA-2__ country codes to their readable names and back again.

##Usage
Fetch with `go get github.com/domodwyer/countrycodes`

```
name, err := countrycodes.ToName("GB")
if err != nil {
	panic("Country code not found :( ")
}

fmt.Printf("GB is the country code for %s", name)
```
