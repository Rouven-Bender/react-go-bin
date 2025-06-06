import React from "react";
import { createRoot } from 'react-dom/client';
import { BrowserRouter } from 'react-router-dom';

import Homepage from "./components/ui/homepage";
import Contentpage from "./components/ui/contentpage.tsx";


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
root.render(
	<BrowserRouter>
		<Application/>
	</BrowserRouter>
)
