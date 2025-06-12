import React from "react";
import { ChangeEvent, useState } from 'react';

export default function Uploadpage() {
	const [file, setFile] = useState<File>();

	const handleFileChange = (e : ChangeEvent<HTMLInputElement>) => {
		if (e.target.files) {
			setFile(e.target.files[0]);
		}
	};

	const uploadFile = () => {
		if (!file) {
			console.log("no file")
			return
		}
		console.log(file)

		fetch("/api/upload", {
			method: "POST",
			body: file,
			headers: {
				'content-type': file.type,
				'x-filename': file.name
			}
		})
	}

	return (
	<div className="flex justify-center items-center h-screen">
	<div className="border-2 border-dotted items-center flex flex-col">
		<input 
			className="pl-4 pr-4 pt-4 block w-full text-sm file:mr-4 file:py-2 file:px-4 file:rounded-md file:border-0 file:text-sm file:font-semibold
				file:bg-pink-50 file:text-pink-700"
			type="file" name="file" onChange={handleFileChange}/><br/>
		<div className="pb-4">{file && `${file.name} - ${file.type}`}</div>
		<button className="pb-4" onClick={uploadFile}>Upload</button>
	</div>
	</div>
	)
}
