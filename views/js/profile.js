const parseJwt = (token) => {

    const base64Url = token.split('.')[1];
    const base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
    const jsonPayload = decodeURIComponent(window.atob(base64).split('').map(function(c) {
        return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2);
    }).join(''));

    return JSON.parse(jsonPayload);
};

window.onload = function() {

    const tokenString = localStorage.getItem('auth');
    const profile = parseJwt(tokenString);
    const username = document.getElementById("profile-username");
    const password = document.getElementById("profile-password");

    htmx.ajax("GET", `/handle/navbar/profile/${tokenString}`, { target: "#navbar-profile", swap: "innerHTML" });
    htmx.ajax("GET", `/handle/profile/info/${tokenString}`, { target: "#profile-info", swap: "innerHTML" });

    username.value = profile.username;
    password.value = profile.password;
};

document.getElementById("profile-form").addEventListener("submit", async (e) => {

    e.preventDefault();

    const profile = parseJwt(localStorage.getItem('auth'));

    const username = document.getElementById("profile-username");
    const password = document.getElementById("profile-password");
    const response = document.getElementById("profile-response");
    const timer = document.getElementById("profile-timer");

    const updateReq = await fetch("/api/user/update", { method: "PUT", body: JSON.stringify({
        "username": username.value,
        "email": profile.email,
        "password": password.value,
        "favorites": profile.favorites,
        "date": profile.date,
        "deleted": false,
        "oldPassword": profile.password
    })});
    const updatedProfile = await updateReq.json();

    if (typeof updatedProfile === "string" && updatedProfile.length > 0) {

        localStorage.setItem('auth', updatedProfile);
        htmx.ajax("GET", `/handle/profile/update/${updatedProfile}`, { target: "#profile-response", swap: "innerHTML" });

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

        htmx.ajax("GET", `/handle/profile/update/null`, { target: "#profile-response", swap: "innerHTML" });

        setTimeout(() => {
            response.innerHTML = "";
        }, 1500);
    };
});

const deleteProfile = async () => {

    const profile = parseJwt(localStorage.getItem('auth'));

    const response = document.getElementById("profile-response");
    const timer = document.getElementById("profile-timer");

    const deleteReq = await fetch("/api/user/delete", { method: "DELETE", body: JSON.stringify({ "username": profile.username, "email": profile.email, "password": profile.password, "deleted": false })});
    const deletedProfile = await deleteReq.json();

    if (deletedProfile === true) {

        localStorage.clear();
        htmx.ajax("GET", `/handle/profile/delete/${deletedProfile}`, { target: "#profile-response", swap: "innerHTML" });

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

        htmx.ajax("GET", `/handle/profile/delete/null`, { target: "#profile-response", swap: "innerHTML" });

        setTimeout(() => {
            response.innerHTML = "";
        }, 1500);
    };
};
