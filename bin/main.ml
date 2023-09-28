open Core
open Lib.FsHelpers
(* open Cmdliner *)
open Sys_unix

let target_dir = match Sys.getenv "HOME" with
    | Some v -> v
    | None -> "./"
;;

let repos = match file_exists target_dir with
    | `Yes -> git_dirs target_dir
    | _ -> []
;;

let () = List.iter repos ~f: print_endline
