Compares the canonical XML structure between two files but ignores the text/data between the tags.

Example usage:
go run main.go -file_a="foo.xml" -file_b="bar.xml"

Example file A:

```
<?xml version="1.0" encoding="UTF-8"?>
<contacts>
    <contact>
        <name>Doe John</name>
        <email>john.doe@example.com</email>
        <phone>123-456-7890</phone>
    </contact>
</contacts>
```

Example file B:

```
<?xml version="1.0" encoding="UTF-8"?>
<contacts>
    <contact>
        <name>Doe John</name>
        <email>doe.john@example.com</email>
        <phoneNumber>123-654-4323</phoneNumber>
    </contact>
</contacts>
```

Output:

![Example Image](https://i.imgur.com/o0LUNDS.png)
