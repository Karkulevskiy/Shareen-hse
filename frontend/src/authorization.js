import axios from "axios";
import {connection,routeEvent} from "./websocket.js"
import { Event } from "./classes/events.js";
import { MyAlert } from "./utils.js";

//HTML формы для входа
const signInHTML = `<div id="blur"></div>
        <form class="mui-form">
            <legend style="margin-top:10px;caret-color:transparent;">Authorization</legend>
            <small style="font-size: 130%;caret-color:transparent;">Enter your login and password</small>
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
            <small style="font-size: 150%;position:relative;top:20px;caret-color:transparent;">No account?</small>
            <br>
            <button style="color:blue;background-color:transparent;border:none;font-size:18px;margin-top:22px;caret-color:transparent;" id="registration-btn">Sign up!</button>
        </form>`

//HTML формы для регистрации
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
        <button style="color:blue;background-color:transparent;border:none;font-size:18px;margin-top:22px" id="registration-btn">Sign in!</button>`

function signHandler(event){ //Обработчик нажатия на кнопку "Войти/Регистрация" в форме
    event.preventDefault();
    const $signbtn = document.querySelector("#submit-btn");
    $signbtn.disabled = true; //Временно морозим нажатую кнопку
    const data = event.target.getElementsByTagName("input");
    let UserData = { //Берём данные с инпута
        "login":data[0].value,
        "password":data[1].value
    }
    if ($signbtn.textContent=="Sign in"){ //Если осуществлялся вход, то кидаем запрос на вход на бэкенд
        let MaxURL = "http://localhost:8080/login";
        axios.post(MaxURL,JSON.stringify(UserData))
        .then(response =>{
            let ans = response.data;
            if (ans.status==200){
                connectWebsocket(ans.payload.otp,UserData.login);
            }
        })
        .catch(error => {
            let ans = error.response;
            if(ans.status==400){
                MyAlert("Wrong login or password","error");
            }
            else if(ans.status==406){
                MyAlert("This user had already authorised","error");
            }
            else if(ans.status==500){
                MyAlert("Internal server error","error");
            }
            console.log("UNEXPECTED ERROR: "+error);
        });
    }
    else{ //Если осуществлялась регистрация, то кидаем запрос на регистрацию на бэкенд
        const MaxURL = "http://localhost:8080/register";
        axios.post(MaxURL,JSON.stringify(UserData))
        .then(response => {
            let ans = response.data
            if (ans.status==200){
                MyAlert("You have successfully signed up!","success");
                showSignInForm(null);
            }
            else if(ans.status==400){
                MyAlert("This login has already been registered!","error")
            }
            else{
                MyAlert("Internal server error","error");
            }
        }).catch(error => {
            console.log("Error: " + error);
        });
    }
    $signbtn.disabled = false; //Снова включаем кнопку
}

export function showSignInForm(event){ //Показать форму входа и добавить обработчики событий на неё
    if (connection.length==1){
        MyAlert("You have already signed in!","error");
    }
    event.preventDefault;
    let $app =  document.querySelector("#app");
    if (event==null){
        document.getElementsByClassName("mui-form").remove();
        document.getElementById("blur").remove();
    }
    else if (event.target.form!=null){
        event.target.form.remove();
        document.getElementById("blur").remove();
    }
    $app.insertAdjacentHTML("beforeend",signInHTML);
    const $btn = document.querySelector("#submit-btn");
    $btn.addEventListener("submit",signHandler);
    const $form = document.querySelector(".mui-form"); //Комментарий
    const $back = document.getElementById("blur");
    $back.addEventListener("click",function(){
        $form.remove();
        $back.remove();
    })
    $form.addEventListener("submit",signHandler);
    document.querySelector("#registration-btn").addEventListener("click",showSighUpForm);
}

function showSighUpForm(event){ //Показать форму регистрации и добавить обработчики событий на неё
    event.preventDefault;
    const $form = event.target.form;
    const $back = document.getElementById("blur");
    $back.addEventListener("click",function(){
        $form.remove();
        $back.remove();
    })
    $form.innerHTML = signUpHTML;
    $form.addEventListener("submit",signHandler);
    document.querySelector("#registration-btn").addEventListener("click",showSignInForm);
}

function connectWebsocket(otp,login){
    if (window["WebSocket"]) {
        // Подключаемся к вебсокету через одноразовый пароль
        const MaxURL = "ws://localhost:8080/ws?otp="+otp;
        conn = new WebSocket(MaxURL);
        // При открытии вебсокета сохраняем данные юзера в localstorage
        conn.onopen = function (event) {
            localStorage.setItem("login",login);
            document.getElementById('blur').click();
            MyAlert("You have successfully signed in!","success");
        }

        conn.onclose = function (event) {
            console.log("Disconnected from WebSocket");
        }

        // При появлении сообщения через вебсокет 
        
        conn.onmessage = function (event) {
            // Парсим сообщение как JSON
            const eventData = JSON.parse(event.data);
            // Отправляем сообщение на определение типа события
            const NewEvent = Object.assign(new Event,eventData);
            routeEvent(NewEvent);
        }
        connection.push(conn);

    } else {
        MyAlert("This browser is not supporting websockets!", "error");
    }
}