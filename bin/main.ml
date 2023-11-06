open Core
open Printf
open Lib.FsHelpers
open Sys_unix
(* open Cmdliner *)

let target_dir =
  match Sys.getenv "HOME" with
  | Some v -> v
  | None -> "./"
;;

let ignore =
  [ "/home/ajlow/.local/share"; "/home/ajlow/.zplug"; "home/ajlow/.nix-defexpr" ]
;;

let repos =
  printf "Target Dir: %s\n-------------------\n" target_dir;
  match file_exists target_dir with
  | `Yes -> git_dirs target_dir ~ignore:ignore
  | _ -> []
;;

let print_entry path = 
    let width = Core_unix.system "tput cols" in
    let rec string n s =
        if n = 0 then "" else s ^ string (n - 1) s
    in
    printf "%s\n" (string (width / 2) "<>");
    printf "%s\n" path;
    printf "%s\n\n" "- - - - - - - - - - - - - - - - - - - -"

let () = List.iter repos ~f:print_entry
