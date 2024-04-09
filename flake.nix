{
  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
  inputs.flake-compat.url = "github:edolstra/flake-compat";
  inputs.flake-compat.flake = false;
  inputs.flake-utils.url = "github:numtide/flake-utils";
  inputs.nix-filter.url = "github:numtide/nix-filter";

  outputs = {
    self,
    nixpkgs,
    flake-compat,
    flake-utils,
    nix-filter,
    } @ inputs:
    flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
        broom = pkgs.buildGoModule {
          pname = "broom";
          version = "0.2.4";

          src = nix-filter.lib {
            root = ./.;
            include = [
              "cmd/"
              "go.mod"
              "go.sum"
              "main.go"
            ];
          };

          vendorHash = "sha256-zNklqGjMt89b+JOZfKjTO6c75SXO10e7YtQOqqQZpnA=";
        };
      in
      {
        devShells = {
          default = pkgs.mkShell {
            packages = [ pkgs.go ];
            inputsFrom = [ self.packages ];
          };
        };
        packages = {
          inherit broom;
        };
        apps = {
          default = {
            type = "app";
            program = "${broom}/bin/broom";
          };
        };
      }
    );
}
