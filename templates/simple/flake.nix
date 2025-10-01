{
  description = "A basic multiplatform compiscript Nix setup";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
    compiscript.url = "github:ElrohirGT/5318008Lang";
  };

  outputs = {
    nixpkgs,
    compiscript,
    ...
  }: let
    # System types to support.
    supportedSystems = ["x86_64-linux" "x86_64-darwin" "aarch64-linux" "aarch64-darwin"];

    # Helper function to generate an attrset '{ x86_64-linux = f "x86_64-linux"; ... }'.
    forAllSystems = nixpkgs.lib.genAttrs supportedSystems;

    # Nixpkgs instantiated for supported system types.
    nixpkgsFor = forAllSystems (system: import nixpkgs {inherit system;});
  in {
    devShells = forAllSystems (system: let
      pkgs = nixpkgsFor.${system};
      compiler = compiscript.outputs.packages.${system}.default;
    in {
      default = pkgs.mkShell {
        packages = [
          # Compiscript compiler
          compiler

          # Other dev packages you want... for example nodejs?
          pkgs.nodejs

          # TODO: Add other packages if you want!
        ];
      };
    });
  };
}
