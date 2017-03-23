# GOML - A YAML manipulation tool written in GO

## Usage

All example will base on the following example yaml:

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

## Get Property

```
$ goml get -file <yaml-file> -prop <path.to.property>
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


## Set Property

**Basic Syntax:**

```
$ goml set -file <yaml-file> -prop <path.to.property> -value <new-value>
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

## Delete Property

**Not Implemented** 
