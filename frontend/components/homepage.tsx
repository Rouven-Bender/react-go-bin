import React from "react";
import getCookie from "../helper/cookies.ts"

export default function Homepage() {
	if (getCookie("authToken") != "") {
		text = "This is my pastebin clone and you seem to have a token"
	} else {
		text = "This is my pastebin clone"
	}
	return (
		<div>
			<TopNav/>
		<div className="flex flex-col justify-center items-center h-screen">
			<h1 className="text-3x1 font-bold underline">
				{text}
			</h1>
		</div>
		</div>
	)
}

function TopNav() {
	return (
	<nav className="block w-full max-w-screen-lg px-4 py-2 mx-auto">
		<div className="border-1 rounded-md">
		<div className="container flex flex-wrap justify-between items-center mx-auto text-slate-800 px-1">
			<a href="/" 
				className="mr-4 block cursor-pointer py-1.5 text-base text-slate-800 font-semibold">
				react-go-bin
			</a>
			<div className="items-center">
				<ul className="flex flex-row gap-2">
					<NavItem link="/login" name="Login" />
					<NavItem link="/upload" name="Upload" />
				</ul>
			</div>
		</div>
		</div>
	</nav>
	)
}

function NavItem({ link, name }) {
	return (
	<li className="flex items-center p-1 text-sm gap-x-2">
		<a href={link} className="flex items-center text-slate-600">
			{name}
		</a>
	</li>
	)
}
