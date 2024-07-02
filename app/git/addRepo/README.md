# add-repo
A utility for cloning repositories with standard name.

By default, this utility expects a url to be specified with 
'--url' and defaults to a clone location of $HOME/repos.
This location can be overridden with '--path'. The repo will
be cloned into a directory of the custom name username_repo.

To make using this tool more ergonomic, create a shell alias
similar to `alias add-repo="toolbox git add-repo --path $HOME/repos --url`.
This allows it to be used like `add-repo [url]`.
