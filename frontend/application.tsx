import React from "react";
import { createRoot } from 'react-dom/client';

import Homepage from "./components/homepage";
import Contentpage from "./components/contentpage.tsx";
import Login from "./components/login.tsx"


function Application() {
	const pathname = document.location.pathname
	if (pathname.length == 1) {
		return <Homepage/>
	}
	if (pathname.split("/")[1] == "login") {
		return <Login/>
	}
	return <Contentpage/>
	
}

//Render your React component instead
const root = createRoot(document.getElementById('app'));
root.render(<Application/>)
