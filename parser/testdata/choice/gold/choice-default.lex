module "module"
[ident] "choice-def"...
{ "{"
container "container"
[ident] "transfer"
{ "{"
choice "choice"
[ident] "how"
{ "{"
default "default"
[string] "interval"
; ";"
case "case"
[ident] "interval"
{ "{"
leaf "leaf"
[ident] "interval"
{ "{"
type "type"
[ident] "uint16"
; ";"
units "units"
[string] "minutes"
; ";"
default "default"
[string] "30"
; ";"
} "}"
} "}"
case "case"
[ident] "daily"
{ "{"
leaf "leaf"
[ident] "daily"
{ "{"
type "type"
[ident] "empty"
; ";"
} "}"
leaf "leaf"
[ident] "time-of-da"...
{ "{"
type "type"
[ident] "string"
; ";"
units "units"
[string] "24-hour-cl"...
; ";"
default "default"
[string] "\"01.00\""
; ";"
} "}"
} "}"
case "case"
[ident] "manual"
{ "{"
leaf "leaf"
[ident] "manual"
{ "{"
type "type"
[ident] "empty"
; ";"
} "}"
} "}"
} "}"
} "}"
} "}"