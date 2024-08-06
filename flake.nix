{
    inputs = {
        nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";

        printPath = {
            url = "github:Ajlow2000/print-path";
            flake = true;
        };
    };

    outputs = { self, nixpkgs, printPath }:
    let
        system = "x86_64-linux";
        pkgs = nixpkgs.legacyPackages.${system};
    in
    {
        packages.${system}.print-path =
            printPath.packages.${system}.default;

        devShell.x86_64-linux =
            pkgs.mkShell {
                buildInputs = with pkgs;[
                    nil
                ];
            };
    };
}



