window.onload = () => {

    const tokenString = localStorage.getItem('auth');

    htmx.ajax("GET", `/handle/navbar/home/${tokenString}`, { target: "#navbar-profile", swap: "innerHTML" });
    htmx.ajax("GET", `/handle/shortcut/${tokenString}`, { target: "#shortcuts", swap: "beforeend" });
    
    setTimeout(() => {
        htmx.ajax("GET", `/handle/username/${tokenString}`, { target: "#terminal-console", swap: "beforeend" });
    }, 600);

    // setTimeout(() => {
    //     htmx.ajax("GET", `/handle/request/new`, { target: "#terminal-console", swap: "beforeend" });
    // }, 1200);
};
