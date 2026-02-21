{
  description = "tkv - Terminal UI for tk ticket graphs";

  inputs = {
    # Use nixpkgs unstable for Go 1.25+ support
    # go.mod requires go 1.25, which isn't in stable nixpkgs yet
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };

        version = "0.14.4";

        # To update vendorHash after go.mod/go.sum changes:
        # 1. Set vendorHash to: pkgs.lib.fakeHash
        # 2. Run: nix build .#tkv 2>&1 | grep "got:"
        # 3. Replace vendorHash with the hash from "got:"
        # Updated to include pgregory.net/rapid and github.com/goccy/go-json dependencies
        # If build fails, use fakeHash method documented above to recalculate
        vendorHash = null;
      in
      {
        packages = {
          tkv = pkgs.buildGoModule {
            pname = "tkv";
            inherit version;

            src = ./.;

            inherit vendorHash;

            subPackages = [ "cmd/bv" ];

            ldflags = [
              "-s"
              "-w"
              "-X github.com/Dicklesworthstone/beads_viewer/pkg/version.version=v${version}"
            ];

            meta = with pkgs.lib; {
              description = "Terminal UI for the Beads issue tracker with graph-aware triage";
              homepage = "https://github.com/adampush/ticket_viewer";
              license = licenses.mit;
              maintainers = [ ];
              mainProgram = "tkv";
              platforms = platforms.unix;
            };
          };

          # Temporary transition alias; keep until rename cutover is fully adopted.
          bv = self.packages.${system}.tkv;
          default = self.packages.${system}.tkv;
        };

        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [
            go
            gopls
            gotools
            go-tools
            delve
          ];

          shellHook = ''
            echo "tkv development environment"
            echo "Go version: $(go version)"
            echo ""
            echo "Available commands:"
            echo "  go build ./cmd/bv  - Build tkv binary"
            echo "  go test ./...      - Run tests"
            echo "  nix build .#tkv    - Build with Nix"
          '';
        };
      }
    );
}
