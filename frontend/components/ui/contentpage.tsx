import React, { useState, useEffect } from "react";

import { CodeBlock } from "./code-block.tsx"

export default function Contentpage() {
	const [data, setData] = useState(null)
	const [content, setContent] = useState(null)
	const uuid = document.location.pathname.split("/")[1];
	const url = "/api/" + uuid
	useEffect(() => {
		fetch(url)
		.then(response => { return response.json() })
		.then(json => {
			if (json.type == 1) {
				fetch(json.data)
				.then(response => {
					return response.text()
				})
				.then(plain => setContent(plain))
				.catch(error => console.error("Error fetching data:", error))
			}
			setData(json)
		})
		.catch(error => console.error("Error fetching data:", error))
	}, []);
	switch (data?.type) {
		case 0: { //console.log("link")
			return (
				<div>
					{data ? <a href={data.data}>{data.data}</a> : ''}
				</div>
			)
		}
		case 1: { //console.log("plain text")
			return (
			<div className="max-w-3xl mx-auto w-full">
				<CodeBlock
					language="txt"
					filename="plain-text"
					tabs={[
							{name: "plain text", code: content, language:"txt"}
					]}
				></CodeBlock>
			</div>
			)
		}
		case 2: { //console.log("image")
			return (
			<div>
					{data ? <img src={data.data}></img> : ''}
			</div>
			)
		}
	}
	return (
		<div>
			<p>there is nothing here besides us</p>
		</div>
	)
}
