window.onload = function() {
    
    const tokenString = localStorage.getItem('auth');

    htmx.ajax("GET", `/handle/navbar/profile/${tokenString}`, { target: "#navbar-profile", swap: "innerHTML" });
};
