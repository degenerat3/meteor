# Client Spec

 - Each client must have a [Universally unique identifier](https://en.wikipedia.org/wiki/Universally_unique_identifier) that will be used by the database for tracking purposes.  This UUID should remain constant accross requests, meaning it should be generated ONCE at registration time, then re-used for each callback.  Creation of a new UUID per request would technically work, but it will cause *massive* bloat in the database, and as such is not suggested.

 - All clients must be able to recieve an action, execute it, then respond to their module with the result.

 - All clients must be able to handle every "Opcode" defined in the MAD documentation (except for the ones labeled TBD, obviously).

 - It is suggested that each client has a defined callback interval/delta, but in cases where that is not implemented (ex: callbacks from shims, dotfiles, etc), those values are still required by the registration endpoint and the value of '0' should be used.