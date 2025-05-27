import React from "react";
import { createRoot } from 'react-dom/client';

function Application() {
	return (
		<div>Got React frontend to compile</div>
	)
}

//Render your React component instead
const root = createRoot(document.getElementById('app'));
root.render(<Application/>)
