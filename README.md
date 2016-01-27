# jsonpath

Find a path to a key in nested JSON structures.

## Example

```
$ cat example.json
{"items":{"item":[{"id":"0001","type":"donut","name":"Cake","ppu":0.55,"batters":{"batter":[{"id":"1001","type":"Regular"},{"id":"1002","type":"Chocolate"},{"id":"1003","type":"Blueberry"},{"id":"1004","type":"Devil's Food"}]},"topping":[{"id":"5001","type":"None"},{"id":"5002","type":"Glazed"},{"id":"5005","type":"Sugar"},{"id":"5007","type":"Powdered Sugar"},{"id":"5006","type":"Chocolate with Sprinkles"},{"id":"5003","type":"Chocolate"},{"id":"5004","type":"Maple"}]}]}}
$ jsonpath topping < example.json
.items.item[0].topping
$ jsonpath type < example.json
.items.item[0].type
.items.item[0].batters.batter[0].type
.items.item[0].batters.batter[1].type
.items.item[0].batters.batter[2].type
.items.item[0].batters.batter[3].type
.items.item[0].topping[0].type
.items.item[0].topping[1].type
.items.item[0].topping[2].type
.items.item[0].topping[3].type
.items.item[0].topping[4].type
.items.item[0].topping[5].type
.items.item[0].topping[6].type
```
