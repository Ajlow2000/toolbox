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

        repoManager = {
            url = "github:Ajlow2000/repo-manager";
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

        zellijSessionManager = {
            url = "github:Ajlow2000/zellij-session-manager";
            flake = true;
        };

        mediaUtilities = {
            url = "github:Ajlow2000/media-utilities";
            flake = true;
        };
    };

    outputs = { self, 
        nixpkgs, 
        printPath, 
        auditDir, 
        repoManager,
        conventionalCommit,
        tmuxSessionManager,
        zellijSessionManager,
        mediaUtilities
    }:
    let
        system = "x86_64-linux";
        pkgs = nixpkgs.legacyPackages.${system};
    in
    {
        packages.${system} = {
            print-path = printPath.packages.${system}.default;
            audit-dir = auditDir.packages.${system}.default;
            repo-manager = repoManager.packages.${system}.default;
            conventional-commit = conventionalCommit.packages.${system}.default;
            tmux-session-manager = tmuxSessionManager.packages.${system}.default;
            zellij-session-manager = zellijSessionManager.packages.${system}.default;
            media-utilities = mediaUtilities.packages.${system}.default;
        };


        devShell.x86_64-linux =
            pkgs.mkShell {
                buildInputs = with pkgs;[
                    nil
                ];
            };
    };
}



