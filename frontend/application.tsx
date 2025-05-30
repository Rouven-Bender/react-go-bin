import React, { useState, useEffect } from "react";
import { createRoot } from 'react-dom/client';

function Homepage() {
	return (
	<body>
		<div class="flex flex-col justify-center items-center h-screen">
			<h1 class="text-3x1 font-bold underline">
				This is my pastebin clone
			</h1>
		</div>
	</body>
	)
}

function Contentpage() {
	const [data, setData] = useState(null)
	const uuid = document.location.pathname.split("/")[1];
	const url = "/api/" + uuid
	useEffect(() => {
		fetch(url)
		.then(response => { return response.json() })
		.then(json => setData(json))
		.catch(error => console.error("Error fetching data:", error))
	}, []);
	return (
		<div>
			{data ? <p>{data.data}</p> : 'Loading...'}
		</div>
	)
}

function Application() {
	const pathname = document.location.pathname
	if (pathname.length == 1) {
		return <Homepage/>
	} else {
		return <Contentpage/>
	}
}

//Render your React component instead
const root = createRoot(document.getElementById('app'));
root.render(<Application/>)
