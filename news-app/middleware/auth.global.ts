import { jwtDecode } from "jwt-decode";

function isValidAuthToken(token: string): boolean {
  const decoded = jwtDecode(token);

  if (!decoded || !decoded.exp) return false;

  const now = Math.floor(Date.now() / 1000);

  return now < decoded.exp;
}

export default defineNuxtRouteMiddleware((to, _) => {
  const loggedIn = useState("loggedIn");
  if (loggedIn.value)
    return;

  const authCookie = useCookie("authToken");
  const isValid = authCookie.value ? isValidAuthToken(authCookie.value) : false;

  loggedIn.value = isValid;

  if (!isValid && to.path !== "/auth") {
    console.log("Rerouting");
    return navigateTo("/auth");
  }
});
