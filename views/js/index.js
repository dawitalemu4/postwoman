window.onload = async function() {
    
    const tokenString = localStorage.getItem('auth');

    if (!tokenString) {
        const navbar = await fetch(`/handle/navbar/empty`);
        console.log("navbar", navbar);
        const username = await fetch(`/handle/username/empty`);
        console.log("username", username);
    } else {
        console.log("hi", tokenString);
        const navbar = await fetch(`/handle/navbar/${tokenString}`);
        console.log("navbar", navbar);
        const username = await fetch(`/handle/username/${tokenString}`);
        console.log("username", username);
    };
};
