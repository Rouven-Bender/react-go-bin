import React from "react";
import getCookie from "../helper/cookies.ts"

export default function Homepage() {
	if (getCookie("authToken") != "") {
		text = "This is my pastebin clone and you seem to have a token"
	} else {
		text = "This is my pastebin clone"
	}
	return (
	<div className="flex flex-col justify-center items-center h-screen">
		<h1 className="text-3x1 font-bold underline">
			{text}
		</h1>
	</div>
	)
}
