(*
** wombats
** author: Ryan King
**
** parse.sats: Definitions for argument parsing
*)

abstype args_type = ptr
typedef args = args_type

datatype command =
  | new of (string)
  | init of (args)
  | build of (args)
  | run of (args)
  | clean of ()
  | install of (string, args)
  | update of (args)
  | package of (string)
  | search of (string, args)

fun get_cmd : {n:nat} (int(n), !argv(n)) -> command