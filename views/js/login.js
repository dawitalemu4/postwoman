window.onload = () => {
    
    const tokenString = localStorage.getItem('auth');

    htmx.ajax("GET", `/handle/navbar/login/${tokenString}`, { target: "#navbar-profile", swap: "innerHTML" });
};

document.getElementById("login-form").addEventListener("submit", async (e) => {
    
    e.preventDefault();

    const email = document.getElementById("login-email").value;
    const password = document.getElementById("login-password").value;
    const response = document.getElementById("login-response");
    const timer = document.getElementById("login-timer");

    const authReq = await fetch("/api/user/auth", { method: "POST", body: JSON.stringify({"email": email, "password": password, "username": "doesntmatter", "date": "doesntmatter", "deleted": false}) });
    const authenticated = await authReq.json();

    if (typeof authenticated === "string" && authenticated.length > 0) {

        localStorage.setItem('auth', authenticated);
        htmx.ajax("GET", `/handle/login/${authenticated}`, { target: "#login-response", swap: "innerHTML" });

        setTimeout(() => {
            timer.innerHTML = "<p>$  redirecting in 3 secs.</p>";
        }, 1000);

        setTimeout(() => {
            timer.innerHTML = "<p>$  redirecting in 2 secs..</p>";
        }, 2000);

        setTimeout(() => {
            timer.innerHTML = "<p>$  redirecting in 1 secs...</p>";
        }, 3000);

        setTimeout(() => {
            window.location.href = "/";
        }, 3500);

    } else {

        htmx.ajax("GET", `/handle/login/null`, { target: "#login-response", swap: "innerHTML" });
        
        setTimeout(() => {
            response.innerHTML = "";
        }, 1500);
    };
});
