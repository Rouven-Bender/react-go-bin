import React from "react";
import { createRoot } from 'react-dom/client';

import Homepage from "./components/homepage"
import Contentpage from "./components/contentpage.tsx"
import Login from "./components/login.tsx"
import Uploadpage from "./components/uploadpage.tsx"
import Accountpage from "./components/accountpage.tsx"
import { getCookie } from "./helper/cookies.ts"

function Application() {
	const pathname = document.location.pathname
	if (pathname.length == 1) {
		return <Homepage/>
	}
	switch (pathname.split("/")[1]) {
		case "login": {
			return <Login/>
		}
		case "upload": {
			if (getCookie("authToken") == undefined) {
				document.location = "/login?from=upload"
				break
			} else {
				return <Uploadpage/>
			}
		}
		case "account": {
			if (getCookie("authToken") == undefined) {
				document.location = "/login?from=account"
				break
			} else {
				return <Accountpage/>
			}
		}
		default: {
			return <Contentpage/>
		}
	}
	
}

//Render your React component instead
const root = createRoot(document.getElementById('app'));
root.render(<Application/>)
