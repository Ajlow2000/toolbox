(** [git_dir] recursively searches through a specificied
    [dir] and returns a the paths to all directories
    that contain a .git/ subdir. *)
let git_dirs dir =
  let rec add_contents accu filename =
    match Sys.file_exists filename with
    | true ->
      (match Filename.basename filename with
       | ".git" -> Filename.dirname filename :: accu
       | _ when Sys.is_directory filename ->
         Sys.readdir filename
         |> Array.map (Filename.concat filename)
         |> Array.fold_left add_contents accu
       | _ -> accu)
    | false -> accu
  in
  add_contents [] dir
;;
