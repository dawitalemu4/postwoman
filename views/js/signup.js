window.onload = function() {
    
    const tokenString = localStorage.getItem('auth');
    
    fetch(`/handle/navbar/${tokenString}`);
    fetch(`//${tokenString}`);
};
