import React from "react";
import { createRoot } from 'react-dom/client';

function Application() {
	return (
		<div>
		<h1 class="text-3x1 font-bold underline">
			Header for tailwind test
		</h1>
		<div>Got React frontend to compile</div>
		</div>
	)
}

//Render your React component instead
const root = createRoot(document.getElementById('app'));
root.render(<Application/>)
