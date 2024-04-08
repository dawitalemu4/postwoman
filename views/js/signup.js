window.onload = () => {

    const tokenString = localStorage.getItem('auth');

    htmx.ajax("GET", `/handle/navbar/signup/${tokenString}`, { target: "#navbar-profile", swap: "innerHTML" });
};

document.getElementById("signup-form").addEventListener("submit", async (e) => {

    e.preventDefault();

    const username = document.getElementById("signup-username").value;
    const email = document.getElementById("signup-email").value;
    const password = document.getElementById("signup-password").value;
    const response = document.getElementById("signup-response");
    const timer = document.getElementById("signup-timer");
    const emailRegex = /.+@.+\..+/;

    if (emailRegex.test(email)) {

        const createRequest = await fetch("/api/user/new", { method: "POST", body: JSON.stringify({"username": username, "email": email, "password": password, "date": `${Date.now()}`, "deleted": false}) });
        const created = await createRequest.json();

        if (typeof created === "string" && created.length > 0) {

            localStorage.setItem('auth', created);
            htmx.ajax("GET", `/handle/signup/${created}`, { target: "#signup-response", swap: "innerHTML" });

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

            htmx.ajax("GET", `/handle/signup/null`, { target: "#signup-response", swap: "innerHTML" });

            setTimeout(() => {
                response.innerHTML = "";
            }, 1500);
        };
    } else {

        response.innerHTML = "<p>$  enter a valid email</p>";

        setTimeout(() => {
            response.innerHTML = "";
        }, 1500);       
    };
});
