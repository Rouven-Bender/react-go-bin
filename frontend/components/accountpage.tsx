import React from "react";
import { useState, useEffect } from 'react';
import { TopNav } from "./homepage.tsx"

export default function Accountpage() {
	const [rowData, setRowData] = useState()

	const copyURI = (evt) => {
		evt.preventDefault();
		navigator.clipboard.writeText(evt.target.getAttribute('href'))
	}

	useEffect(() => {
		fetch("/api/account").then(response => {
			return response.json()
		}).then(json => {
			setRowData(json)
		}).catch(error => console.error("Error fetching data:", error))
	}, []) // the ,[] is importent because without it this useEffect just repeats endlessly

	return (
		<div>
		<TopNav/>
			<div className="mx-5">
			<table className="mx-auto border-spacing-5">
				<thead>
					<tr>
						<th>Id</th>
						<th className="pl-5">Link</th>
					</tr>
				</thead>
				<tbody>
				{rowData?.map((row, idx) => {
				return (
				<tr key={row.id}>
					<td>{row.id}</td>
					<td className="pl-5"><a href={document.location.host+"/"+row.id} onClick={copyURI}>Link</a></td>
				</tr>
				)})}
				</tbody>
			</table>
			</div>
		</div>
	)
}
