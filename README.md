# go-struct-registry

_Experimental, and unmaintained. Use at your own risk._

A simple "global" Go struct registry, and an accompanying code generator to register structs for dynamic initialisation.

---

`go-struct-registry` was conceived out of an attempt at building a composable platform for bridging services with some level of type safety. Passing messages from producer to consumer for example, is not often a type safe operation (though, that is something Avro intends to solve). It is now left as remnants of an experiment into Go's type introspection, and reflection capabilities.
