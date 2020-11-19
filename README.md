# CSV Columns
Easily extract and access columns from a csv file



## FindColumns()

Keep in mind that the indexes in the results have the same sequence as in the `find_this` variable:

```go
var clean bool = true
var find_this []string = []string { "Name", "Date", "Counter" }
var header_row []string = []string { "Counter", "Name", "Date" }

c, i, _ := csvtools.FindColumns(find_this, header_row, clean)
```

Result will be:

```go
[ "Name", "Date", "Counter" ], [ 1, 2, 0 ]
```


### find_this
**Type**: `[]string`   
A slice of columns to find. If _empty_, module will take all the columns available in the `header_row`


### header_row
**Type**: `[]string`   
A row from csv file that you think is a header


### clean
**Type**: `bool`   
Exclude non-existent columns from the results

```go
// `clean` = false
[ "Name", "NonExistent", "Counter" ], [ 1, -1, 0 ]

// `clean` = true
[ "Name", "Counter" ], [ 1, 0 ]
```



## IndexByName()

Returns a column index by it's name. If column does not exists, returns `-1`

```go
i := csvtools.IndexByName(column_name, mapped_list)
```

### column_name
**Type**: `string`   
Name of the column you are looking for


### mapped_list
**Type**: `map[string]int`   
A row from csv file that you consider a header
