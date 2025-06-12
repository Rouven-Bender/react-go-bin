import React from "react";

export default function Login() {
	const lifetime = 2 * 60 * 60
	function submit(formData) {
		let key = formData.get("key")
		fetch("/api/login", {
			method: "POST",
			body: JSON.stringify({
				key: key
			}),
			headers: {
				"Content-type": "application/json"
			}
		})
		.then(response => {return response.json()})
		.then(json => {
			document.cookie = "authToken="+json.token+";max-age=" + lifetime
			if (document.location.search != "") {
				let paras = document.location.search.split("=")
				if (paras[0] = "?from"){
					document.location = "/"+paras[1]
					return
				}
			}
			document.location = "/"
		})
	}

	return (
		<div className="flex justify-center items-center h-screen">
		<div className="border-dotted border-2 rounded-full">
			<form action={submit}>
				<label className="pl-5">Secret Key:</label>
				<input className="pl-5" type="password" maxLength="30" name="key"/>
				<button className="pl-5 pr-5" type="submit">Log in</button>
			</form>
		</div>
		</div>
	)
}
