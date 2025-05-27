const esbuild = require("esbuild");

esbuild
	.build({
		entryPoints: ["./frontend/Application.tsx"],
		outdir: "./static/",
		bundle: true,
		minify: true
	})
	.then(() => console.log("build complete!"))
	.catch(() => process.exit(1));
