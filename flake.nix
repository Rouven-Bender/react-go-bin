{
  description = "golang pastebin with react frontend";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
  };

	outputs = { self, nixpkgs }:
	let
		system = "x86_64-linux";
		pkgs = nixpkgs.legacyPackages.${system};
	in
	{
		devShells.${system}.default  = pkgs.mkShell
		{
			buildInputs = [
				pkgs.go
				pkgs.yarn
				pkgs.tailwindcss_4
				pkgs.sqlite
				pkgs.nodejs_24
			];

			shellHook = ''
				alias run="yarn run build && go run . serve"
				echo 'use "run" to build and launch the server'
			'';
		};
	};
}
