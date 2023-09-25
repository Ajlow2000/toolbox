open Lib

let dir = "./foo"


let () = match Sys.file_exists dir with
    | false -> Printf.printf "Specified directory (\"%s\") does not exist\n" dir
    | true -> match FsHelpers.dir_is_empty dir with
        | true -> print_endline "Directory is empty"
        | false -> List.iter print_endline (FsHelpers.dir_contents dir)






