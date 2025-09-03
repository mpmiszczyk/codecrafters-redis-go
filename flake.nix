{
  description = "A Redis implementation in Go";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
        
        # Go module setup
        goModule = pkgs.buildGoModule {
          pname = "redis-go";
          version = "0.1.0";
          src = ./.;
          vendorHash = null;
        };
      in
      {
        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [
            go_1_24
            gopls
            gotools
            redis
            redli
          ];
          
          shellHook = ''
            echo "Redis Go development environment"
            echo "Available commands:"
            echo "  - go run app/main.go"
            echo "  - go build -o redis-server app/*.go"
            echo "  - ./your_program.sh"
          '';
        };
        
        packages = {
          default = goModule;
          redis-go = goModule;
        };
        
        apps.default = {
          type = "app";
          program = "${goModule}/bin/redis-go";
        };
      });
}
