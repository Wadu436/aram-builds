export function createAuthHeaders(username: string, password: string) {
  // Create headers
  let headers = new Headers();
  headers.set("Authorization", "Basic " + btoa(username + ":" + password));
  return headers;
}

const SPECIAL_CHAR_REGEX = /[^a-zA-Z]/g;
export function canonicalizeString(str: string) {
  return str.toLowerCase().replace(SPECIAL_CHAR_REGEX, "");
}
