import React from "react";
import { getCookie, deleteCookie } from "../helper/cookies.ts"

export default function Homepage() {
	return (
		<div>
			<TopNav/>
		<div className="flex flex-col justify-center items-center h-screen">
			<h1 className="text-3x1 font-bold underline">
				This is my pastebin clone
			</h1>
		</div>
		</div>
	)
}

export function TopNav() {
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
					<AccountButton/>
					<NavItem link="/upload" name="Upload" />
				</ul>
			</div>
		</div>
		</div>
	</nav>
	)
}

export function AccountButton() {
	const logout = () => {
		deleteCookie("authToken")
	};

	const openModal = () => {
		document.getElementById("account-modal").style.display = "flex"
		document.getElementById("account-modal").showModal();
	};
	const closeModal = () => {
		document.getElementById("account-modal").style.display = "none"
		document.getElementById("account-modal").close();
	};

	if (getCookie("authToken") == "" || getCookie("authToken") == undefined) {
		return (
			<NavItem link="/login" name="login" />
		)
	}
	return (
		<li className="flex items-center p-1 text-sm gap-x-2">
			<button className="flex items-center text-slate-600" onClick={openModal}>
				Account
			</button>
			<dialog id="account-modal" className="flex flex-col hidden ml-auto mr-auto mt-auto mb-auto min-w-1/2 min-h-1/5">
				<button id="account-modal-close-btn" className="ml-auto" onClick={closeModal}>X</button>
				<a href="/" onClick={logout}>Logout</a>
			</dialog>
		</li>
	)
}

function NavItem({ link, name, onClick }) {
	if (onClick == undefined) {
	return (
		<li className="flex items-center p-1 text-sm gap-x-2">
			<a href={link} className="flex items-center text-slate-600">
				{name}
			</a>
		</li>
	)
	}
	return (
	<li className="flex items-center p-1 text-sm gap-x-2">
		<a href={link} onClick={onClick} className="flex items-center text-slate-600">
			{name}
		</a>
	</li>
	)
}
