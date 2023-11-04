open Core
open Printf
open Lib.FsHelpers

(* open Cmdliner *)
open Sys_unix

let target_dir =
  match Sys.getenv "HOME" with
  | Some v -> v
  | None -> "./"
;;

let ignore =
  [ "/home/ajlow/.local/share"; "/home/ajlow/.zplug"; "home/ajlow/.nix-defexpr" ]
;;

let () = printf "Target Dir: %s\n-------------------\n" target_dir

let repos =
  match file_exists target_dir with
  | `Yes -> git_dirs target_dir ~ignore
  | _ -> []
;;

let () = List.iter repos ~f:print_endline
