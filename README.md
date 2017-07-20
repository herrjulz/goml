# GOML - A CRUD tool for YAML files

With `goml` you can:

- **C** reate YAML properties (option `set`)
- **R** etrieve YAML properties (option `get`)
- **U** pdate YAML properties (option `set`)
- **D** elete YAML properties (option `delete`)

Additionally, you can **transfer** properties from one YAML to another YAML

## Installation

### OS X

```
$ wget -O /usr/local/bin/goml https://github.com/JulzDiverse/goml/releases/download/v0.2.0/goml-darwin-amd64 && chmod +x /usr/local/bin/goml
```

**Using Homebrew:**

```
$ brew tab julzdiverse/tools  
$ brew install goml
```

### Linux

```
$ wget -O /usr/bin/goml https://github.com/JulzDiverse/goml/releases/download/v0.2.0/goml-linux-amd64 && chmod +x /usr/bin/goml
```

## Usage

All examples are based on the following yaml:

```
map:
  foo:
    bar: value

array:
- one
- two
- three

mapArray:
- id: one
  name: foo
- id: two
  name: bar
```

## Retrieve Properties with `get`

```
$ goml get --file <yaml-file> --prop <path.to.property>
```

### Examples:

**Maps**

Get value from map:

```
$ goml get -f sample.yml -p map.foo.bar
// returns value
```
**Arrays**

Get value from array:

```
$ goml get -f sample.yml -p array.1
// returns two
```

Get value from array which contains maps:

```
$ goml get -f sample.yml -p array.0.name
// return foo
```

**Map Arrays**

Get value from array which contains maps by an identifier:

```
$ goml get -f sample.yml -p array.id:two.name
// returns bar
```


## Create/Update Properties with `set`

You can use the `set` option to either set or update properties. If an property for a valid path does not exist, goml will add it for you. If an property exists the `set` option will upate this property with the provided value.

Another useful thing you can do with `set` is to set/update a key (eg. ssh private key) from a file.  

**Basic Syntax:**

```
$ goml set --file <yaml-file> --prop <path.to.property> --value <new-value>
$ goml set --file <yaml-file> --prop <path.to.property> --key <key-file>
```

### Examples:

**Maps**

Update value to a map:

```
$ goml set -f sample.yml -p map.foo.bar -v newValue
// will update the value 'value' of 'bar' to 'newValue'
```

Add value to a map:

```
$ goml set -f sample.yml -p map.foo.newProp -v value
// will add a property 'newProp' with value 'newValue'
```

**Arrays**

Update value to an array by index:

```
$ goml set -f sample.yml -p array.0 -v newValue
// this will update the array value 'one' with 'newValue'
```

Update value to an array by identifier:

```
$ goml set -f sample.yml -p array.:three -v newValue
// this will update the array value 'three' with 'newValue'
```

Add value to array

```
$ goml set -f sample.yml -p array.+ -v newValue
// this will add the value 'newValue' to the array
```

**Map Arrays**

Update a property by index:

```
$ goml set -f sample.yml -p mapArray.0.name -v julz
// this will update the property 'name' with value 'julz'
// for the first entry in the array
```

Update a property by identifier:

```
$ goml set -f sample.yml -p mapArray.id:one.name -v julz
// this will update the property 'name' with value 'julz'
// for the entry that has id 'one'
```

**Set key from file `--key, -k`**

```
$ goml set -f sample.yml -p mapArray.id:one.name -k keyfile
```

## Delete Properties with `delete`

```
$ goml delete --file <yaml-file> --prop <path.to.property>
```

### Examples:

**Maps**

Delete value from map:

```
$ goml delete -f sample.yml -p map.foo.bar
// deletes value
```
**Arrays**

Delete value from array:

```
$ goml delete -f sample.yml -p array.1
// deletes two
```

Delete value from array which contains maps:

```
$ goml delete -f sample.yml -p array.0.name
// deletes foo
```

**Map Arrays**

Delete value from array which contains maps by an identifier:

```
$ goml delete -f sample.yml -p array.id:two.name
// deletes bar
```

## Transfer Properties with `transfer`

Transfer is the same as `set`, with the difference that you specify a destination file and property instead of a value:

```
$ goml transfer --file <yaml-file> --prop <path.to.property> --df <destination-file> --dp <destination-property>
```

*Note:*
- The syntax for the source property is the same as for  `set`
- The syntax for the destination property is the same as for `get`
