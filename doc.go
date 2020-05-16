/*

Package swgoh and it's subfolders contains implementations to parse
and interact with EA's game "Star Wars Galaxy of Heroes".


Obtaining game data

Game developers allow certain data to be otainable from associated
API (Application Programming Interface) developed and hosted by third-parties.

The two main API providers are https://swgoh.gg/api and https://api.swgoh.help,
each with specific goodies and behavior but both allowing public data
from each player account to be queried.

This package contains client-side Go implementations for both APIs
as well as some developer friendly goodies such as built-in cache
using etcd BoltDB, helper methods to filer characters and mods,
helper methods to "guess" character aliases and "slang" generally
used by players, and a few others.


Reference implementation

The cmd/swgoh program provided with the source code contains also
a reference implementation that allows one to query data from the
CLI (command-line interface).


Learn more

To learn more about this implementation, please join our Discord
server channel at the SWGOH.HELP server:

	https://discord.gg/7VvSPEZ

*/
package swgoh
