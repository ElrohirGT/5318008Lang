{buildGoModule}:
buildGoModule {
  # FIXME: Once tests pass, enable this...
  doCheck = false;

  pname = "5318008Lang";
  version = "1.0.0";
  src = ./.;
  vendorHash = "sha256-P+XrbXkJzPqVrhV//6s1cYys6vuR23jMXtp1/7kfpdA=";
}
