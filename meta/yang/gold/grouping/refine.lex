module "module"
[ident] "refine"
{ "{"
namespace "namespace"
[string] "\"\""
; ";"
prefix "prefix"
[string] "\"\""
; ";"
revision "revision"
[revision] "0"
; ";"
grouping "grouping"
[ident] "x"
{ "{"
leaf "leaf"
[ident] "s"
{ "{"
description "descriptio"...
[string] "\"orig\""
; ";"
type "type"
[ident] "string"
; ";"
} "}"
container "container"
[ident] "t"
{ "{"
description "descriptio"...
[string] "\"orig II\""
; ";"
leaf "leaf"
[ident] "u"
{ "{"
description "descriptio"...
[string] "\"orig III\""
; ";"
type "type"
[ident] "string"
; ";"
} "}"
} "}"
} "}"
container "container"
[ident] "a"
{ "{"
uses "uses"
[ident] "x"
{ "{"
refine "refine"
[path] "s"
{ "{"
description "descriptio"...
[string] "\"refined\""
; ";"
} "}"
refine "refine"
[path] "t/u"
{ "{"
description "descriptio"...
[string] "\"refined I"...
; ";"
} "}"
} "}"
} "}"
container "container"
[ident] "b"
{ "{"
uses "uses"
[ident] "x"
; ";"
} "}"
} "}"
