open Lib.FsHelpers

let target_dir = match Sys.getenv_opt "HOME" with
    | Some v -> v
    | None -> "./"

let git_repos = match Sys.file_exists target_dir with
    | false -> []
    | true -> git_dirs target_dir

let () = List.iter print_endline (git_repos)
