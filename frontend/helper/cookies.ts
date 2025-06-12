export default function getCookie(cname) {
	return document.cookie
		.split("; ")
		.find((row) => row.startsWith(cname+"="))
		?.split("=")[1];
}
