(** [dir_is_empty dir] is true, if [dir] contains no files except
 * "." and ".."
 *)
let dir_is_empty dir =
    Array.length (Sys.readdir dir) = 0
;;

(** [dir_contents] returns the paths of all regular files that are
 * contained in [dir]. Each file is a path starting with [dir].
  *)
let dir_contents dir =
    let rec loop result worklist = match worklist with
        | f::fs when Sys.is_directory f ->
            Sys.readdir f
            |> Array.to_list
            |> List.map (Filename.concat f)
            |> List.append fs
            |> loop result
        | f::fs -> loop (f::result) fs
        | []    -> result
    in
    loop [] [dir]
;;

let dir_contents_fast dir =
    let rec add_contents accu filename =
        if Sys.is_directory filename then
            Sys.readdir filename
            |> Array.map (Filename.concat filename)
            |> Array.fold_left add_contents accu
        else filename :: accu
    in add_contents [] dir
;;

            (* match filename with *)
            (* | ".git" -> filename :: accu *)
            (* | _ -> Array.fold_left add_contents accu *)
let git_dirs dir =
    let rec add_contents accu filename =
        if Sys.is_directory filename then
            match Sys.readdir
    in add_contents [] dir
