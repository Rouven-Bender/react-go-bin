import React, { useState, useEffect } from "react";

import { CodeBlock } from "./code-block.tsx"

export default function Contentpage() {
	const [data, setData] = useState(null)
	const [content, setContent] = useState(null)
	const uuid = document.location.pathname.split("/")[1];
	useEffect(() => {
		fetch("/api/lookup/"+uuid)
		.then(response => { return response.json() })
		.then(json => {
			if (json.type == 1) {
				fetch("/userdata/"+json.id)
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
			window.location = data?.data
			break
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
			<div className="flex flex-col pb-3">
				<div className="pt-3">
					<div className="max-w-9/10 max-h-1/3 mx-auto object-contain border-solid border-2">
						{data ? <img className="p-1" src={"/userdata/"+data.id}></img> : ''}
					</div>
				</div>
				<div className="flex max-w-9/10 items-right mx-auto pt-4">
					{data ? <a href={"/userdata/"+data.id} download={data.data}>Download File</a> : ''}
				</div>
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
