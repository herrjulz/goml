# GOML - A CRUD tool for YAML files

With `goml` you can:

- **C** reate YAML properties (option `set`)
- **R** etrieve YAML properties (option `get`)
- **U** pdate YAML properties (option `set`)
- **D** elete YAML properties (option `delete`)

Additionally, you can **transfer** properties from one YAML to another YAML

## Installation

### OS X

```bash
$ wget -O /usr/local/bin/goml https://github.com/JulzDiverse/goml/releases/download/v0.5.0/goml-darwin-amd64 && chmod +x /usr/local/bin/goml
```

**Using Homebrew:**

```bash
$ brew tap julzdiverse/tools  
$ brew install goml
```

### Linux

```bash
$ wget -O /usr/bin/goml https://github.com/JulzDiverse/goml/releases/download/v0.5.0/goml-linux-amd64 && chmod +x /usr/bin/goml
```

## Usage

All examples are based on the following yaml:

```bash
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

```bash
$ goml get --file <yaml-file> --prop <path.to.property>
```

### Examples:

**Maps**

Get value from map:

```bash
$ goml get -f sample.yml -p map.foo.bar
# returns value
```
**Arrays**

Get value from array:

```bash
$ goml get -f sample.yml -p array.1
# returns two
```

**Map Arrays**

Get value from array which contains maps:

```bash
$ goml get -f sample.yml -p mapArray.0.name
# returns foo
```

Get value from array which contains maps by an identifier:

```bash
$ goml get -f sample.yml -p mapArray.id:two.name
# returns bar
```


## Create/Update Properties with `set`

You can use the `set` option to either set or update properties. If an property for a valid path does not exist, goml will create it for you. If an property exists the `set` option will upate this property with the provided value.

Another useful thing you can do with `set` is to set/update a key (eg. ssh private key) from a file.  

**Basic Syntax:**

```bash
$ goml set --file <yaml-file> --prop <path.to.property> --value <new-value>
$ goml set --file <yaml-file> --prop <path.to.property> --key <key-file>
```

Alternatively, you can provide the value directly with the property string using `=`, like this:

```bash
$ goml set --file <yaml-file> --prop <path.to.property=value>
```

_Note: The `--value|-v` option has precedence_

**The `dry-run` option**

By default `set` updates a yaml file in place. With the `dry-run` option you can print the result to `stdout`. To do so provide the `--dry-run` option on a `goml set` call:

```bash
$ goml set --file <yaml-file> --prop <path.to.property> --value <new-value> --dry-run
```

This can be useful to experiment with `goml` or to pipe the results into a new file or other yaml tools.

### Examples:

**Maps**

Update value to a map:

```bash
$ goml set -f sample.yml -p map.foo.bar -v newValue
# or
$ goml set -f sample.yml -p map.foo.bar=newValue
# will update the value 'value' of 'bar' to 'newValue'
```

Add value to a map:

```bash
$ goml set -f sample.yml -p map.foo.newProp -v value
# or
$ goml set -f sample.yml -p map.foo.newPropv=value
# will add a property 'newProp' with value 'newValue'
```

**Arrays**

Update value to an array by index:

```bash
$ goml set -f sample.yml -p array.0 -v newValue
# or
$ goml set -f sample.yml -p array.0=newValue
# this will update the array value 'one' with 'newValue'
```

Update value to an array by identifier:

```bash
$ goml set -f sample.yml -p array.:three -v newValue
# or
$ goml set -f sample.yml -p array.:three=newValue
# this will update the array value 'three' with 'newValue'
```

Add value to array

```bash
$ goml set -f sample.yml -p array.+ -v newValue
# or
$ goml set -f sample.yml -p array.+=newValue
# this will add the value 'newValue' to the array
```

**Map Arrays**

Update a property by index:

```bash
$ goml set -f sample.yml -p mapArray.0.name -v julz
# or
$ goml set -f sample.yml -p mapArray.0.name=julz
# this will update the property 'name' with value 'julz'
# for the first entry in the array
```

Update a property by identifier:

```bash
$ goml set -f sample.yml -p mapArray.id:one.name -v julz
# or
$ goml set -f sample.yml -p mapArray.id:one.name=julz
# this will update the property 'name' with value 'julz'
# for the entry that has id 'one'
```

Add a map to an array:

```bash
# as goml creates every key provided in the property parameter it is as easy as:
$ goml set -f sample.yml -p mapArray.newKey:newValue
```

_Note that no value parameter is required to add a new map entry_

**Set key from file `--key, -k`**

```bash
$ goml set -f sample.yml -p mapArray.id:one.name -k keyfile
```

## Delete Properties with `delete`

```bash
$ goml delete --file <yaml-file> --prop <path.to.property>
```

### Examples:

**Maps**

Delete value from map:

```bash
$ goml delete -f sample.yml -p map.foo.bar
# deletes value
```
**Arrays**

Delete value from array:

```bash
$ goml delete -f sample.yml -p array.1
# deletes two
```

Delete value from array which contains maps:

```bash
$ goml delete -f sample.yml -p array.0.name
# deletes foo
```

**Map Arrays**

Delete value from array which contains maps by an identifier:

```bash
$ goml delete -f sample.yml -p array.id:two.name
# deletes bar
```

## Transfer Properties with `transfer`

Transfer is the same as `set`, with the difference that you specify a destination file and property instead of a value:

```bash
$ goml transfer --file <yaml-file> --prop <path.to.property> --df <destination-file> --dp <destination-property>
```

*Note:*
- The syntax for the source property is the same as for  `set`
- The syntax for the destination property is the same as for `get`
