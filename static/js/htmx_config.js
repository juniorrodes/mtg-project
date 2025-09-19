htmx.config.responseHandling = [
    {code: "404", swap: true},
    {code:"[23]..", swap: true}, // 200 & 300 responses are non-errors and are swapped            
]
