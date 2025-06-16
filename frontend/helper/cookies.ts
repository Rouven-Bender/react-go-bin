export function getCookie(cname) {
	return document.cookie
		.split("; ")
		.find((row) => row.startsWith(cname+"="))
		?.split("=")[1];
}
export function deleteCookie(cname) {
	document.cookie = cname + "=;Max-Age=0"
}
