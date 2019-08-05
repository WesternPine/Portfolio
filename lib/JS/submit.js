function init(){
  var button = document.getElementById('submit');
  button.addEventListener('click', validate);
}

function validate(){
  var button = document.getElementById('submit');
  
  if(!validateFields()){
    button.classList.remove("modifiedButton");
    button.style.cursor = "pointer";
    button.addEventListener('click', validate);
    return false;
  }
  
  var name = document.getElementById('name').value;
  var email = document.getElementById('email').value;
  var subject = document.getElementById('subject').value;
  var message = document.getElementById('message').value;
  var status = document.getElementById('status');
  
  var url = "https://wasternpine.dev:81/forms";
  var data = '{'
  + '"name" : "' + name + '",'
  + '"email" : "' + email + '",'
  + '"subject" : "' + subject + '",'
  + '"message" : "' + message + '"}'
  
  var xmlhttp;
  if(window.XMLHttpRequest){
    xmlhttp = new XMLHttpRequest();
  } else {
    xmlhttp = new ActiveXObject("MicrosoftXMLHTTP");
  }
  
  xmlhttp.onreadystatechange=function(){
    if(xmlhttp.readyState==4){
      var obj = xmlhttp.responseText;
      if(obj != ""){
        if(obj == "true"){
          //clear fields 
          document.getElementById('name').value = "";
          document.getElementById('email').value = "";
          document.getElementById('subject').value = "";
          document.getElementById('message').value = "";
          
          status.innerHTML = "Your message was sent!";
          button.style.cursor = "not-allowed";
          return true;
        }
      }
      //false or something...
      status.innerHTML = "An error occured. Please try again later.";
      button.classList.remove("modifiedButton");
      button.style.cursor = "pointer";
      button.addEventListener('click', validate);
    }
  }
  
  xmlhttp.open("POST", url, true);
  xmlhttp.setRequestHeader("Access-Control-Allow-Origin", "true");
  xmlhttp.setRequestHeader("Content-Type","application/x-www-form-urlencoded");
  xmlhttp.setRequestHeader("Accept", "application/json");
  
  xmlhttp.send(data)
}

function validateFields(){
  var button = document.getElementById('submit');
  
  button.removeEventListener('click', validate);
  button.classList.add("modifiedButton");
  button.style.cursor = "wait";
  
  var name = document.getElementById('name').value;
  var email = document.getElementById('email').value;
  var subject = document.getElementById('subject').value;
  var message = document.getElementById('message').value;
  var status = document.getElementById('status');
  
  status.innerHTML = "Processing...";
  
  if(name.length < 3) {
    status.innerHTML = "Your name must be at least 3 characters.";
    return false;
  }
  
  if(email.indexOf("@") == -1 || email.length < 6){
    status.innerHTML = "Please enter valid E-Mail.";
    return false;
  }
  
  if(subject.length < 2) {
    status.innerHTML = "Please enter a subject.";
    return false;
  }
  
  if(message.length < 2) {
    status.innerHTML = "Your message must be at least 15 characters.";
    return false;
  }
  
  if(message.length > 2500) {
    status.innerHTML = "Your message must not exceed 2,500 characters.";
    return false;
  }
  
  return true;
}

init();