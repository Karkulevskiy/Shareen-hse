import axios from "axios";
import { show } from "./utils.js";
import {routeEvent} from "./websocket.js"

const signInHTML = `<div id="blur" onclick=""></div>
        <form class="mui-form">
            <legend style="margin-top:10px">Authorisation</legend>
            <small style="font-size: 130%;">Enter your login and password</small>
            <div class="mui-textfield mui-textfield--float-label">
                <input type="text" placeholder="Login">
                <label>Login</label>
            </div>
            <div class="mui-textfield mui-textfield--float-label">
                <input placeholder="Password" type="text">
                <label>Password</label>
            </div>
            <button type="submit" class="mui-btn mui-btn--raised mui-btn--danger" id="submit-btn">Sign in</button>
            <br>
            <small style="font-size: 150%;position:relative;top:20px">No account?</small>
            <br>
            <button style="color:blue;background-color:transparent;border:none;font-size:18px;margin-top:22px" id="registration-btn">Sign up!</button>
        </form>`

        const signUpHTML = `<legend style="margin-top:10px">Registration</legend>
        <small style="font-size: 130%;">Enter your login and password</small>
        <div class="mui-textfield mui-textfield--float-label">
            <input type="text" placeholder="Login">
            <label>Login</label>
        </div>
        <div class="mui-textfield mui-textfield--float-label">
            <input placeholder="Password" type="text">
            <label>Password</label>
        </div>
        <button type="submit" class="mui-btn mui-btn--raised mui-btn--danger" id="submit-btn">Sign up</button>
        <br>
        <small style="font-size: 150%;position:relative;top:20px">Already have an account?</small>
        <br>
        <button style="color:blue;background-color:transparent;border:none;font-size:18px;margin-top:22px">Sign in!</button>`

function signHandler(event){
    event.preventDefault();
    const $signbtn = document.querySelector("#submit-btn");
    $signbtn.disabled = true;
    const data = event.target.getElementsByTagName("input");
    let UserData = {
        "login":data[0].value,
        "password":data[1].value
    }
    if ($signbtn.value=="Sign in"){
        let MaxURL = "http://localhost:8080/login";
        const headers = {
            'Content-Type': 'application/json',
            "Access-Control-Allow-Origin":"*"
        }
        axios.post(MaxURL,JSON.stringify(UserData),{
            headers:headers
        })
        .then(response =>{
            let ans = JSON.parse(response);
            if (ans.status==200){
                connectWebsocket(ans.payload.otp);
            }
            else{
                alert("Ошибка!");
            }
        })
        .catch(error => {
            console.log("Pizdec:" + error);
        });
    }
    else{
        const MaxURL = "http://localhost:8080/register";
        axios.post(MaxURL,JSON.stringify(UserData))
        .then(response => {
            if (JSON.parse(response).status==200){
                localStorage.setItem("login",UserData.login);
                alert("Кайф братишка!");
            }
            else{
                alert("Ошибка!")
            }
<<<<<<< HEAD
        })
        .catch(error => {
=======
        }).catch(error => {
>>>>>>> 5ca3b151a8f7ca2024a567212eab486f7dfd1cfe
            console.log("Pizdec:" + error);
        });
    }
}

export function showSignInForm(event){
    const $app = document.querySelector("#app");
    $app.innerHTML+=signInHTML;
    const $btn = document.querySelector("#submit-btn");
    $btn.addEventListener("submit",signHandler);
    const $form = document.querySelector(".mui-form"); //Комментарий
    const $back = document.getElementById("blur");
    show("block");
    $back.addEventListener("click",function(){
        show("none");
    })
    $form.addEventListener("submit",signHandler);
    document.querySelector("#registration-btn").addEventListener("click",showSighUpForm);
}

function showSighUpForm(event){
    event.preventDefault;
    const form = event.target.form;
    form.innerHTML = signUpHTML;

}

function connectWebsocket(otp){
    if (window["WebSocket"]) {
        // Connect to websocket using OTP as a GET parameter
        const MaxURL = "ws://localhost:8080/ws?otp="+otp;
        conn = new WebSocket(MaxURL);

        // Onopen
        conn.onopen = function (event) {
            console.log("Connected to WebSocket")
        }

        conn.onclose = function (event) {
            console.log("Disconnected from WebSocket");
        }

        // Add a listener to the onmessage event
        conn.onmessage = function (event) {
            console.log(event);
            // parse websocket message as JSON
            const eventData = JSON.parse(event.data);
            // Assign JSON data to new Event Object
            const NewEvent = Object.assign(new Event, eventData);
            // Let router manage message
            routeEvent(NewEvent);
        }

    } else {
        alert("Not supporting websockets");
    }
}