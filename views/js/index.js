window.onload = () => {

    const tokenString = localStorage.getItem('auth');

    htmx.ajax("GET", `/handle/navbar/home/${tokenString}`, { target: "#navbar-profile", swap: "innerHTML" });
    htmx.ajax("GET", `/handle/username/${tokenString}`, { target: "#terminal-console", swap: "innerHTML" });
};
