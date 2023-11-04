open Core
open Sys_unix

(** [git_dir] recursively searches through a specificied
    [target_dir] and returns a the paths to all directories
    that contain a .git/ subdir. while ignoring all paths
    specified by [ignore]*)
let git_dirs target_dir ~ignore =
  let rec add_contents accu f =
    match file_exists f with
    | `Unknown | `No -> accu
    | `Yes ->
      (match List.mem ignore f ~equal:String.equal with
       | true -> accu
       | false ->
         (match Filename.basename f with
          | ".git" -> Filename.dirname f :: accu
          | _ ->
            (match is_directory f with
             | `Yes ->
               readdir f
               |> Array.map ~f:(Filename.concat f)
               |> Array.fold ~f:add_contents ~init:accu
             | _ -> accu)))
  in
  add_contents [] target_dir
;;
