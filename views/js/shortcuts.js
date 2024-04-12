const currentPage = window.location.pathname;

const shortuctKeys = {
    "empty": "KeyJ",
    "history": "KeyH",
    "toggleFavorite": "KeyF",
    "viewFavorites": "KeyV",
    "hideRequest": "KeyD",
    "closeModal": "KeyQ",
    "home": "KeyH",
    "login": "KeyL",
    "signup": "KeyS",
    "profile": "KeyP",
    "delete": "KeyD",
    "logout": "KeyL"
};

const homeKeys = [
    shortuctKeys["empty"],
    shortuctKeys["history"],
    shortuctKeys["toggleFavorite"],
    shortuctKeys["viewFavorites"],
    shortuctKeys["hideRequest"],
    shortuctKeys["closeModal"],
    shortuctKeys["login"],
    shortuctKeys["signup"],
    shortuctKeys["profile"],
    shortuctKeys["logout"]
];

const loginKeys = [
    shortuctKeys["home"],
    shortuctKeys["signup"]
];

const signupKeys = [
    shortuctKeys["home"],
    shortuctKeys["login"]
];

const profileKeys = [
    shortuctKeys["delete"],
    shortuctKeys["home"],
    shortuctKeys["logout"]
];

const executeHomeShortcuts = (shortcut) => {

    const loggedIn = localStorage.getItem("auth");

    if (shortcut === shortuctKeys["empty"]) {
        emptyForm();
    } else if (shortcut === shortuctKeys["history"]) {
        toggleHistoryList();
    } else if (shortcut === shortuctKeys["toggleFavorite"]) {
        toggleFavoriteItem();
    } else if (shortcut === shortuctKeys["viewFavorites"]) {
        toggleFavoritesList();
    } else if (shortcut === shortuctKeys["hideRequest"]) {
        hideRequest();
    } else if (shortcut === shortuctKeys["closeModal"]) {
        document.getElementById("history-modal").style.display = "none";
        document.getElementById("favorites-modal").style.display = "none";
    } else if (shortcut === shortuctKeys["login"] && loggedIn === null) {
        window.location.href = "/login";
    } else if (shortcut === shortuctKeys["signup"] && loggedIn === null) {
        window.location.href = "/signup";
    } else if (shortcut === shortuctKeys["profile"] && loggedIn != null) {
        window.location.href = "/profile";
    } else if (shortcut === shortuctKeys["logout"] && loggedIn != null) {
        localStorage.removeItem("auth");
        window.location.href = "/";
    };
};

const executeLoginShortcuts = (shortcut) => {
    
    if (shortcut === shortuctKeys["home"]) {
        window.location.href = "/";
    } else if (shortcut === shortuctKeys["signup"]) {
        window.location.href = "/signup";
    };
};

const executeSignupShortcuts = (shortcut) => {
    
    if (shortcut === shortuctKeys["home"]) {
        window.location.href = "/";
    } else if (shortcut === shortuctKeys["login"]) {
        window.location.href = "/login";
    };
};

const executeProfileShortcuts = (shortcut) => {

    if (shortcut === shortuctKeys["delete"]) {
        deleteProfile();
    } else if (shortcut === shortuctKeys["home"]) {
        window.location.href = "/";
    } else if (shortcut === shortuctKeys["logout"]) {
        localStorage.removeItem("auth");
        window.location.href = "/";
    };
};

document.addEventListener("keydown", (e) => {

    if (currentPage === "/") {

        for (let i = 0; i < homeKeys.length; i++) {
            if ((e.metaKey || e.ctrlKey) && e.altKey && e.code === homeKeys[i]) {
                executeHomeShortcuts(homeKeys[i]);
            };
        };
    } else if (currentPage === "/login") {

        for (let i = 0; i < loginKeys.length; i++) {
            if ((e.metaKey || e.ctrlKey) && e.altKey && e.code === loginKeys[i]) {
                executeLoginShortcuts(loginKeys[i]);
            };
        };
    } else if (currentPage === "/signup") {

        for (let i = 0; i < signupKeys.length; i++) {
            if ((e.metaKey || e.ctrlKey) && e.altKey && e.code === signupKeys[i]) {
                executeSignupShortcuts(signupKeys[i]);
            };
        };
    } else if (currentPage === "/profile") {

        for (let i = 0; i < profileKeys.length; i++) {
            if ((e.metaKey || e.ctrlKey) && e.altKey && e.code === profileKeys[i]) {
                executeProfileShortcuts(profileKeys[i]);
            };
        };
    };
});
