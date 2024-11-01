{
    inputs = {
        nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";

        printPath = {
            url = "github:Ajlow2000/print-path";
            flake = true;
        };

        auditDir = {
            url = "github:Ajlow2000/audit-dir";
            flake = true;
        };

        addRepo = {
            url = "github:Ajlow2000/add-repo";
            flake = true;
        };

        conventionalCommit = {
            url = "github:Ajlow2000/conventional-commit";
            flake = true;
        };

        tmuxSessionManager = {
            url = "github:Ajlow2000/tmux-session-manager";
            flake = true;
        };
    };

    outputs = { self, 
        nixpkgs, 
        printPath, 
        auditDir, 
        addRepo,
        conventionalCommit,
        tmuxSessionManager
    }:
    let
        system = "x86_64-linux";
        pkgs = nixpkgs.legacyPackages.${system};
    in
    {
        packages.${system} = {
            print-path = printPath.packages.${system}.default;
            audit-dir = auditDir.packages.${system}.default;
            add-repo = addRepo.packages.${system}.default;
            conventional-commit = conventionalCommit.packages.${system}.default;
            tmux-session-manager = tmuxSessionManager.packages.${system}.default;
        };


        devShell.x86_64-linux =
            pkgs.mkShell {
                buildInputs = with pkgs;[
                    nil
                ];
            };
    };
}



