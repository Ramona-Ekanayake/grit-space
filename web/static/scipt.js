document.addEventListener("DOMContentLoaded", () => {
    const signInBtn = document.querySelector(".signin-btn");
    
    signInBtn.addEventListener("click", (event) => {
        event.preventDefault();
        
        const username = document.getElementById("username").value;
        const password = document.getElementById("password").value;
        
        if (!username || !password) {
            alert("Please enter both username and password.");
        } else {
            alert(`Welcome, ${username}!`);
            // Here, you can add actual authentication logic
        }
    });
});
