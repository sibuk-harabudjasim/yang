module "module"
[ident] "recurse-2"
{ "{"
uses "uses"
[ident] "g1"
; ";"
grouping "grouping"
[ident] "g1"
{ "{"
container "container"
[ident] "a"
{ "{"
uses "uses"
[ident] "g1"
; ";"
} "}"
container "container"
[ident] "b"
{ "{"
uses "uses"
[ident] "g2"
; ";"
} "}"
} "}"
grouping "grouping"
[ident] "g2"
{ "{"
container "container"
[ident] "c"
{ "{"
uses "uses"
[ident] "g2"
; ";"
} "}"
} "}"
} "}"