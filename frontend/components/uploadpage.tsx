import React from "react";
import { ChangeEvent, useState } from 'react';
import { Tab, Tabs, TabList, TabPanel } from 'react-tabs';

export default function Uploadpage() {
	return (
		<div className="flex justify-center items-center h-screen">
		<div>
			<link rel="stylesheet" href="/cdn/react-tabs.css"/>
			<Tabs>
				<TabList>
					<Tab>file</Tab>
					<Tab>link</Tab>
					<Tab>paste text</Tab>
				</TabList>
				<TabPanel>
					<Fileupload/>
				</TabPanel>
				<TabPanel>
					<p className="pt-r">link</p>
				</TabPanel>
				<TabPanel>
					<p className="pt-4">paste text</p>
				</TabPanel>
			</Tabs>
		</div>
		</div>
	)
}

function Fileupload() {
	const [file, setFile] = useState<File>();
	const [uuid, setUUID] = useState(null);
	const [linkcopyed, setLinkCopyed] = useState(false);
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
		fetch("/api/upload", {
			method: "POST",
			body: file,
			headers: {
				'content-type': file.type,
				'x-filename': file.name
			}
		})
		.then(response => {return response.json()})
		.then(json => {
			setUUID(json.uuid)
		})
	}

	const copyURI = (evt) => {
		evt.preventDefault();
		navigator.clipboard.writeText(evt.target.getAttribute('href')).then(() =>{ setLinkCopyed(true) })
	}

	if (linkcopyed) {
		linkfieldtext = "copied to clipboard"
	} else {
		linkfieldtext = "copy to clipboard"
	}
	host = document.location.host

	return (
	<div className="pt-4">
	<div className="border-2 border-dotted items-center flex flex-col">
		<input 
			className="pl-4 pr-4 pt-4 block w-full text-sm file:mr-4 file:py-2 file:px-4 file:rounded-md file:border-0 file:text-sm file:font-semibold
				file:bg-pink-50 file:text-pink-700"
			type="file" name="file" onChange={handleFileChange}/><br/>
		<div className="pb-4">{file && `${file.name} - ${file.type}`}</div>
		<button className="pb-4" onClick={uploadFile}>Upload</button>
		{uuid ? <a className="px-4 pb-4" href={host+"/"+uuid} onClick={copyURI} id="link">{linkfieldtext}</a> : ''}
	</div>
	</div>
	)
}
