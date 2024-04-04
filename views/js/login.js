window.onload = function() {
    
    const tokenString = localStorage.getItem('auth');

    fetch(`/handle/navbar/${tokenString}`);
};

document.body.addEventListener('htmx:afterRequest', function(e) {
    const res = e.detail.xhr.response;
    if (res.status === 200) {
        localStorage.setItem('auth', res.response);
        fetch(`/handle/login/${res}`);
        // display in dom or use htmx
    } else {
        // incorrect login creds
    };
});

