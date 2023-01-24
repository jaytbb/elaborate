function hello() {
    document.getElementById("greeting").innerHTML = "Hello, JavaScript!";
    console.log("meow :3");
}

function click() {
    document.getElementById("greeting").innerHTML = "Hello, mouse click!";
}

const button = document.getElementById("greeting-button");

button.addEventListener("click", click); 
setTimeout(hello, 2000);
