{ pkgs ? (
    let
      inherit (builtins) fetchTree fromJSON readFile;
      inherit ((fromJSON (readFile ./flake.lock)).nodes) nixpkgs gomod2nix;
    in
    import (fetchTree nixpkgs.locked) {
      overlays = [
        (import "${fetchTree gomod2nix.locked}/overlay.nix")
      ];
    }
  )
}:

pkgs.buildGoApplication {
  pname = "tfmigrate";
<<<<<<< Updated upstream
  version = "0.3.23";
  pwd = ./.;
  src = ./.;
  modules = ./gomod2nix.toml;
  nativeBuildInputs = [ pkgs.terraform ];
=======
  version = "0.4.1";
  pwd = ./.;
  src = ./.;
  # CGO_ENABLED = 0;
  #nativeBuildInputs = [musl];
  # ldflags = [
  # #  # https://words.filippo.io/shrink-your-go-binaries-with-this-one-weird-trick/
  #   "-s -w"
  # #  "-X main.commit=${COMMIT} -X main.date=${DATE} -X main.version=${VERSION}"
  # ];
  # nativeBuildInputs = [ pkgs.terraform ];
>>>>>>> Stashed changes
}
